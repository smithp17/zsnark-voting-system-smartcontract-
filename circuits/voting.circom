pragma circom 2.0;

// Simplified Voting Circuit
// This circuit proves:
// 1. The voter knows a secret
// 2. Their vote is valid (0 or 1)
// 3. Generates a nullifier to prevent double voting

template VotingProof() {
    // Public inputs
    signal input nullifier;
    
    // Private inputs (secret)
    signal input secret;
    signal input vote; // 0 or 1
    
    // Outputs
    signal output nullifierHash;
    signal output voteValid;
    
    // Constraint 1: Vote must be 0 or 1
    // This is enforced by: vote * (1 - vote) = 0
    vote * (1 - vote) === 0;
    
    // Constraint 2: Generate nullifier hash
    // Simple hash: nullifierHash = secret + nullifier
    // (In production, use proper cryptographic hash)
    signal tempHash;
    tempHash <== secret + nullifier;
    nullifierHash <== tempHash;
    
    // Constraint 3: Vote is valid
    voteValid <== 1;
}

component main {public [nullifier]} = VotingProof();