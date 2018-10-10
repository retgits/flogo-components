// Package globalcache implements a wrapper around go-cache for Flogo. go-cache is an in-memory key:value
// store/cache similar to memcached that is suitable for applications running on a single machine.
// Its major advantage is that, being essentially a thread-safe map[string]interface{} with expiration
// times, it doesn't need to serialize or transmit its contents over the network.
package globalcache

// The imports used for this activity
import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	cache "github.com/patrickmn/go-cache"
)

const (
	// Parameter representing the input value 'key'
	ivKey = "key"
	// Parameter representing the input value 'value'
	ivValue = "value"
	// Parameter representing the input value 'action'
	ivAction = "action"
	// Parameter representing the input value 'expiryTime'
	ivExpiryTime = "expiryTime"
	// Parameter representing the input value 'purgeTime'
	ivPurgeTime = "purgeTime"
	// Parameter representing the input value 'loadset'
	ivLoadset = "loadset"
	// Parameter representing the output value 'result'
	ovResult = "result"
	// CacheName represents the name of the cache as it exists in the Flogo engine
	CacheName = "GlobalCache"
	// DefaultExpiryTime is the default expiry time if no cache was created
	DefaultExpiryTime = 5
	// DefaultPurgeTime is the default purge time if no cache was created
	DefaultPurgeTime = 10
)

// log is the default package logger
var log = logger.GetLogger("activity-cache")

// GlobalCache is an activity that governs a single engine cache
type GlobalCache struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &GlobalCache{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *GlobalCache) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *GlobalCache) Eval(context activity.Context) (done bool, err error) {
	// Get the action
	action := context.GetInput(ivAction).(string)

	switch action {
	case "INITIALIZE":
		// Get the input values
		expiryTime := context.GetInput(ivExpiryTime).(int)
		purgeTime := context.GetInput(ivPurgeTime).(int)

		// Execute the initialize action
		var c *cache.Cache
		c = initializeCache(expiryTime, purgeTime)

		// Load the cache with data if needed
		loadset := context.GetInput(ivLoadset).(string)
		if len(loadset) > 0 {
			var cacheData map[string]interface{}
			if err := json.Unmarshal([]byte(loadset), &cacheData); err != nil {
				log.Errorf("couldn't load cache with data: %s", err.Error())
				return false, fmt.Errorf("couldn't load cache with data: %s", err.Error())
			}
			for key, value := range cacheData {
				// Execute the set action
				set(c, key, value.(string))
			}
		}

		// Set the output context
		context.SetOutput(ovResult, "")
		return true, nil
	case "SET":
		// Get the input values
		key := context.GetInput(ivKey).(string)
		value := context.GetInput(ivValue).(string)

		// Initialize if the cache doesn't exist
		var c *cache.Cache
		val, ok := data.GetGlobalScope().GetAttr(CacheName)
		if ok {
			c = val.Value().(*cache.Cache)
		} else {
			log.Infof("cache doesn't exist yet, it will be created first with default settings of expiryTime [%d] minutes and purgeTime [%d] minutes", DefaultExpiryTime, DefaultPurgeTime)
			c = initializeCache(DefaultExpiryTime, DefaultPurgeTime)
		}

		// Execute the set action
		set(c, key, value)

		// Set the output context
		context.SetOutput(ovResult, "")
		return true, nil
	case "GET":
		// Get the input values
		key := context.GetInput(ivKey).(string)

		// Fail if the cache doesn't exist
		var c *cache.Cache
		val, ok := data.GetGlobalScope().GetAttr(CacheName)
		if ok {
			c = val.Value().(*cache.Cache)
		} else {
			log.Error("cache doesn't exist")
			return false, fmt.Errorf("cache doesn't exist")
		}

		// Execute the Get action
		cacheVal, found := get(c, key)

		// Set the output context
		if found {
			context.SetOutput(ovResult, cacheVal)
			return true, nil
		}
		log.Infof("No cache entry was found for [%s]", key)
		context.SetOutput(ovResult, "")
		return true, nil
	case "DELETE":
		// Get the input values
		key := context.GetInput(ivKey).(string)

		// Fail if the cache doesn't exist
		var c *cache.Cache
		val, ok := data.GetGlobalScope().GetAttr(CacheName)
		if ok {
			c = val.Value().(*cache.Cache)
		} else {
			log.Error("cache doesn't exist")
			return false, fmt.Errorf("cache doesn't exist")
		}

		// Execute the delete action
		delete(c, key)

		context.SetOutput(ovResult, "")
		return true, nil
	default:
		log.Errorf("action [%s] does not exist in GlobalCache", action)
		return false, fmt.Errorf("action [%s] does not exist in GlobalCache", action)
	}
}

// initializeCache initializes the cache with expiration time in minutes, and which purges expired items every set amount of minutes
func initializeCache(expiryTime int, purgeTime int) *cache.Cache {
	newCache := cache.New(time.Duration(expiryTime)*time.Minute, time.Duration(purgeTime)*time.Minute)
	data.GetGlobalScope().AddAttr(CacheName, data.TypeAny, newCache)
	log.Infof("Created cache with expiryTime [%d] minutes and purgeTime [%d] minutes", expiryTime, purgeTime)
	return newCache
}

// set adds a new entry to the cache with a default expiration time
func set(c *cache.Cache, key string, value string) {
	c.Set(key, value, cache.DefaultExpiration)
}

// get retrieves an entry from the cache
func get(c *cache.Cache, key string) (interface{}, bool) {
	return c.Get(key)
}

// delete removes an entry from the cache
func delete(c *cache.Cache, key string) {
	c.Delete(key)
}
