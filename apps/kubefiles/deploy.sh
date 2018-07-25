#!/bin/bash

# Description: Demo script to help deploy Flogo apps to Kubernetes
# Author: retgits <https://github.com/retgits>
# Last Updated: 2018-07-19

#--- Variables ---
# These are the variables you want to update. 
# The root directory to download all artifacts to. It defaults to the current working directory, but you can change it if you want
ROOTDIR=`pwd`
# The docker hub echo is used to connect to Docker Hub and push images so a Kubernetes instance will be able to download the images
DOCKERHUBUSER=retgits

#--- Start of the script ---
# You do not want to modify below this line, unless you really want to obviously :)

#--- Variables ---
VERSION=0.0.2
WORKDIR=flogo-on-kubernetes
GOPTH="$(go env GOPATH)"

#--- Function to print the header ---
header () {
    clear
    echo "********************************************************************************"
    echo "* Flogo on Kubernetes Demo Script                                              *"
    echo "*                                                               Version: $VERSION *"
    echo "********************************************************************************"
    echo ""
}

#--- Function to validate the prerequisites ---
validate () {
    header
    echo "Checking cURL"
    curl --version

    echo ""
    echo "Checking Go"
    go version

    echo ""
    echo "Checking dep"
    dep version
    
    echo ""
    echo "Checking Flogo"
    flogo version
    
    echo ""
    echo "Checking docker"
    docker version --format "{{.Client.Version}}"
    
    echo ""
    echo "Checking kubectl"
    kubectl version
}

