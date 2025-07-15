# Put your money where your mouth is

PYMWYMI is a an app which allows users to place wagers amongst eachother, using a solidty smart contract as escrow. Winners can only withdraw their 'winnings' if a unanimous decision is reached.

This idea came about after a number of 1v1 basketball games of mine ended with me never receiving the cash I should have won :(

*currently only USDC on Base supported

## dependencies:
gotestsum

##  Api
Go backend using Goethereum client and MongoDB

## /apps/web
Nextjs frontend using wagmi

## /contracts/wager_escrow
The solidity escrow smart contract

### testing the backend
gotestsum --format short-verbose ./services/auth/
gotestsum --format short-verbose ./...

# Backend
- Golang backend with MongoDB
- NOTE! we cannot currently scale horozontally. there are race conditions with any operations that update challenge status.
eg: accepting a challenge, voting in a challenge etc. if challenge status changes, certain operations should not be possible.
To prioritize consistency over availablity here, I am using channels to ensure only one goroutine can update the challenge at a given time.
**future me will most likely use kafka or another queue based solution but this is ok for now
