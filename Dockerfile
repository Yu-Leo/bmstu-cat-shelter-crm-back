FROM golang:1.19

EXPOSE 9000

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download

COPY cmd /app/cmd
COPY config /app/config
COPY docs /app/docs
COPY internal /app/internal
COPY pkg /app/pkg

RUN go build ./cmd/app/main.go

COPY database.db /app/database.db

CMD [ "./main" ]
