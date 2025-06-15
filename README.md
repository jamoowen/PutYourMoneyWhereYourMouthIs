# Put your money where your mouth is


## dependencies:
gotestsum

## Challenge Api


## Challenge Service


## Mobile react-native expo app


### testing the backend
gotestsum --format short-verbose ./services/auth/
gotestsum --format short-verbose ./...

# Backend
- Golang backend with MongoDB
- NOTE! we cannot currently scale horozontally. there are race conditions with any operations that update challenge status.
eg: accepting a challenge, voting in a challenge etc. if challenge status changes, certain operations should not be possible.
To prioritize consistency here, I am using channels to ensure only one goroutine can update the challenge at a given time.
