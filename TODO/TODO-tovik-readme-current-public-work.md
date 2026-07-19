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

## Task

- [x] tovik.1 Update the stale public-work pointer in `README.md`.
- [x] tovik.2 Correct outdated capability, message-envelope, and storage descriptions.
- [x] tovik.3 Keep the README change narrow and reviewable.
- [x] tovik.4 Preserve referential transparency and event replay in the README.
