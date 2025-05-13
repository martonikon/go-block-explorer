package explorer

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/martonikon/go-block-explorer/blockchain"
)

func GetBlocks(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(filepath.Join("data", "chain.json"))
	if err != nil {
		http.Error(w, "Failed to read blockchain data", http.StatusInternalServerError)
		return
	}

	var blocks []blockchain.Block
	if err := json.Unmarshal(data, &blocks); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blocks)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(filepath.Join("data", "chain.json"))
	if err != nil {
		http.Error(w, "Failed to read blockchain data", http.StatusInternalServerError)
		return
	}

	var blocks []blockchain.Block
	if err := json.Unmarshal(data, &blocks); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusInternalServerError)
		return
	}

	// Gather all transactions
	var allTxs []blockchain.Transaction
	for _, block := range blocks {
		allTxs = append(allTxs, block.Transactions...)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allTxs)

}
