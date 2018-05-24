# flogo-components
Collection custom built flogo components

## Components

### Activities
* [addtodate](activity/addtodate): Adds a specified number of days to either the current date or a chosen date
* [amazons3](activity/amazons3): Activities for interacting with Amazon Simple Storage Service (S3)
* [amazonsqssend](activity/amazonsqssend): Send a message using Amazon Simple Queue Service (SQS)
* [amazonssm](activity/amazonssm): Activities for interaction with the Paramter Store of Amazon Simple Storage Manager (SSM)
* [commandparser](activity/commandparser): Parses a commandline string into separate arguments
* [downloadfile](activity/downloadfile): Downloads a file
* [dynamodbinsert](activity/dynamodbinsert): Insert a record in an Amazon DynamoDB
* [dynamodbquery](activity/dynamodbquery): Execute a query against an Amazon DynamoDB
* [githubissues](activity/githubissues): Get the GitHub issues for an authenticated user
* [gzip](activity/gzip): Activities for reading and writing of gzip format compressed files
* [ifttt webhook](activity/iftttwebhook): Activity to send a webhook request to IFTTT
* [randomnumber](activity/randomnumber): Generate a random unique number between the min and max value
* [randomstring](activity/randomstring): Generate a random string consisting with the length you specify
* [tomlreader](activity/tomlreader): Reads and queries a TOML file
* [trellocard](activity/trellocard): Create a new Trello card in your account
* [writetofile](activity/writetofile): Writes the input of the activity to a file

### Apps
* [invoiceservice](apps/invoiceservice): A simple service listening to requests on a port exposed as environment variable, sending back random data and leveraging the [paymentservice](apps/paymentservice)
* [kubefiles](apps/kubefiles): Files to deploy the [invoiceservice](apps/invoiceservice) and [paymentservice](apps/paymentservice) to Kubernetes
* [paymentservice](apps/paymentservice): A simple service listening to requests on a port exposed as environment variable, sending back a random date
* [serverless-demo](apps/serverless-demo): A collection of three Flogo apps deployed on AWS Lambda, where one app queries a DynamoDB, one app queries a MySQL instance and one app collects that data and presents it using an API Gateway
* [tci-combinator-app](apps/tci-combinator-app): An API spec and a Flogo app that work in TIBCO Cloud Integation that communicates with the apps from the [serverless-demo](apps/serverless-demo)
