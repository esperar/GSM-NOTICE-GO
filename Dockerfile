FROM golang:1.20-buster AS gobuilder
WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go.mod download

EXPOSE 3070

COPY ./ ./

RUN go build

CMD ["/goboard"]