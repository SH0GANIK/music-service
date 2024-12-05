FROM golang:1.23 as builder

# Env variables
ENV GOOS linux
ENV CGO_ENABLED 0

# Work directory
WORKDIR     /app
# Installing dependencies
COPY . .

RUN go mod download
RUN go build -o music-service ./cmd/music-service/main.go

FROM alpine:3.16 as production

# Copying built assets from builder
COPY --from=builder /app/music-service .
COPY .env .
COPY config config
# Exposing server port
EXPOSE 8080
# Starting our application
CMD [ "./music-service" ]