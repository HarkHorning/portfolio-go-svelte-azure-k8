FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy all source files
COPY . .

# Regenerate go.sum with Go 1.23 compatible dependencies
# This is needed because local go.sum was generated with Go 1.25 pre-release
RUN go mod tidy

# Build the Go application
# CGO_ENABLED=0: Disable C bindings for a fully static binary (no external dependencies)
# GOOS=linux: Target Linux OS (required even on Windows/Mac build hosts)
# -a: Force rebuilding of packages that are already up-to-date
# -installsuffix cgo: Add suffix to package directory to keep separate from default builds
# -ldflags='-w -s': Strip debug information and symbol table to reduce binary size
#   -w: Omit DWARF symbol table (debugging info)
#   -s: Omit symbol table and string table
# -o server: Output the compiled binary as 'server'
# ./cmd/server: Path to the main package (entry point)
RUN CGO_ENABLED=0 GOOS=linux go build \
    -a \
    -installsuffix cgo \
    -ldflags='-w -s' \
    -o server \
    ./cmd/server

# -----------------------------------------------------------------------------
# STAGE 2: Runtime
# -----------------------------------------------------------------------------
# Purpose: Create a minimal production image with only the compiled binary
# Base Image: Alpine Linux (~5MB) - one of the smallest Linux distributions
# Alternative: 'scratch' (empty image) for even smaller size, but no shell/debugging
# -----------------------------------------------------------------------------
FROM alpine:3.19

# Install runtime dependencies
# - ca-certificates: Required for making HTTPS requests to external services
#   (Azure Blob Storage, external APIs, etc.)
# - tzdata: Timezone data for proper time handling in logs and scheduling
RUN apk --no-cache add ca-certificates tzdata

# Create a non-root user for security
# Running as non-root prevents potential container escape vulnerabilities
# - addgroup: Create a new group named 'appgroup'
# - adduser: Create a new user 'appuser' in 'appgroup'
#   -S: Create a system user (no password, no home directory modifications)
#   -G: Specify the group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory for the application
WORKDIR /app

# Copy the compiled binary from the builder stage
# --from=builder: Reference the 'builder' stage defined above
# This is the key to multi-stage builds - only copy what's needed
COPY --from=builder /app/server .

# Change ownership of the application to the non-root user
# This ensures the appuser can execute the binary
RUN chown -R appuser:appgroup /app

# Switch to the non-root user
# All subsequent commands and the container runtime will use this user
USER appuser

# Document which port the application listens on
# EXPOSE doesn't publish the port - it's documentation for users/orchestrators
# The actual port mapping happens at runtime (docker run -p 8080:8080)
EXPOSE 8080

# Set environment variables
# GIN_MODE=release: Run Gin framework in production mode (disables debug logging)
# You can override these at runtime with: docker run -e GIN_MODE=debug
ENV GIN_MODE=release

# Health check configuration
# Docker/Kubernetes use this to determine if the container is healthy
# - --interval=30s: Check every 30 seconds
# - --timeout=3s: Fail if check takes longer than 3 seconds
# - --start-period=5s: Grace period for container startup
# - --retries=3: Mark unhealthy after 3 consecutive failures
# Note: Requires a /health endpoint in your Go API
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Define the command to run when the container starts
# Using array syntax (exec form) is preferred over shell form because:
#   - Signals are passed directly to the process (proper shutdown handling)
#   - No shell process wrapper (slightly more efficient)
CMD ["./server"]
