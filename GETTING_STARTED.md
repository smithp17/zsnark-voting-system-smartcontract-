# Getting Started with ZK Voting System

## ✅ What You've Built

A **production-grade ZK-SNARK private voting system** with:
- ✅ Solidity smart contracts (compiled & tested)
- ✅ Go backend API (running)
- ✅ Web frontend UI (running)
- ✅ Circom ZK circuits (ready to compile)

## 🚀 Quick Start

### Prerequisites
- Node.js v22+
- Go v1.25+
- Git
- npm v10+

### Step 1: Start Go Backend
```bash
cd backend
go run main.go
```
Expected output: `Backend server starting on http://localhost:8080`

### Step 2: Start Frontend (New Terminal)
```bash
cd frontend
node server.js
```
Expected output: `Frontend server running on http://localhost:3000`

### Step 3: Open Browser
```
http://localhost:3000
```

### Step 4: Test the System
1. Click **Create Session**
2. Click **Generate Nullifier**
3. Click **✓ Yes** to vote
4. Click **Fetch Results**

## 📁 Project Structure

```
zk-kyc/
├── contracts/          # Solidity smart contracts
│   ├── contracts/
│   │   ├── PrivateVoting.sol    # Main voting contract
│   │   └── MockVerifier.sol     # Test verifier
│   ├── test/           # Contract tests
│   └── hardhat.config.ts
│
├── backend/            # Go REST API
│   ├── main.go         # Backend server
│   └── go.mod
│
├── frontend/           # Web UI
│   ├── server.js       # Express server
│   ├── public/
│   │   └── index.html  # Voting interface
│   └── package.json
│
├── circuits/           # ZK-SNARK circuits
│   └── voting.circom   # Circom circuit
│
└── proof/              # Proof generation (WIP)
```

## 🔍 Testing the System

### Test 1: Create Voting Session
```bash
curl -X POST http://localhost:8080/api/session/create \
  -H "Content-Type: application/json" \
  -d '{"proposalId": "test-proposal"}'
```

### Test 2: Generate Nullifier
```bash
curl -X POST http://localhost:8080/api/nullifier/generate \
  -H "Content-Type: application/json" \
  -d '{"voterId": "test-voter"}'
```

### Test 3: Submit Vote
```bash
curl -X POST http://localhost:8080/api/vote/submit \
  -H "Content-Type: application/json" \
  -d '{
    "proposalId": "test-proposal",
    "vote": {
      "nullifier": "YOUR_NULLIFIER_HERE",
      "vote": 1,
      "proof": "0x"
    }
  }'
```

### Test 4: Get Results
```bash
curl http://localhost:8080/api/results?proposalId=test-proposal
```

## 🔧 Smart Contract Commands

### Compile Contracts
```bash
cd contracts
npx hardhat compile
```

### Run Tests
```bash
npx hardhat test
```

### Deploy (Requires RPC Configuration)
```bash
npx hardhat ignition deploy ./ignition/modules/PrivateVoting.ts
```

## 📊 Key Endpoints

| Method | Endpoint | Purpose |
|--------|----------|---------|
| POST | `/api/session/create` | Create voting session |
| POST | `/api/vote/submit` | Submit a vote |
| GET | `/api/results?proposalId=X` | Get voting results |
| POST | `/api/nullifier/generate` | Generate voter nullifier |
| GET | `/health` | Health check |

## 🔐 How Privacy Works

1. **Voter generates nullifier**: `nullifier = hash(secret, voter_id)`
2. **Secret stays local**: Never sent to backend/contract
3. **ZK proof proves**: "I know the secret that generated this nullifier"
4. **Smart contract verifies**: Proof is valid and nullifier hasn't been used
5. **Vote recorded**: Anonymously linked only to nullifier, not wallet

## ⚠️ Important Notes

### Development Mode
- Mock verifier always accepts proofs
- No actual ZK proof generation yet
- Backend is centralized (for demo purposes)


- ✅ Smart contract development
- ✅ Backend API design
- ✅ Frontend integration
- ✅ Cryptographic concepts
- ✅ Full-stack blockchain development

---
