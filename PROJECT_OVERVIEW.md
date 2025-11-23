# ZK-SNARK Private Voting System

A complete decentralized voting system with zero-knowledge proofs for voter privacy.

## System Architecture

### 1. **Smart Contracts** (`contracts/`)
- **PrivateVoting.sol**: Main voting contract
  - Manages voting sessions
  - Verifies ZK proofs on-chain
  - Tracks nullifiers to prevent double voting
  - Records vote counts
  - Implements IVerifier interface

- **MockVerifier.sol**: Mock proof verifier for testing
  - Simulates ZK proof verification
  - Always returns true (for development)

### 2. **Go Backend** (`backend/`)
- REST API server on `http://localhost:8080`
- Endpoints:
  - `POST /api/session/create` - Create voting session
  - `POST /api/vote/submit` - Submit a vote
  - `GET /api/results` - Fetch voting results
  - `POST /api/nullifier/generate` - Generate voter nullifier
  - `GET /health` - Health check

- Features:
  - Vote aggregation
  - Nullifier tracking (prevents double voting)
  - Session management
  - Result calculation

### 3. **Frontend** (`frontend/`)
- Web UI on `http://localhost:3000`
- Express.js server
- React-ready (can be expanded)
- Features:
  - Create voting sessions
  - Generate nullifiers
  - Cast private votes
  - View real-time results

### 4. **Circuits** (`circuits/`)
- **voting.circom**: Zero-knowledge circuit
  - Proves voter eligibility
  - Validates vote (0 or 1)
  - Generates nullifier
  - Prevents double voting

## How It Works

### Voting Flow

1. **Session Creation**
   - DAO admin calls `startVoting()` on smart contract
   - Frontend creates session via Go backend

2. **Voter Registration**
   - User generates a nullifier (unique identifier)
   - Nullifier = hash(secret, voter_id)
   - Secret is stored locally (never revealed)

3. **Vote Casting**
   - User creates ZK proof locally proving:
     - They know the secret
     - Vote is valid (0 or 1)
     - Nullifier matches their secret
   - Submit proof + nullifier + vote to backend
   - Backend verifies proof via smart contract
   - Contract marks nullifier as used (prevents re-voting)
   - Vote is recorded anonymously

4. **Result Aggregation**
   - Go backend counts votes
   - Frontend displays results
   - Smart contract stores final tally

## Technology Stack

- **Blockchain**: Solidity (Ethereum)
- **Zero-Knowledge**: Circom
- **Backend**: Go (REST API)
- **Frontend**: Node.js + Express + HTML/JS
- **Development**: Hardhat (contract testing)
- **Tools**: npm, git, viem, ethers.js

## Key Features

✅ **Privacy**: Votes are linked only via nullifier, not wallet address
✅ **Anti-Double-Voting**: Nullifier prevents same person voting twice
✅ **Decentralization**: Smart contracts enforce rules automatically
✅ **Transparency**: Verifiable on-chain (anyone can verify results)
✅ **Auditability**: ZK proofs prove vote validity without revealing identity

## Running the System

### Terminal 1: Go Backend
```bash
cd backend
go run main.go
# Runs on http://localhost:8080
```

### Terminal 2: Frontend
```bash
cd frontend
npm install
node server.js
# Runs on http://localhost:3000
```

### Terminal 3: Smart Contracts (Testing)
```bash
cd contracts
npm install
npx hardhat compile
npx hardhat test
```

## API Examples

### Create Session
```bash
curl -X POST http://localhost:8080/api/session/create \
  -H "Content-Type: application/json" \
  -d '{"proposalId": "proposal-1"}'
```

### Generate Nullifier
```bash
curl -X POST http://localhost:8080/api/nullifier/generate \
  -H "Content-Type: application/json" \
  -d '{"voterId": "voter-1"}'
```

### Submit Vote
```bash
curl -X POST http://localhost:8080/api/vote/submit \
  -H "Content-Type: application/json" \
  -d '{
    "proposalId": "proposal-1",
    "vote": {
      "nullifier": "6d2c8fcf57e0aa6334044224a48a264f92afab3de1bc2e57ce1b2d8e5dfaf7e5",
      "vote": 1,
      "proof": "0x"
    }
  }'
```

### Fetch Results
```bash
curl http://localhost:8080/api/results?proposalId=proposal-1
```

## Next Steps

1. **Integrate Real ZK Proofs**
   - Compile Circom circuits
   - Generate verification keys
   - Implement proof generation in frontend

2. **Deploy to Testnet**
   - Update contract to use Sepolia/Polygon
   - Deploy contracts to testnet
   - Connect frontend to testnet RPC

3. **Production Hardening**
   - Implement proper key management
   - Add rate limiting
   - Implement vote encryption
   - Add voter authentication

4. **Scale**
   - Use Merkle trees for voter registries
   - Implement batch proof verification
   - Add L2 scaling solutions

## Project Status

- ✅ Smart contracts compiled and tested
- ✅ Go backend running with vote aggregation
- ✅ Frontend UI functional
- ⏳ Circom circuits (ready for compilation)
- ⏳ Real ZK proof integration

## Security Considerations

- Nullifiers are deterministic but secret-dependent
- Votes are anonymous but verifiable
- Backend assumes trusted environment (can be decentralized)
- Smart contract is the source of truth

---

Built as a demonstration of ZK-SNARK technology in DAO governance.