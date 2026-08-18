# TE-fujur: Grid-CBOR And DAG-CBOR Superset Boundaries

TE ID: TE-fujur

Date: 2026-08-16

Decision status: needs DF under [DR-tuzob](../../DR/DR-tuzob-grid-cbor-dag-cbor-superset.md)

Related work: [TODO-gilog, tasks gilog.6 through gilog.9](../../TODO/TODO-gilog-grid-cbor-multicodec.md)

## Decision Under Test

Which Grid-CBOR encoding rules and decoded representation can simultaneously:

1. accept every valid DAG-CBOR item unchanged when that item is placed in a
   protocol-defined Grid slot;
2. preserve the exact DAG-CBOR bytes through a Grid-CBOR codec round trip;
3. permit Grid-CBOR values outside the DAG-CBOR data model, including arbitrary
   CBOR tags and map-key types;
4. retain the Grid envelope and pCID framing already specified for PromiseGrid;
   and
5. provide explicit behavior for deterministic writing, generic CID-link
   discovery, malformed input, mixed versions, and long-term recovery?

This TE tests alternatives. It does not select one or supersede an existing
Decision Intent record. The resulting choice remains open in DR-tuzob.

## Scope

This TE covers the byte-level Grid-CBOR profile, the relationship between
Grid-CBOR and DAG-CBOR, and the generic decoded representation used before a
pCID-specific handler interprets protocol-owned slots.

It affects:

- the public Grid-CBOR multicodec specification;
- generic Grid envelope codecs and validators;
- IPLD link discovery over Grid-CBOR blocks;
- protocol payloads that are already valid DAG-CBOR;
- future conformance fixtures; and
- the generic envelope work that precedes the identity-genesis A02 product-code
  release barrier.

It does not choose an identity protocol, identity-event slot count, identity
authorization rule, package API, production pCID constant, signature scheme, or
handler registration design. Authentication remains distinct from application
authorization.

## Evidence And Decision Status

The following sources are inputs, not conclusions:

- The registered [Grid CBOR tag specification](../grid-cbor-tag-spec.md) fixes
  the outer `grid` tag, the tagged array, and tag-42 pCID in slot 0. It delegates
  every later slot's meaning and format to the pCID.
- [DI-vogod](../../TODO/TODO-gilog-grid-cbor-multicodec.md) currently requires
  RFC 8949 Core Deterministic Encoding for complete Grid envelopes, arbitrary
  tags and map keys, valid visible tag-42 CIDs, and generic processing without
  application authorization. It is a locked historical/current decision, not a
  premise that this TE may silently edit.
- [DI-nipap](../../TODO/TODO-gilog-grid-cbor-multicodec.md) keeps internal
  coordination references out of the public specification. It affects how a
  later public edit cites provenance, not the encoding result.
- The current [public Grid-CBOR multicodec specification](../grid-cbor-multicodec-spec.md)
  requires shortest floating-point encodings and exact-byte round trips. It is
  an implementation target derived from the DIs, and its rules are tested here.
- The earlier cdint-grid `TE-natod` recommended RFC 8949 Core Deterministic
  Encoding after treating DAG-CBOR as a narrower alternative. That conclusion
  did not test unchanged DAG-CBOR as a Grid slot, so it is prior analysis rather
  than an answer to this TE.
