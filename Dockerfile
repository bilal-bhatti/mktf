# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.22 AS builder

WORKDIR /app

COPY . ./

RUN rm -rf vendor
RUN go mod tidy
RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -o /mktf

FROM gcr.io/distroless/static-debian12 AS release

WORKDIR /

COPY --from=builder /mktf /mktf

ENTRYPOINT ["/mktf"]
