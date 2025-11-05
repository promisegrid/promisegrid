class: center, middle

# PromiseGrid

**A Decentralized Computing System for Collaborative Governance**

*Consensus-based computing, communications, and governance*

---

## What is PromiseGrid?

PromiseGrid is a **decentralized computer** — to computation what the Internet is to communication.

- Owned and operated by users, not central entities
- Scalable and resilient, growing as users join  
- Addresses tragedy of the commons through collaborative governance
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

## Kernel Internal Architecture: PromiseBase Layers

The PromiseGrid kernel is structured in three layers, with PromiseBase providing the foundation:

```
       ┌──────────────────────────────────┐
       │    Protocol Handlers & Runtimes  │
       │  (WASM, Containers, VMs, Bare)   │
       └──────────────────┬───────────────┘
                          │
       ┌──────────────────┴───────────────┐
       │  Dispatcher(s) & Capability Mgmt │
       │  (Request routing, scheduling)   │
       └──────────────────┬───────────────┘
                          │
       ┌──────────────────┴───────────────┐
       │      PromiseBase                 │
       │  (KV Store, transactions, state) │
       └──────────────────────────────────┘
```

**PromiseBase:** Content-addressable key-value store with transactional semantics

**Dispatcher:** Routes messages to handlers, manages resources, enforces capabilities

**Runtimes:** Execute protocol handlers in isolated sandboxes

---

## PromiseBase: Key-Value Store Foundation

PromiseBase provides distributed, replicated transactional key-value storage at the kernel's core.

**Key Properties:**
- **Distributed** — Data partitioned across nodes
- **Replicated** — Fault tolerance through redundancy
- **Transactional** — ACID semantics for consistency
- **Content-Addressed** — Keys are CIDs (hashes of content)
- **Pure Functional** — Referential transparency enables caching

**Storage Layout:**
```
keys:    [CID₁][CID₂][CID₃]...[CIDₙ]
values:  [val₁][val₂][val₃]...[valₙ]
```

**Values stored:**
- Protocol handler code (WASM modules)
- Application code and libraries
- Grid configuration and metadata
- Cached message responses
- Capability tokens and proofs

---

## Dispatcher Layer: Request Routing & Scheduling

The dispatcher sits above PromiseBase, routing incoming messages to appropriate handlers and managing resource allocation.

**Core Responsibilities:**

.left-column[
### Message Processing
- Parse CBOR 5-element messages
- Validate protocol tags & CIDs
- Extract signatures & claims
- Verify cryptographic proofs
]

.right-column[
### Capability Management
- Track node capabilities
- Enforce access control
- Resolve handler CIDs
- Schedule execution
- Monitor resource budgets
]

**Dispatcher Architecture:**
```
Incoming Message → Validation → CID Resolution → 
  ↓
  Handler Dispatch → Resource Accounting → 
  ↓
  Response Generation → Signature → Return
```

---

## Dispatcher: Inter-Function Communication

The dispatcher enables efficient communication patterns between functions and nodes:

**Local Execution** — Handler on same host uses shared memory or Unix sockets

**Remote Invocation** — Handler on different node uses network transport

**Nested Handlers** — Messages contain other messages for higher-order composition

**Hybrid Selection** — Runtime auto-selects optimal communication method

.smaller[
```
Message from A to B:
├─ Is B on same host?
│  ├─ Yes → Use local buffer/Unix socket (microseconds)
│  └─ No → Serialize & network transport (milliseconds)
└─ Is response cached?
   ├─ Yes → Return from cache (nanoseconds)
   └─ No → Execute & store result
```
]

---

## Runtime Layer: Execution Environments

Runtimes execute protocol handler code in isolated sandboxes, sitting atop the dispatcher.

**Supported Runtime Types:**

| Runtime | Sandbox | Use Case | Performance |
|---------|---------|----------|-------------|
| **WASM** | VM | Browsers, mobile, edge | Near-native |
| **Container** | OS processes | Complex apps, legacy code | Native |
| **VM** | Hypervisor | Strong isolation, multi-OS | Native |
| **Bare Metal** | Kernel namespace | Dedicated resources | Native |

