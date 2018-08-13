.PHONY: deps

#--- Variables ---
DOCKERHUBUSER=retgits
DOCKERTAG=latest
IPADDR=xxx.xxx.xxx.xxx

#--- Clean up the dist folder ---
clean:
	rm -rf invoiceservice

#--- Clean up the Kubernetes deployment ---
clean-kube:
	kubectl delete deployment invoice-svc
	kubectl delete svc invoice-svc

#--- Build the Flogo app ---
build-app:
	flogo create -f invoiceservice.json invoiceservice
	cd invoiceservice && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 flogo build -e
	cd invoiceservice && CGO_ENABLED=0 flogo build -e

#--- Build the docker container ---
build-docker:
	cp Dockerfile invoiceservice/bin/Dockerfile
	cd invoiceservice/bin && docker build -t $(DOCKERHUBUSER)/invoiceservice:$(DOCKERTAG) .

#--- Run the container ---
run-docker:
	docker run --rm -it -e PAYMENTSERVICE=http://$(IPADDR):9998/api/expected-date/:id -p 9999:8080 $(DOCKERHUBUSER)/invoiceservice:$(DOCKERTAG)

#--- Run the app on Kubernetes ---
run-kube:
	kubectl apply -f invoice-svc.yml