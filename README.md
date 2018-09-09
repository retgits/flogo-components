# flogo-components
My collection of custom built flogo components

## Components

### Activities
* [addtodate](activity/addtodate): Add a specified number of days to either the current date or a chosen date
* [amazons3](activity/amazons3): Upload or Download files from Amazon Simple Storage Service (S3)
* [amazonses](activity/amazonses): Sends emails using Amazon Simple Email Service (SES)
* [amazonsqssend](activity/amazonsqssend): Send a message using Amazon Simple Queue Service (SQS)
* [awsssm](activity/amazonssm): Store and Retrieve parameters from the Parameter Store in Amazon Simple Systems Manager (SSM)
* [commandparser](activity/commandparser): Parses a commandline string into separate arguments
* [downloadfile](activity/downloadfile): Download a file
* [dynamodbinsert](activity/dynamodbinsert): Insert an object into Amazon DynamoDB
* [dynamodbquery](activity/dynamodbquery): Query objects from Amazon DynamoDB
* [envkey](activity/envkey): Get environment variables or the provided fallback value
* [githubissues](activity/githubissues): Get the GitHub issues assigned to an authenticated user
* [gzip](activity/gzip): Read and write gzip format compressed files
* [ifttt webhook](activity/iftttwebhook): Send webhook requests to IFTTT
* [mashtoken](activity/mashtoken): Get a login token for TIBCO Cloud Mashery
* [null](activity/null): An activity that does nothing (useful for branching out right after the trigger)
* [pubnubpublisher](activity/pubnubpublisher): An activity that publishes messages to PubNub
* [queryparser](acitivity/queryparser): Parse a query string into name/value pairs
* [randomnumber](activity/randomnumber): Generate a random unique number between the min and max value
* [randomstring](activity/randomstring): Generate a random string consisting with the length you specify
* [readfile](activity/readfile): Reads a file
* [tomlreader](activity/tomlreader): Reads and queries a TOML file
* [trellocard](activity/trellocard): Create a new Trello card
* [writetofile](activity/writetofile): Write to a file

### Triggers
* [grpctrigger](trigger/grpctrigger): A trigger to receive gRPC messages
* [pubnubsubscriber](trigger/pubnubsubscriber): A trigger to receive messages from PubNub

### Apps
* [invoiceservice](apps/invoiceservice): A simple service listening to requests on a port exposed as environment variable, sending back random data and leveraging the [paymentservice](apps/paymentservice)
* [github-lambda](https://github.com/retgits/github-lambda): A Flogo powered Lambda function to get new GitHub issues (based on the Flogo Go API)
* [kubefiles](apps/kubefiles): Files to deploy the [invoiceservice](apps/invoiceservice) and [paymentservice](apps/paymentservice) to Kubernetes
* [paymentservice](apps/paymentservice): A simple service listening to requests on a port exposed as environment variable, sending back a random date
* [serverless-demo](apps/serverless-demo): A collection of three Flogo apps deployed on AWS Lambda, where one app queries a DynamoDB, one app queries a MySQL instance and one app collects that data and presents it using an API Gateway
* [tci-combinator-app](apps/tci-combinator-app): An API spec and a Flogo app that work in TIBCO Cloud Integation that communicates with the apps from the [serverless-demo](apps/serverless-demo)
* [trello-lambda](https://github.com/retgits/trello-lambda): A Flogo powered Lambda function to create new Trello cards (based on the Flogo Go API)