#--- Function to deploy the flow based apps to Kubernetes ---
deploy_flows () {
    header
    echo "This script will help you set up a Flogo demo on Kubernetes and will pause after"
    echo "each step, explaining what has happened. After the script has paused you'll see"
    echo "[ENTER], meaning you can hit the enter key to continue."
    echo "If you're ready to get started hit [ENTER] now!"
    read -p ""

    echo "Let's talk about the scenario first..."
    echo "Lets create a set of microservices which will manage invoices for a company."
    echo "The idea was nicely explained on https://hackernoon.com/getting-started-with-microservices-and-kubernetes-76354312b556"
    echo "using Node.js code, so we'll do it with Flogo and we'll create:"
    echo " - A front end invoices service to return information about invoices"
    echo " - A back end expected date service that’ll tell us when an invoice is likely to be paid"
    echo "[ENTER]"
    read -p ""

    rm -rf $ROOTDIR/$WORKDIR
    mkdir -p $ROOTDIR/$WORKDIR
    echo "We've just created a directory called \"$WORKDIR\" which we'll use to"
    echo "download all artifacts in"
    echo "[ENTER]"
    read -p ""

    cd $ROOTDIR/$WORKDIR
    curl -o invoiceservice.json https://raw.githubusercontent.com/retgits/flogo-components/master/apps/invoiceservice/invoiceservice.json
    curl -o paymentservice.json https://raw.githubusercontent.com/retgits/flogo-components/master/apps/paymentservice/paymentservice.json
    echo ""
    echo "The demo will consist of two apps deployed to a Kubernetes cluster. We'll use"
    echo "docker images that already exist, but we do want to provide you with all the"
    echo "code and artifacts to show the apps in the Flogo Web UI. To do so, we just"
    echo "downloaded the two JSON files that represent the Flogo apps"
    echo " - invoiceservice.json"
    echo " - paymentservice.json"
    echo "[ENTER]"
    read -p ""

    echo "Now we'll create the Flogo app for the invoice service"
    echo "to do so, we're running the command"
    echo "flogo create -f invoiceservice.json invoiceservice"
    echo "[ENTER]"
    read -p ""
    flogo create -f invoiceservice.json invoiceservice

    echo ""
    echo "Great! The first app is done, now on to the second one."
    echo "For that one we'll run"
    echo "flogo create -f paymentservice.json paymentservice"
    echo "[ENTER]"
    read -p ""
    flogo create -f paymentservice.json paymentservice

    echo ""
    echo "Now, let's build the executables for the first app."
    echo "We'll build two executables. One that will be able to run natively"
    echo "on your machine and one that will be a Linux executable to run in a"
    echo "docker container. To do so we'll execute:"
    echo "flogo build -e"
    echo "env GOOS=linux flogo build -e"
    echo "[ENTER]"
    read -p ""
    cd $ROOTDIR/$WORKDIR/invoiceservice
    env GOOS=linux flogo build -e
    flogo build -e

    echo ""
    echo "One down and one to go! We'll do the same thing for the second service."
    echo "Again we'll execute:"
    echo "flogo build -e"
    echo "env GOOS=linux flogo build -e"
    echo "[ENTER]"
    read -p ""
    cd $ROOTDIR/$WORKDIR/paymentservice
    env GOOS=linux flogo build -e
    flogo build -e

    echo ""
    echo "Now on to Docker! We'll create docker images based on Alpine Linux."
    echo "We'll put a Dockerfile in the bin/linux_amd64 folder and run the command"
    echo "docker build -t $DOCKERHUBUSER/invoiceservice ."
    echo "[ENTER]"
    read -p ""
    cd $ROOTDIR/$WORKDIR/invoiceservice/bin/linux_amd64
    dockerfile="FROM alpine:latest
    RUN apk update && apk add ca-certificates
    ENV HTTPPORT=8080 \ 
        PAYMENTSERVICE=bla
    ADD invoiceservice .
    EXPOSE 8080
    CMD ./invoiceservice"
    echo "$dockerfile" > Dockerfile
    docker build -t $DOCKERHUBUSER/invoiceservice .

    echo ""
    echo "And we'll do the same for the second service as well."
    echo "We'll put a Dockerfile in the bin/linux_amd64 folder and run the command"
    echo "docker build -t $DOCKERHUBUSER/paymentservice ."
    echo "[ENTER]"
    read -p ""
    cd $ROOTDIR/$WORKDIR/paymentservice/bin/linux_amd64
    dockerfile="FROM alpine:latest
    RUN apk update && apk add ca-certificates
    ENV HTTPPORT=8080
    ADD paymentservice .
    EXPOSE 8080
    CMD ./paymentservice"
    echo "$dockerfile" > Dockerfile
    docker build -t $DOCKERHUBUSER/paymentservice .

    echo ""
    echo "Right now you have two new docker images available in your registry, which you"
    echo "could use to test the whole scenario as well. If you want to do that run the"
    echo "below commands in separate terminal windows"
    echo "docker run --rm -it -p 9999:8080 $DOCKERHUBUSER/paymentservice"
    echo "docker run --rm -it -p 9998:8080 -e PAYMENTSERVICE=http://<YOUR IP>:9999/api/expected-date/:id $DOCKERHUBUSER/invoiceservice"
    echo "for more information and sample messages check out"
    echo "invoiceservice: https://hub.docker.com/r/retgits/invoiceservice/"
    echo "paymentservice: https://hub.docker.com/r/retgits/paymentservice/"
    echo "[ENTER]"
    read -p ""

    echo ""
    echo "Sidestep... you can push your docker containers to Docker Hub (assuming you"
    echo "have an account for it) by running"
    echo "docker push $DOCKERHUBUSER/paymentservice:latest"
    echo "docker push $DOCKERHUBUSER/invoiceservice:latest"
    echo "[ENTER]"
    read -p ""

    echo ""
    echo "The last step is to deploy to Kubernetes. To do that we need to download two"
    echo "additional files."
    echo "[ENTER]"
    read -p ""
    cd $ROOTDIR/$WORKDIR
    curl -o invoice-svc.yml https://raw.githubusercontent.com/retgits/flogo-components/master/apps/kubefiles/invoice-svc.yml
    curl -o payment-svc.yml https://raw.githubusercontent.com/retgits/flogo-components/master/apps/kubefiles/payment-svc.yml

    echo ""
    echo "The payment-svc.yml file will create a deployment and a service resource"
    echo "in your Kubernetes cluster. For that we'll use an existing docker image called"
    echo "retgits/paymentservice, which is the same as the one you just built. If you want"
    echo "you can update the yaml file before pressing a key..."
    echo "[ENTER]"
    read -p ""
    kubectl apply -f payment-svc.yml

    echo ""
    echo "You now have a docker container that is accessible as a service on Kubernetes"
    echo "The payment service is accessible on port 80 of the cluster IP address"
    echo "that was assigned to it. If you want to try it out look for the CLUSTER-IP"
    echo "of the payment-svc in the output below"
    kubectl get services
    echo "and run kubectl run curl --image=radial/busyboxplus:curl -i --tty"
    echo "which will start a new buxybox terminal in your cluster. From there you can"
    echo "run curl <CLUSTERIP>/api/expected-date/3456 which should return something like"
    echo "{\"expectedDate\":\"2018-02-26\",\"id\":\"3456\"}"
    echo "[ENTER]"
    read -p ""

    echo ""
    echo "The second service we'll make available using the type: LoadBalancer which means"
    echo "that you can access it from outside your Kubernetes cluster. For this we'll use"
    echo "the existing container retgits/invoiceservice. If you open the invoice.yml file"
    echo "you'll see that there is an environment variable called PAYMENTSERVICE (line 24)"
    echo "which points to the DNS entry for the payment service. This way we can makes"
    echo "updates and potentially move the payment service around without having to update"
    echo "this service."
    echo "[ENTER]"
    read -p ""
    kubectl apply -f invoice-svc.yml

    echo ""
    echo "All done! You now have two Flogo apps running on a Kubernetes cluster which you"
    echo "invoke by sending a curl message to the Kubernetes IP address or localhost if"
    echo "you're running Docker for Mac. You can execute a command like"
    echo "curl `minikube service invoice-svc --url`/api/invoices/1234 which will return something like"
    echo "{\"amount\":1162,\"balance\":718,\"currency\":\"USD\",\"expectedPaymentDate\":\"2018-03-02\",\"id\":\"1234\",\"ref\":\"INV-1234\"}"
    echo ""
    success "Happy Kube-ing!!"
}

