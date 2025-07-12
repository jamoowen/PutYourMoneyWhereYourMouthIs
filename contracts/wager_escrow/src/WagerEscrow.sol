// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin-contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin-contracts/token/ERC20/IERC20.sol";
import "@openzeppelin-contracts/utils/ReentrancyGuard.sol";

contract WagerEscrow is ReentrancyGuard {
    using SafeERC20 for IERC20;

    uint256 public constant MAX_PARTICIPANTS = 10;
    uint256 public COMMISSION_BASIS_POINTS;
    address public owner;
    mapping(address => uint256) public commissionBalances;

    enum WagerStatus {
        Created,
        Pending,
        Cancelled,
        Completed,
        Claimed
    }

    mapping(address => IERC20) public supportedTokens;
    address[] public supportedTokensArray;

    struct Participant {
        address walletAddress;
        uint256 stake;
        bool hasClaimed;
    }

    struct Wager {
        address winner;
        address token;
        uint256 totalStake;
        uint256 requiredStake;
        WagerStatus status;
        mapping(address => Participant) participants;
        address[] participantArray;
    }

    mapping(uint256 => Wager) private wagers;
    uint256[] wagerArray;
    uint256 private nextWagerId;

    event WagerCreated(uint256 indexed id, address indexed creator);
    event WagerAccepted(uint256 indexed id, address indexed participant, uint256 amount);
    event WagerStatusUpdated(uint256 indexed id, WagerStatus status);
    event WagerDeleted(uint256 indexed id);
    event WinnerSet(uint256 indexed id, address indexed winner);
    event Claim(address indexed participant, uint256 amount);
    event CompletedWagersDeleted(uint256[] indexed ids);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this");
        _;
    }

    constructor(uint256 _basisPoints) {
        owner = msg.sender;
        COMMISSION_BASIS_POINTS = _basisPoints;
    }

    function addToken(address _token) external onlyOwner {
        supportedTokens[_token] = IERC20(_token);
        supportedTokensArray.push(_token);
    }

    function getSupportedTokens() public view returns (address[] memory tokens, address[] memory erc20Contracts) {
        uint256 length = supportedTokensArray.length;
        tokens = new address[](length);
        erc20Contracts = new address[](length);

        for (uint256 i = 0; i < length; i++) {
            address token = supportedTokensArray[i];
            tokens[i] = token;
            erc20Contracts[i] = address(supportedTokens[token]);
        }

        return (tokens, erc20Contracts);
    }

    function setCommissionBasisPoints(uint256 _basisPoints) external onlyOwner {
        require(_basisPoints > 0 && _basisPoints <= 10_000, "Basis points must be between 1 and 10,000");
        COMMISSION_BASIS_POINTS = _basisPoints;
    }

    function createWager(address[] calldata _participants, uint256 _stake, address _token)
        external
        returns (uint256)
    {
        require(address(supportedTokens[_token]) != address(0), "Invalid token");
        require(_stake > 0, "Stake must be greater than 0");
        require(_participants.length + 1 <= MAX_PARTICIPANTS, "Too many participants");

        supportedTokens[_token].safeTransferFrom(msg.sender, address(this), _stake);

        uint256 id = nextWagerId++;
        Wager storage wager = wagers[id];
        wagerArray.push(id);

        wager.requiredStake = _stake;
        wager.totalStake = _stake;
        wager.status = WagerStatus.Created;
        wager.token = _token;

        wager.participants[msg.sender] = Participant(msg.sender, _stake, false);
        wager.participantArray.push(msg.sender);

        for (uint256 i = 0; i < _participants.length; i++) {
            wager.participants[_participants[i]] = Participant(_participants[i], 0, false);
            wager.participantArray.push(_participants[i]);
        }

        emit WagerCreated(id, msg.sender);
        return id;
    }

    function acceptWager(uint256 id, uint256 _stake, address _token) external {
        Wager storage wager = wagers[id];
        Participant storage p = wager.participants[msg.sender];

        require(wager.status == WagerStatus.Created, "Wager not active");
        require(wager.token == _token, "Token mismatch");
        require(p.walletAddress != address(0), "Not a participant");
        require(p.stake == 0, "Already accepted");
        require(_stake >= wager.requiredStake, "Insufficient stake");

        supportedTokens[_token].safeTransferFrom(msg.sender, address(this), _stake);

        p.stake = _stake;
        wager.totalStake += _stake;

        bool allParticipantsHaveStaked = true;
        for (uint256 i = 0; i < wager.participantArray.length; i++) {
            if (wager.participants[wager.participantArray[i]].stake < wager.requiredStake) {
                allParticipantsHaveStaked = false;
                break;
            }
        }
        if (allParticipantsHaveStaked) {
            wager.status = WagerStatus.Pending;
            emit WagerStatusUpdated(id, WagerStatus.Pending);
        }

        emit WagerAccepted(id, msg.sender, _stake);
    }

    function setWagerStatus(uint256 id, WagerStatus status) external onlyOwner {
        wagers[id].status = status;
        emit WagerStatusUpdated(id, status);
    }

    function setWinner(uint256 id, address winner) external onlyOwner {
        require(wagers[id].participants[winner].walletAddress != address(0), "Winner not a participant");
        wagers[id].winner = winner;
        emit WinnerSet(id, winner);
    }

    function claimRefund(uint256 id) external nonReentrant {
        Wager storage wager = wagers[id];
        Participant storage p = wager.participants[msg.sender];

        require(wager.status == WagerStatus.Cancelled, "Wager not cancelled");
        require(!p.hasClaimed, "Already claimed");
        require(wager.totalStake >= p.stake, "Claims not allowed");

        uint256 refund = p.stake;
        wager.totalStake -= refund;
        p.hasClaimed = true;
        IERC20(wager.token).safeTransfer(msg.sender, refund);
        emit Claim(msg.sender, refund);

        if (wager.totalStake == 0) {
            wager.status = WagerStatus.Claimed;
            emit WagerStatusUpdated(id, WagerStatus.Claimed);
        }
    }

    function claimWinnings(uint256 id) external nonReentrant {
        Wager storage wager = wagers[id];
        Participant storage p = wager.participants[msg.sender];

        require(wager.winner == msg.sender, "Not winner");
        require(wager.status == WagerStatus.Completed, "Wager not completed");
        require(!p.hasClaimed, "Already claimed");

        uint256 commission = (wager.totalStake * COMMISSION_BASIS_POINTS) / 10_000;
        uint256 amount = wager.totalStake - commission;
        p.hasClaimed = true;
        wager.status = WagerStatus.Claimed;
        IERC20(wager.token).safeTransfer(msg.sender, amount);
        commissionBalances[wager.token] += commission;

        emit Claim(msg.sender, amount);
        emit WagerStatusUpdated(id, WagerStatus.Claimed);
    }

    function getWager(uint256 id)
        external
        view
        returns (
            WagerStatus status,
            address winner,
            uint256 totalStake,
            uint256 requiredStake,
            address[] memory participantArray
        )
    {
        Wager storage wager = wagers[id];
        return (
            wager.status,
            wager.winner,
            wager.totalStake,
            wager.requiredStake,
            wager.participantArray
        );
    }

    function getParticipant(uint256 id, address user)
        external
        view
        returns (address walletAddress, uint256 stake, bool hasClaimed)
    {
        Participant storage p = wagers[id].participants[user];
        return (p.walletAddress, p.stake, p.hasClaimed);
    }

    function deleteCompletedWagers() external onlyOwner {
        uint256[] memory completedWagers = new uint256[](wagerArray.length);
        uint256[] memory wagerArrayCopy = wagerArray;
        delete wagerArray;

        uint256 completedCounter = 0;
        for (uint256 i = 0; i < wagerArrayCopy.length; i++) {
            uint256 id = wagerArrayCopy[i];
            Wager storage wager = wagers[id];
            if (wager.status == WagerStatus.Claimed) {
                for (uint256 j = 0; j < wager.participantArray.length; j++) {
                    delete wager.participants[wager.participantArray[j]];
                }
                delete wagers[id];
                completedWagers[completedCounter] = id;
                completedCounter++;
            } else {
                wagerArray.push(id);
            }
        }

        if (wagerArray.length < wagerArrayCopy.length) {
            uint256[] memory completedWagersTrimmed = new uint256[](completedCounter);
            for (uint256 i = 0; i < completedCounter; i++) {
                completedWagersTrimmed[i] = completedWagers[i];
            }
            emit CompletedWagersDeleted(completedWagersTrimmed);
        }
    }

    function withdrawCommission(address _token) external onlyOwner {
        require(_token != address(0), "Invalid token");
        uint256 amount = commissionBalances[_token];
        require(amount > 0, "Nothing to withdraw");
        commissionBalances[_token] = 0;
        IERC20(_token).safeTransfer(owner, amount);
    }

    receive() external payable {}
}
