# DR-tuzob: Grid-CBOR And DAG-CBOR Superset Boundaries

DR-ID: DR-tuzob

Date: 2026-08-16 21:21:38 -0700

Asked by: stevegt@t7a.org (Steve Traugott)

State: open

Question: Which encoding and decoded-representation alternatives identified by
[TE-fujur](../docs/thought-experiments/TE-fujur-grid-cbor-dag-cbor-superset.md)
should define Grid-CBOR's one-way DAG-CBOR compatibility while preserving native
arbitrary CBOR tags and map-key types?

Why this blocks progress: The current public Core Deterministic float rule
rejects or rewrites valid DAG-CBOR binary64 floats that have shorter exact CBOR
representations. A codec implementation, conformance fixtures, public-spec
correction, and dependent generic identity-envelope work cannot safely proceed
until the base encoding and decoded representation are selected and recorded in
a superseding DI.

Affects: `TODO/TODO-gilog-grid-cbor-multicodec.md`,
`docs/grid-cbor-multicodec-spec.md`, generic Grid-CBOR codecs and validators,
IPLD link projection, and cdint-grid identity-genesis generic-envelope planning

Unblocks: `gilog.7`, `gilog.8`, `gilog.9`, and the generic-envelope dependency
that precedes identity-genesis A02 product code

Waiting on: stevegt@t7a.org (Steve Traugott)

Decision: Pending. TE-fujur leaves encoding alternatives E2, E3, and E4 for the
first DF question. The representation choice follows from surviving R1, R2, and
any still-eligible R4. No alternative is selected by this DR.

Linked DI: None yet. The selected result requires a new DI that explicitly
supersedes the incompatible part of DI-vogod and preserves DI-nipap's public
provenance rule.

Related commits: PromiseGrid analysis base
`7b515985809236a7d3ddae7872bf4d698fe7d824`; cdint-grid coordination correction
`872fb79`

Last updated: 2026-08-16 21:21:38 -0700

## Event Log

### 2026-08-16 21:21:38 -0700

Opened after the user clarified that Grid-CBOR is a one-way superset for
composition: every valid DAG-CBOR item must remain byte-for-byte valid in a Grid
slot, while Grid-CBOR may also contain native values outside DAG-CBOR. Filed the
neutral TE-fujur comparison. Decision status remains `needs DF`.
