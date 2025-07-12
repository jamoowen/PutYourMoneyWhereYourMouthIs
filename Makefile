# Go variables
GO_CMD=go
GO_BUILD_FLAGS=-ldflags="-s -w"

# Directories for Go services
SERVICES_DIR=services
API_NAME=pymwymi
MOBILE_DIR=apps/mobile
FRONTEND_DIR=apps/web

# Name of the Go binary (you can add a variable for each service if needed)
GO_BINARY=myapp

# Default target (build everything)
.DEFAULT_GOAL := help

web-dev:
	cd apps/web && npm run dev


run-api:
	@echo "Running challenge api..."
	cd $(SERVICES_DIR)/$(API_NAME) && $(GO_CMD) run ./cmd/api

# ---- DOCKER ----
# Build Docker images for backend services
docker-build:
	@echo "Building Docker images for backend services..."
	docker build -t $(GO_BINARY) $(BACKEND_DIR)

deploy-contract:
	cd contracts/wager_escrow && \
	./deploy/deploy.sh && \
	echo "copying abi to go backend /contracts dir" && \
	cd ../.. && \
	cp -f contracts/wager_escrow/out/WagerEscrow.sol/WagerEscrow.json ./services/pymwymi/contracts/WagerEscrow.json && \
	cp -f contracts/wager_escrow/out/WagerEscrow.sol/WagerEscrow.json ./apps/web/src/contracts/WagerEscrow.json && \
	echo "generating go contract from abi" && \
	cd services/pymwymi/contracts && \
	jq .abi WagerEscrow.json > WagerEscrow.abi && \
	abigen --abi WagerEscrow.abi --pkg contracts --type WagerEscrow --out wager_escrow.go
