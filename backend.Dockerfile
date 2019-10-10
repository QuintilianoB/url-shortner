FROM golang:1.12 as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN GO111MODULE=on go get -d -v ./... \
 && CGO_ENABLED=0 GO111MODULE=on go build -o /go/bin/app

FROM golang:1.12-alpine3.9
COPY --from=build /go/bin/app /go/app

EXPOSE 8000/TCP

RUN mkdir -p /go/store/sql/
COPY .env /go/.env
COPY store/sql/initial.sql /go/store/sql/initial.sql

ENTRYPOINT ["/go/app"]