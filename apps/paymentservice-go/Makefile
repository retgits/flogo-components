.PHONY: deps

#--- Variables ---
DOCKERHUBUSER=retgits
DOCKERTAG=latest

#--- Get the dependencies ---
deps: 
	go get -u ./...

#--- Clean up the dist folder ---
clean:
	rm -rf dist

#--- Clean up the Kubernetes deployment ---
clean-kube:
	kubectl delete deployment payment-go-svc
	kubectl delete svc payment-go-svc

#--- Build the Flogo app ---
build-app:
	go generate
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o dist/paymentservice-go

#--- Build the docker container ---
build-docker:
	cp Dockerfile dist/Dockerfile
	cp swagger.json dist/swagger.json
	cd dist && docker build -t $(DOCKERHUBUSER)/paymentservice-go:$(DOCKERTAG) .

#--- Run the container ---
run-docker:
	docker run --rm -it -p 9998:8080 $(DOCKERHUBUSER)/paymentservice-go:$(DOCKERTAG)

#--- Run the app on Kubernetes ---
run-kube:
	kubectl apply -f payment-go-svc.yml