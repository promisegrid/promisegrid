# Repository Guidelines

## Project Structure & Module Organization
- Keep prose drafts, presentation decks, and small repo tools in purpose-named top-level paths; avoid `internal/` and `pkg/`.
- Keep planning artifacts in the root `TODO/` directory. Optional root `DR/` and `docs/thought-experiments/` directories hold DR and TE records when needed.
- Do not commit local state files (for example `.grok`, `.grok.lock`) or generated binaries.

## Build, Test, and Development Commands
- Run Go tests from the relevant tool module, for example `tools/mint-handle/`.
- Keep Go code `gofmt`-clean.

## Agent Instruction Architecture (Required)
- `AGENTS.md` is the canonical home for repo-wide workflow, review, and vocabulary rules.
- Keep repo-wide rules here; keep private runtime notes out of committed files.

## Decision-First Specification and Compliance Protocol (Required)
- Decision-first means decisions must be locked before coding; it does not forbid pre-decision analysis such as required thought experiments.
- The agent must collect and lock user decisions before making any code edits for a task.
- Locked decisions must be recorded as Decision Intent Log entries in the relevant `TODO/TODO-<handle>-<slug>.md` file(s) with clear intent and rationale.
- Existing numeric TODO and DI records are legacy records and remain valid until migrated.
- The agent must ask decision questions up front in a single intake round whenever possible.
- Required decision categories are architecture, design/behavior, implementation approach, function naming, variable naming, and file/path decisions.
- The agent must ask these as multiple-choice questions whenever practical.
- When a thought experiment (TE) is required, the agent must complete the TE before asking final DF questions. TEs narrow alternatives; DF questions and answers lock the decision before implementation.
- Thought experiments (TEs) are analysis artifacts; Decision Intent (DI) entries are the separate records that capture the locked decision after DF is resolved.

## Thought Experiment Protocol (Required)
- Before locking any non-trivial decision that will require DF questions and answers, the agent must run a thought experiment (TE) if multiple plausible designs remain.
- A TE happens before final DF questions. Its purpose is to narrow the design space so DF questions and answers are informed by explicit scenario analysis.
- The agent must not collapse a TE into a short opinion or recommendation. The agent must explicitly model concrete scenarios and consequences.
- Each new TE must have a unique proquint handle in the format `TE-<handle>`, where `<handle>` is minted by `tools/mint-handle` from the global TODO/TE/DR/DI handle namespace. Existing pre-upgrade TE handles and prior aliases remain historical records.
- The TE doc filename must be `TE-<handle>-<slug>.md` and live under `docs/thought-experiments/`, for example: `docs/thought-experiments/TE-mumuv-naming-reconciliation.md`. The slug is informational and may be edited; the proquint handle is permanent.

### TE Intake Requirements
- Before locking decisions or asking final DF questions, the agent must identify:
  - the decision being tested,
  - the candidate alternatives,
  - the assumptions and threat/trust model,
  - the scope and systems affected.
- If the TE relates to an existing TODO, the agent must reference the TODO handle and subtask handle (for example, `fonuz.1`).

### TE Execution Requirements
- Each TE must evaluate the same decision across multiple concrete scenarios.
- Scenarios must include, when relevant:
  - normal operation,
  - failure/corruption/incomplete writes,
  - concurrent actors or mixed-version nodes,
  - long-horizon evolution and migration,
  - trust-boundary changes,
  - scale effects (storage, bandwidth, CPU, operational complexity).
- The agent must compare alternatives under the same assumptions instead of switching assumptions mid-analysis.
- The agent must state what each alternative makes easier, what it makes harder, and what new obligations it creates.

### TE Authoring Conventions
- Named actors follow the cryptography-literature alphabetical convention. Use Alice, Bob, Carol, Dave, Ellen, Frank, and so on for cooperative actors; use Mallory for adversaries. Use this convention in TEs, scenario analyses, tabletop simulations, DR/DI prose, and worked examples in specs or docs when named actors are useful. Do not invent ad-hoc names when the convention fits.

### TE Output to DF
- After the TE, the agent must identify:
  - rejected alternatives,
  - surviving alternatives,
  - unresolved questions that still require user choice,
  - any new naming/path/runtime decisions exposed by the TE.
