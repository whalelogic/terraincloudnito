# ---- Application Build ----
# ---- docker build -t templ-go . ----
FROM golang:1.25.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your source code
COPY . .

# Build the Go app.
# -o /app/main specifies the output file.
# -ldflags="-w -s" strips debug information, making the binary smaller.
# CGO_ENABLED=0 disables Cgo, creating a static binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main -ldflags="-w -s" .


# ---- Final Stage ----
# Use a minimal, non-root base image for the final container.
# "distroless" images contain only your application and its runtime dependencies.
# It's more secure and much smaller than a full OS like Alpine or Ubuntu.
FROM gcr.io/distroless/static-debian11

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the 'builder' stage
COPY --from=builder /app/main .

# Copy static assets and the posts data file.
COPY ./static ./static
COPY ./allposts.json .

# Expose the port your app runs on.
# This should match the port you use in your Go code (e.g., http.ListenAndServe(":8080", mux))
EXPOSE 8080

# The command to run when the container starts
CMD ["/app/main"]
