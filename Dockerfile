FROM golang:1-alpine

WORKDIR /app
COPY . .

RUN go build -o jaytube-search-recommendation-service cmd/jaytube_search_recommendation_service/main.go

CMD ["./jaytube-search-recommendation-service"]