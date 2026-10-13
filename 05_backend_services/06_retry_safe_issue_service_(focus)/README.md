# 047. Retry-Safe Issue Service

Phase: Reliability, APIs, and backend services. Local order: 06. Follow the [master plan](../../README.md).

**[Focus — breadth]** Use the session guide to select the core and a few relevant variants. You do not need every checkbox before moving on.

[Session order and progress](../../STUDY_PLAN.md)

## Core

Expose an authorized issue service with idempotent create, versioned updates, and pagination.

Before coding, write your API, invariants, error contract, ownership model, and complexity target. For concurrency, identify shared state, synchronization, cancellation, and shutdown behavior.

## Ordered variants / checks

- [ ] Supplied prompt, 75 min
- [ ] Complete HTTP API
- [ ] Forbidden versus missing policy
- [ ] Concurrent same-key creates
- [ ] Key/payload conflicts
- [ ] Stable cursor pagination
- [ ] Duplicate retries
- [ ] Contract tests

## Language passes

Follow the assigned session scope in one language. Full variants and other languages are later depth passes. Keep a few meaningful correctness checks.

- [ ] [Python](python) core with focused checks; optional variants later.
- [ ] [Go](go) core with focused checks; optional variants later.
- [ ] [TypeScript](typescript) core with focused checks; optional variants later.

## Acceptance and review

- [ ] Core happy path, empty/minimal input, invalid input, and exact boundaries pass.
- [ ] Relevant failure paths preserve invariants and do not leave partial mutations.
- [ ] If concurrent: deterministic contention test, bounded work/memory, cancellation, and shutdown pass. Otherwise mark this item N/A in your notes.
- [ ] Responses do not unintentionally expose mutable internal state or another user's data.
- [ ] Explain time/space costs, design choices, and remaining limits without reading code.
- [ ] Repeat the core from scratch under a time limit; record duration and missed edge cases.

## Attempt notes

| Date | Language | Core / variant | Minutes | Missed invariant or bug | Next review |
| --- | --- | --- | --- | --- | --- |
