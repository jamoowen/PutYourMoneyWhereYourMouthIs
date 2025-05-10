// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// import "openzeppelin-contracts/token/ERC20/utils/SafeERC20.sol";

// import "openzeppelin-contracts/token/ERC20/utils/SafeERC20.sol";
// contracts/challenge_escrow/lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol
// import "openzeppelin-contracts/token/ERC20/IERC20.sol";
import "@openzeppelin-contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin-contracts/token/ERC20/IERC20.sol";

contract ChallengeEscrow {
    using SafeERC20 for IERC20;

    uint256 public constant MAX_PARTICIPANTS = 10;
    address public owner;
    uint256 public commissionBalance;

    enum ChallengeStatus {
        Created,
        Pending,
        Cancelled,
        Completed,
        Claimed
    }

    mapping(address => IERC20) public supportedTokens;

    struct Participant {
        address walletAddress;
        uint256 stake;
        bool hasClaimed;
    }

    struct Challenge {
        address winner;
        address token;
        uint256 totalStake;
        uint256 requiredStake;
        ChallengeStatus status;
        mapping(address => Participant) participants;
        address[]  participantArray;
    }

    mapping(uint256 => Challenge) private challenges;
    uint256[]  challengeArray;
    uint256 private nextChallengeId;

    error ChallengeNotActive();
    error InvalidToken();
    error NotInvited();
    error AlreadyAccepted();
    error InsufficientDeposit();

    error NotAccepted();
    error NotCancelled();
    error NotCompleted();

    error AlreadyClaimed();
    error ClaimsNotAllowed();
    error NotOwner();

    event ChallengeCreated(uint256 indexed id, address indexed creator);
    event ChallengeAccepted(uint256 indexed id, address indexed participant, uint256 amount);
    event ChallengeStatusUpdated(uint256 indexed id, ChallengeStatus status);
    event ChallengeDeleted(uint256 indexed id);
    event WinnerSet(uint256 indexed id, address indexed winner);
    event Claim(address indexed participant, uint256 amount);
    event CompletedChallengesDeleted(uint256[] indexed ids);

    modifier onlyOwner() {
        if (msg.sender != owner) revert NotOwner();
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    function addToken(address _token) external onlyOwner {
        supportedTokens[_token] = IERC20(_token);
    }

    function createChallenge(address[] calldata _participants, uint256 _stake, address _token)
        external
        returns (uint256)
    {
        require(supportedTokens[_token] != IERC20(address(0)), "Token not supported");
        require(_stake > 0, "Stake must be greater than 0");
        require(_participants.length < 10, "Max 10 participants");

        // Transfer stake from the caller to the contract
        supportedTokens[_token].safeTransferFrom(msg.sender, address(this), _stake);

        uint256 id = nextChallengeId++;
        Challenge storage challenge = challenges[id];
        challengeArray.push(id);

        challenge.requiredStake = _stake;
        challenge.totalStake = _stake;
        challenge.status = ChallengeStatus.Created;
        challenge.token = _token;

        // first add the creator
        challenge.participants[msg.sender] = Participant(msg.sender, 0, false);
        challenge.participantArray.push(msg.sender);

        for (uint256 i = 0; i < _participants.length; i++) {
            challenge.participants[_participants[i]] = Participant(_participants[i], 0, false);
            challenge.participantArray.push(_participants[i]);
        }

        emit ChallengeCreated(id, msg.sender);
        return id;
    }

    // perhaps we want another function to allow people to top up their stakes
    function acceptChallenge(uint256 id, uint256 _stake, address _token) external {
        Challenge storage challenge = challenges[id];
        Participant storage p = challenge.participants[msg.sender];

        if (challenge.status != ChallengeStatus.Created) revert ChallengeNotActive();
        if (challenge.token != _token) revert InvalidToken();
        if (p.walletAddress == address(0)) revert NotInvited();
        if (p.stake != 0) revert AlreadyAccepted();
        if (_stake < challenge.requiredStake) revert InsufficientDeposit();

        supportedTokens[_token].safeTransferFrom(msg.sender, address(this), _stake);

        p.stake = _stake; // not +=, they can only accept once
        challenge.totalStake += _stake;

        // if everyone has staked the required amount then we're good and the challenge moves to pending
        bool allParticipantsHaveStaked = true;
        for (uint256 i = 0; i < challenge.participantArray.length; i++) {
            if (challenge.participants[challenge.participantArray[i]].stake < challenge.requiredStake) {
                allParticipantsHaveStaked=false;
                break;
            } 
        }
        if (allParticipantsHaveStaked) {
            challenge.status = ChallengeStatus.Pending;
            emit ChallengeStatusUpdated(id, ChallengeStatus.Pending);
        }

        emit ChallengeAccepted(id, msg.sender, _stake);
    }

    function setChallengeStatus(uint256 id, ChallengeStatus status) external onlyOwner {
        challenges[id].status = status;
        emit ChallengeStatusUpdated(id, status);
    }

    function setWinner(uint256 id, address winner) external onlyOwner {
        challenges[id].winner = winner;
        emit WinnerSet(id, winner);
    }

    function claimRefund(uint256 id) external {
        Challenge storage challenge = challenges[id];
        Participant storage p = challenge.participants[msg.sender];

        if (challenge.status != ChallengeStatus.Cancelled) revert NotCancelled();
        if (p.hasClaimed) revert AlreadyClaimed();
        if (challenge.totalStake<p.stake) revert ClaimsNotAllowed();

        p.hasClaimed = true;
        uint256 refund = p.stake;
        challenge.totalStake -= refund;
        payable(msg.sender).transfer(refund);
        emit Claim(msg.sender, refund);

        if (challenge.totalStake == 0){
            challenge.status = ChallengeStatus.Claimed;
            emit ChallengeStatusUpdated(id, ChallengeStatus.Claimed);
        }

    }
    
    function claimWinnings(uint256 id) external {
        Challenge storage challenge = challenges[id];
        Participant storage p = challenge.participants[msg.sender];

        if (challenge.status != ChallengeStatus.Completed) revert NotCompleted();
        if (p.hasClaimed) revert AlreadyClaimed();

        challenge.status = ChallengeStatus.Claimed;

        uint256 amount = challenge.totalStake;
        p.hasClaimed = true;
        payable(msg.sender).transfer(amount);
        challenge.status = ChallengeStatus.Claimed;
        emit Claim(msg.sender, amount);
        emit ChallengeStatusUpdated(id, ChallengeStatus.Claimed);
    }

    function getChallenge(uint256 id) external view returns (
        ChallengeStatus status,
        address winner,
        uint256 totalStake,
        uint256 requiredStake,
        address[] memory participantArray
    ) {
        Challenge storage challenge = challenges[id];
        return ( challenge.status, challenge.winner, challenge.totalStake, challenge.requiredStake, challenge.participantArray);
    }

    function getParticipant(uint256 id, address user) external view returns (
        address walletAddress,
        uint256 stake,
        bool hasClaimed 
    ) {
        Participant storage p = challenges[id].participants[user];
        return (p.walletAddress, p.stake, p.hasClaimed);
    }

    function deleteCompletedChallenges() external onlyOwner {
        uint256[] memory completedChallenges = new uint256[](challengeArray.length);
        uint256[] memory challengeArrayCopy = challengeArray;
        delete challengeArray;

        uint256 completedCounter = 0;
        for (uint256 i = 0; i < challengeArrayCopy.length; i++) {
            if (challenges[challengeArrayCopy[i]].status == ChallengeStatus.Claimed) {
                delete challenges[challengeArrayCopy[i]];
                completedChallenges[completedCounter] = challengeArrayCopy[i];
                completedCounter++;
            } else {
                challengeArray.push(challengeArrayCopy[i]);
            }       
        }

        if (challengeArray.length < challengeArrayCopy.length) {
            uint256[] memory completedChallengesTrimmed = new uint256[](completedCounter);
            for (uint256 i = 0; i < completedCounter; i++) {
                completedChallengesTrimmed[i] = completedChallenges[i];
            }
            emit CompletedChallengesDeleted(completedChallengesTrimmed);
        }
        
    }

    function withdrawCommission() external onlyOwner {
        uint256 amount = commissionBalance;
        commissionBalance = 0;
        payable(owner).transfer(amount);
    }

    receive() external payable {}
}
