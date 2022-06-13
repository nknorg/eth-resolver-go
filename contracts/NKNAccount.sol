// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract NKNAccount {
    mapping(address => string) ETHMappings;

    function set(string memory addr) public {
        require(bytes(addr).length != 0, "addr length must be longer than 0");
        ETHMappings[msg.sender] = addr;
    }

    function del() public {
        string memory addr = ETHMappings[msg.sender];
        require(bytes(addr).length != 0);
        delete ETHMappings[msg.sender];
    }

    function getAddr() public view returns (string memory) {
        string memory addr = ETHMappings[msg.sender];
        require(bytes(addr).length != 0);
        return addr;
    }

    function queryAddr(address ethAddr) public view returns (string memory) {
        string memory addr = ETHMappings[ethAddr];
        require(bytes(addr).length != 0);
        return addr;
    }
}
