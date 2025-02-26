# Ethereum Indexer

I set out to build a Retrieval-Augmented Generation (RAG) chatbot that leverages the Ethereum blockchain as its knowledge base. The most straightforward approach was integrating my application with Etherscan, but its free tier imposes a 100,000 request-per-day limit, which quickly becomes a bottleneck.

Coincidentally, at work, I've built similar systems designed to index and process data from Kafka topics. Inspired by that experience, I decided to take on the challenge of building a lightweight Ethereum indexer that pulls blockchain data directly, stores it in PostgreSQL, and enables structured SQL queries for efficient retrieval. This allows my chatbot to query blockchain data seamlessly without relying on third-party APIs.

**Should you use in production?**

Probably not. If you need a robust, production-grade solution, I’d recommend using established services like Etherscan, Alchemy, Infura, or The Graph, which offer highly scalable and optimized indexing solutions.

This project was built purely for fun and experimentation. If you're interested in playing around with Ethereum data locally, I’d suggest running this indexer with a testnet like Sepolia or Holesky.

## Architecture
![Architecture](https://i.imgur.com/1Z2Z2Zz.png)

## Tech Stack
- [Golang](https://golang.org/)
  - [Gorilla Mux](github.com/gorilla/mux)
  - [Gorm](github.com/go-gorm/gorm)
  - [ZeroLog](github.com/rs/zerolog/log)
  - [Ethclient](github.com/ethereum/go-ethereum/ethclient)
  - [Templ](github.com/a-h/templ) w/ [Tailwindcss](https://tailwindcss.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Redis](https://redis.io/)
- [Htmx](https://htmx.org/)
---

<img src="https://i.imgur.com/N0CHCSd.gif" alt="Demo" />

## Roadmap
- [x] Add authentication for the API
- [ ] Add an Etherscan-Like UI to view indexed data
- [ ] Setup my own Ethereum node using Geth/Erigon
- [ ] Add Support for other EVM chains

## Setup

1. Clone the repository
```bash
  git clone github.com/conceptcodes/eth-indexer-go.git
  cd eth-indexer-go
```

2. Install dependencies
```bash
  go mod tidy
```

3. Copy the .env.example file to .env and update the values
```bash
  cp .env.example .env
```

4. Run the application
```bash
  docker-compose up --build
```

## Usage

**Health Check**

```bash
curl -X GET \
  https://localhost:8080/api/v1/health/alive
```
```json
{
    "message": "Server is running",
    "data": null
}
```

---

**Get Transaction by Hash**

```bash
curl -X GET \
  https://localhost:8080/api/v1/tx/0x4d72ea05d79c6098420d1a134749881a1e2966422d8712c2b6b8b436343cdab4
```
```json
{
    "message": "Found Transaction with id 0x4d72ea05d79c6098420d1a134749881a1e2966422d8712c2b6b8b436343cdab4.",
    "data": {
        "hash": "0x4d72ea05d79c6098420d1a134749881a1e2966422d8712c2b6b8b436343cdab4",
        "from": "0x89798BD94DA6bb9d3e7712dD3B08de633f43A629",
        "to": "0x800eC0D65adb70f0B69B7Db052C6bd89C2406aC4",
        "amount": "",
        "gas_price": "13240306042",
        "gas_limit": 26849,
        "gas_used": 0,
        "nonce": 0,
        "block_number": 7624880,
        "value": "0"
    }
}
```

---

**Get Block by Number**

```bash
curl -X GET \
  http://localhost:8080/api/v1/block/7624880'
```
```json
{
    "message": "Found Block with id 7624880.",
    "data": {
        "number": 7624880,
        "hash": "0x64e58f944dc4b2a166e25ef0abafcf63719e31e5e7617b24e413e5c9855fef9f",
        "parent_hash": "0x41c8e615b864339f5e264878b63e499e20229b5fbbcd338546e18157d7a36f39",
        "size": 159275,
        "miner": "0x0000000000000000000000000000000000000000",
        "timestamp": 1738507392,
        "transactions": [
            {
                "hash": "0xee647627ce6387817fc9129164257e1c70ec32131c494ea0ad0ef8c82f9d51d6",
                "from": "0x0000000000000000000000000000000000000000",
                "to": "0xd0dbB2486e3Fbf371D1B2e35Fa330bef6529d2f4",
                "amount": "",
                "gas_price": "22691913422",
                "gas_limit": 400000,
                "gas_used": 0,
                "nonce": 0,
                "block_number": 7624880,
                "value": "0"
            },
            {
                "hash": "0xc390761d8b1c4cbed7285428595c6818bc13c85b3e2ff943c335bc5c08938623",
                "from": "0x2CdA41645F2dBffB852a605E92B185501801FC28",
                "to": "0xc36f77136058EaC4c199933126f4306f4f1997d0",
                "amount": "",
                "gas_price": "100000000000",
                "gas_limit": 21000,
                "gas_used": 0,
                "nonce": 0,
                "block_number": 7624880,
                "value": "50000000000000000"
            },
            ...
        ],
    }
}
```

## Resources
https://www.youtube.com/watch?v=WgBab6kamtg