**Security Properties:**
- Memory isolation (bounds checking in WASM)
- No privileged instruction access
- Capability-based resource access
- Deterministic execution for verification

---

## Runtime Handler Execution Flow

```
Handler CID Resolved
        │
        ▼
Lookup/Load in Runtime (cache or fetch from PromiseBase)
        │
        ▼
Allocate Sandbox Resources (memory, CPU time, network quota)
        │
        ▼
Execute Handler with Input (pure function semantics)
        │
        ▼
Monitor Resource Consumption (preempt if budgets exceeded)
        │
        ▼
Capture Output/Response
        │
        ▼
Cache Result in PromiseBase
        │
        ▼
Return to Dispatcher
```

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

Each node: Independent kernel, replicated state, cached content
Coordination: Capability tokens, message signatures, merge consensus
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
Issuer → Creates CWT-based capability including function hash (CID)
         ↓
Holder → Calls function via CID (presents capability)
         ↓
Issuer → Fulfills, revokes, or defers until later
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
- Source code merges
- Race condition resolution
- Application data conflicts  
- Organizational disputes
- Community decisions

---


## Five-Element CBOR Message Structure

PromiseGrid messages use a **5-element CBOR array** providing self-identification, routing, isolation, payload transport, and cryptographic verification.

```
[
  "grid",              // Protocol tag
  protocol_cid,        // Handler routing  
  grid_cid,            // Instance isolation
  cwt_payload,         // Payload data
  signature            // Cryptographic proof
]
```

**Technologies:**
- **CBOR** = Concise Binary Object Representation (RFC 8949)
- **CID** = Content Identifier (based on multihash)
- **CWT** = CBOR Web Token (RFC 8392)
- **COSE** = Cryptographic Object Signing (RFC 8152/9052)

---

## Element 1: Protocol Tag

```
[ "grid", ... ]
   ↑
   Self-identifying protocol marker
```

**Purpose:** Immediate protocol recognition

**Implementation:** UTF-8 text string "grid" (5 bytes with CBOR header)

**Benefits:**
- Fast message validation before parsing
- Human-readable in debugging
- Network analysis tool recognition
- Future protocol variant support ("grid2", "grid-alt")

---

## Element 2: Protocol CID

```
[ "grid", protocol_cid, ... ]
           ↑
           Hash of protocol handler code
```

**Purpose:** Routes message to correct protocol handler

**Key Properties:**
- Content-addressed protocol implementation
- Exact version specification (not just version numbers)
- Automatic code distribution and caching
- Multiple protocol versions coexist

**Example:** SHA2-256 CID ≈ 34-36 bytes

---

## Element 3: Grid Instance CID

```
[ "grid", protocol_cid, grid_cid, ... ]
                         ↑
                         Network namespace identifier
```

**Purpose:** Isolates communication within specific grid instances

**Enables:**
- Multiple independent grids on shared infrastructure
- Test/development/production separation
- Private organizational grids
- Industry or geographic-specific grids
- Capability scope isolation

**Security:** Cross-grid message injection prevention

---

## Element 4: CWT Payload

```
[ "grid", protocol_cid, grid_cid, cwt_payload, ... ]
                                   ↑
                                   CBOR Web Token claims
```

**Purpose:** Application-specific message content

**Standard Claims:**
- Issuer (iss), Subject (sub), Audience (aud)
- Expiration (exp), Not-before (nbf), Issued-at (iat)
- CWT identifier (cti)

**Custom Claims:** Application-defined via IANA registry

**Proof-of-Possession:** Confirmation claim (cnf) with COSE keys

---

## Element 5: Signature

```
[ "grid", protocol_cid, grid_cid, cwt_payload, signature ]
                                                ↑
                                                COSE signature structure
```

**Purpose:** Cryptographic authenticity and integrity

**COSE Signature Contains:**
- Protected headers (algorithm, countersignatures)
- Unprotected headers (key identifier, certificates)
- Signature value over all preceding elements

**Algorithms:** ECDSA, EdDSA, RSA-PSS variants

**Security:** End-to-end protection over untrusted networks

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

## Message Processing Pipeline

