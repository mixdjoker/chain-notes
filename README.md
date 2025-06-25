# Chain Notes

**⚠️PROOF OF CONCEPT⚠️**

**Chain Notes** is a secure, blockchain-powered notebook system designed for storing sensitive information with version control and strong encryption. Inspired by Git and modern distributed systems, Chain Notes ensures that your notes are tamper-proof, versioned, and accessible only to their rightful owners.

## Key Features

- 🧱 Blockchain-Backed History  
Every change is committed to an immutable blockchain, providing transparent and verifiable history.

- 🔐 End-to-End Encryption  
Notes are encrypted on the client side. Only those with the correct decryption key can view the content.

- 🛰 Distributed Storage  
Content is stored using IPFS/Filecoin or other decentralized storage backends for maximum availability.

- ⚡ Message-Driven Architecture  
The system communicates through NATS for modularity, scalability, and fault-tolerance.

- 🛠 Built with Go and Rust  
Backend services are written in Go with performance-critical components implemented in Rust.

## Use Cases

- Private knowledge management

- Secure journaling and note-taking

- Encrypted collaborative documentation

- Blockchain-based changelog or audit log systems

## Getting Started

Coming soon: setup guide, running the services locally with Docker, and usage examples.

## Architecture

```
                         ┌──────────────┐
                         │  Web / CLI   │
                         └─────┬────────┘
                               │
                               ▼
                        ┌──────────────┐
                        │  Commit API  │ ◄──── User sends encrypted content + metadata
                        └─────┬────────┘
                              │
               ┌──────────────┼──────────────┐
               ▼                              ▼
      ┌────────────────┐           ┌───────────────────┐
      │  CommitService │           │  StorageService   │
      └──────┬─────────┘           └─────────┬─────────┘
             │                               │
             ▼                               ▼
     ┌───────────────┐              ┌────────────────────┐
     │  BlockchainDB │              │  IPFS/Filecoin/etc │
     └───────────────┘              └────────────────────┘
```

Each note is:

1. Encrypted locally.
1. Wrapped in a signed commit.
1. Stored in IPFS or another backend.
1. Registered on-chain with its hash and metadata.
1. Replicated and accessible based on trust and key ownership.

## How It Works

1. Create a Note  
You write a note. It’s encrypted using a symmetric key.

2. Commit & Sign  
The encrypted note and metadata (e.g. parent hash) are bundled into a Git-style commit. It’s signed with your private key.

3. Store Content  
The encrypted payload is pushed to distributed storage (e.g. IPFS). A content address (CID) is returned.

4. Publish to Chain  
The commit hash, CID, and signature are written to a blockchain-like ledger through the CommitService.

5. Verify & Fetch  
Anyone with the right key can fetch the commit chain, verify signatures, and decrypt the note contents.

## License

This project is licensed under the MIT License. See the LICENSE file for details.