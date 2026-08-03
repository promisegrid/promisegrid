# Grid CBOR Multicodec

## Decision Intent Log

ID: DI-rizuz
Date: 2026-08-02 18:25:12
Status: active
Author: stevegt@t7a.org (Steve Traugott)
Decision: Define `grid-cbor` as the structural codec for complete PromiseGrid Grid envelopes encoded with RFC 8949 Section 4.2.1 Core Deterministic Encoding. Propose draft multicodec code `0x1027` in the `ipld` category. The same deterministic envelope bytes may be identified by generic `cbor/0x51`; changing the CID codec changes identity and decoder dispatch, not serialization. A pCID may constrain and interpret protocol-defined values but may not select another base encoding for inline envelope values.
Intent: Give independent implementations one reproducible representation, allow codec-selected validators to discover recursively visible tag-42 CID links, and preserve arbitrary valid CBOR tags and map-key forms without making protocol interpretation a prerequisite for structural decoding.
Constraints: The outer value is the registered `grid` CBOR tag over an array whose slot 0 is a tag-42 pCID. Use preferred shortest forms, definite lengths, bytewise-lexicographic map-key ordering, canonical NaN and infinity encodings, and preserve negative zero and other CBOR-level distinctions. Do not normalize Unicode or reduce tagged or floating-point values based on application semantics. Every recursively visible valid tag-42 CID is an unlabeled graph edge; do not scan opaque byte strings, infer retention, or automatically fetch targets. Unknown tags remain valid. Malformed tag-42 values, duplicate map keys, and non-deterministic encodings are invalid. The exact proposal code remains provisional until merged by the registry; existing CIDs are never relabeled. The public specification and registry proposal must contain no private application data.
Affects: `TODO/TODO-gilog-grid-cbor-multicodec.md`, `TODO/TODO.md`, `docs/grid-cbor-multicodec-spec.md`, `README.md`, the external `multiformats/multicodec` draft registration proposal

ID: DI-vogod
Date: 2026-08-02 19:02:43
Status: active
Author: stevegt@t7a.org (Steve Traugott)
Decision: Revise the public `grid-cbor` specification to define the proposed codec's deterministic Grid envelope encoding, visible tag-42 CID validity, and recommended decoder rejection behavior. Remove the separate claims about generic `cbor/0x51` using the same serialization, pCIDs not selecting another inline base encoding, generic software recording each visible CID as an edge, and existing CIDs never being relabeled. Remove the duplicate data-item definition and conformance vectors. Use `SHOULD`, rather than `MUST`, for rejecting malformed or nonconforming input.
Intent: Keep the public codec proposal concise and avoid freezing compatibility, graph-index, migration, and conformance-vector policy in this specification before those concerns are settled elsewhere.
Constraints: Continue to require RFC 8949 Core Deterministic Encoding for complete Grid envelopes, permit arbitrary valid CBOR tags and map-key types, require every visible tag-42 value to contain a valid CID, and prohibit trusting or executing a message merely because its Grid format is valid. The Grid CBOR tag specification remains the source for the envelope shape. Code `0x1027` remains provisional, and the public specification must contain no private application data.
Affects: `TODO/TODO-gilog-grid-cbor-multicodec.md`, `docs/grid-cbor-multicodec-spec.md`, the external `multiformats/multicodec` draft registration proposal
Supersedes: DI-rizuz

ID: DI-nipap
Date: 2026-08-02 19:07:05
Status: active
Author: stevegt@t7a.org (Steve Traugott)
Decision: Keep the public Grid CBOR multicodec specification free of internal DR, DI, and TODO references. Preserve decision provenance in the repository's internal coordination records and Git history rather than linking those records from the public specification.
Intent: Keep the public standards document self-contained and avoid exposing internal project-coordination machinery to its readers.
Constraints: This changes provenance presentation only; it does not change the specification's technical requirements. Public standards and public PromiseGrid documents may remain linked. Internal DR, DI, and TODO records remain the source of truth for project decisions. This is the user's explicit exception to the repo rule that settled statements in documents cite a DI.
Affects: `TODO/TODO-gilog-grid-cbor-multicodec.md`, `docs/grid-cbor-multicodec-spec.md`

## Task

- [x] gilog.1 Record the deterministic Grid CBOR representation and proposed
  multicodec registration decision.
- [x] gilog.2 Publish the generic structural codec specification.
- [x] gilog.3 Link the codec specification from the README.
- [ ] gilog.4 Propose the draft `grid-cbor` row to the multicodec registry.
- [ ] gilog.5 Track registry review without treating an unmerged proposal as an
  assigned code.
