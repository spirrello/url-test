FROM golang:1.12.7 as builder

LABEL maintainer="Stefano Pirrello <spirrello@gmail.com>"

ENV GO111MODULE=on

WORKDIR /go/src/url-test

COPY go.mod ./

RUN go mod download

COPY . .

#WORKDIR /go/src/kapture/services/kapture-api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /url-test


####### Start new stage from scratch ########

FROM scratch

WORKDIR /

COPY --from=builder /url-test .

ENTRYPOINT ["/url-test"]
