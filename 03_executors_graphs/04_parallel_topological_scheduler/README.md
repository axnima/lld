# 023. Parallel topological scheduler with semaphore

Phase: Executors, dependency graphs, and pipelines. Local order: 04. Follow the [master plan](../../README.md).

## Core

Execute a DAG with a fixed worker pool; release dependents only after prerequisites finish.

Before coding, write your API, invariants, error contract, ownership model, and complexity target. For concurrency, identify shared state, synchronization, cancellation, and shutdown behavior.

## Ordered variants / checks

- [ ] Semaphore signals ready work and mutex protects scheduler state
- [ ] Validate cycles before execution
- [ ] Exactly-once dispatch
- [ ] No-ready-but-running versus deadlock
- [ ] Task failure blocks descendants
- [ ] Shutdown wakes all workers

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
