# Go variables
GO_CMD=go
GO_BUILD_FLAGS=-ldflags="-s -w"

# Directories for Go services
SERVICES_DIR=services
CHALLENGE_API_NAME=pymwymi
CHALLENGE_SERVICE_NAME=challenge-service
MOBILE_DIR=apps/mobile
FRONTEND_DIR=apps/web

# Name of the Go binary (you can add a variable for each service if needed)
GO_BINARY=myapp

# Default target (build everything)
.DEFAULT_GOAL := help

web-dev:
	cd apps/web && npm run dev


# Run all Go services (backend)
run-challenge-api:
	@echo "Running challenge api..."
	$(GO_CMD) run  ./$(SERVICES_DIR)/$(CHALLENGE_API_NAME)/cmd/api/main.go

# ---- DOCKER ----
# Build Docker images for backend services
docker-build:
	@echo "Building Docker images for backend services..."
	docker build -t $(GO_BINARY) $(BACKEND_DIR)

