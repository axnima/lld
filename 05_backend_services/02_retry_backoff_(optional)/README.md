# 043. Retry and backoff [Optional — depth]

Phase: Reliability, APIs, and backend services. Local order: 02. Follow the [master plan](../../README.md).

**[Optional — depth]** Skip this standalone implementation during the first pass; related concepts may appear in a focus session.

[Session order and progress](../../STUDY_PLAN.md)

## Core

Retry transient failures under a total deadline.

Before coding, write your API, invariants, error contract, ownership model, and complexity target. For concurrency, identify shared state, synchronization, cancellation, and shutdown behavior.

## Ordered variants / checks

- [ ] Exponential backoff
- [ ] Capped jitter with seeded RNG
- [ ] Retry classification
- [ ] Cancellation
- [ ] Retry budget
- [ ] Idempotent versus unsafe operation

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
