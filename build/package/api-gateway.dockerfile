# Build container
# -----------------
FROM golang:1.19-alpine AS builder
RUN apk update && apk add --no-cache gcc musl-dev git

RUN mkdir /app
WORKDIR /app
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-w -s' -a -o /bin/api-gateway /app/cmd/api-gateway

# Deployment container
# ----------------------
FROM alpine
RUN apk update

COPY --from=builder /bin/api-gateway /myapp/

EXPOSE 4001
CMD ["/myapp/api-gateway"]