- Final DF questions must be framed from the surviving alternatives identified by the TE. The agent must not ask broad DF questions that ignore TE results.

### TE Artifacts
- The agent must track required TEs in the relevant `TODO/TODO-<handle>-<slug>.md`.
- For each completed TE, the agent must write a verbatim copy of the thought experiment into a standalone file under `docs/thought-experiments/`.
- The doc filename must begin with the TE ID and then use a descriptive suffix.
- The doc must stand on its own and include:
  - title,
  - TE ID,
  - decision under test,
  - assumptions,
  - alternatives,
  - scenario analysis,
  - conclusions,
  - implications for the repo's open TODOs and pending DIs.

### TE Decision Rules
- A TE does not by itself lock a decision.
- After the TE, the agent must either:
  - ask the user to choose among the surviving alternatives, or
  - recommend one surviving alternative and clearly state why the others were rejected.
- After user choice is resolved, the agent must record the locked result via the existing DI process before implementation.
- If a TE exposes a new ambiguity, dependency, or naming/path decision, the agent must stop and resolve that before implementation.

### TE Final Handoff Requirements
- In the final response for TE work, the agent must include:
  - which TE was completed,
  - the TE ID,
  - the doc path under `docs/thought-experiments/`,
  - the surviving alternatives,
  - the recommended conclusion or the exact DF question that remains for the user.
- Hard gate: for decisions that require a TE, work is incomplete until the TE doc exists and the resulting decision status is explicit (`needs DF`, `locked`, or `deferred`).

### TE Editing Policy (Required)

- Treat TE files as durable analysis records. Do not rewrite filed TE history for style cleanup.
- Use a `## Refinements` section for navigational updates or resolved follow-up notes.
- Use a new superseding TE for material changes to a prior TE's conclusion, scope, or decision impact.

### Naming Decisions (Required)
- The agent must not invent function names or variable names that are not already covered by locked naming decisions.
- If naming is not covered, the agent must stop and ask multiple-choice naming options before continuing.

### File/Path Decisions (Required)
- Path approvals are mandatory for all touched paths:
  - repo-changed files (create/rename/move/delete),
  - runtime touched paths (read/write/delete), including input files, output files, DB files, caches, fixtures, and temporary test files.
- The agent must ask path approvals one path at a time via multiple-choice questions.
- Path-question order must be dependency order.
- Each path question must include: action, exact path (or approved dynamic pattern ID), purpose, class (`prod-code | prod-data | test | temp`), and lifecycle intent.
- Temporary test paths require explicit approval and an explicit cleanup plan before handoff.
- Dynamic/runtime-generated paths must be approved by pattern, with:
  - allowed root bounds,
  - allowed actions,
  - concrete examples.
- The agent must ask one multiple-choice approval per dynamic path pattern.
- If any unapproved runtime path appears, the agent must stop and ask before continuing.

### Decision Lock and Stop Rule
- The agent must produce a Decision Lock summary with decision IDs before code edits begin.
- The agent must not proceed if any required decision is missing, ambiguous, or conflicting.
- The agent must stop and ask immediately if a new decision need appears during implementation.
- The agent must not assume defaults for locked categories unless the user explicitly approves defaults.

### Compliance Ownership (Agent)
- The agent must treat user decisions as authoritative and implement to those decisions.
- The agent must run a compliance self-review before finalizing and must fix all non-compliance before handoff.
- Hard gate: work is incomplete until compliance is PASS, or the user explicitly approves an exception.
- The user should not need to manually inspect diffs to determine compliance.

### Required final handoff artifacts
- `Decision Compliance: PASS/FAIL`
- Decision Matrix mapping each locked decision ID to implementation evidence.
- Inline diff annotations in the form `path:line -> decision_id -> rationale`.
- Runtime Path Touch Matrix listing each approved runtime path/pattern, action used, and where it is implemented/validated.
- `Exceptions:` listing only user-approved deviations.
- Every non-trivial behavior change must include intent provenance per existing DI requirements.

## Coding Style & Naming Conventions
- Use object-oriented design with structs and methods; avoid large functions and global state.
- Follow generally accepted object oriented design patterns.
- Keep Go code `gofmt`-clean; package names should be short and lower-case.
- Prefer focused edits over broad refactors unless required.
- Add and maintain explanatory comments for non-obvious logic.
- Use `git mv` for file moves/renames to preserve history.

