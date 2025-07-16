// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;
//
// import "forge-std/Test.sol";
// import "../src/WagerEscrow.sol";
// import "@openzeppelin-contracts/token/ERC20/ERC20.sol";
//
// contract MockToken is ERC20 {
//     constructor() ERC20("MockToken", "MTK") {
//         _mint(msg.sender, 1_000_000 ether);
//     }
// }
//
// contract WagerEscrowCreateTest is Test {
//     WagerEscrow public escrow;
//     MockToken public token;
//
//     address player1 = address(0x1);
//     address player2 = address(0x2);
//
//     function setUp() public {
//         // Deploy token and escrow
//         token = new MockToken();
//         escrow = new WagerEscrow(500); // 5% commission
//
//         // Add token to supported list
//         escrow.addToken(address(token));
//
//         // Transfer tokens to players
//         token.transfer(player1, 1000 ether);
//         token.transfer(player2, 1000 ether);
//
//         // Approve escrow contract
//         vm.prank(player1);
//         token.approve(address(escrow), type(uint256).max);
//
//         vm.prank(player2);
//         token.approve(address(escrow), type(uint256).max);
//     }
//
//     function testCreateWager() public {
//         address ;
//         opponents[0] = player2;
//
//         // Create wager as player1
//         vm.prank(player1);
//         uint256 wagerId = escrow.createWager(opponents, 100 ether, address(token));
//
//         // Check returned wagerId
//         assertEq(wagerId, 0); // first wager should be id 0
//
//         // Check wager data
//         (
//             WagerEscrow.WagerStatus status,
//             address winner,
//             uint256 totalStake,
//             uint256 requiredStake,
//             address[] memory participantAddresses
//         ) = escrow.getWager(wagerId);
//
//         assertEq(uint8(status), uint8(WagerEscrow.WagerStatus.Created));
//         assertEq(winner, address(0));
//         assertEq(totalStake, 100 ether);
//         assertEq(requiredStake, 100 ether);
//         assertEq(participantAddresses.length, 2);
//         assertEq(participantAddresses[0], player1);
//         assertEq(participantAddresses[1], player2);
//     }
// }
