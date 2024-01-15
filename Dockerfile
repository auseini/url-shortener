FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o shortener ./server 

EXPOSE 8080

RUN chmod +x ./shortener

CMD ["./shortener"]
