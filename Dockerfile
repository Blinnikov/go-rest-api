FROM golang:1.16

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
# RUN go build -o main .
RUN go build -o main -v ./cmd/apiserver

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

RUN mkdir configs
RUN cp /build/configs/apiserver.toml ./configs

# Export necessary port
#EXPOSE 3000

# Command to run when starting the container
CMD ["/dist/main"]