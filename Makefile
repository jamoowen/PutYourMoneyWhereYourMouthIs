# Go variables
GO_CMD=go
GO_BUILD_FLAGS=-ldflags="-s -w"

# Directories for Go services
SERVICES_DIR=services
MATCH_API_NAME=match-api
MATCH_SERVICE_NAME=match-service
MOBILE_DIR=apps/mobile
FRONTEND_DIR=apps/web

# Name of the Go binary (you can add a variable for each service if needed)
GO_BINARY=myapp

# Default target (build everything)
.DEFAULT_GOAL := help

# ---- BACKEND ----
# Build all Go services (backend)
build-backend:
	@echo "Building Go backend services..."
	$(GO_CMD) build $(GO_BUILD_FLAGS) -o $(BACKEND_DIR)/$(GO_BINARY) $(BACKEND_DIR)

# Run all Go services (backend)
start-backend:
	@echo "Running Go backend services..."
	$(BACKEND_DIR)/$(GO_BINARY)

# Run all Go services (backend)
run-match-api:
	@echo "Running match api..."
	$(GO_CMD) run  $(SERVICES_DIR)/$(MATCH_API_NAME)/cmd/main.go

# Run all Go services (backend)
run-match-service:
	@echo "Running match service..."
	$(GO_CMD) run  $(SERVICES_DIR)/$(MATCH_SERVICE_NAME)/cmd/main.go

# Run tests for Go services
test-backend:
	@echo "Running Go tests..."
	$(GO_CMD) test -v ./...

# Lint Go services
lint-backend:
	@echo "Running linters for Go services..."
	golangci-lint run $(BACKEND_DIR)

# Format Go code
fmt-backend:
	@echo "Formatting Go code..."
	gofmt -s -w $(BACKEND_DIR)

# ---- MOBILE (Expo) ----
# Build the Expo app
build-mobile:
	@echo "Building Expo app..."
	cd $(MOBILE_DIR) && expo build

# Run the Expo app locally
run-mobile:
	@echo "Running Expo app..."
	cd $(MOBILE_DIR) && expo start

# Run tests for Expo app
test-mobile:
	@echo "Running tests for Expo app..."
	cd $(MOBILE_DIR) && jest

# Lint the Expo app
lint-mobile:
	@echo "Running linters for Expo app..."
	cd $(MOBILE_DIR) && eslint .

# Format Expo app code
fmt-mobile:
	@echo "Formatting Expo app code..."
	cd $(MOBILE_DIR) && prettier --write .

# ---- DOCKER ----
# Build Docker images for backend services
docker-build:
	@echo "Building Docker images for backend services..."
	docker build -t $(GO_BINARY) $(BACKEND_DIR)

# Run Docker containers for backend services
docker-run:
	@echo "Running Docker containers for backend services..."
	docker run -d -p 8080:8080 $(GO_BINARY)

# ---- CLEAN ----
# Clean up compiled binaries (backend)
clean:
	@echo "Cleaning up compiled binaries..."
	rm -f $(BACKEND_DIR)/$(GO_BINARY)

# ---- DEPLOY ----
# Build and deploy backend services and Expo app
deploy: build-backend build-mobile
	@echo "Deploying backend services and Expo app..."
