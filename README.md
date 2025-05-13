# Go Blockchain Explorer 🔍

A minimal blockchain ledger viewer built in Go to showcase networking, crypto basics, and Go’s concurrency model.

---

## 🚀 Features

- View block information at `/blocks`
- View all transactions at `/txs`
- Serves mock blockchain data from a local JSON file
- Built using Go’s `net/http` package

---

## 🛠 Tech Stack

- Language: **Go**
- No frameworks—just standard library
- JSON I/O for mock blockchain data

---

## 📂 Project Structure
go-block-explorer/
│
├── main.go # Entry point, sets up routes
├── go.mod # Module info
├── blockchain/ # Data models (Block, Transaction)
│ └── blockchain.go
├── explorer/ # HTTP route handlers
│ └── routes.go
├── data/
│ └── chain.json # Mock blockchain data
└── README.md


## 🚴 Getting Started

```bash
git clone https://github.com/martonikon/go-block-explorer.git
cd go-block-explorer
go run main.go


Visit:
http://localhost:8080/blocks → list of all blocks

http://localhost:8080/txs → list of all transactions