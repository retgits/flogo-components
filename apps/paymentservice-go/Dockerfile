FROM alpine:latest
RUN apk update && apk add ca-certificates
ENV HTTPPORT=8080
ADD paymentservice-go .
ADD swagger.json .
EXPOSE 8080
CMD ./paymentservice-go