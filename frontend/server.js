const express = require('express');
const axios = require('axios');
const path = require('path');

const app = express();
const PORT = 3000;
const BACKEND_URL = 'http://localhost:8080';

app.use(express.json());
app.use(express.static('public'));

// Serve HTML page
app.get('/', (req, res) => {
  res.sendFile(path.join(__dirname, 'public', 'index.html'));
});

// Create voting session
app.post('/api/session/create', async (req, res) => {
  try {
    const { proposalId } = req.body;
    const response = await axios.post(`${BACKEND_URL}/api/session/create`, {
      proposalId
    });
    res.json(response.data);
  } catch (error) {
    console.error('Error creating session:', error.message);
    res.status(500).json({ error: 'Failed to create session' });
  }
});

// Generate nullifier
app.post('/api/nullifier/generate', async (req, res) => {
  try {
    const { voterId } = req.body;
    const response = await axios.post(`${BACKEND_URL}/api/nullifier/generate`, {
      voterId
    });
    res.json(response.data);
  } catch (error) {
    console.error('Error generating nullifier:', error.message);
    res.status(500).json({ error: 'Failed to generate nullifier' });
  }
});

// Submit vote
app.post('/api/vote/submit', async (req, res) => {
  try {
    const { proposalId, vote } = req.body;
    const response = await axios.post(`${BACKEND_URL}/api/vote/submit`, {
      proposalId,
      vote
    });
    res.json(response.data);
  } catch (error) {
    console.error('Error submitting vote:', error.message);
    res.status(500).json({ error: 'Failed to submit vote' });
  }
});

// Get results
app.get('/api/results/:proposalId', async (req, res) => {
  try {
    const { proposalId } = req.params;
    const response = await axios.get(`${BACKEND_URL}/api/results`, {
      params: { proposalId }
    });
    res.json(response.data);
  } catch (error) {
    console.error('Error fetching results:', error.message);
    res.status(500).json({ error: 'Failed to fetch results' });
  }
});

// Health check
app.get('/health', (req, res) => {
  res.json({ status: 'Frontend healthy' });
});

app.listen(PORT, () => {
  console.log(`Frontend server running on http://localhost:${PORT}`);
  console.log(`Backend connected to ${BACKEND_URL}`);
});