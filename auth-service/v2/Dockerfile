# Stage 1: Build modules
FROM golang:1.22 as modules-build

COPY ../../modules/v2/common /modules/common
COPY ../../modules/v2/database/d7mysql /modules/database/d7mysql
COPY ../../modules/v2/database/d7redis /modules/database/d7redis
COPY ../../scripts /scripts

# Stage 2 : Build main app
FROM golang:1.22 as main-build

COPY --from=modules-build /modules /modules

COPY --from=modules-build /scripts /app
COPY . /app

# Copy go workspace script
RUN chmod +x /app/install_modules.sh && /app/install_modules.sh auth
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
EXPOSE 3001

## Set ENV
#ENV DB_USERNAME=pointer \
#    DB_PASSWORD=1234 \
#    DB_HOST=127.0.0.1 \
#    DB_PORT=3306 \
#    DB_DATABASE=todopoint
#

# Command to run the executable
CMD ["./main"]