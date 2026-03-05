# NeuroCache

> **Memory for AI systems. Speed for everything else.**

NeuroCache is a high-performance, in-memory data platform built for AI systems, real-time applications, and streaming workloads. It unifies the capabilities of a cache, an event stream, and a vector store into a single, blazing-fast engine — so you never have to stitch together Redis, Kafka, and a vector database again.

---

## Vision

Modern AI and real-time applications suffer from infrastructure sprawl. Teams manage multiple specialized systems — a cache for speed, a message broker for events, a vector database for AI memory — each with its own ops burden, latency overhead, and consistency challenges.

NeuroCache solves this with **one unified engine**:

| Instead of          | Use NeuroCache |
| ------------------- | -------------- |
| Redis               | KV Engine      |
| Apache Kafka        | Stream Engine  |
| Pinecone / Weaviate | Vector Memory  |

---

## Core Features (MVP)

### 1️⃣ KV Engine — Cache

A fast key-value store for session caching, API caching, and rate limiting.

```
SET key value
GET key
DEL key
TTL key
```

### 2️⃣ Vector Memory — AI Feature

Store and search high-dimensional embeddings for RAG pipelines, semantic search, and AI agent memory.

```
VSET key vector
VSEARCH vector k
```

### 3️⃣ Stream Engine — Real-Time Events

A lightweight event streaming engine for analytics, notifications, and AI event pipelines.

```
XADD stream message
XREAD stream
```

---

## Performance Targets

| Metric       | Target      |
| ------------ | ----------- |
| Latency      | < 1ms       |
| Throughput   | 5M ops/sec  |
| Storage      | In-memory   |
| Scalability  | Multi-core  |
| Availability | Replication |

---

## System Architecture

```
                Client
                   │
            TCP/HTTP Server
                   │
            Command Parser
                   │
               Router
        ┌──────────┼──────────┐
        │          │          │
   KV Engine   Vector      Streams
               Engine
```

Each engine runs on sharded workers, isolated per-core for maximum throughput and predictable latency.

---

## Concurrency Model

NeuroCache does **not** use a global lock. Instead, each shard owns its memory, runs on a dedicated goroutine, and processes commands sequentially — eliminating lock contention entirely.

```
               Router
                 │
      ┌──────────┼──────────┐
      │          │          │
   Shard 1    Shard 2    Shard 3
   Worker     Worker     Worker
```

**Benefits:**

- No lock contention
- High parallelism
- Predictable latency

---

## Communication Protocol

NeuroCache speaks **RESP** — the same protocol used by Redis. Drop-in compatibility means you can use existing Redis clients with zero changes.

Example wire format:

```
*3
$3
SET
$4
name
$4
saad
```

---

## Project Structure

```
neurocache/
├── cmd/
│   └── server/
├── internal/
│   ├── network/
│   ├── protocol/
│   ├── router/
│   ├── shard/
│   └── engine/
│       ├── kv/
│       ├── vector/
│       └── stream/
├── pkg/
├── configs/
└── go.mod
```

> **Rule:** one package = one responsibility.

---

## Technology Stack

| Layer         | Technology                              |
| ------------- | --------------------------------------- |
| Language      | Go (Golang)                             |
| Networking    | `net` package (epoll/kqueue internally) |
| Vector Search | HNSW / ANN _(planned)_                  |

---

## Roadmap

- [x] KV Engine with TTL support
- [x] RESP protocol compatibility
- [x] Sharded concurrency model
- [ ] Vector Memory with HNSW indexing
- [ ] Stream Engine with consumer groups
- [ ] Replication & persistence layer
- [ ] Cluster mode

---

## License

MIT
