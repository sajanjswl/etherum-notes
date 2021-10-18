//SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.9;

contract MySmartContract {
    function Hello() public pure returns (string memory) {
        return "Hello World";
    }
    function Greet(string memory str) public pure returns (string memory) {
        return str;
    }
}