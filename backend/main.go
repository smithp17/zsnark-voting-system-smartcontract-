package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// VoteData represents a single vote
type VoteData struct {
	Nullifier string `json:"nullifier"`
	Vote      int    `json:"vote"` // 1 for yes, 0 for no
	Proof     string `json:"proof"`
}

// VotingSession tracks votes for a proposal
type VotingSession struct {
	ProposalID string
	Votes      []VoteData
	Results    map[string]int
	mu         sync.Mutex
}

var sessions = make(map[string]*VotingSession)
var sessionMu sync.Mutex

// CreateSession creates a new voting session
func createSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ProposalID string `json:"proposalId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	sessionMu.Lock()
	defer sessionMu.Unlock()

	session := &VotingSession{
		ProposalID: req.ProposalID,
		Votes:      []VoteData{},
		Results:    map[string]int{"yes": 0, "no": 0},
	}

	sessions[req.ProposalID] = session

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":    "Session created",
		"proposalId": req.ProposalID,
	})
}

// SubmitVote receives a vote and stores it
func submitVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ProposalID string   `json:"proposalId"`
		Vote       VoteData `json:"vote"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	sessionMu.Lock()
	session, exists := sessions[req.ProposalID]
	sessionMu.Unlock()

	if !exists {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	session.mu.Lock()
	defer session.mu.Unlock()

	// Check if nullifier already used (prevent double voting)
	for _, v := range session.Votes {
		if v.Nullifier == req.Vote.Nullifier {
			http.Error(w, "Nullifier already used", http.StatusBadRequest)
			return
		}
	}

	session.Votes = append(session.Votes, req.Vote)

	// Update results
	if req.Vote.Vote == 1 {
		session.Results["yes"]++
	} else {
		session.Results["no"]++
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Vote recorded",
		"status":  "success",
	})
}

// GetResults retrieves voting results
func getResults(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	proposalID := r.URL.Query().Get("proposalId")
	if proposalID == "" {
		http.Error(w, "Missing proposalId", http.StatusBadRequest)
		return
	}

	sessionMu.Lock()
	session, exists := sessions[proposalID]
	sessionMu.Unlock()

	if !exists {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	session.mu.Lock()
	defer session.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"proposalId": proposalID,
		"results":    session.Results,
		"totalVotes": len(session.Votes),
	})
}

// GenerateNullifier creates a unique nullifier for a voter
func generateNullifier(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		VoterID string `json:"voterId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Simple nullifier generation (in production, use proper ZK techniques)
	hash := sha256.Sum256([]byte(req.VoterID))
	nullifier := hex.EncodeToString(hash[:])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"nullifier": nullifier,
	})
}

// Health check endpoint
func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
	})
}

func main() {
	// Register endpoints
	http.HandleFunc("/health", health)
	http.HandleFunc("/api/session/create", createSession)
	http.HandleFunc("/api/vote/submit", submitVote)
	http.HandleFunc("/api/results", getResults)
	http.HandleFunc("/api/nullifier/generate", generateNullifier)

	port := ":8080"
	fmt.Printf("Backend server starting on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
