# Register Grid CBOR Tag

## Decision Intent Log

ID: DI-gatiz
Date: 2026-07-19 21:22:32
Status: active
Author: stevegt@t7a.org (Steve)
Decision: Register the PromiseGrid grid envelope as CBOR tag `0x67726964` / `1735551332`, use `docs/grid-cbor-tag-spec.md` as the public semantics document, include an IANA submission appendix, and use `Steve Traugott <ietf-poc@t7a.org>` as the point of contact.
Intent: Make the outer grid envelope self-identifying with a CBOR semantic tag whose tag-number bytes spell `grid`, while preserving pCID-selected protocol slots and the existing tag 42 CID semantics.
Constraints: Do not reuse tag 42; keep tag 42 for pCID. Do not submit the IANA request automatically. Re-check the IANA CBOR Tags registry before manual submission. Keep README and slide edits narrow. Do not alter `.gitignore` or add standard project metadata links.
Affects: docs/grid-cbor-tag-spec.md; docs/thought-experiments/TE-tavud-grid-cbor-tag-registration.md; README.md; slides/intro/README.md; TODO/TODO.md; TODO/TODO-gidur-register-grid-cbor-tag.md

## Task

- [x] gidur.1 Record the tag-registration thought experiment as `TE-tavud`.
- [x] gidur.2 Create `docs/grid-cbor-tag-spec.md` with the tag semantics and IANA appendix.
- [x] gidur.3 Update the README's pCID-selected grid message section to point at the tag spec.
- [x] gidur.4 Update canonical intro slide source that still described `"grid"` as a CBOR text-string marker.
- [ ] gidur.5 Manually submit the IANA CBOR Tags registration request after the spec URL is public.
- [ ] gidur.6 Update the spec and README after IANA confirms the assigned tag.
