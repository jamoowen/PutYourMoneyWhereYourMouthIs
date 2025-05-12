// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {ChallengeEscrow} from "../src/ChallengeEscrow.sol";

contract DeployChallengeEscrowScript is Script {
    function setUp() public {}

    function run() public {
        vm.startBroadcast();

        uint256 basisPoints = 100;
        ChallengeEscrow escrow = new ChallengeEscrow(basisPoints);

        console.log("Deployed ChallengeEscrow at:", address(escrow));

        vm.stopBroadcast();
    }
}
