// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MockVerifier {
    function verify(bytes calldata proof, uint256[] calldata inputs) external pure returns (bool) {
        return true; // Mock verifier always returns true for testing
    }
}