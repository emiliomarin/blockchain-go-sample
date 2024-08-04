// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

contract Migrations {
    address public owner = msg.sender;
    uint public lastCompletedMigration;

    modifier restricted() {
        require(msg.sender == owner);
        _;
    }

    function setCompleted(uint completed) public restricted {
        lastCompletedMigration = completed;
    }
}