# 046. Optimistic concurrency control

Phase: Reliability, APIs, and backend services. Local order: 05. Follow the [master plan](../../README.md).

## Core

Update resources using expected versions and reject lost updates.

Before coding, write your API, invariants, error contract, ownership model, and complexity target. For concurrency, identify shared state, synchronization, cancellation, and shutdown behavior.

## Ordered variants / checks

- [ ] Compare-and-swap
- [ ] HTTP ETag/If-Match
- [ ] Concurrent deletes
- [ ] Bounded retry
- [ ] Multi-entity consistency
- [ ] ABA discussion

## Language passes

- [ ] [Python](python/) core and variants, with tests.
- [ ] [Go](go/) core and variants, with tests.
- [ ] [TypeScript](typescript/) core and variants, with tests.

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
