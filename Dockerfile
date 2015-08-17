FROM golang

ADD . /go/src/github.com/twanies/weavebox-docker-example
RUN go install -v github.com/twanies/weavebox-docker-example

ENTRYPOINT ["/go/bin/weavebox-docker-example"]

EXPOSE 3000
