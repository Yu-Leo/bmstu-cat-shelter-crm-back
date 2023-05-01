FROM golang:1.19

EXPOSE 9000

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download

COPY . .

RUN go build ./cmd/app/main.go

CMD [ "./main" ]
