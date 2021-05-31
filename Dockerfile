FROM golang:1.16 as builder

# Set necessary environmet variables needed for our image
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o apiserver -v ./cmd/apiserver


FROM alpine
WORKDIR /app

COPY --from=builder /build/apiserver .

# Export necessary port
#EXPOSE 3000

# RUN apk update && apk upgrade \
# && apk add --no-cache ca-certificates

# to trust root certificate if we're going to call other services from this one
# RUN update-ca-certificates
ENTRYPOINT ./apiserver