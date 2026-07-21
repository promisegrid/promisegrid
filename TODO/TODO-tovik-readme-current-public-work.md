# Refresh README Current Public Work

## Decision Intent Log

ID: DI-budan
Date: 2026-07-19 13:09:55
Status: active
Author: stevegt@t7a.org (Steve)
Decision: Refresh the root README as the public PromiseGrid overview, pointing current public work at wire-lab, promisegrid-dev-guide, and grid-examples while correcting stale protocol, capability, and storage language.
Intent: Keep new readers and collaborators oriented toward the current public work without rewriting the whole README.
Constraints: Keep edits minimal and diffable; do not alter .gitignore; do not make broad prose normalization changes; treat promisegrid-dev-guide as a current public entry point even if it is temporarily unavailable during implementation.
Affects: README.md; TODO/TODO.md; TODO/TODO-tovik-readme-current-public-work.md

ID: DI-refar
Date: 2026-07-19 13:41:34
Status: active
Author: stevegt@t7a.org (Steve)
Decision: Preserve referential-transparency and event-replay characteristics in the README storage and merge descriptions.
Intent: Keep the README's current pCID and sparse-CAS framing without losing deterministic interpretation, replayable histories, and event-sourced rebuild properties.
Constraints: State these properties as protocol-scoped characteristics rather than a global claim that all grid activity is pure or side-effect-free.
Affects: README.md; TODO/TODO-tovik-readme-current-public-work.md

ID: DI-pizov
Date: 2026-07-19 20:03:36
Status: active
Author: stevegt@t7a.org (Steve)
Decision: Plan the remaining README cleanup as four incremental passes: status/reader path, top-level claim tightening, architecture boundary cleanup, and public overview cleanup.
Intent: Improve the README for new public readers without turning the next edit into a full rewrite or adding standard project metadata links.
Constraints: Keep changes incremental and diffable; do not add license, governance, code-of-conduct, security, maintainership, or documentation-index links; keep README as a public overview rather than a comprehensive protocol spec.
Affects: README.md; TODO/TODO.md; TODO/TODO-tovik-readme-current-public-work.md

ID: DI-totar
Date: 2026-07-20 20:24:25
Status: superseded
Author: stevegt@t7a.org (Steve)
Decision: Keep historical enabling milestones visible in the README while replacing the long milestones checklist with a compact roadmap summary.
Intent: Preserve the public story of why the work is now possible without keeping transient outreach and project-management detail in the README.
Constraints: Keep the edit incremental and diffable; do not add standard project metadata links; do not move the detailed checklist into another file in this pass.
Affects: README.md; TODO/TODO-tovik-readme-current-public-work.md

ID: DI-fitut
Date: 2026-07-20 20:45:45
Status: active
Author: stevegt@t7a.org (Steve)
Decision: Preserve the README milestones as a checklist-style Foundations and Roadmap section while tightening only the public-facing content that belongs there.
Intent: Keep historical milestones, active work, in-progress dogfooding, and future application targets visible because they explain the work's trajectory and enabling context.
Constraints: Do not remove the milestone/checklist structure; trim transient outreach, obsolete names, and private coordination detail only where needed; do not add standard project metadata links.
Affects: README.md; TODO/TODO-tovik-readme-current-public-work.md
Supersedes: DI-totar

ID: DI-nanaj
Date: 2026-07-21 10:15:34
Status: superseded
Author: stevegt@t7a.org (Steve)
Decision: Clarify README architecture boundaries while keeping closure-based and signed-token capability models as live pCID-selected options.
Intent: Make the public overview more precise about issuer authority, pCID ownership, local trust, non-global merge semantics, sparse-CAS roles, and replay boundaries without freezing public APIs.
Constraints: Keep edits narrow and diffable; preserve both capability models; align grid-envelope wording with docs/grid-cbor-tag-spec.md; do not add standard project metadata links.
Affects: README.md; TODO/TODO-tovik-readme-current-public-work.md

ID: DI-jajir
Date: 2026-07-21 10:52:33
Status: active
Author: stevegt@t7a.org (Steve)
Decision: Treat the current README architecture wording as the completed tovik.7 outcome, including the closure capability model, signed-token POC model, pCID envelope wording, VCS-like merge consensus framing, and sparse-CAS replay language.
Intent: Keep the TODO/DI record aligned with the user-edited README instead of forcing the README back to the earlier broader architecture-boundary plan.
Constraints: Do not rewrite the README to match DI-nanaj; preserve the current README's merge-as-consensus framing and sparse-CAS wording; keep edits minimal and diffable.
Affects: README.md; TODO/TODO-tovik-readme-current-public-work.md
Supersedes: DI-nanaj

## Task

