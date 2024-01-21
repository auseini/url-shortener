FROM golang:1.21-alpine AS base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o shortener ./server 

RUN chmod +x ./shortener

EXPOSE 8080

FROM scratch

COPY --from=base /app/shortener .


CMD ["./shortener"]
