FROM golang:1.12 as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN GO111MODULE=on go get -d -v ./... \
 && CGO_ENABLED=0 GO111MODULE=on go build -o /go/bin/app

FROM golang:1.12-alpine3.9
COPY --from=build /go/bin/app /go/app

RUN apk update \
 && apk add mysql-client curl

ENTRYPOINT ["/go/app"]

EXPOSE 8000/TCP