- [x] tovik.1 Update the stale public-work pointer in `README.md`.
- [x] tovik.2 Correct outdated capability, message-envelope, and storage descriptions.
- [x] tovik.3 Keep the README change narrow and reviewable.
- [x] tovik.4 Preserve referential transparency and event replay in the README.
- [x] tovik.5 Increment 1: add a current-status section and a compact reader path.
  - [x] tovik.5.1 Add a short `## Current status` section near the top after the current public-work repo paragraph.
  - [x] tovik.5.2 Split status into exactly three explicit buckets: `Implemented now`, `Active prototype work`, and `Planned design`.
  - [x] tovik.5.3 Use the status section to reduce overclaiming by separating implemented behavior, architecture direction, and long-term vision.
  - [x] tovik.5.4 Add a compact `## How to explore` section for technical readers.
  - [x] tovik.5.5 In `## How to explore`, point readers to `wire-lab`, `grid-examples`, and `promisegrid-dev-guide` in reading order.
  - [x] tovik.5.6 State that this root repo is primarily an overview repo if no repo-local runnable demo is provided here.
  - [x] tovik.5.7 Avoid adding demo commands unless they are repo-local and verified during the edit.
  - [x] tovik.5.8 Include where to read next for protocol details.
- [x] tovik.6 Increment 2: tighten top-level claims without broad rewrites.
  - [x] tovik.6.1 Revise only the opening paragraphs enough to clarify what PromiseGrid is.
  - [x] tovik.6.2 State who PromiseGrid is for.
  - [x] tovik.6.3 State what problem PromiseGrid addresses first.
  - [x] tovik.6.4 State why the promise-based design is different from conventional distributed systems.
  - [x] tovik.6.5 Keep the `decentralized computer` framing if it still fits the surrounding prose.
  - [x] tovik.6.6 Distinguish long-term governance vision from current prototype work.
  - [x] tovik.6.7 Replace broad or hype-like claims with concrete, testable examples where a small local edit can do so.
  - [x] tovik.6.8 Keep line wrapping stable and avoid rewriting unrelated sections.
- [x] tovik.7 Increment 3: align capability, pCID, merge, and sparse-CAS architecture wording with the current README.
  - [x] tovik.7.1 Keep both capability models visible: `Capability as Closure` and signed capability promises used by current public POCs.
  - [x] tovik.7.2 In the closure model, describe a capability token as signed closure code addressed by hash under protocols that use that model.
  - [x] tovik.7.3 In the signed-token model, describe bearer tokens, identity-bound tokens, and redemption for more-specific identity-bound tokens.
  - [x] tovik.7.4 Keep token shape protocol-specific and cite CWT/COSE-shaped objects as examples, not frozen APIs.
  - [x] tovik.7.5 Refine `pCID-selected grid messages` to include `grid([42(pCID), ...])`, IPLD tag 42, pCID-owned slots, and pCID-defined trust/proof/replay semantics.
  - [x] tovik.7.6 Preserve the README's statement that a pCID is not a peer address, app address, message type, route, or operation code.
  - [x] tovik.7.7 Refine `Merge-as-Consensus Model` around VCS-like continual consensus formation.
  - [x] tovik.7.8 Clarify that protocol-specific merge inputs can include event streams, branch heads, or object DAGs.
  - [x] tovik.7.9 Make clear that a cross-reference can be evidence without requiring an immediate or full merge.
  - [x] tovik.7.10 Refine `Sparse CAS and local trust` to distinguish mutable or rebuildable local data from immutable CAS source objects.
  - [x] tovik.7.11 Preserve referential-transparency and event-replay wording using signed envelopes, append-only event logs, or parent-linked object DAGs from CAS.
- [x] tovik.8 Increment 4: preserve the public roadmap milestones while excluding standard project metadata links.
  - [x] tovik.8.1 Keep a checklist-style `## Foundations and Roadmap` section that preserves historical milestones.
  - [x] tovik.8.2 In the milestones, distinguish completed foundations, active or in-progress work, and future application targets.
  - [x] tovik.8.3 Keep detailed milestone material in the main README when it explains historical enabling context, current public work, dogfooding, or future grid application targets.
  - [x] tovik.8.4 Keep the current neutral contributing section and do not restore personal contact details, named outreach targets, or transient meeting status.
  - [x] tovik.8.5 Keep sponsorship brief and neutral unless the user asks to move or remove it.
  - [x] tovik.8.6 Do not add standard project metadata links: no license link, governance link, contribution guide link, code of conduct link, security disclosure link, maintainership policy link, or documentation index link.
  - [x] tovik.8.7 Keep the README project-generic and avoid private deployment, customer-specific, or transient coordination context.
  - [x] tovik.8.8 Clean up trailing whitespace and obvious scanability problems in touched lines only; do not normalize style or rewrap unrelated prose.
