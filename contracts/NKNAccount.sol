// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract NKNAccount {
    struct NKNAddress {
        string identifier;
        bytes32 publicKey;
    }
    mapping(address => NKNAddress) ETHToNKNAddr;

    function set(string memory identifier, bytes32 publicKey) public {
        require(bytes(identifier).length <= 64, "identifier length must be no longer than 64");
        ETHToNKNAddr[msg.sender] = NKNAddress({identifier: identifier, publicKey: publicKey});
    }

    function del() public {
        NKNAddress memory result = ETHToNKNAddr[msg.sender];
        require(result.publicKey != bytes32(0));
        delete ETHToNKNAddr[msg.sender];
    }

    function getNKNAddr() public view returns (NKNAddress memory) {
        NKNAddress memory result = ETHToNKNAddr[msg.sender];
        require(result.publicKey != bytes32(0));
        return result;
    }

    function getNKNAddr(address publicKey) public view returns (NKNAddress memory) {
        NKNAddress memory result = ETHToNKNAddr[publicKey];
        require(result.publicKey != bytes32(0));
        return result;
    }
}
