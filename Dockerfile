FROM golang:1.16.3-alpine3.13

ENV GO111MODULE=on

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

WORKDIR /build/cmd/website-visitor-stats

RUN go build -o main .

WORKDIR /dist

RUN cp /build/cmd/website-visitor-stats/main .

EXPOSE 8080

CMD ["/dist/main"]