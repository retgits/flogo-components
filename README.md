# flogo-components
Collection custom built flogo components

## Components

### Activities
* [addtodate](activity/addtodate): Add a specified number of days to either the current date or a chosen date
* [amazons3](activity/amazons3): Upload or Download files from Amazon Simple Storage Service (S3)
* [amazonsqssend](activity/amazonsqssend): Send a message using Amazon Simple Queue Service (SQS)
* [awsssm](activity/amazonssm): Store and Retrieve parameters from the Parameter Store in Amazon Simple Systems Manager (SSM)
* [commandparser](activity/commandparser): Parses a commandline string into separate arguments
* [downloadfile](activity/downloadfile): Download a file
* [dynamodbinsert](activity/dynamodbinsert): Insert an object into Amazon DynamoDB
* [dynamodbquery](activity/dynamodbquery): Query objects from Amazon DynamoDB
* [githubissues](activity/githubissues): Get the GitHub issues assigned to an authenticated user
* [gzip](activity/gzip): Read and write gzip format compressed files
* [ifttt webhook](activity/iftttwebhook): Send webhook requests to IFTTT
* [null](activity/null): An activity that does nothing (useful for branching out right after the trigger)
* [queryparser](acitivity/queryparser): Parse a query string into name/value pairs
* [randomnumber](activity/randomnumber): Generate a random unique number between the min and max value
* [randomstring](activity/randomstring): Generate a random string consisting with the length you specify
* [tomlreader](activity/tomlreader): Reads and queries a TOML file
* [trellocard](activity/trellocard): Create a new Trello card
* [writetofile](activity/writetofile): Write to a file

### Apps
* [invoiceservice](apps/invoiceservice): A simple service listening to requests on a port exposed as environment variable, sending back random data and leveraging the [paymentservice](apps/paymentservice)
* [kubefiles](apps/kubefiles): Files to deploy the [invoiceservice](apps/invoiceservice) and [paymentservice](apps/paymentservice) to Kubernetes
* [paymentservice](apps/paymentservice): A simple service listening to requests on a port exposed as environment variable, sending back a random date
* [serverless-demo](apps/serverless-demo): A collection of three Flogo apps deployed on AWS Lambda, where one app queries a DynamoDB, one app queries a MySQL instance and one app collects that data and presents it using an API Gateway
* [tci-combinator-app](apps/tci-combinator-app): An API spec and a Flogo app that work in TIBCO Cloud Integation that communicates with the apps from the [serverless-demo](apps/serverless-demo)
