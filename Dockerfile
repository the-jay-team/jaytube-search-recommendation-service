FROM golang:1-alpine

WORKDIR /app
COPY . .

RUN go build .

CMD ["./jaytube-search-recommendation-service"]