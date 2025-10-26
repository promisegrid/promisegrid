class: center, middle, inverse

# PromiseGrid

**A Decentralized Computing System for Collaborative Governance**

*Consensus-based computing, communications, and governance*

---

## What is PromiseGrid?

PromiseGrid is a **decentralized computer** — to computation what the Internet is to communication.

- Owned and operated by users, not central entities
- Scalable and resilient, growing as users join  
- Addresses tragedy of the commons through automated governance
- Built on proven technologies: WebAssembly, WASI, cryptography, LLMs

Key insight: The same consensus and governance mechanisms that manage the grid can govern organizations and communities using it.

---

## The Problem: Tragedy of the Commons

Individual users acting in self-interest can collectively harm shared resources.

**Resource types affected:**
- Organization well-being
- Physical resources
- Global resources
- Digital resources  
- Community trust

**PromiseGrid solution:** Automate resource management using community-defined algorithms to align individual incentives with collective good.

---

## User-Visible Features

.left-column[
### Accessibility
- Run in browser tabs
- Mobile devices
- Raspberry Pi
- Command line
- URL-loaded apps

### Control
- Choose your software version
- Own your node
- No forced upgrades
]

.right-column[
### Composability
- Mixed-language modules
- Modular applications
- Reusable libraries
- Plugin architecture

### Reliability
- Grid survives failures
- Data replication
- Automatic fallback
- Decentralized operation
]

---

## Admin-Visible Features

### Infrastructure Management & DevOps Benefits

.left-column[
### Operational Excellence
- **Zero centralized hosting costs** — each org runs own nodes
- **Automatic scaling** — grid grows with demand
- **Self-healing infrastructure** — nodes rejoin seamlessly
- **Cross-organization resource sharing** — efficient utilization
- **Infrastructure as Code** — declarative config management
]

.right-column[
### DevOps Integration
- **Bare-metal orchestration** — disk image + config management
- **Container support** — Docker/OCI integration
- **CI/CD automation** — built-in deployment pipelines
- **Observability** — capability graph tracking
- **Disaster recovery** — automated replication and backup
- **Multi-environment consistency** — VMs, containers, WASM, bare-metal
]

---

## Organizational Features

### Benefits to Organizations & Communities

.left-column[
### Autonomy & Sovereignty
- **Independent operation** — no vendor lock-in
- **Data ownership** — complete control of assets
- **Governance control** — define your own rules
- **Supply chain resilience** — decentralized participation
- **Regulatory compliance** — no middleman data sharing
]

.right-column[
### Collaboration & Scale
- **Multi-org federation** — coordinate without centralization
- **Reduced costs** — no SaaS subscriptions
- **Faster innovation** — deploy directly to grid
- **Democratic governance** — built-in consensus mechanisms
- **Community preservation** — content persists independently
]

---

## Architecture Overview

```
           ┌─────────────────────────────────────┐
           │         Applications Layer          │
           │  (Modules, Services, Containers)    │
           └──────────────────┬──────────────────┘
                              │
                    ┌─────────┴─────────┐
                    │                   │
             WASM Runtime         Other Runtimes
            (Browser/CLI)        (VM, Container)
                    |                   │
                    └─────────┬─────────┘
                              │
                ┌─────────────┴─────────────┐
                │                           │
                │  PromiseGrid Kernel       │
                │  (Syscall-like services)  │
                │                           │
                └───────────────────────────┘
```

The kernel abstracts away underlying infrastructure while providing unified security, consensus, and governance services.

---

## Distributed Grid Architecture

Multiple autonomous nodes, each running their own kernel instance, forming a coordinated decentralized network:

```
    ┌──────────────┐          ┌──────────────┐
    │    Node A    │          │    Node B    │
    │  ┌────────┐  │          │  ┌────────┐  │
    │  │ Kernel │◄─┼──────────┼─►│ Kernel │  │
    │  └────────┘  │          │  └────────┘  │
    │  Apps/Data   │          │  Apps/Data   │
    └──────────────┘          └──────────────┘
            ▲                          ▲
            │                          │
            └──────────────┬───────────┘
                           │
                    ┌──────────────┐
                    │   Node C     │
                    │  ┌────────┐  │
                    │  │ Kernel │  │
                    │  └────────┘  │
                    │  Apps/Data   │
                    └──────────────┘

Messages flow between nodes via capability tokens
Each node maintains its own cache and state
```

---

## Content-Addressable Code with CIDs

Every piece of code and data is addressed by its **Content Identifier (CID)**, derived from cryptographic hashing.

**CID Structure (based on Multihash):**
- **Version**: CIDv0 or CIDv1
- **Codec**: Data format (dag-pb, raw, dag-json)
- **Multihash**: Algorithm + hash length + hash digest
- **Multibase**: Encoding (base58, base32, base64)

**Benefits:**
- Same standard used by **IPFS** and **Bluesky (AT Protocol)**
- Unique, unambiguous addresses
- Future-proof hash algorithm agility
- Automatic deduplication  
- Tamper-proof integrity verification

Example CID: `bafyreigmitjgwhpx2vgrzp7knbqdu2ju5ytyibfybll7tfb7eqjqujtd3y`

---

## Capability-as-Promise Model

Capabilities represent **promises** — in both the JavaScript and Promise Theory sense.