#--- Function to clean up the flow based apps ---
cleanup_flows () {
    header
    echo "This script will help you tear down the apps deployed using \"deploy-flows\""
    echo "If you're not ready to continue, press ctrl+c otherwise press [ENTER] to remove" 
    echo "the apps"
    read -p ""
    kubectl delete deployments,services -l run=payment-svc
    kubectl delete deployments,services -l run=invoice-svc
}

#--- Function to deploy the Go based apps to Kubernetes ---
deploy_golang () {
    header
    echo "This script will help you set up a Flogo demo on Kubernetes and will pause after"
    echo "each step, explaining what has happened. After the script has paused you'll see"
    echo "[ENTER], meaning you can hit the enter key to continue."
    echo "If you're ready to get started hit [ENTER] now!"
    read -p ""

    echo "Let's talk about the scenario first..."
    echo "Lets create a set of microservices which will manage invoices for a company."
    echo "The idea was nicely explained on https://hackernoon.com/getting-started-with-microservices-and-kubernetes-76354312b556"
    echo "using Node.js code, so we'll do it with Flogo and we'll create:"
    echo " - A front end invoices service to return information about invoices"
    echo " - A back end expected date service that’ll tell us when an invoice is likely to be paid"
    echo "[ENTER]"
    read -p ""

    go get -d https://github.com/retgits/flogo-components
    echo "We've just downloaded the code to \"$GOPTH/src/github.com/retgits/flogo-components/apps\". In there are the two apps"
    echo "that we'll use to build and deploy to Kubernetes"
    echo "[ENTER]"
    read -p ""

    echo ""
    echo "The demo will consist of two Flogo apps built with the Go API deployed to a Kubernetes cluster."
    echo "We'll use docker images that already exist, but we do want to provide you with all the"
    echo "code and artifacts. If you want to open up the code and check, feel free to do so!"
    echo "The go files are called 'main.go' and are in the \"$GOPTH/src/github.com/retgits/flogo-components/apps\" directory"
    echo "The first app is in the paymentservice-go directory"
    echo "The second app in the the invoiceservice-go directory"
    echo "[ENTER]"
    read -p ""

    echo "Now we'll build the app for the invoice service"
    echo "to do so, we're running a few commands"
    echo "go get -u ./..."
    echo "go generate"
    echo "GOOS=linux go build"
    echo "[ENTER]"
    read -p ""
    cd $GOPTH/src/github.com/retgits/flogo-components/apps/invoiceservice-go
    go get -u ./...
    go generate
    GOOS=linux go build

    echo ""
    echo "Great! The first app is done, now on to the second one."
    echo "For that one we'll run"
    echo "go get -u ./..."
    echo "go generate"
    echo "GOOS=linux go build"
    echo "[ENTER]"
    read -p ""
    cd $GOPTH/src/github.com/retgits/flogo-components/apps/paymentservice-go
    go get -u ./...
    go generate
    GOOS=linux go build

    echo ""
    echo "Now on to Docker! We'll create docker images based on Alpine Linux."
    echo "We'll put a Dockerfile in the bin/linux_amd64 folder and run the command"
    echo "docker build -t $DOCKERHUBUSER/invoiceservice ."
    echo "[ENTER]"
    read -p ""
    cd $GOPTH/src/github.com/retgits/flogo-components/apps/invoiceservice-go
    dockerfile="FROM alpine:latest
    RUN apk update && apk add ca-certificates
    ENV HTTPPORT=8080 \ 
        PAYMENTSERVICE=bla
    ADD invoiceservice-go .
    EXPOSE 8080
    CMD ./invoiceservice-go"
    echo "$dockerfile" > Dockerfile
    docker build -t $DOCKERHUBUSER/invoiceservice-go .

    echo ""
    echo "And we'll do the same for the second service as well."
    echo "We'll put a Dockerfile in the bin/linux_amd64 folder and run the command"
    echo "docker build -t $DOCKERHUBUSER/paymentservice ."
    echo "[ENTER]"
    read -p ""
    cd $GOPTH/src/github.com/retgits/flogo-components/apps/paymentservice-go
    dockerfile="FROM alpine:latest
    RUN apk update && apk add ca-certificates
    ENV HTTPPORT=8080
    ADD paymentservice-go .
    EXPOSE 8080
    CMD ./paymentservice-go"
    echo "$dockerfile" > Dockerfile
    docker build -t $DOCKERHUBUSER/paymentservice-go .

    echo ""
    echo "Right now you have two new docker images available in your registry, which you"
    echo "could use to test the whole scenario as well. If you want to do that run the"
    echo "below commands in separate terminal windows"
    echo "docker run --rm -it -p 9999:8080 $DOCKERHUBUSER/paymentservice-go"
    echo "docker run --rm -it -p 9998:8080 -e PAYMENTSERVICE=http://<YOUR IP>:9999/api/expected-date/:id $DOCKERHUBUSER/invoiceservice-go"
    echo "for more information and sample messages check out"
    echo "invoiceservice: https://hub.docker.com/r/retgits/invoiceservice-go/"
    echo "paymentservice: https://hub.docker.com/r/retgits/paymentservice-go/"
    echo "[ENTER]"
    read -p ""

    echo ""
    echo "Sidestep... you can push your docker containers to Docker Hub (assuming you"
    echo "have an account for it) by running"
    echo "docker push $DOCKERHUBUSER/paymentservice-go:latest"
    echo "docker push $DOCKERHUBUSER/invoiceservice-go:latest"
    echo "[ENTER]"
    read -p ""

    echo ""
    echo "The last step is to deploy to Kubernetes. To do that we need to download two"
    echo "additional files."
    echo "[ENTER]"
    read -p ""
    curl -o $GOPTH/src/github.com/retgits/flogo-components/apps/invoiceservice-go/invoice-go-svc.yml https://raw.githubusercontent.com/retgits/flogo-components/master/apps/kubefiles/invoice-go-svc.yml
    curl -o $GOPTH/src/github.com/retgits/flogo-components/apps/paymentservice-go/payment-go-svc.yml https://raw.githubusercontent.com/retgits/flogo-components/master/apps/kubefiles/payment-go-svc.yml

    echo ""
    echo "The payment-go-svc.yml file will create a deployment and a service resource"
    echo "in your Kubernetes cluster. For that we'll use an existing docker image called"
    echo "retgits/paymentservice-go, which is the same as the one you just built. If you want"
    echo "you can update the yaml file before pressing a key..."
    echo "[ENTER]"
    read -p ""
    cd $GOPTH/src/github.com/retgits/flogo-components/apps/paymentservice-go
    kubectl apply -f payment-go-svc.yml

    echo ""
    echo "You now have a docker container that is accessible as a service on Kubernetes"
    echo "The payment service is accessible on port 80 of the cluster IP address"
    echo "that was assigned to it. If you want to try it out look for the CLUSTER-IP"
    echo "of the payment-go-svc in the output below"
    kubectl get services
    echo "and run kubectl run curl --image=radial/busyboxplus:curl -i --tty"
    echo "which will start a new buxybox terminal in your cluster. From there you can"
    echo "run curl <CLUSTERIP>/api/expected-date/3456 which should return something like"
    echo "{\"expectedDate\":\"2018-02-26\",\"id\":\"3456\"}"
    echo "[ENTER]"
    read -p ""

    echo ""
    echo "The second service we'll make available using the type: LoadBalancer which means"
    echo "that you can access it from outside your Kubernetes cluster. For this we'll use"
    echo "the existing container retgits/invoiceservice-go. If you open the invoice.yml file"
    echo "you'll see that there is an environment variable called PAYMENTSERVICE (line 24)"
    echo "which points to the DNS entry for the payment service. This way we can makes"
    echo "updates and potentially move the payment service around without having to update"
    echo "this service."
    echo "[ENTER]"
    read -p ""
    cd $GOPTH/src/github.com/retgits/flogo-components/apps/invoiceservice-go
    kubectl apply -f invoice-go-svc.yml

    echo ""
    echo "All done! You now have two Flogo apps running on a Kubernetes cluster which you"
    echo "invoke by sending a curl message to the Kubernetes IP address or localhost if"
    echo "you're running Docker for Mac. You can execute a command like"
    echo "curl `minikube service invoice-go-svc --url`/api/invoices/1234 which will return something like"
    echo "{\"amount\":1162,\"balance\":718,\"currency\":\"USD\",\"expectedPaymentDate\":\"2018-03-02\",\"id\":\"1234\",\"ref\":\"INV-1234\"}"
    echo ""
    success "Happy Kube-ing!!"
}

