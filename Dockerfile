FROM golang:1.15.4-alpine3.12
RUN mkdir /app

ADD . /app
WORKDIR /app/cmd/http

RUN go mod tidy
# RUN go mod download

RUN go build -o main .

CMD ["/app/cmd/server/restfull"]

EXPOSE 9000