.smaller[
**Stage 1: Structure Validation**
- Parse CBOR 5-element array
- Reject malformed messages early

**Stage 2: Protocol Tag Check**
- Verify "grid" tag
- Reject misrouted messages

**Stage 3: CID Validation**
- Extract protocol_cid and grid_cid
- Fetch unknown handlers/configs

**Stage 4: Payload Processing**
- Decode CWT claims
- Validate timestamps and proof-of-possession

**Stage 5: Signature Verification**
- Reconstruct Sig_structure
- Verify cryptographic signature
- Pass to protocol handler only on success
]

---
## Implementation Patterns

**Content-Addressable Code** — Every function has a unique CID derived from its content

**Pure Functions** — Protocol handlers observe referential transparency

**Nested Messages** — Messages can contain messages (higher-order functions)

**Distributed Cache** — Nodes cache protocol handlers and responses

**Automatic Code Distribution** — Unknown CIDs trigger content fetch

**Graceful Degradation** — Legacy protocol support via CID versioning

---

## Distributed Network Architecture

.smaller.diagram[
```
Node A                  Node B                  Node C
┌────────┐             ┌────────┐             ┌────────┐
│Protocol│             │Protocol│             │Protocol│
│Handlers│             │Handlers│             │Handlers│
├────────┤             ├────────┤             ├────────┤
│ Kernel │◄────────────┤ Kernel │◄────────────┤ Kernel │
│ Cache  │────────────►│ Cache  │────────────►│ Cache  │
└────────┘             └────────┘             └────────┘
    ▲                      ▲                      ▲
    │                      │                      │
    └──────────────────────┴──────────────────────┘
         Content-Addressed Message Routing
         (Each kernel validates and routes 
          based on CIDs and capabilities)
```
]

**Properties:** Peer-to-peer communication, distributed cache, autonomous nodes, capability-based security across boundaries

---

## Technology Stack

**Execution Environment**
- WebAssembly (WASM) virtual machines
- Container orchestration support
- Bare metal deployment capability

**Data Formats**
- CBOR (RFC 8949) - Binary encoding
- CWT (RFC 8392) - Token structure
- COSE (RFC 8152/9052) - Cryptography
- CID/Multihash - Content addressing

**Security**
- Capability-based access control
- Content-addressed code verification
- End-to-end message signatures

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

## Use Cases

**Community Systems** — Collaborative editing, chat, video, event scheduling

**DevOps Automation** — Container orchestration, configuration management, infrastructure as code

**Enterprise Applications** — ERP, inventory, project management, accounting

**Version Control** — Distributed repositories with issue tracking and code review

**Governance** — Organizational decision-making using consensus mechanisms

**Resource Management** — Preventing tragedy of commons in shared systems

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

---

## Get Involved

**Watch & Contribute**
- Star repos on GitHub
- Join community discussions
- Submit pull requests

**Community Channels**
- CSWG meetings (weekly)
- IoP collaboration
- Direct contact via GitHub/LinkedIn/Twitter

**Current Focus**
- WASI target development
- Example application ports
- Self-hosting milestone

---

## Sponsorship

Development supported by:

**C D International Technology, Inc.**
[www.cdint.com](http://www.cdint.com)

**TerraLuna, LLC**
[www.t7a.org](http://www.t7a.org)

---

class: center, middle

## Questions?

### PromiseGrid: Decentralized Computing for Collaborative Governance

---

## References

.smaller[
**CBOR & Message Format**
- RFC 8949: Concise Binary Object Representation
- RFC 8392: CBOR Web Token (CWT)
- RFC 8152/9052: CBOR Object Signing and Encryption (COSE)
- RFC 8747: CWT Proof-of-Possession

**Content Addressing**
- Multiformats Project: multihash, multicodec, multibase
- IPFS CID Specification
- Content-Addressable Storage systems

**Tragedy of Commons Examples**
- Harvard Business School: Sustainability Issues
- NCBI: Climate Change as Commons Problem
- Dummies.com: Ten Real-Life Examples

**Kernel Architecture**
- SPIN: Extensible Microkernel for Application-specific Operating Systems
- Windows Kernel Architecture
- CWASI: WebAssembly Runtime Shim for Inter-function Communication
]

