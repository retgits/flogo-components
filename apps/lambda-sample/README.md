# Lambda Sample

This repository contains a template Flogo app, built with the Go API of Project Flogo, which can be deployed to AWS Lambda

## Building this app
1. Update the code in `function.go` if needed

2. Run the following commands to build the executable
```bash
# Get all dependencies
$ go get -u ./...
# Generate the Flogo metadata
$ go generate ./...
# Build an executable
$ env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/*.go
```

3. Zip the executable

4. Upload it to your Lambda function

_This would obviously be even easier to deploy using [AWS SAM](https://github.com/awslabs/aws-sam-cli) or [Serverless Framework](https://serverless.com/framework/)_

## More information
If you're looking for a more in-depth overview of how the app is built, check out the [lab](https://tibcosoftware.github.io/flogo/labs/serverless/)