```
Issuer → Creates closure with function hash (CID)
         ↓
Holder → Calls closure via CID (presents capability)
         ↓
Issuer → Fulfills, revokes, or promises later
```

**Key principles:**
1. Agents can only make promises about their own behavior
2. Agent should not impose obligations on others

This design mirrors both Promise Theory and healthy organizational behavior.

---

## Merge-as-Consensus Model

Consensus formation mirrors version control **merge** operations.

**How it works:**
- Multiple participants propose versions
- Merge function reconciles differences
- Success: single coherent version
- Failure: escalate to LLM, then human

**Applies at all levels:**
- Race condition resolution
- Application data conflicts  
- Organizational disputes
- Community decisions

---

## PromiseGrid Universal Protocol

**Message structure using CIDs:**

```
┌─────────────────────────────────────┐
│  Capability Token (function CID)    │
├─────────────────────────────────────┤
│  Payload (arguments/nested msgs)    │
└─────────────────────────────────────┘
```

**Key characteristics:**
- No version numbers in header
- Token doubles as protocol version
- Supports nested messages and hyperkernels
- CID-based function addressing
- Extensible by design

---

## Decentralized Cache

Distributed content store using directed graph database, indexed by CIDs.

**How it works:**
1. Node receives message (capability CID + arguments)
2. Checks local cache for CID+args combination
3. Cache hit → reply immediately  
4. Cache miss → forward to capable nodes
5. Receive response → cache locally + forward to sender

**Advantage:** Pure functions with referential transparency enable aggressive caching across the network.

---

## Technology Stack

| Component | Technology |
|-----------|-----------|
| **Execution** | WebAssembly (WASM), containers, VMs |
| **System Interface** | WASI (WebAssembly System Interface) |
| **Storage** | Content-addressable with CIDs (Multihash-based) |
| **Security** | Capability-based access control |
| **Consensus** | Promise Theory + Merge algorithms |
| **Language AI** | Large language models (LLMs) |
| **Configuration** | Infrastructure as Code (IaC) |

---

## Use Cases

.pull-left[
### Community Systems
- Makerspaces
- Non-profits
- Open-source projects
- Cooperative organizations

### Enterprise
- Distributed microservices
- Multi-org collaboration
- Supply chain coordination
]

.pull-right[
### Development Tools
- Version control (CID-based)
- Collaborative editors
- Package management
- CI/CD pipelines

### Infrastructure
- Container orchestration
- Resource scheduling
- Configuration management
- DevOps automation
]

---

## Project Roadmap

.left-column[
### Completed
- Community Systems WG
- LLM bootstrap tools
- Promise Theory foundation
- Architecture design

### In Progress  
- WASI target implementation
- Proof-of-concept code
- Example applications
]

.right-column[
### Upcoming
- Example application ports
- Community infrastructure
- Enterprise ERP tools
- Self-hosting migration

### Long-term
- Full mainnet launch
- Ecosystem maturity
- Global adoption
]

---

## Why PromiseGrid Matters

**Problem:** Centralized systems create single points of failure and concentrations of power.

**Vision:** A computing infrastructure where organizations and communities maintain autonomy while achieving coordination and scale.

**Impact:** Could fundamentally restructure how software, governance, and resource management work at scale.

---

## How to Get Involved

.left-column[
### For Developers
- Explore the codebase
- Contribute modules
- Port applications
- Test in browsers/CLI
]

.right-column[
### For Communities
- Join CSWG
- Shape governance
- Define consensus rules
- Participate in design
]

**Contact:** [Steve Traugott](https://github.com/stevegt) · [Twitter](https://twitter.com/stevegt) · [LinkedIn](https://linkedin.com/in/stevegt/)

---

## Current Work

### Primary Projects
- **PromiseGrid Core:** [github.com/promisegrid/promisegrid](https://github.com/promisegrid/promisegrid)
- **Grid POC:** [github.com/promisegrid/grid-poc](https://github.com/promisegrid/grid-poc)
- **PromiseBase:** [github.com/stevegt/promisebase](https://github.com/stevegt/promisebase)
- **Collab Editor:** [github.com/computerscienceiscool/collab-editor](https://github.com/computerscienceiscool/collab-editor)
- **Grid CLI:** [github.com/stevegt/grid-cli](https://github.com/stevegt/grid-cli)
- **Grokker (LLM Tool):** [github.com/stevegt/grokker](https://github.com/stevegt/grokker)
- **ToyGrid (Microfrontend PoC):** [github.com/stevegt/toygrid](https://github.com/stevegt/toygrid)

### Community & Governance
- **CSWG Home:** [cswg.infrastructures.org](https://cswg.infrastructures.org)
- **Maker Community Practices:** [mcp.infrastructures.org](https://mcp.infrastructures.org)
- (spinoff from) **Nation of Makers:** [nationofmakers.us](https://www.nationofmakers.us/community-systems-working-group)

### Related Work and Background

- Mark Burgess. [Promise Theory](https://markburgess.org/promises.html)
- WebAssembly Foundation. [WASM Specification](https://webassembly.org/)
- WASI Development Group. [WebAssembly System Interface](https://wasi.dev/)
- Multiformats Project. [Multihash Specification](https://multiformats.io/multihash/)
- Protocol Labs. [IPFS Content Addressing](https://docs.ipfs.tech/concepts/content-addressing/)
- Bluesky Protocol. [AT Protocol CID Usage](https://atproto.blue/)

