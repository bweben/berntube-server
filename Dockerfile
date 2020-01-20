FROM golang:1.13

WORKDIR /go/src/app

COPY . .

RUN go build -o main .

EXPOSE 5000
EXPOSE 5001

CMD ["./main"]