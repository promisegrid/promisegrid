class: center, middle, inverse

# PromiseGrid

### Consensus-Based Computing, Communications, and Governance

---

## Current Work

**Grid Proof of Concept**
- [github.com/promisegrid/grid-poc](https://github.com/promisegrid/grid-poc)
- [github.com/stevegt/promisebase](https://github.com/stevegt/promisebase)
- [github.com/computerscienceiscool/collab-editor](https://github.com/computerscienceiscool/collab-editor)

**Community & Documentation**
- [cswg.infrastructures.org](https://cswg.infrastructures.org)
- [www.nationofmakers.us/community-systems-working-group](https://www.nationofmakers.us/community-systems-working-group)
- [www.internetofproduction.org](https://www.internetofproduction.org)

---

## The Problem: Tragedy of the Commons

**What is it?** When individuals acting in self-interest deplete shared resources, harming the collective good.

**Real-World Examples:**
- **Overfishing**: Grand Banks cod fishery collapse decimated fishing communities
- **Climate Change**: Atmospheric carbon as an unmanaged commons
- **Groundwater**: California Central Valley aquifer depletion
- **Ocean Pollution**: Pacific garbage gyres from inadequate waste management
- **Traffic Congestion**: Individual driving choices creating gridlock
- **Digital Commons**: Bandwidth and server resources in shared systems

---

## The Solution: Automated Governance

PromiseGrid addresses tragedy of the commons through:

**Automated Resource Management** — Algorithms ensure efficient, equitable resource use without manual intervention.

**Community-Driven Algorithms** — Resource rules set by community consensus, not centralized authorities.

**Computational Governance** — Computer-based systems integrate decision-making directly into infrastructure.

**Transparent Operations** — All participants can verify fair resource allocation.

---

## User-Visible Features

**Decentralized Ownership** — Users own and control their own nodes; no centralized hosting costs.

**Universal Access** — Run applications from browser tabs, phones, Raspberry Pi devices, or command lines.

**Application Flexibility** — Load apps from URLs, execute locally, compose from mixed-language modules.

**User Control** — Choose software versions independently; no forced upgrades.

**Always Available** — Grid remains operational even when parts go offline.

**Fine-Grained Security** — Capability-based access control for precise permissions.

---

## Admin-Visible Features

**Scalable Infrastructure** — Grid grows organically as nodes join; no capacity planning required.

**Zero Central Costs** — No centralized hosting infrastructure to maintain.

**Simplified Deployment** — Applications deploy via content addressing; no complex orchestration.

**Self-Healing Systems** — Automatic recovery from node failures through replication.

**Policy as Code** — Infrastructure management follows DevOps principles.

**Resource Metering** — Built-in tracking prevents resource exhaustion attacks.

**Observability** — Comprehensive monitoring and debugging across distributed nodes.

---

## Organizational Features

**Cost Reduction** — Eliminate centralized hosting expenses and infrastructure overhead.

**Improved Resilience** — Distributed architecture prevents single points of failure.

**Governance Integration** — Same consensus mechanisms manage both infrastructure and organizational decisions.

**Collaborative Tools** — Built-in support for shared editing, communication, and project management.

**Vendor Independence** — Open protocols prevent lock-in to proprietary platforms.

**Audit Trail** — Content addressing provides cryptographic verification of all operations.

**Community Alignment** — Technology structure mirrors organizational values of decentralization.

---

## Message Structure Overview

PromiseGrid messages use a **5-element CBOR array** providing self-identification, routing, isolation, payload transport, and cryptographic verification.

.diagram[
```
[
  "grid",              // Protocol tag
  protocol_cid,        // Handler routing  
  grid_cid,            // Instance isolation
  cwt_payload,         // Payload data
  signature            // Cryptographic proof
]
```
]

**CBOR** = Concise Binary Object Representation (RFC 8949)
**CID** = Content Identifier (based on multihash)
**CWT** = CBOR Web Token (RFC 8392)

---

## Element 1: Protocol Tag

.diagram[
```
[ "grid", ... ]
   ↑
   Self-identifying protocol marker
```
]

**Purpose:** Immediate protocol recognition

**Implementation:** UTF-8 text string "grid" (5 bytes with CBOR header)

**Benefits:**
- Fast message validation before parsing
- Human-readable in debugging
- Network analysis tool recognition
- Future protocol variant support ("grid2", "grid-alt")

---

## Element 2: Protocol CID

.diagram[
```
[ "grid", protocol_cid, ... ]
           ↑
           Hash of protocol handler code
```
]

**Purpose:** Routes message to correct protocol handler

**Key Properties:**
- Content-addressed protocol implementation
- Exact version specification (not just version numbers)
- Automatic code distribution and caching
- Multiple protocol versions coexist

**Example:** SHA2-256 CID ≈ 34-36 bytes

---

## Element 3: Grid Instance CID

.diagram[
```
[ "grid", protocol_cid, grid_cid, ... ]
                         ↑
                         Network namespace identifier
```
]

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

.diagram[
```
[ "grid", protocol_cid, grid_cid, cwt_payload, ... ]
                                   ↑
                                   CBOR Web Token claims
```
]

**Purpose:** Application-specific message content

**Standard Claims:**
- Issuer (iss), Subject (sub), Audience (aud)
- Expiration (exp), Not-before (nbf), Issued-at (iat)
- CWT identifier (cti)

**Custom Claims:** Application-defined via IANA registry

**Proof-of-Possession:** Confirmation claim (cnf) with COSE keys

---

## Element 5: Signature

.diagram[
```
[ "grid", protocol_cid, grid_cid, cwt_payload, signature ]
                                                ↑
                                                COSE signature structure
```
]

**Purpose:** Cryptographic authenticity and integrity

**COSE Signature Contains:**
- Protected headers (algorithm, countersignatures)
- Unprotected headers (key identifier, certificates)
- Signature value over all preceding elements

**Algorithms:** ECDSA, EdDSA, RSA-PSS variants

**Security:** End-to-end protection over untrusted networks

---

## Message Security Properties

**Multi-Layer Defense:**

**Protocol/Instance Isolation** — Tags and CIDs prevent context confusion

**Version Verification** — Protocol CID ensures compatible implementations

**Authentication** — CWT claims prove sender identity

**Integrity** — COSE signatures detect tampering

**Replay Prevention** — Timestamp and nonce claims

**Proof-of-Possession** — Confirmation claims prove key ownership

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

## Architecture Overview

.diagram[
```
        Application Layer
        -----------------
        Pure Functions
        (Protocol Handlers)
               |
        Kernel Layer
        ---------------
        Message Routing
        Capability Management
        Cache Coordination
               |
        Network Layer
        --------------
        TCP/UDP/QUIC/WebSocket
        Content Distribution
```
]

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

## Implementation Patterns

**Content-Addressable Code** — Every function has a unique CID derived from its content

**Pure Functions** — Protocol handlers observe referential transparency

**Nested Messages** — Messages can contain messages (higher-order functions)

**Distributed Cache** — Nodes cache protocol handlers and responses

**Automatic Code Distribution** — Unknown CIDs trigger content fetch

**Graceful Degradation** — Legacy protocol support via CID versioning

---

## Use Cases

**Community Systems** — Collaborative editing, chat, video, event scheduling

**DevOps Automation** — Container orchestration, configuration management, infrastructure as code

**Enterprise Applications** — ERP, inventory, project management, accounting

**Version Control** — Distributed repositories with issue tracking and code review

**Governance** — Organizational decision-making using consensus mechanisms

**Resource Management** — Preventing tragedy of commons in shared systems

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
]
