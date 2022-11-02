# Build environment
# -----------------
FROM golang:1.19-alpine AS builder
RUN apk update && apk add --no-cache gcc musl-dev git

RUN mkdir /app
WORKDIR /app
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o /bin/api-service /app/cmd/api-service

# Deployment environment
# ----------------------
FROM alpine
RUN apk update

COPY --from=builder /bin/api-service /myapp/

EXPOSE 4000
CMD ["/myapp/api-service"]