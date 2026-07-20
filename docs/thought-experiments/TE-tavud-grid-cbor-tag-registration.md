# Grid CBOR Tag Registration Thought Experiment

TE ID: TE-tavud

## Decision Under Test

Choose the public CBOR tag strategy for PromiseGrid grid envelopes before locking the IANA registration request in DI-gatiz.

## Assumptions

- The outer envelope remains pCID-selected: slot 0 is tag 42 carrying the protocol CID, and the pCID-selected spec owns the remaining slots.
- The tag must identify a PromiseGrid grid envelope without making tag 42 do two jobs.
- The user wants the registered tag to spell `grid` when the tag number bytes are rendered as ASCII.
- The IANA request is prepared in the repo but submitted manually.

## Alternatives

- A: Request the lowest available FCFS CBOR tag.
- B: Request `0x67726964` / `1735551332`, whose four tag-number bytes spell `grid`.
- C: Keep a CBOR text string `"grid"` as the first array element and do not register an outer CBOR tag.
- D: Register a fixed legacy five-element array shape.

## Scenario Analysis

Normal decoding: Alice receives a tagged CBOR value. Alternative B gives Alice a standard CBOR semantic tag and keeps the diagnostic alias `grid(...)` tied to one numeric value. Alternative A is more compact but does not satisfy the ASCII spell-out requirement. Alternative C is readable but is not a CBOR tag registration. Alternative D over-specifies slots that now belong to pCID-selected protocols.

Replay and signature stability: Bob signs or replays byte-identical envelopes. Alternative B keeps the outer tag bytes stable and leaves deterministic encoding and proof coverage to the pCID-selected protocol. Alternative C preserves a visible string but mixes protocol recognition with the application array. Alternative D makes future pCID evolution harder because slot meanings would be registered too early.

Mixed-version nodes: Carol's old decoder may not understand the new semantic tag. With Alternative B, Carol can reject the unknown tag or inspect the tagged array under explicit compatibility rules. Alternative D creates more failure modes because old and new slot definitions can both look like five-element arrays.

Registry collision: Dave checks the IANA CBOR Tags registry before submission. Alternative B has a specific requested value, so the implementation must stop if `1735551332` is assigned before filing. Alternative A avoids that specific risk but loses the `grid` byte pattern.

Trust boundary: Mallory sends a syntactically valid `grid` tag with an untrusted pCID. None of the alternatives should make the message trusted. Alternative B keeps this clear: the outer tag classifies the envelope, while validation, authorization, and execution remain pCID-defined and locally trusted.

Long-horizon evolution: Ellen introduces a new protocol with different payload or proof slots. Alternative B supports this because only the outer envelope and slot 0 are common. Alternative D would require updating or bypassing the registered shape.

## Conclusions

Reject A because it does not meet the ASCII spell-out requirement. Reject C because it does not register `grid` as a CBOR semantic tag. Reject D because it freezes obsolete five-element slot semantics. Use B: register `0x67726964` / `1735551332` as the outer `grid` tag over a pCID-selected array.

## Implications For Open Work

- DI-gatiz locks the chosen tag value, spec path, point of contact, and manual-submission workflow.
- TODO-gidur tracks the local spec and documentation work plus the later manual IANA submission.
- Public slide/source prose should stop presenting `"grid"` as the current CBOR text-string marker.
