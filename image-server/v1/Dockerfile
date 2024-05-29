# Stage 1: Build modules
FROM golang:1.22 as modules-build

COPY ../../modules/v2/common /modules/common
COPY ../../modules/database/d7mysql /modules/database/d7mongo

COPY ../../scripts/install_modules.sh install_modules.sh

RUN chmod +x install_modules.sh && ./install_modules.sh

# Stage 2 : Build main app
FROM golang:1.22 as main-build

COPY --from=modules-build /modules /modules

COPY /server/auth-service/v2 /app

# Copy go workspace script
COPY ../../scripts/init_workspace.sh /app/init_workspace.sh
RUN chmod +x /app/init_workspace.sh && /app/init_workspace.sh

WORKDIR /app
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/main main.go

# Stage 2: Create a minimal Docker image with the built binary
FROM scratch

# Copy the built binary from the builder stage
COPY --from=main-build /bin/main .

# Expose the port the app runs on
EXPOSE 3003

# Command to run the executable
CMD ["./main"]