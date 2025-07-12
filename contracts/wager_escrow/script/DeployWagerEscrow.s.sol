// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {WagerEscrow} from "../src/WagerEscrow.sol";

contract DeployWagerEscrowScript is Script {
    function setUp() public {}

    function run() public {
        vm.startBroadcast();

        uint256 basisPoints = 100;
        WagerEscrow escrow = new WagerEscrow(basisPoints);

        console.log("Deployed WagerEscrow at:", address(escrow));

        vm.stopBroadcast();
    }
}
