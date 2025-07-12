#!/bin/bash

# Load environment variables
source .env

echo "building and deploying contract..."
forge clean
forge build
# Deploy with Foundry script
forge script script/DeployWagerEscrow.s.sol:DeployWagerEscrowScript \
  --rpc-url $ALCHEMY_BASE_SEPOLIA_URL \
  --broadcast \
  --private-key $PRIVATE_KEY \
  > latest_deployment.txt
