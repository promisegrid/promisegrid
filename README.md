# PromiseGrid

PromiseGrid is a consensus-based computing, communications, and
governance system.  It's designed specifically to address the problems
of collaborative work and leadership that plague organizations and
communities, including non-profits, corporations, and governments.

PromiseGrid is to computation what the Internet is to communication --
the Internet is a decentralized network; PromiseGrid is a
decentralized computer.  Owned and operated by its users rather than
any single legal entity, the grid is scalable and resilient, growing
as users join. 

Consensus and governance mechanisms are implemented as basic system
calls and exposed to higher-level code, applications, and ultimately
the users. This means that the same mechanisms that govern the grid
can also be used to govern the organizations and communities that use
it.

This document is a high-level overview of the project.

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

PromiseGrid addresses ToC by providing an equitable platform for
collaboration and shared resource management. The grid's design is
based on several principles from [Promise Theory](https://en.wikipedia.org/wiki/Promise_theory) by Mark Burgess.
These principles provide a foundation for a decentralized system that
can not only govern itself but also provide governance as a service to
its hosted organizations and communities.

Decentralization, if designed in a way that explicitly prevents ToC,
brings several intrinsic benefits that are essential for the
development of resilient, adaptive, and equitable systems:

1. **Resiliency**: Decentralized systems distribute their operations
   across multiple nodes or entities. This dispersion increases the
   system's ability to withstand attacks, failures, and crises,
   ensuring continuous operation under various conditions.

2. **Crisis Preparedness and Response**: In the event of a crisis,
   decentralized systems can adapt and respond more quickly than
   centralized ones. Each node or entity can react based on local
   conditions and needs, enabling a more flexible and effective
   response to emergencies.

3. **Alignment with Self-Interest**: Decentralization acknowledges the
   reality that individuals and organizations act in their own
   self-interest. By designing systems that expect and accommodate
   this self-interest, we can create healthier and more robust
   systems. It avoids the pitfalls of over-reliance on either altruism
   or centralized control.

PromiseGrid implements "computational governance":  This term refers
to the use of computer-based systems and algorithms to facilitate and
automate decision-making processes and governance structures. This
enables more efficient, transparent, and participatory forms of
governance within organizations, networks, and broader society.
Through computational governance, tasks such as consensus finding,
resource allocation, and policy enforcement can be implemented in a
decentralized and programmatic fashion, reducing the need for
bureaucracy while improving the scalability and adaptability of
governance mechanisms.

## Features

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
- [PromiseGrid Universal Protocol](#promisegrid-universal-protocol)
- [Merge-as-consensus Model](#merge-as-consensus-model)
- [Decentralized cache](#decentralized-cache)

## Milestones

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
    - [ ] WebAssembly System Interface (WASI) (2020s)
    - [x] Large language models (2020s)
- [ ] Convene/contact 
    - [x]  Community Systems Working Group [CSWG](http://cswg.infrastructures.org/)
    - [x]  Internet of Production [IoP](http://www.internetofproduction.org/)
    - [ ]  [Worknetting](https://en.wikipedia.org/wiki/Worknet) community
    - [ ]  [Mark Burgess](http://markburgess.org/)
    - [ ] ...
- [x] Write LLM-based bootstrap tooling
    - [x]  [Grokker](http://github.com/stevegt/grokker)
- [ ] Write WASI target
- [ ] Develop/port example applications to WASI target
    - [ ] Development tools
        - [ ] Grokker 
        - [ ] Multi-agent production
        - [ ] Native LLM/ML 
        - [ ] Grid protocol analyzer
        - [ ] WASI kernel debugger
        - [ ] ...
    - [ ] DevOps tools to manage underlying infrastructure
        - [ ] Container orchestration
        - [ ] Disk image management
        - [ ] Configuration management 
    - [ ] Community systems
        - [ ] Collaborative text editor
        - [ ] Chat
        - [ ] Video conferencing
        - [ ] Event scheduling/management
        - [ ] Virtual world
        - [ ] Membership management
    - [ ] Enterprise Resource Planning (ERP)
        - [ ] Inventory management
        - [ ] Project and task management
        - [ ] Bill of materials
        - [ ] Production scheduling
        - [ ] Accounting 
        - [ ] Quoting
        - [ ] Order entry
        - [ ] Shipping
        - [ ] Invoicing
        - [ ] Payroll
            - [ ] Time tracking 
        - [ ] Facilities management
            - [ ] IoT device management
        - [ ] ...
    - [ ] Version control
        - [ ] issue tracker
        - [ ] wiki
        - [ ] code review
        - [ ] CI/CD
        - [ ] code search
        - [ ] execute from repo
        - [ ] large blobs
        - [ ] tool to import from `git fast-export` 
    - [ ] ...
- [ ] Self-hosting
    - [ ] Migrate remaining community systems to PromiseGrid
    - [ ] Migrate remaining repositories from github to PromiseGrid


## Architecture

The core PromiseGrid code operates as a decentralized kernel and
presents syscall-like services to applications.  It acts as a "sandbox
orchestrator", regardless of the sandbox technology employed; the grid supports
container, VM, WASM, or bare metal environments.

For WASM, for example, the grid takes advantage of the WebAssembly
virtual machine now in all major browsers, offering services to WASM
modules. The kernel also runs natively on server nodes and can be used
to run applications from the command line or as daemons.

The grid can execute containerized applications, either within WASI as
a machine emulator (e.g.,
[container2wasm](https://github.com/ktock/container2wasm)) or natively
as a container orchestrator similar to Kubernetes or Docker Swarm.

For bare-metal applications, the grid can manage the deployment of
disk images and configuration files to servers, and can manage the
execution of applications on those servers -- a bare-metal server is
just another "sandbox".  This allows the grid to be used for
configuration management, following DevOps and Infrastructure as Code
(IoC) principles.

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
as of this writing it appears that:

- An issuer creates a capability token by creating a
  [closure](https://en.wikipedia.org/wiki/Closure_(computer_programming))
  containing a function that will fulfill the promise, and then
  hashing the closure's code.  The hash is both the capability token
  and the address of the closure.
- A holder calls the closure (sends a message to the closure's
  address).  When called, the closure can elect to fulfill the
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
- The kernel could maintain a graph of promises and their fulfillment,
  and could use this graph to manage access control and to facilitate
  consensus formation and conflict resolution among participants, or
  these functions could be delegated to hosted modules or to the
  original issuer.
- Likewise, reputation -- how well a participant fulfills promises --
  could be tracked by the kernel or delegated.

Which of these models is in use at any given time is a matter of
protocol, and could be determined in a given session by the
counterparties.

### PromiseGrid Universal Protocol

The low-level protocol of the grid, both on the wire and within the
kernel, is designed for extensibility.

A function call is a message.

A message consists of a capability token followed by a payload.
Because a token is the hash of the function that will fulfill the
promise, a message always starts with the hash of the function that
will fulfill the promise.  

In other words, a message is a function call.  The message payload is
the arguments to that function.  

A response to a message is another message.

A message can contain one or more messages in its payload.  This
is roughly analogous to passing a function as an argument in a
functional programming language.

There are no version number or other metadata fields in a message
header before the token.  The token is the address of the function,
and doubles as a protocol version hash.  

Because messages can be nested, kernels can be nested.  This is what
allows the grid itself to be extensible and to evolve -- a hyperkernel
can route messages to other kernels, and so on.

### Merge-as-Consensus Model

The grid's consensus formation and conflict resolution mechanisms are
based on a model semantically similar to the
[merge](https://git-scm.com/docs/git-merge) function of a version
control system.

Merges are a form of consensus formation.  A merge is a function that
takes two or more versions of a document and produces a new version
that incorporates the changes from all of the input versions.  How
this should be done is application-specific, and we expect that the
grid will support multiple merge functions for different types of
data, each as content-addressable code.

Merge conflicts occur when a merge function cannot fulfill its promise
to produce a new version that incorporates the changes from all of the
input versions.  This is a form of consensus failure.  We expect
applications to use heuristics that include a cascade of merge
functions, ultimately falling back to LLMs and finally to human
intervention if necessary.

We expect this same model to be applicable at very low levels, e.g.
resolving race conditions caused by concurrent writes to a single
resource, and at very high levels, e.g. resolving disputes between
participants in an organization or community.

### Decentralized Cache

The grid includes a decentralized cache that is used to store and
replicate code and data across the network.  The cache is a
content-addressable store on each node.  

Grid functions are pure functions -- they observe [referential
transparency](https://en.wikipedia.org/wiki/Referential_transparency),
always producing the same output from the same inputs. 

The cache is a directed graph database.  Vertices are typically either
hashes or arguments -- a message, when received, can be decomposed,
with the leading hash (the capability token, function address) used as
a vertex value, the first argument used as a destination vertex, and
so on, eventually leading to the components of the response message.

This design allows the cache to store responses to function calls,
even when the calling messages or responses are nested.

When a node's kernel receives a message, it first checks its cache for
the token and arguments.  If the call is in the cache, the kernel
replies with the cached response.  If the message is not in the cache,
the kernel encapsulates and forwards the message to any other node it
has capability tokens for.  When the response is received, the kernel
stores the response in the cache and forwards it to the original
sender.

## Contributing

We welcome collaborators.  It's still early days; the core team is
currently meeting as a group weekly as well as 1:1 in between, and
we'll soon be standing up community systems (email, chat, etc.).  In
the meantime, if you're interested in joining a mind-bending open
source project that could change the world, watch this repo and
contact Steve Traugott on
[github](https://www.github.com/stevegt), 
[twitter](https://www.twitter.com/stevegt), or
[linkedin](https://www.linkedin.com/in/stevegt/).

## Sponsorship

PromiseGrid development is supported by 
[C D International Technology, Inc.](http://www.cdint.com/) and 
[TerraLuna, LLC](http://www.t7a.org).