- The [DAG-CBOR specification](https://ipld.io/specs/codecs/dag-cbor/spec/)
  defines the established IPLD codec. Its data model restricts maps to string
  keys and semantic tags to tag 42 links. It uses definite lengths, preferred
  integer and length encodings, deterministic map order, and binary64 for every
  float. It rejects NaN, infinities, and negative zero.
- RFC 8949 Section 4.2.1 Core Deterministic Encoding instead requires a float to
  use the shortest width that represents the value exactly. Therefore the two
  profiles differ even for values accepted by both.
- The cdint-grid POCs retain exact raw slots and build disposable IPLD
  projections for traversal. They demonstrate one implementation technique;
  they are experimental evidence, not a production contract.

## Locked Compatibility Premises

The user supplied these premises for this TE. Alternatives are evaluated
against them rather than using the earlier recommendation as the answer.

1. Grid-CBOR permits any CBOR tag.
2. For every valid DAG-CBOR item `D`, it MUST be possible to form a valid
   Grid-CBOR message by placing `42(pCID)` in slot 0, placing the exact bytes of
   `D` in slot 1, wrapping the slots in the `grid` tag, and supplying the
   required Grid array and tag headers.
3. A Grid-CBOR codec round trip MUST accept that message and reproduce the
   original DAG-CBOR bytes in slot 1.
4. Removing the outer `grid` tag from a Grid-CBOR message MAY leave a value that
   another use can treat as DAG-CBOR, but this is not guaranteed. A valid
   Grid-CBOR message may contain non-42 tags, integer or compound map keys, and
   other values outside DAG-CBOR.
5. The relationship is one-way inclusion of DAG-CBOR payloads, not a claim that
   every unwrapped Grid-CBOR value is valid DAG-CBOR.

“Round trip” in premise 3 means exact bytes, not merely an equal application
value. A CID commits to bytes, so replacing a valid DAG-CBOR binary64 float with
a shorter float changes the block and any CID over it.

## Assumptions And Trust Model

Alice is a protocol writer. She already has a valid DAG-CBOR payload and places
it unchanged in slot 1 of a Grid message.

Bob is an independent protocol writer. He constructs the same abstract payload
from values rather than copying Alice's bytes.

Carol is a generic Grid storage and routing agent. She recognizes Grid framing
and visible CID links but may not have the pCID specification.

Dave is a pCID-specific application agent. After generic validation, he
interprets the protocol-owned slots.

Ellen is a recovery operator decades later. She has retained blocks, codec
specifications, and conformance fixtures, but may not have the original
implementation.

Mallory supplies malformed, ambiguous, deeply nested, or resource-exhausting
CBOR and tries to make different implementations derive different values or
links.

The exact CAS block bytes and their CIDs are durable evidence. A decoded object,
IPLD projection, map index, or link index is disposable only when the system
states how it is derived from those bytes, how it detects staleness, and how it
is rebuilt. Recognizing a valid envelope or signature does not authorize Dave's
application action.

## Neutrality Method

This TE uses the following safeguards against selecting the answer in the setup:

1. Only the five compatibility premises above are fixed as new requirements.
2. Existing DIs, specifications, earlier TEs, libraries, and POCs are classified
   as evidence with their actual status; none is treated as a conclusion merely
   because it exists.
3. Alternative names describe mechanics rather than desirability.
4. Fixture categories, scenarios, and questions are fixed before comparing the
   alternatives.
5. Every alternative is evaluated under the same trust model and scenarios.
6. No weighted score is used. A tradeoff is not converted into a rejection
   without identifying the failed premise or the existing requirement that
   would have to change.
7. Each alternative receives an explicit account of what it makes easier, what
   it makes harder, and what obligations it creates.
8. Conclusions distinguish incompatibility with a locked premise from conflict
   with an existing decision and from an unresolved preference.

## Encoding Alternatives

### E1: RFC 8949 Core Deterministic Encoding For The Complete Envelope

Apply the current public Grid-CBOR rule recursively to every item in the
envelope. Integers, lengths, tags, and floats use their RFC 8949 preferred
shortest encodings; maps use bytewise lexical order of encoded keys.

Makes easier:

- reuse of fxamacker's `CoreDetEncOptions` for writers;
- one uniform validation rule over the whole envelope; and
- conformance with the present public Grid-CBOR text.

Makes harder:

- unchanged composition with DAG-CBOR when a binary64 DAG-CBOR float has an
  exact binary16 or binary32 representation; and
- use of existing DAG-CBOR bytes as opaque-by-structure inline slots.

Creates obligations:

- either reject such a valid DAG-CBOR slot or normalize it and change its bytes;
- supersede the new compatibility premise if retained; and
- explain how the resulting CID change is compatible with the claimed round
  trip.

### E2: DAG-CBOR-Derived Deterministic Rules Across The CBOR Domain

Use DAG-CBOR's representation rules wherever they apply, then extend those
rules to additional CBOR tags, map-key types, simple values, and floating-point
values outside DAG-CBOR. In this alternative, deterministic key order extends
to the encoded bytes of every key, and every accepted float has one specified
width. One possible closure uses binary64 for all floats because that is the
DAG-CBOR overlap rule; the treatment of non-finite and negative-zero values
would still need to be stated.

Makes easier:

- a single profile whose DAG-CBOR subset is byte-compatible by construction;
- deterministic construction by independent writers once the extension is
  complete; and
- direct representation of arbitrary Grid tags and key types.

Makes harder:

- reuse of an unchanged RFC 8949 Core Deterministic encoder;
- reuse of an unchanged DAG-CBOR decoder for the extended domain; and
- compact float encoding when binary64 is chosen for all accepted floats.

Creates obligations:

- specify ordering for non-string and compound keys;
- specify non-DAG simple values, NaN payloads, infinities, and negative zero;
- implement a Grid validator rather than relabeling an existing DAG-CBOR codec;
  and
- publish fixtures for both the overlap and extensions.

### E3: DAG-CBOR Overlap Rules Plus Grid Rules Outside The Overlap

Require any recursively represented item that is valid DAG-CBOR to retain its
DAG-CBOR representation. Define separate deterministic Grid rules for values or
structures outside the DAG-CBOR data model. For example, a finite DAG-CBOR float
stays binary64, while an otherwise forbidden non-finite float is governed by a
Grid-specific rule. A map with a non-string key is governed by the Grid key
ordering rule while any DAG-CBOR payload nested as its value remains unchanged.

Makes easier:

- exact inclusion of the DAG-CBOR subset;
- independent optimization of Grid-only representations; and
- use of tests that separate compatibility fixtures from extension fixtures.

Makes harder:

- explaining the boundary between an overlapping item and a Grid-only
  structure;
- streaming encoders that must know which rule governs a container; and
- avoiding subtle context-dependent canonicalization rules.

Creates obligations:

- define the boundary recursively and without circularity;
- specify deterministic encodings for every accepted Grid-only value;
- prove that one abstract Grid value cannot be classified two ways; and
- maintain cross-boundary conformance fixtures.

### E4: Multiple Accepted CBOR Representations With Exact Bytes Retained

Accept more than one valid representation of a CBOR value. The received bytes
remain the authoritative representation and round-trip output uses those bytes.
Writers may have a recommended deterministic profile, but the Grid-CBOR validity
rule itself does not require all writers to produce the same bytes.

Makes easier:

- exact acceptance and reproduction of existing DAG-CBOR and broader CBOR;
- use of an input-preserving parser; and
- migration between writer profiles without rewriting old blocks.

Makes harder:

- deriving a unique CID from only an abstract value;
- comparing semantic equality without also comparing bytes; and
- detecting representational malleability in signatures over decoded values.

Creates obligations:

- make signatures and CAS identity explicitly bind exact envelope bytes or an
  separately specified signing representation;
- retain raw bytes for every authoritative block;
- state whether duplicate keys and other ambiguous representations are valid;
- distinguish byte identity from value equivalence in APIs; and
- decide whether the existing independent-writer convergence requirement is
  being superseded.

### E5: DAG-CBOR Wrapper Values For CBOR Features Outside DAG-CBOR

Represent an arbitrary CBOR tag, non-string map key, or unsupported simple value
as an ordinary DAG-CBOR map or list carrying a type discriminator and content.
The stored wire value remains DAG-CBOR even though it describes a richer CBOR
value.

Makes easier:

- direct reuse of DAG-CBOR codecs and IPLD nodes;
- ordinary IPLD traversal over the stored representation; and
- one deterministic wire profile.

Makes harder:

- carrying an actual arbitrary CBOR tag or key in the Grid-CBOR wire item;
- distinguishing user maps from wrapper maps without reserved schema; and
- exact round trip to an input that used the native CBOR feature.

Creates obligations:

- define an escaping and collision-free wrapper schema;
- change the meaning of “Grid-CBOR permits any CBOR tag” from native acceptance
  to modeled representation; and
- explain why native valid Grid-CBOR inputs may be rewritten.

### E6: pCID-Selected Serialization Of Protocol-Owned Slots

Use the outer Grid framing rule for slot 0, then let the pCID select the encoding
profile for later slots. A DAG-CBOR-owning pCID can require exact DAG-CBOR in
slot 1; another pCID can choose a different CBOR profile.

Makes easier:

- protocol-specific optimization and schema control;
- exact DAG-CBOR use by protocols that request it; and
- staged addition of other serializations.

Makes harder:

- generic validation or reconstruction when the pCID is unavailable;
- determining where one slot ends without a shared structural CBOR parser; and
- generic link discovery before pCID resolution.

Creates obligations:

- retain a common framing grammar capable of locating raw slots;
- retain every pCID specification needed for later reconstruction;
- define behavior for unknown pCIDs and mixed versions; and
- avoid treating protocol parsing as generic authorization.

## Decoded-Representation Alternatives

The encoding decision does not by itself choose how software represents an
accepted block. These alternatives are evaluated separately and can be paired
with compatible encoding alternatives.

### R1: Retained Block And Raw Slots With A Derived IPLD Projection

Keep the complete block and each slot's byte range authoritative. Build a
separate, disposable IPLD node for traversal and visible tag-42 link discovery.

Makes easier:

- exact byte round trips even when a host-language value model is narrower;
- bounded pCID dispatch over raw protocol-owned slots; and
- recovery by rebuilding projections from CAS blocks.

Makes harder:

- APIs because callers must distinguish raw authority from projected values;
- value edits, which require deliberate re-encoding rather than mutating a
  projection; and
- projections of compound keys and unknown tags.

Creates obligations:

- name the projection as non-authoritative;
- define its derivation, version/frontier, stale detection, and rebuild path;
- validate raw bytes before trusting projected links; and
- prevent normalized projection values from being signed or re-encoded as if
  they were the original block.

### R2: Reversible Grid-Specific Node Model Or IPLD ADL

Decode into a node model that directly represents all accepted CBOR distinctions
and offers an IPLD-compatible logical view, using a Grid-specific data model or
an IPLD Advanced Data Layout where needed.

Makes easier:

- traversal and editing through one explicit model;
- reversible access to tags and non-string map keys; and
- typed APIs that do not use unstructured host-language interface values.

Makes harder:

- implementation and interoperability because standard IPLD nodes do not
  directly contain arbitrary tags or non-string map keys;
- proving exact byte preservation when the encoding alternative permits more
  than one representation; and
- adoption by generic IPLD tools that do not know the ADL.

Creates obligations:

- specify the complete logical and representation models;
- define link behavior and ADL loading without pCID authority;
- retain representation choices needed for byte-exact re-encoding; and
- supply cross-language conformance fixtures.

### R3: Generic Host-Language CBOR Values Followed By Deterministic Re-Encoding

Decode into generic maps, lists, numbers, tags, and scalars provided by a CBOR
library, then invoke a deterministic encoder to reproduce bytes.

Makes easier:

- small initial implementations;
- reuse of familiar library APIs; and
- ordinary application access to common CBOR values.

Makes harder:

- preservation of tags that a library assigns built-in semantics;
- compound map keys that are not comparable host-language map keys;
- preservation of integer width, float width, NaN payload, and map entry bytes;
  and
- distinguishing omitted information from intentional normalization.

Creates obligations:

- prove lossless coverage for every accepted Grid-CBOR feature;
- configure or replace built-in tag conversions;
- retain raw representations for distinctions absent from the generic model;
  and
- reject inputs that the chosen value model cannot preserve.

### R4: Retained Raw Block Without A Generic Traversal Projection

Validate the outer Grid shape and retain the exact block. Defer all later-slot
parsing and link discovery to pCID-specific handlers.

Makes easier:

- byte-exact storage and relay;
- a small generic envelope implementation; and
- support for unknown rich CBOR values.

Makes harder:

- recursively finding visible tag-42 links before pCID support exists;
- generic graph indexing and retention analysis; and
- diagnostics for malformed nested structures.

Creates obligations:

- change or supersede the current generic visible-link discovery requirement;
- specify opaque handling for unknown pCIDs; and
- ensure that handler absence never becomes accidental authorization.

## Fixed Fixture Categories

Every encoding and representation alternative is evaluated against the same
fixture categories:

1. DAG-CBOR scalars, lists, and string-key maps without floats.
2. DAG-CBOR finite floats that require binary64 by specification but are exactly
   representable in binary16 or binary32, including `1.5`.
3. Valid tag-42 links and malformed tag-42 values.
4. Unknown and nested non-42 tags.
5. Maps with integer, byte-string, float, tagged, array, map, and mixed key types.
6. NaN, infinities, negative zero, simple values, and full-width integers.
7. Opaque byte strings, including tag-24 content and bytes that resemble tag 42.
8. Unknown pCIDs and operation without network or protocol-spec access.
9. Duplicate keys, misordered maps, overlong arguments, invalid UTF-8,
   indefinite lengths, truncation, trailing items, and excessive depth or item
   counts.
10. Independent writers, mixed codec versions, IPLD traversal, extraction of a
    slot as DAG-CBOR, corruption recovery, migration, rollback, and large
    blocks.

## Scenario Analysis

### Scenario 1: Unchanged DAG-CBOR Composition

Alice has the valid DAG-CBOR encoding of float `1.5`. DAG-CBOR emits the float as
binary64: `fb3ff8000000000000`. She places those bytes in slot 1 after the
tag-42 pCID and adds the deterministic Grid array and outer tag headers.

- E1 rejects the slot as non-CoreDet or re-encodes it as binary16 `f93e00`.
  Either behavior fails locked premises 2 and 3.
- E2 accepts and preserves it because the DAG-CBOR subset uses the DAG-CBOR
  width rule.
- E3 accepts and preserves it because this item is inside the DAG-CBOR overlap.
- E4 accepts and retains it as one valid representation.
- E5 can carry the float as DAG-CBOR, but it cannot accept every native rich-CBOR
  extension unchanged in later fixtures.
- E6 accepts it only when the pCID rule selects DAG-CBOR or the common base rule
  independently guarantees DAG-CBOR acceptance.

R1 and R4 preserve the slot by retaining its byte range. R2 can preserve it if
its representation model records the original encoding or if the selected
encoding has only one valid representation. R3 loses the binary64-width fact
when it stores only the numeric value and then applies CoreDet.

This is a direct counterexample, not a preference: E1 and unqualified R3 cannot
satisfy the locked exact-composition premises as stated.

### Scenario 2: Ordinary DAG-CBOR Map And Link

Alice places a DAG-CBOR map containing a string key and a valid tag-42 CID in
slot 1. Carol validates the envelope without understanding its pCID, finds the
visible link, and extracts slot 1.

E1, E2, E3, and E4 can all accept the non-float fixture. E5 also carries it
without a wrapper. E6 depends on either pCID availability or a shared fallback.
R1 and R2 can expose the link while preserving extraction. R3 works for this
narrow fixture if tag 42 is configured as a link rather than normalized away.
R4 preserves the bytes but does not supply Carol's generic link discovery.

The fixture does not distinguish E1 from E2 or E3; float and extension fixtures
are required to avoid drawing a conclusion from only the common subset.

### Scenario 3: Arbitrary Tags And Map Keys

Alice sends a Grid-CBOR slot containing an unknown tag around a map with an
integer key and an array key. Carol does not know the tag or pCID but must retain
the exact block. No bytes inside an opaque byte string are interpreted as links.

E1, E2, E3, and E4 can natively cover the value if their map-order and tag rules
are complete. E5 changes the native tags and keys into wrappers, so it does not
satisfy premise 1 as native wire acceptance and does not round-trip this input.
E6 can cover it only through an available pCID rule or a shared generic parser.

R1 retains the authoritative bytes and may project non-string maps as pair lists
and unknown tags as wrapper nodes, but that projection is not reversible and
must not replace the block. R2 can be reversible if its representation model is
complete. R3 is not generally lossless: generic decoding may apply built-in tag
semantics, and a compound key may not be usable as a host-language map key. R4
retains the bytes but supplies no generic traversal view.

This scenario rejects E5 and unqualified R3 against the locked rich-CBOR
premise. It leaves R1 and R2 viable with different obligations.

### Scenario 4: Independent Writers And CID Convergence

Alice copies an existing valid representation. Bob starts from the same abstract
value and independently encodes it. Both want to predict whether their blocks
have the same CID.

E1, E2, and E3 each can provide one representation if their rules are complete,
although E1 already fails the DAG-CBOR float premise. E2 has one uniform
extension profile. E3 must prove that its overlap boundary does not yield two
classifications. E4 does not guarantee convergence because both representations
may be valid. E5 converges for its wrapper model but fails native rich-CBOR
acceptance. E6 converges only when both writers retain and implement the same
pCID specification and version.

R1 and R4 do not themselves determine writer output; they preserve whichever
bytes the encoding rule accepts. R2 can support deterministic construction when
paired with E2 or E3. R3 can appear convergent while silently normalizing an
input, which is not sufficient for exact round trip.

Independent-writer convergence is an existing objective in DI-vogod's ancestry,
but the new premises do not alone prove that it must outrank acceptance of
multiple exact representations. E4 therefore conflicts with the existing
deterministic policy rather than failing the five newly locked premises. Keeping
E4 would require a superseding DI and exact-byte signature rules.

### Scenario 5: Unknown pCID And Offline Recovery

Carol and, decades later, Ellen possess a Grid block and the Grid-CBOR codec
specification but cannot resolve its pCID. They need to validate framing,
reproduce bytes, and identify structural tag-42 links without executing the
protocol.

E1, E2, E3, and E4 have pCID-independent base parsing. E5 also parses without a
pCID but changes the accepted value domain. E6 cannot fulfill the whole scenario
unless it includes a pCID-independent base parser and default encoding; once it
does, that base parser is substantively one of E1 through E4.

R1 supplies byte recovery and a rebuildable link projection. R2 can supply both
if its ADL specification and implementation survive. R3 depends on the exact
library behavior and may not reconstruct rich inputs. R4 supplies byte recovery
but not generic link discovery.

E6 is therefore not a complete replacement for a generic base codec under the
current unknown-pCID requirement. A pCID can still narrow allowed slot values
and define their semantics after generic validation.

### Scenario 6: Malformed And Resource-Exhausting Input

Mallory sends duplicate map keys, a misordered deterministic map, an overlong
integer, malformed tag 42, invalid UTF-8, an indefinite container, a truncated
item, trailing items, and excessive nesting.

E1, E2, and E3 can reject nonconforming representations before pCID dispatch,
but E2 and E3 require a complete new validator specification. E4 must separately
decide which alternate encodings are harmless diversity and which are ambiguous
or unsafe. It can still forbid duplicate keys, malformed tag 42, invalid UTF-8,
truncation, trailing items, and resource-limit violations while permitting more
than one non-ambiguous representation. E5 inherits DAG-CBOR's narrow rejection
rules for its wrapper wire form. E6 risks inconsistent validation when a pCID is
missing or implementations support different protocol versions.

R1 should validate the raw bytes before constructing the projection. R2 must
bound representation and ADL construction. R3 must not rely on a generic decoder
that accepts and normalizes a representation the validator was meant to reject.
R4 can enforce structural and resource limits but cannot validate nested link
semantics if later slots remain entirely opaque.

No alternative removes the need for explicit byte, depth, item-count, and
allocation limits. Resource limits are implementation policy and do not grant
application authority.

### Scenario 7: Mixed Versions, Migration, And Rollback

Alice writes with version 1. Bob upgrades to version 2. Carol stores both blocks,
and Ellen later rolls a reader back to version 1.

E1 has broad existing CBOR tooling but cannot accept the locked DAG-CBOR float
fixture. Moving from E1 to E2 or E3 changes the accepted/produced bytes and
requires retaining the old decoder for existing CIDs. E2 provides one new
profile. E3 permits the Grid-only closure to evolve only by minting a new codec
or otherwise preventing old and new rules from claiming the same codec identity.
E4 naturally retains old representations but needs version-independent rules for
which representations are safe. E5 requires a wrapper-schema version. E6 moves
version retention into every pCID and increases the number of specifications
needed for recovery.

R1 can rebuild a current projection from every retained block when the validator
version is known. R2 requires retaining ADL versions. R3 is vulnerable to library
upgrade normalization changes. R4 has the smallest generic migration surface
but loses generic graph behavior.

Rollback must never overwrite or relabel old blocks. New validators can add
support while old codec specifications and fixtures remain addressable.

### Scenario 8: Scale And Streaming

Carol ingests large Grid blocks with deep but bounded structures and many visible
links.

E1, E2, and E3 allow streaming validation, deterministic writing, and link
extraction, although E3 may need container lookahead to classify the overlap.
E4 can stream-copy authoritative bytes and validate in one pass. E5 uses
additional wrapper bytes. E6 may block on pCID lookup unless generic framing is
independent.

R1 can retain the input while building only the projection needed by the caller,
but a full projection duplicates memory. It should allow streaming link events
or a bounded projection. R2 may allocate a complete representation/ADL tree.
R3 creates generic values and can allocate heavily for compound structures. R4
can stream validation and storage with the least decoded state, at the cost of
link discovery.

All choices need declared maximum bytes, nesting, items, and links. A disposable
link index needs a derivation version and block frontier so stale entries can be
detected and rebuilt from CAS.

### Scenario 9: Corruption And Rebuild

One stored block fails its CID hash check, or Carol's derived link index is
partially written.

Every encoding alternative must reject a block whose bytes do not match its CID;
canonicalization cannot repair durable evidence in place. R1 and R2 can rebuild
derived views from verified blocks. R3 cannot recover a lost representation
choice from only normalized values. R4 has no generic index to corrupt, but a
pCID-specific index still needs the same rebuild discipline.

The recovery frontier is the set of CID-verified Grid blocks admitted by the
selected codec version. Derived data records the codec/projection version and
the processed block frontier. Recovery discards or replaces incomplete derived
state and replays verified blocks; it never rewrites the authoritative blocks.

### Scenario 10: Trust Boundary And Authorization

Mallory creates a structurally valid Grid-CBOR envelope with a valid pCID and
signature but asks Dave to perform an unauthorized action.

All encoding and representation alternatives can establish only byte validity,
framing, CID links, and inputs to signature verification. Dave still applies the
pCID-defined authentication semantics and separate local authorization policy.
No alternative gains authority merely by using DAG-CBOR, IPLD, deterministic
encoding, or a reversible node model.

## Comparison Matrices

### Encoding

| Alternative | Exact DAG-CBOR slot | Native arbitrary tags/keys | One representation per value | Unknown-pCID base processing | Principal obligation |
| --- | --- | --- | --- | --- | --- |
| E1 | No for binary64 floats with shorter exact forms | Yes | Yes | Yes | Would need to change a locked premise |
| E2 | Yes if extension preserves DAG rules | Yes | Yes | Yes | Specify the full extended profile |
| E3 | Yes | Yes | Yes if classification is unambiguous | Yes | Specify and prove the overlap boundary |
| E4 | Yes | Yes | No | Yes | Bind authority to exact retained bytes |
| E5 | Yes for DAG values, no for native rich-CBOR inputs | No; wrappers instead | Yes | Yes | Changes native rich-CBOR meaning |
| E6 | Only when selected or backed by a common base | Protocol-dependent | Protocol-dependent | No without a common base | Retain and resolve pCID rules |

### Decoded Representation

| Alternative | Exact bytes | Rich CBOR coverage | Generic link traversal | Rebuildable from CAS | Principal obligation |
| --- | --- | --- | --- | --- | --- |
| R1 | Yes, through retained raw bytes | Yes in authority; projection may narrow | Yes through derived view | Yes | Keep projection explicitly non-authoritative |
| R2 | Yes if representation choices are retained | Yes if model is complete | Yes | Yes | Specify and implement reversible model/ADL |
| R3 | No in general | Library-dependent | Sometimes | Not from normalized values | Prove losslessness or narrow accepted input |
| R4 | Yes | Yes as opaque bytes | No generic traversal | Yes | Supersede current generic-link requirement |

## Results

### Alternatives Rejected By The Locked Premises

- E1 is rejected because the valid DAG-CBOR binary64 encoding of a value such as
  `1.5` is not RFC 8949 CoreDet's shortest encoding. Rejecting or shortening it
  violates the exact slot-composition and round-trip premises.
- E5 is rejected as the complete Grid-CBOR encoding because wrapper maps do not
  accept and reproduce native arbitrary CBOR tags and map keys.
- R3 is rejected as a complete decoded representation because generic
  host-language values do not, without an additional raw syntax tree, preserve
  every arbitrary tag, compound key, and representation choice. Adding such a
  tree turns it into R1 or R2.

### Alternatives That Do Not Replace The Generic Base

- E6 remains valid as a layer in which a pCID narrows slot values and defines
  meaning, signatures, or embedded formats. It is not a complete base codec
  while unknown-pCID structural processing and exact DAG-CBOR composition are
  required.

### Surviving Encoding Alternatives

- E2: one DAG-CBOR-derived deterministic profile extended across Grid's broader
  CBOR domain.
- E3: exact DAG-CBOR overlap rules plus separately specified deterministic rules
  for Grid-only values.
- E4: multiple accepted representations with exact retained bytes as authority.

E4 satisfies the five new compatibility premises but conflicts with the current
independent-writer deterministic policy. It survives for DF because choosing
between one-representation convergence and accepted-representation diversity is
a policy decision, not a fact established by the premises.

### Surviving Representation Alternatives

- R1: retained block/raw slots plus a derived IPLD projection.
- R2: a reversible Grid-specific node model or IPLD ADL.
- R4: raw block only, but only if a later DI supersedes generic visible-link
  discovery.

R1 has direct experimental evidence in cdint-grid but is not thereby selected.
R2 has larger specification and implementation obligations but may provide a
cleaner unified API. R4 has the smallest generic model but conflicts with the
current link-discovery behavior. These tradeoffs require DF after the encoding
choice narrows what representation state must be retained.

## Exact DF Queue

The first unresolved question is:

> Should Grid-CBOR use (E2) one DAG-CBOR-derived deterministic representation
> across its full CBOR domain, (E3) exact DAG-CBOR rules for the overlap plus
> separately specified deterministic rules for Grid-only values, or (E4)
> multiple accepted representations whose exact input bytes remain
> authoritative?

Only one question should be asked at a time. After that answer, ask the
representation question using only R1, R2, and any still-eligible R4. Later DF
must close map ordering, non-DAG floats/simple values, duplicate-key policy,
resource limits, migration treatment of existing CoreDet blocks, and public
specification wording. Those are not silently decided here.

## Ninik Review

Classification: alignment.

The comparison uses exact CID-addressed bytes as durable authority, keeps pCID
semantics above generic framing, and treats IPLD projections and indexes as
disposable views with a stated derivation frontier and rebuild path. It checks
the simpler existing mechanisms—DAG-CBOR, RFC 8949 CBOR, CAS retention, raw
slots, and IPLD projection—against every alternative rather than assuming that
custom machinery wins.

The TE introduces no registry, mutable identity authority, recovery authority,
or application authorization rule. E2 and E3 would require a thin Grid-specific
validator over shared CBOR/CID primitives. E4 would require an input-preserving
validator and exact-byte authority. R1 reuses the POC's CAS-first projection
pattern; R2 must justify its added durable specification; R4 must explicitly
trade away current generic traversal. None is selected by this classification.

Decision status remains `needs DF` under DR-tuzob. There is no Ninik departure
to approve merely from running this comparison.

## Bias Audit

PASS.

- The alternatives were named by mechanism and received the same fixtures.
- The matrices report hard-premise failures separately from conflicts with
  existing decisions.
- Experimental availability was not treated as proof of architectural fitness.
- Existing CoreDet decisions were not treated as winners or silently discarded.
- No numerical weights or preference-loaded score selected a survivor.
- The conclusions follow explicit counterexamples and leave policy tradeoffs for
  DF.

## Implications For TODOs And DIs

1. TODO-gilog should track this completed analysis, DR-tuzob, the subsequent DF,
   and a later public-spec/conformance revision.
2. DI-vogod's CoreDet rule conflicts with the locked exact DAG-CBOR float
   composition premise. It remains historical/current authority until the user
   selects a survivor and a new DI explicitly supersedes the affected rule.
3. DI-nipap still governs the later public specification: internal TE/DR/DI
   references remain in coordination records rather than the public document.
4. No public specification, production code, pCID constant, identity vector, or
   production key should change from this TE alone.
5. The identity-genesis A01 worker may cite this TE and DR as a generic-envelope
   dependency, but the A02 product-code release barrier remains unchanged.

## References

- [RFC 8949: Concise Binary Object Representation](https://www.rfc-editor.org/rfc/rfc8949)
- [DAG-CBOR specification](https://ipld.io/specs/codecs/dag-cbor/spec/)
- [IPLD Data Model](https://ipld.io/docs/data-model/)
- [PromiseGrid Grid CBOR tag specification](../grid-cbor-tag-spec.md)
- [Current PromiseGrid Grid-CBOR multicodec specification](../grid-cbor-multicodec-spec.md)
- [TODO-gilog and Decision Intent Log](../../TODO/TODO-gilog-grid-cbor-multicodec.md)
- [DR-tuzob](../../DR/DR-tuzob-grid-cbor-dag-cbor-superset.md)
