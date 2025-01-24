FROM golang:1.22.6

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/apps/api

RUN go build -o main .

EXPOSE 3000

CMD ["/app/apps/api/main"]
