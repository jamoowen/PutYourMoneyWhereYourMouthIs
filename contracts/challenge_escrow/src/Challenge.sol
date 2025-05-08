
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// import "openzeppelin-contracts/token/ERC20/utils/SafeERC20.sol";

// import "openzeppelin-contracts/token/ERC20/utils/SafeERC20.sol";
// contracts/challenge_escrow/lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol
// import "openzeppelin-contracts/token/ERC20/IERC20.sol";
import  "@openzeppelin-contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin-contracts/token/ERC20/IERC20.sol";


contract ChallengeEscrow {
    using SafeERC20 for IERC20
    
    address public owner;
    uint256 public commissionBalance;

    enum ChallengeStatus { Created, Cancelled, Completed, Claimed }

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
    }

    mapping(uint256 => Challenge) private challenges;
    uint256 private nextChallengeId;

    error NotAccepted();
    error AlreadyWithdrawn();
    error WithdrawalsNotAllowed();
    error NotOwner();

    event ChallengeCreated(uint256 indexed id, address indexed creator);
    event ChallengeAccepted(uint256 indexed id, address indexed participant, uint256 amount);
    event ChallengeStatusUpdated(uint256 indexed id, ChallengeStatus status);
    event WinnerSet(uint256 indexed id, address indexed winner);
    event Withdrawal(address indexed participant, uint256 amount);

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

    function createChallenge(address[] calldata _participants, uint256 _stake, address _token) external returns (uint256) {
        require(supportedTokens[_token] != IERC20(address(0)), "Token not supported");
        require(_stake > 0, "Stake must be greater than 0");

        // Transfer stake from the caller to the contract
        supportedTokens[_token].safeTransferFrom(msg.sender, address(this), _stake);

        uint256 id = nextChallengeId++;
        Challenge storage ch = challenges[id];
        ch.requiredStake = _stake;
        ch.totalStake = _stake;
        ch.status = ChallengeStatus.Created;
        ch.token = _token;

        // first add the creator
        ch.participants[msg.sender] = Participant(msg.sender, 0,  false);

        for (uint i = 0; i < _participants.length; i++) {
            ch.participants[_participants[i]] = Participant(_participants[i], 0,  false);
        }

        emit ChallengeCreated(id, msg.sender);
        return id;
    }

    function acceptChallenge(uint256 id) external payable {
        Challenge storage ch = challenges[id];
        Participant storage p = ch.participants[msg.sender];

        require(ch.status == ChallengeStatus.Created, "Challenge is not active");
        require(p.deposit == 0, "Already accepted");
        require(msg.value >= ch.requiredDeposit, "Insufficient deposit");

        p.accepted = true;
        p.deposit = msg.value;
        ch.pot += msg.value;

        emit ChallengeAccepted(id, msg.sender, msg.value);
    }

    // function setChallengeStatus(uint256 id, ChallengeStatus status) external onlyOwner {
    //     challenges[id].status = status;
    //     emit ChallengeStatusUpdated(id, status);
    // }

    // function setWinner(uint256 id, address winner) external onlyOwner {
    //     challenges[id].winner = winner;
    //     emit WinnerSet(id, winner);
    // }

    // function withdraw(uint256 id) external {
    //     Challenge storage ch = challenges[id];
    //     Participant storage p = ch.participants[msg.sender];

    //     if (!p.accepted) revert NotAccepted();
    //     if (p.withdrawn) revert AlreadyWithdrawn();

    //     p.withdrawn = true;

    //     if (ch.status == ChallengeStatus.Cancelled) {
    //         uint256 refund = p.deposit;
    //         p.deposit = 0;
    //         payable(msg.sender).transfer(refund);
    //         emit Withdrawal(msg.sender, refund);

    //     } else if (ch.status == ChallengeStatus.Completed && msg.sender == ch.winner) {
    //         uint256 winnerAmount = (ch.pot * 99) / 100;
    //         uint256 commission = ch.pot - winnerAmount;
    //         commissionBalance += commission;

    //         ch.pot = 0; // prevent double withdraw
    //         payable(msg.sender).transfer(winnerAmount);
    //         emit Withdrawal(msg.sender, winnerAmount);
    //     } else {
    //         revert WithdrawalsNotAllowed();
    //     }
    // }

    // function getParticipants(uint256 id) external view returns (address[] memory) {
    //     return challenges[id].participantList;
    // }

    // function getChallenge(uint256 id) external view returns (
    //     address creator,
    //     address winner,
    //     uint256 pot,
    //     uint256 requiredDeposit,
    //     ChallengeStatus status
    // ) {
    //     Challenge storage ch = challenges[id];
    //     return (ch.creator, ch.winner, ch.pot, ch.requiredDeposit, ch.status);
    // }

    // function getParticipant(uint256 id, address user) external view returns (
    //     bool accepted,
    //     uint256 deposit,
    //     bool withdrawn
    // ) {
    //     Participant storage p = challenges[id].participants[user];
    //     return (p.accepted, p.deposit, p.withdrawn);
    // }

    function withdrawCommission() external onlyOwner {
        uint256 amount = commissionBalance;
        commissionBalance = 0;
        payable(owner).transfer(amount);
    }

    receive() external payable {}
}

