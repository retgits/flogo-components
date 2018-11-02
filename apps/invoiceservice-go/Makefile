.PHONY: deps

#--- Variables ---
DOCKERHUBUSER=retgits
DOCKERTAG=latest
IPADDR=xxx.xxx.xxx.xxx

#--- Get the dependencies ---
deps: 
	go get -u ./...

#--- Clean up the dist folder ---
clean:
	rm -rf dist

#--- Clean up the Kubernetes deployment ---
clean-kube:
	kubectl delete deployment invoice-go-svc
	kubectl delete svc invoice-go-svc

#--- Build the Flogo app ---
build-app:
	go generate
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o dist/invoiceservice-go

#--- Build the docker container ---
build-docker:
	cp Dockerfile dist/Dockerfile
	cp swagger.json dist/swagger.json
	cd dist && docker build -t $(DOCKERHUBUSER)/invoiceservice-go:$(DOCKERTAG) .

#--- Run the container ---
run-docker:
	docker run --rm -it -e PAYMENTSERVICE=http://$(IPADDR):9998/api/expected-date/ -p 9999:8080 $(DOCKERHUBUSER)/invoiceservice-go:$(DOCKERTAG)

#--- Run the app on Kubernetes ---
run-kube:
	kubectl apply -f invoice-go-svc.yml