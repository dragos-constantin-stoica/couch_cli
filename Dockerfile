FROM golang:1.21 as builder

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
#RUN go mod init couch_cli && go get github.com/rivo/tview 
RUN go mod download && go mod verify
COPY . .
RUN go build -o /usr/local/bin/couch_cli

#CMD ["app"]

FROM alpine:latest as production

COPY --from=builder /usr/local/bin/couch_cli .
CMD ./couch_cli
