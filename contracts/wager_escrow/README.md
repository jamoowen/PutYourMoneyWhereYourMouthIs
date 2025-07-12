## To deploy

from root of monorepo:
1. make deploy-contract
2. add the newly deployed address (in the latest_deployment.txt) to the constants files in backend and frontend


------- old
(you might first need forge clean)
1. forge build
2. ./deploy/deploy.sh
3. copy the abi from out/ChallengeEscrow.sol/ChallengeEscrow.json -> go backend && frontend locations
------- old

backend location: contracts/


## Most recent deploy:
