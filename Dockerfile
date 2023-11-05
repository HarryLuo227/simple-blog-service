# Build stage
FROM golang:1.21 AS builder
WORKDIR /go/src/simple-blog-service
COPY . .
RUN make

# Final stage
FROM ubuntu:22.04
WORKDIR /root
COPY --from=builder /go/src/simple-blog-service .
EXPOSE 8000
ENTRYPOINT [ "./simple-blog-service" ]