## Diff Discipline (Required)
- Keep changes minimal and directly tied to the user's request or locked decision.
- Do not rewrite files from scratch, arbitrarily rewrap lines, reorder unrelated sections, or normalize prose style unless the user explicitly asks for that cleanup.
- This applies to `AGENTS.md`, slides, white papers, drafts, TODO files, code, and all other repo files.
- For public prose artifacts, edits must be traceable to a concrete user request or locked decision, not general polish.

## Runtime Artifact Hygiene (Required)
- Never put temporary test files, Go cache directories, build caches, or other runtime artifacts in this repo.
- Runtime artifacts must go under `/tmp` subdirectories.
- Configure tools with temp/cache paths under `/tmp` when they would otherwise write into the working tree.

## Error Handling Policy (Required)
- Never use `|| true` in scripts, templates, or make recipes. Always inspect
  command exit codes explicitly with `if/else` branches and handle each outcome.
- For non-fatal cleanup/diagnostics steps, record command status (exit code and
  logs) explicitly; do not fail silently.
- In Go code, never ignore errors with `_ = ...`; handle, propagate, or report
  errors explicitly.
- Run `errcheck ./...` and keep it passing for Go changes.

## DR/DI Source-of-Truth Protocol (Required)
- In this repo, DR and DI logs are the primary source of truth for decisions and open questions.
- Documents and code are outputs of that process and must link back to DR/DI records.
- Person identity in DR/DI records must use full email with label format: `user@example.com (FirstName)`.
- In DRs, `Asked by` and person-valued `Waiting on` fields must use that format.
- In DIs, `Author` must use that format.
- A DI `Author` is the decision-maker, not merely the recorder. See
  the git config for this repo to detect the author's name.
- A settled statement in docs (or critical logic in code comments) must cite at least one DI ID.
- An unresolved question or uncertainty must cite at least one DR ID.
- If an unresolved question has no DR yet, create a DR before finalizing the change.
- New coordination artifacts use a single proquint handle namespace so TODO, TE, DR, and DI references do not depend on timestamp or integer allocation.
- Existing numeric TODO and DI records remain valid legacy records until migrated.

## Public Artifact Provenance (Required)
- Do not put DI, DR, TODO, or TE references in slides; keep public decks visually clean.
- In white papers, use GitHub Markdown footnote tags for DI, DR, TODO, or TE references.
- White papers with DI, DR, TODO, or TE footnotes must include a `## References` section at the bottom of the document.
- The `## References` section must cite the appropriate repo files using relative file paths so readers can find the detailed record.


## Comment Preservation Protocol (Required)
- Never remove existing code comments unless they are replaced in the same patch by equal-or-better explanatory comments near the same logic.
- When rewriting or refactoring code, port old explanatory intent first, then improve wording.
- If a touched non-trivial code block has no comments, add explanatory comments.
- Do not treat shorter comments as better unless they preserve all important intent.
- For any non-trivial behavior change, include a behavior-level comment with:
  - `Intent:` a short, clear rationale (a sentence or a few; no hard cap if more is needed for clarity).
  - `Source:` a DI ID in the format `DI-<handle>`.
  - `<handle>` is minted by `tools/mint-handle` and is globally unique across TODO, TE, DR, and DI owners.
  - Optional: TODO file/section reference for faster lookup.
- If a comment must be dropped with no replacement, stop and ask the user before proceeding.
- Before editing a file, review existing comments in that file.
- Maintain a `## Decision Intent Log` at the top of relevant `TODO/TODO-<handle>-<slug>.md` files.
- Treat DI logs as append-only history. Do not rewrite or delete prior entries.
- When intent evolves, add a new DI entry and set `Supersedes: <old-di-id>`.
- DI entries must include:
  - `ID: DI-<handle>`
  - `Date: YYYY-MM-DD HH:MM:SS`
  - `Status: active|superseded`
  - `Decision:`
  - `Intent:`
  - `Constraints:`
  - `Affects:`
  - `Supersedes:` (optional)
- After editing, run a comment-delta audit on each touched code file using: `git diff -U0 -- <file> | rg -n '^-\\s*//|^-\\s*/\\*|^\\+\\s*//|^\\+\\s*/\\*'`.
- Resolve all removed-comment lines before finalizing unless explicit user approval was given.
- In the final response, include:
  - `Comment audit: PASS/FAIL`, with file list.
  - `Intent provenance audit: PASS/FAIL`, listing files with behavior changes and DI sources.
