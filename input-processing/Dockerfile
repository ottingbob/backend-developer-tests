# Builder image
FROM golang:1.14.13-alpine3.12 as builder
RUN mkdir /build
WORKDIR /build
ADD main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o input-processor .

# Application image
FROM alpine:3.12.3
COPY --from=builder /build/input-processor .

# Copy a file here to test when dockerRun recipe is ran in the Makefile
# COPY input_5gb.txt .

ENTRYPOINT ["./input-processor"]
