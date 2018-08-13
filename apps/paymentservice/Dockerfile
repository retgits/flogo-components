FROM alpine:latest
RUN apk update && apk add ca-certificates
ENV HTTPPORT=8080
ADD paymentservice .
EXPOSE 8080
CMD ./paymentservice