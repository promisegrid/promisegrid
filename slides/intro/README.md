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

**PromiseGrid solution:** Automate resource management using community-defined algorithms.

---

## Core Features

.left-column[
### Ownership
- Users own nodes
- No centralized hosting

### Flexibility
- Browser tabs
- Mobile phones
- Raspberry Pi
- Command line
]

.right-column[
### Resilience
- Grid survives node failures
- Decentralized operation
- Content replication

### Security
- Capability-based access
- Fine-grained control
- Cryptographic integrity
]

---

## Architecture Overview

```
┌─────────────────────────────────────┐
│     PromiseGrid Kernel              │
│  (Decentralized, Syscall-like)      │
└──────────┬──────────────────────────┘
           │
    ┌──────┴──────┬──────────┬────────┐
    │             │          │        │
  WASM       Containers    VMs    Bare-metal
 (Browser)   (Docker)   (KVM/VirtualBox)  (Physical)
```

The kernel acts as a sandbox orchestrator, managing execution across multiple environments while maintaining a unified security and governance model.

---

## Content-Addressable Code

Every piece of code and data is addressed by its **cryptographic hash**, not location or name.

**Benefits:**
- Unique, unambiguous addresses
- Automatic deduplication
- Global namespace
- Unlimited storage (256/512-bit hashes)
- Tamper-proof integrity verification

Implementation: Uses [Multiformats](https://multiformats.io/multihash/) for future-proof hash agility.

---

## Capability-as-Promise Model

Capabilities represent **promises** — in both the JavaScript and Promise Theory sense.

```
Issuer → Creates closure with function hash
         ↓
Holder → Calls closure (presents capability)
         ↓
Issuer → Fulfills, revokes, or promises later
```

**Key principles:**
1. Resources cannot issue promises for other resources
2. Autonomous parts should not impose obligations

This design mirrors both Promise Theory and healthy organizational behavior.

---

## Merge-as-Consensus Model

Consensus formation mirrors version control **merge** operations.

**How it works:**
- Multiple participants propose versions
- Merge function reconciles differences
- Success: single coherent version
- Failure: human intervention (escalate to LLM)

**Applies at all levels:**
- Race condition resolution
- Application data conflicts  
- Organizational disputes
- Community decisions

---

## PromiseGrid Universal Protocol

**Message structure:**

```
┌─────────────────────────────────────┐
│  Capability Token (function hash)   │
├─────────────────────────────────────┤
│  Payload (arguments/nested msgs)    │
└─────────────────────────────────────┘
```

**Key characteristics:**
- No version numbers in header
- Token doubles as protocol version
- Supports nested messages
- Extensible by design
- Enables nested kernels (hyperkernels)

---

## Decentralized Cache

Distributed content store using directed graph database.

**How it works:**
1. Node receives message (token + arguments)
2. Checks local cache for token+args combination
3. Cache hit → reply immediately  
4. Cache miss → forward to capable nodes
5. Receive response → cache + forward to sender

**Advantage:** Pure functions with referential transparency enable aggressive caching across the network.

---

## Technology Stack

| Component | Technology |
|-----------|-----------|
| **Execution** | WebAssembly (WASM), containers, VMs |
| **System Interface** | WASI (WebAssembly System Interface) |
| **Storage** | Content-addressable with Multihash |
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
- Version control
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

## References

## Current Work 

### Primary Projects
- **PromiseGrid Core:** [github.com/promisegrid/promisegrid](https://github.com/promisegrid/promisegrid)
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

