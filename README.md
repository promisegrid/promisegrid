# PromiseGrid

PromiseGrid is an experimental decentralized computing and coordination
architecture for groups, communities, and organizations that need shared
software, shared infrastructure, and shared decisions without handing
control to one operator.

It starts from a common failure mode in collaborative systems: shared
resources, authority, and software execution become fragile when they
depend on central operators or unclear promises between participants.

PromiseGrid models capabilities, protocol actions, storage, and
governance as promises that agents can make, keep, delegate, revoke,
replay, and audit.  The long-term goal is a decentralized computer:
computation, communication, and governance run by the participants who
rely on them.  Current executable work is still prototype work, as
summarized below.

This document is a high-level overview of the project.

The current public work in progress is split across multiple repos,
including:
- [wire-lab](https://github.com/promisegrid/wire-lab), for protocol,
kernel, runtime, storage, and application experiments;
- [grid-examples](https://github.com/ciwg/grid-examples), for early
prototypes of runnable example applications.

## Goals

While still experimental, PromiseGrid is intended to support the
following long-term philosophy and goals:

*PromiseGrid is to computation what the Internet is to communication --
the Internet is a decentralized network; PromiseGrid is a
decentralized computer.  Owned and operated by its users rather than
any single legal entity, the grid is scalable and resilient, growing
as users join.*

*Consensus and governance mechanisms are implemented as basic system
features and exposed to higher-level code, applications, and ultimately
the users. This means that the same mechanisms that govern the grid
can also be used to govern the organizations and communities that use
it.*

*PromiseGrid implements computational governance by integrating
computer-based systems and algorithms to automate and facilitate
decision-making and governance. This mechanism makes governance more
efficient, transparent, and participatory within organizations and
broader society.*

## Current status

- **Implemented now:** this repo is the public overview and planning
  trail.  It contains this README, TODO/DI tracking, the intro slide
  material, and small repo tools.
- **Active prototype work:** current executable work is happening in
  [wire-lab](https://github.com/promisegrid/wire-lab) and
  [grid-examples](https://github.com/ciwg/grid-examples).

## How to explore

Start here for the high-level model.  Then read
[wire-lab](https://github.com/promisegrid/wire-lab) for the current
protocol and runtime design trail, inspect or run
[grid-examples](https://github.com/ciwg/grid-examples) for example
applications, and use
[promisegrid-dev-guide](https://github.com/ciwg/promisegrid-dev-guide)
for guide material as it becomes available.

This repo is primarily an overview repo.  It does not provide the
main runnable demo path; protocol details live in `wire-lab` and in
example-local docs under `grid-examples`.

## Why this matters

[Tragedy of the commons](https://en.wikipedia.org/wiki/Tragedy_of_the_commons) 
is a fundamental problem in corporate, community, and national
governance, and it is a problem that PromiseGrid is designed to
address.

Tragedy of the commons (ToC) is a situation in a shared-resource
system where individual users, acting independently according to their
own self-interest, behave contrary to the common good by depleting or
spoiling that resource through their collective action.  The resource
in question may be the well-being of the organization or community as
a whole, or it may be a shared physical, natural, or digital resource.

PromiseGrid addresses ToC by automating resource management and governance, ensuring efficient and equitable use. It leverages algorithms set by its community for this autonomous operation.

## Feature Goals

- Users own their own nodes.
- The grid grows as nodes join.
- A node can be as small as a browser tab, a phone, or a Raspberry Pi. 
- Organizations and communities join the grid without centralized
  hosting costs.
- Applications can be loaded from a URL and executed in a browser tab.
- Applications can be executed from a command line.
- Users control the choice and version of the software they run.
- Applications can be composed from multiple modules in mixed
  languages.
- The grid remains operational even if parts go offline.
- The grid uses capability-based security for fine-grained access
  control.
- The same consensus, conflict resolution, and governance mechanisms
  that manage the grid are available to hosted
  applications and users.

## Technology 

### Prior art

- [WebAssembly](https://webassembly.org/) (WASM)
- [WebAssembly System Interface](https://wasi.dev/) (WASI)
- [Large language models](https://en.wikipedia.org/wiki/Large_language_model) (LLMs)
- [Infrastructure as Code](https://en.wikipedia.org/wiki/Infrastructure_as_Code) (IoC)
- [Container Orchestration](https://en.wikipedia.org/wiki/Container_orchestration)
- [Capability-based security](https://en.wikipedia.org/wiki/Capability-based_security)
- [Promise Theory](https://en.wikipedia.org/wiki/Promise_theory)
- [Content-addressable storage](https://en.wikipedia.org/wiki/Content-addressable_storage)

### PromiseGrid-specific concepts

- [Content-addressable code](#content-addressable-code)
- [Capability-as-Promise Model](#capability-as-promise-model)
- [pCID-selected grid messages](#pcid-selected-grid-messages)
- [Merge-as-consensus Model](#merge-as-consensus-model)
- [Sparse CAS and local trust](#sparse-cas-and-local-trust)


## Foundations and Roadmap

For the items below, checked means done, "." means in progress, "e"
means in progress by an external third party.

- [x] Enabling technologies
    - [x] Lessons learned from centralized computing 
        - [x] Mainframe (1960s)
        - [x] Networked personal computers (1990s)
        - [x] Infrastructure as Code (1990s)
        - [x] Grid computing (2000s)
        - [x] DevOps (2000s)
    - [x] Strong open-source cryptography (1990s)
    - [x] Virtualization (2000s)
    - [x] Containerization (2010s)
    - [x] WebAssembly (WASM) (2020s)
    - [e] WebAssembly System Interface (WASI) (2020s)
    - [x] Large language models (2020s)
- [x] Write LLM-based bootstrap tooling
- [x] Write POC and example code
  - [grid-poc](https://github.com/promisegrid/grid-poc)
  - [wire-lab](https://github.com/promisegrid/wire-lab)
  - [promisegrid-dev-guide](https://github.com/ciwg/promisegrid-dev-guide)
  - [grid-examples](https://github.com/ciwg/grid-examples)
  - ...
- [x] Write grid protocol specification
- [x] Register 'grid' CBOR tag
- [.] Write draft RFC
- [.] Dogfood grid apps internally in a manufacturing environment
- [.] Develop/port example applications to grid
    - [ ] Development tools
        - [.] multi-user coding environment
        - [ ] Native LLM/ML 
        - [.] Grid protocol analyzer
        - [ ] ...
    - [ ] DevOps tools to manage underlying infrastructure
        - [.] Container orchestration
        - [.] Disk image management
        - [.] Configuration management
    - [ ] Community systems
        - [.] Collaborative text editor
        - [ ] Chat
        - [ ] Video conferencing
        - [ ] Event scheduling/management
        - [ ] Virtual world
        - [ ] Membership management
    - [.] Enterprise Resource Planning (ERP)
        - [ ] Inventory management
        - [ ] Project and task management
        - [ ] Bill of materials
        - [.] Production scheduling
        - [.] Accounting
        - [ ] Quoting
        - [.] Order entry
        - [.] Shipping
        - [.] Invoicing
        - [ ] Payroll
            - [ ] Time tracking 
        - [ ] Facilities management
        - [.] IoT device management
        - [ ] ...
    - [.] Version control
        - [x] basic VCS features
        - [.] issue tracker
        - [ ] wiki
        - [ ] code review
        - [ ] CI/CD
        - [ ] code search
        - [ ] execute from repo
        - [ ] large blobs
        - [.] git import/export
    - [ ] ...
- [ ] Self-hosting
    - [ ] Migrate remaining CSWG systems to PromiseGrid
    - [ ] Migrate remaining repositories from github to PromiseGrid

## Architecture

### Architectural Goals

*The core PromiseGrid code operates as a set of decentralized kernel
roles or modules and presents syscall-like message-based services to
applications.  It acts as a "sandbox orchestrator", regardless of the
sandbox technology employed; the grid supports container, VM, WASM, or
bare metal environments.*

*For WASM, for example, the grid takes advantage of the WebAssembly
virtual machine now in all major browsers, offering services to WASM
modules. The kernel also runs natively on server nodes and can be used
to run applications from the command line or as daemons.*

*The grid can execute containerized applications, either within WASI as
a machine emulator (e.g.,
[container2wasm](https://github.com/ktock/container2wasm)) or natively
as a container orchestrator similar to Kubernetes or Docker Swarm.*

*For bare-metal applications, the grid can manage the deployment of
disk images and configuration files to servers, and can manage the
execution of applications on those servers -- a bare-metal server is
just another "sandbox".  This allows the grid to be used for
configuration management, following DevOps and Infrastructure as Code
(IoC) principles.*

### Content-addressable Code

The grid uses content-addressable storage for code and data.  Both
code and data are addressed by their content, not by location or name.
This allows the grid to store and execute code and access data from
any node in the network, and to replicate code and data across the
network as needed.  This is a key feature of the grid's scalability
and resilience.

To future-proof the grid, the kernel does not hardcode the hash size
or hash function.  PromiseGrid utilizes the Multihash format specified by the
[Multiformats project](https://multiformats.io/multihash/). 

Because a grid address is a cryptographic hash of the content, every
piece of code, whether a large binary or a small library function, has
a unique address that can be used to reference it unambiguously from
other code.

Using a 256 or 512 bit hash creates an address space that is
large enough to uniquely address 
[every atom in the observable universe](https://en.wikipedia.org/wiki/Large_numbers).
This means that the code and data storage of the grid is
effectively unlimited for practical purposes; every piece of code ever
written could in theory be stored and addressed in the grid.

### Capability-as-Promise Model

The PromiseGrid kernel uses [capability-based security](https://en.wikipedia.org/wiki/Capability-based_security) 
for access control.  We also consider capabilities to be a form of
promise, both in the sense of javascript-style async promises and in
the sense of promises as described in [Promise
Theory](https://en.wikipedia.org/wiki/Promise_theory) by Mark Burgess.

In our model, a capability token represents a promise that the issuer
will either break or fulfill at a later time.  Revocation of a
capability token is a form of breaking the promise.  The holder can
present the token to the issuer as a request to fulfill the promise --
the response will be a fulfillment, a revocation, or a further promise
to fulfill at a later time.  

This recognition that capabilities are like promises is key to the
design integrity of a decentralized system.  As in Promise Theory,
these principles hold:

1. A resource cannot issue capability tokens promising access to
   another resource.
2. In a system composed of autonomous parts, a resource sending
   unsolicited directives to another, imposing obligations, should
   expect poor results.

These principles are consistent with behavior and dysfunction observed
in human organizations and distributed systems.

While we expect some evolution of this model as we develop the grid,
as of this writing it appears that one or both of the following models
will be in use, depending on the protocol selected by the pCID of the
messages used:

#### Capability-as-Closure Model

- In protocols that use the "Capability as Closure" model, an issuer creates a capability token by creating a
  [closure](https://en.wikipedia.org/wiki/Closure_(computer_programming))
  containing a function that will fulfill the promise, and then
  hashing and signing the closure's code.
- A holder calls the closure (sends a message to the closure's
  hash address).  When called, the closure can elect to fulfill the
  promise, revoke the capability token, or issue a further promise.
- Closures can be nested; the choice of whether to revoke or fulfill
  the promise could be handled by a security screening closure, for
  instance, before forwarding the request to the original issuer. This
  could help simplify issuer code.
    - In order to avoid breaking principle (1) above, the issuer would
      send the original promise to the security screener, so the
      security screener could wrap it in a further promise. This
      avoids the case where an issuer would otherwise be making a
      promise for the security screener to fulfill.
- The receiver could maintain a graph of promises and their fulfillment,
  and could use this graph to manage access control and to facilitate
  consensus formation and conflict resolution among participants.

#### Capability-as-Promise Model

Another model, used by the current public POCs, treats signed
capability tokens as issuer promises carried and redeemed under
pCID-selected protocol rules:

- An issuer signs a capability token describing what it promises to do,
  under what terms, and for whom.
- A holder presents that token back to the issuer, or to a protocol
  participant the issuer has explicitly named.
- Expiry, replay rejection, revocation, and trust updates are local
  consequences of the promises each agent chose to rely on.
- Issuers might create bearer tokens or identity-bound tokens.  Bearer
  tokens might be redeemable for more-specific identity-bound tokens.
- The exact token shape is protocol-specific.  Recent POCs use
  CWT/COSE-shaped objects, for example.

### pCID-selected grid messages

The low-level protocol of the grid, both on the wire and within the
kernel, is designed for extensibility.

A grid message is a compact CBOR envelope:

```text
1735551332([42(pCID), ...protocol-defined-slots])
```

The registered specification for the outer `grid` CBOR tag uses
`1735551332` (`0x67726964`), whose tag-number bytes spell `grid` in
ASCII.  See the [grid CBOR tag
specification](docs/grid-cbor-tag-spec.md) for the IANA registration
document and the `grid(...)` diagnostic alias.

Diagnostic notation may render the same value as
`grid([42(pCID), ...protocol-defined-slots])`.  Slot 0 carries the
pCID wrapped in CBOR tag 42, the IPLD CID tag.  The outer `grid` tag
identifies the envelope; the inner IPLD CID tag preserves the standard
CID representation for the protocol spec.

The pCID selects the parser and owns the following slots: payload
shape, proof and signature semantics, parent or reference links, nested
objects, and any protocol-specific meaning.

The outer `grid` tag does not make a message trusted or executable by
itself.  Trust, authorization, replay handling, signatures, proofs, and
resource limits are defined by the pCID-selected protocol and enforced
by local policy.

The pCID is not a peer address, app address, message type, route, or
operation code.  Those meanings belong inside the pCID-defined payload
or inside nested protocol objects.  This keeps the outer envelope small
and lets protocols evolve without requiring central registration.

A response to a message is another message.  Messages can contain or
refer to other messages, so kernels and higher-level protocols can be
composed without giving up exact byte-level protocol identity.

### Merge-as-Consensus Model

Merges are a form of consensus formation.  A merge is a function that
takes two or more versions of a set of data, producing a new version
that incorporates the changes from all of the input versions.  In a
version control system, a [merge](https://git-scm.com/docs/git-merge)
is part of the process of reaching agreement or consensus about the
contents of a file or set of files.

PromiseGrid protocols can use a similar process for continual
consensus formation.  How this should be done is protocol-specific,
and can involve event streams, branch heads, or object DAGs.  The
current POCs have so far evolved a VCS-like CAS model, where events
(commits) on timelines (branches) can be linked for reference, or
entire timelines merged. A cross-reference to another event or object
can be evidence without requiring an immediate or full merge.

This model works best when application history is represented as
immutable events or parent-linked object DAGs.  A merge can then replay
the same content-addressed inputs under the same protocol rules before
deciding where human judgment is needed.

Merge conflicts occur when a merge function cannot fulfill its promise
to produce a new version that incorporates the changes from all of the
input versions.  This is a form of consensus failure.  We expect
applications to use heuristics that include a cascade of merge
functions, ultimately falling back to LLMs and finally to human
intervention if necessary.

We expect this same model to be applicable at very low levels, e.g.
resolving race conditions caused by concurrent writes to a single
resource, and at very high levels, e.g. forming consensus among
participants in an organization or community.

### Sparse CAS and local trust

The storage model is a sparse content-addressed store on each
node, with CIDs as the CAS keys.  As of this writing, it appears that
the most likely arrangement will be that most CAS objects link to earlier CAS objects, forming a DAG of timelines, likely per-agent.

This is similar in spirit to git's object and branch model, but more
generalized for arbitrary application objects and protocols.  We also
expect that most objects will chunked before storage using a
content-defined chunking algorithm in order to better support large
files and streaming data.

Mutable or rebuildable local data is separate from immutable CAS source
objects.  Local indexes, replication state, cached query results, and
application-level projections may summarize, cache, or reconcile CAS
objects, but they do not change the content identified by a CID.

In at least some cases, a given pCID may specify that a protocol
participant must preserve referential transparency, making event
replay practical.  A node can rebuild application state by replaying
signed envelopes, append-only event logs, or parent-linked object DAGs
from CAS.

Each agent keeps the objects it chooses to keep or has promised to
retain: code, data, signed envelopes, parent-linked message DAGs,
capability tokens, CAR payloads, and application objects.  Other agents
may promise to store, serve, or transfer those objects.

## Contributing

Start with [wire-lab](https://github.com/promisegrid/wire-lab) for the
current design trail,
[grid-examples](https://github.com/ciwg/grid-examples) for a few
runnable example applications, and
[promisegrid-dev-guide](https://github.com/ciwg/promisegrid-dev-guide)
for where we're building guide material.  If you're interested in
closer collaboration, watch this repo or contact any of the
developers in the [ciwg
repositories](https://github.com/orgs/ciwg/repositories) on github.

## Sponsorship

PromiseGrid development is supported in part by
[C D International Technology, Inc.](http://www.cdint.com/) and 
[TerraLuna, LLC](http://www.t7a.org).
