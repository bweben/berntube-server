FROM golang:1.13

WORKDIR /go/src/app

COPY . .

RUN go get -d -v
RUN go install -d -v

EXPOSE 5000
EXPOSE 5001

CMD ["berntube-server"]