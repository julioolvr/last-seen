FROM golang:1.8
WORKDIR /go/src/app
COPY ./server.go .

RUN go-wrapper download
RUN go-wrapper install
EXPOSE 8080

CMD ["go-wrapper", "run", "server.go"]