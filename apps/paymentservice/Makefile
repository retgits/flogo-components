.PHONY: deps

#--- Variables ---
DOCKERHUBUSER=retgits
DOCKERTAG=0.2

#--- Clean up the dist folder ---
clean:
	rm -rf paymentservice

#--- Clean up the Kubernetes deployment ---
clean-kube:
	kubectl delete deployment payment-svc
	kubectl delete svc payment-svc

#--- Build the Flogo app ---
build-app:
	flogo create -f paymentservice.json paymentservice
	cd paymentservice && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 flogo build -e
	cd paymentservice && CGO_ENABLED=0 flogo build -e

#--- Build the docker container ---
build-docker:
	cp Dockerfile paymentservice/bin/Dockerfile
	cd paymentservice/bin && docker build -t $(DOCKERHUBUSER)/paymentservice:$(DOCKERTAG) .

#--- Run the container ---
run-docker:
	docker run --rm -it -p 9998:8080 $(DOCKERHUBUSER)/paymentservice:$(DOCKERTAG)

#--- Run the app on Kubernetes ---
run-kube:
	kubectl apply -f payment-go-svc.yml