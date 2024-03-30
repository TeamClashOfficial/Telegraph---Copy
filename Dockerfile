FROM golang:1.18 AS development

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 7044

RUN CGO_ENABLED=0 GOOS=linux go build -a -o bridge .

# Production stage
FROM alpine:latest AS production

WORKDIR /

COPY .env .

COPY --from=development /app/bridge /bridge

EXPOSE 7044

ENTRYPOINT ["/bridge"]
