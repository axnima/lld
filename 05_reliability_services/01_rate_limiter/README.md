# 042. Rate Limiter

Phase: Reliability, APIs, and backend services. Local order: 01. Follow the [master plan](../../README.md).

## Core

Guard a simulated LLM gateway with per-user tier limits and two interchangeable policies.

Before coding, write your API, invariants, error contract, ownership model, and complexity target. For concurrency, identify shared state, synchronization, cancellation, and shutdown behavior.

## Ordered variants / checks

- [ ] Supplied prompt, 45 min
- [ ] Fixed window
- [ ] Sliding window log
- [ ] Sliding window counter
- [ ] Token bucket
- [ ] Leaky bucket
- [ ] Atomic check-and-consume
- [ ] Retry-after
- [ ] Per-user/global combination
- [ ] Bounded inactive-key cleanup

Existing notes: [original rate-limiter sketch](../../rate_limiter/rate_limiter.md). Preserve it as your initial design; compare its assumptions with the implementations.

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
