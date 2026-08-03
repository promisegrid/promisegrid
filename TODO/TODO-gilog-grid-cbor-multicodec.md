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

## Task

- [x] gilog.1 Record the deterministic Grid CBOR representation and proposed
  multicodec registration decision.
- [x] gilog.2 Publish the generic structural codec specification.
- [x] gilog.3 Link the codec specification from the README.
- [ ] gilog.4 Propose the draft `grid-cbor` row to the multicodec registry.
- [ ] gilog.5 Track registry review without treating an unmerged proposal as an
  assigned code.
