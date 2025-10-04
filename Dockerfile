############################
# STEP 1: Build executable
############################
FROM golang:1.25-alpine AS builder

# Install git and CA certificates for fetching dependencies
RUN apk update && apk add --no-cache git ca-certificates upx && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# Set working directory
WORKDIR /app

# Copy go mod files first for caching dependencies
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the executable with optimizations
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o /app/estimator && \
    upx --best --lzma /app/estimator


############################
# STEP 2: Build final image
############################
FROM scratch

# Import CA certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import user and group files
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy the executable
COPY --from=builder /app/estimator /estimator

# Use appuser
USER appuser:appuser

# Expose port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/estimator"]
