package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"

	"github.com/TIBCOSoftware/flogo-lib/app"
	"github.com/TIBCOSoftware/flogo-lib/engine"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("main-engine")
var cpuprofile = flag.String("cpuprofile", "", "Writes CPU profiling for the current process to the specified file")
var memprofile = flag.String("memprofile", "", "Writes memory profiling for the current process to the specified file")
var (
	cp app.ConfigProvider
)

func main() {

	if cp == nil {
		// Use default config provider
		cp = app.DefaultConfigProvider()
	}

	app, err := cp.GetApp()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Println(fmt.Sprintf("Failed to create CPU profiling file due to error - %s", err.Error()))
			os.Exit(1)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	e, err := engine.New(app)
	if err != nil {
		log.Errorf("Failed to create engine instance due to error: %s", err.Error())
		os.Exit(1)
	}

	err = e.Start()
	if err != nil {
		log.Errorf("Failed to start engine due to error: %s", err.Error())
		os.Exit(1)
	}

	exitChan := setupSignalHandling()

	code := <-exitChan

	e.Stop()

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			fmt.Println(fmt.Sprintf("Failed to create memory profiling file due to error - %s", err.Error()))
			os.Exit(1)
		}

		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			fmt.Println(fmt.Sprintf("Failed to write memory profiling data to file due to error - %s", err.Error()))
			os.Exit(1)
		}
		f.Close()
	}

	os.Exit(code)
}

func setupSignalHandling() chan int {

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exitChan := make(chan int, 1)
	select {
	case s := <-signalChan:
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			exitChan <- 0
		default:
			logger.Debug("Unknown signal.")
			exitChan <- 1
		}
	}
	return exitChan
}
