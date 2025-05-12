#!/bin/bash

# Load environment variables
source .env

# Deploy with Foundry script
forge script script/DeployChallengeEscrow.s.sol:DeployChallengeEscrowScript \
  --rpc-url $ALCHEMY_BASE_SEPOLIA_URL \
  --broadcast \
  --private-key $PRIVATE_KEY