#--- Function to clean up the Go based apps ---
cleanup_golang () {
    header
    echo "This script will help you tear down the apps deployed using \"deploy-golang\""
    echo "If you're not ready to continue, press ctrl+c otherwise press [ENTER] to remove" 
    echo "the apps"
    read -p ""
    kubectl delete deployments,services -l run=payment-go-svc
    kubectl delete deployments,services -l run=invoice-go-svc
}

case "$1" in 
    "deploy-flows")
        deploy_flows
        ;;
    "cleanup-flows")
        cleanup_flows
        ;;
    "deploy-golang")
        deploy_golang
        ;;
    "cleanup-golang")
        cleanup_golang
        ;;
    "validate")
        validate
        ;;
    *)
        echo "The target {$1} you want to execute doesn't exist"
        echo ""
        echo "Usage:"
        echo "./`basename "$0"` [target]"
        echo ""
        echo "Possible targets are"
        echo "  deploy-flows  : Deploy the apps that were created using the Flogo Web UI to Kubernetes"
        echo "  cleanup-flows : Remove the apps and services created using the deploy-flows option"
        echo "  deploy-golang : Deploy the apps that were created using Flogo as a lib to Kubernetes"
        echo "  cleanup-golang: Remove the apps and services created using the deploy-golang option"
        echo "  validate      : Checks if the prerequisites for this script are met"
esac