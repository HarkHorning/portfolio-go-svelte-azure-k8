# =============================================================================
# FRONTEND DOCKERFILE - Svelte + TypeScript Application
# =============================================================================
# This is a multi-stage Docker build for the SvelteKit frontend.
# Multi-stage builds allow us to:
#   1. Use Node.js to install dependencies and build the application
#   2. Copy only the static build output to a lightweight Nginx server
#   3. Reduce final image size from ~1GB to ~25MB
# =============================================================================

# -----------------------------------------------------------------------------
# STAGE 1: Dependencies
# -----------------------------------------------------------------------------
# Purpose: Install node modules separately for better caching
# Base Image: Node.js Alpine for smaller size
# -----------------------------------------------------------------------------
FROM node:22-alpine AS deps

# Set working directory
WORKDIR /app

# Copy only package files first
# This layer is cached unless package.json or package-lock.json change
COPY package.json package-lock.json* ./

# Install dependencies
# --frozen-lockfile: Fail if lockfile is out of sync (ensures reproducible builds)
# If no package-lock.json exists, fall back to regular npm install
RUN if [ -f package-lock.json ]; then \
        npm ci --frozen-lockfile; \
    else \
        npm install; \
    fi

# -----------------------------------------------------------------------------
# STAGE 2: Builder
# -----------------------------------------------------------------------------
# Purpose: Build the SvelteKit application into static files
# Base Image: Same Node.js version for consistency
# -----------------------------------------------------------------------------
FROM node:22-alpine AS builder

# Set working directory
WORKDIR /app

# Copy dependencies from the deps stage
# This avoids reinstalling node_modules if only source code changed
COPY --from=deps /app/node_modules ./node_modules

# Copy all source files
COPY . .

# Set Node environment to production
# This optimizes the build output and excludes dev-only code
ENV NODE_ENV=production

# Build the SvelteKit application
# This runs 'vite build' which outputs to the 'build' directory
# SvelteKit with adapter-static generates static HTML/CSS/JS files
RUN npm run build

# -----------------------------------------------------------------------------
# STAGE 3: Runtime
# -----------------------------------------------------------------------------
# Purpose: Serve the static files with Nginx
# Base Image: Nginx Alpine (~25MB) - production-grade web server
# Why Nginx over Node:
#   - More efficient for serving static files
#   - Lower memory footprint
#   - Better caching and compression out of the box
#   - Industry standard for production deployments
# -----------------------------------------------------------------------------
FROM nginx:alpine AS runtime

# Remove default Nginx static content
# The default 'Welcome to nginx' page is not needed
RUN rm -rf /usr/share/nginx/html/*

# Copy the built static files from the builder stage
# SvelteKit with adapter-static outputs to 'build' directory
COPY --from=builder /app/build /usr/share/nginx/html

# Copy custom Nginx configuration
# This configures:
#   - SPA routing (fallback to index.html for client-side routing)
#   - Gzip compression
#   - Cache headers for static assets
#   - Security headers
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Document the exposed port
# Nginx listens on port 80 by default, but we'll use 3000 to match docker-compose
EXPOSE 3000

# Health check for container orchestration
# Verifies the web server is responding
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:3000/ || exit 1

# Start Nginx in the foreground
# -g 'daemon off;': Keeps Nginx in foreground so Docker can track the process
# This is required for Docker to properly manage the container lifecycle
CMD ["nginx", "-g", "daemon off;"]
