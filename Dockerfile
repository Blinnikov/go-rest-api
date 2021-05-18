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
RUN mkdir configs
COPY --from=builder /build/configs/apiserver.toml /app/configs/

# Export necessary port
#EXPOSE 3000

ENTRYPOINT ./apiserver