- Hard gate: behavior-changing work is incomplete unless comments preserve intent and include DI provenance.
- Do not remove comments or documentation; update them if outdated or incorrect.

### Comment + DI Examples
- Comment format example:
  - `// Intent: Keep context resolution stable across workspace scans to avoid target drift between plan and run. Source: DI-vapoj`
- Decision Intent Log entry template (for TODO files):
  - `ID: DI-<handle>`
  - `Date: YYYY-MM-DD HH:MM:SS`
  - `Status: active`
  - `Decision: <what was decided>`
  - `Intent: <short clear rationale>`
  - `Constraints: <hard limits, dependencies, assumptions>`
  - `Affects: <paths, modules, commands, docs>`
  - `Supersedes: <old DI ID, optional>`

# DR Records

The DR/ directory stores Decision Request (DR) records for coordination work.

Rules:
- One DR per file.
- DR files are append-only event logs.
- Keep TODO files as snapshots; link TODOs to DR files for open questions.
- Person identity format: `user@example.com (FirstName)`.

Recommended file naming:
- `DR-<handle>-<slug>.md`, where `<handle>` is minted by `tools/mint-handle`.

Required DR fields:
- `DR-ID`
- `Date`
- `Asked by` (person identity format above)
- `State` (`open | decided | blocked | implemented | closed`)
- `Question`
- `Why this blocks progress`
- `Affects` (repos/files/components)
- `Unblocks` (TODO IDs/tasks)
- `Waiting on` (person identity format above, or DI ID)
- `Decision` (filled when decided)
- `Linked DI`
- `Related commits`
- `Last updated`

Reference pattern:
- From TODO files: `../DR/<filename>.md`


## TODO Tracking
- Maintain a `./TODO/` directory for tracking tasks and plans.
- Maintain a `./TODO/TODO.md` file that lists small tasks and the other TODO files.
- New TODO files use `TODO-<handle>-<slug>.md`, where `<handle>` is minted by `tools/mint-handle`.
- Existing numeric TODO files remain valid legacy records until migrated.
- Sort `TODO.md` by priority, not filename.
- When completing a TODO, mark it as done by checking it off.
- Within a TODO file, include numbered checkboxes for subtasks where helpful.

## Testing Guidelines
- Use Go's standard `testing` package with deterministic tests.
- Avoid network calls in tests unless explicitly required and documented.
- Put test temp files and Go caches under `/tmp`, not inside this repo.

## Commit & Pull Request Guidelines
- Treat a line containing only `commit` as: add and commit all changes with an AGENTS-compliant message.
- Use short, imperative, capitalized commit subjects.
- Summarize changes per file in commit bodies.
- Stage files explicitly (avoid `git add .` / `git add -A`).
- Do not open GitHub pull requests unless the user explicitly asks for one.
- Do not force-push repo branches unless the user explicitly authorizes the exception.
- Do not commit local state files, generated binaries, credentials, tokens, signing keys, or other secrets.

## Glossary
- **TE**: Thought Experiment. Analysis doc under `docs/thought-experiments/TE-<handle>-<slug>.md`. The handle is a proquint minted by `tools/mint-handle`.
- **DR**: Decision Request. Open question or decision-tracking record under `DR/DR-<handle>-<slug>.md`, where `<handle>` is minted by `tools/mint-handle`.
- **DI**: Decision Intent. Locked decision record inside a `## Decision Intent Log` in a TODO file. New DI ID format is `DI-<handle>`, where `<handle>` is minted by `tools/mint-handle`.
- **DF**: Decision Framing. The multiple-choice intake round used to lock a decision after any required TE narrows the alternatives.
- **TODO**: Task-tracking file under `TODO/TODO-<handle>-<slug>.md`, where `<handle>` is minted by `tools/mint-handle`. Existing numeric TODO files are legacy records until migrated.
- **pCID**: Protocol CID. A pCID is the content hash of a spec document that defines a wire protocol; it is analogous to a TCP/UDP port number with no central registry because the spec hash is the port number. A pCID is not the hash of a particular message, payload, or promise body.
