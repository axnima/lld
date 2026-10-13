# 39. Bloom filter project defense

[Track and completion rules](../../README.md) · [Official reference](https://pkg.go.dev/testing)

## Goal

Rebuild the Go Bloom filter service and justify every mathematical and operational claim.

## Ordered practice checklist

- [ ] Sizing m/k from n/p
- [ ] MurmurHash3 known vectors
- [ ] 128-bit implementation
- [ ] Double hashing
- [ ] Concurrent bit updates
- [ ] No false negatives absent reset
- [ ] Empirical false positives
- [ ] Insertion attempts versus distinct inserts
- [ ] Saturation
- [ ] Named filter deletion races
- [ ] Batch limits
- [ ] Docker/Compose metrics stack

## Deliverables and acceptance

- [ ] Record pinned versions, prerequisites, resource needs, credentials/roles, and compatibility decisions.
- [ ] Add minimal runnable config/code plus synthetic fixtures; document exact setup and health verification commands.
- [ ] Demonstrate the goal end to end with assertions or integration tests.
- [ ] Inject a relevant failure from the checklist; save sanitized symptoms, diagnosis, fix, and recovery evidence.
- [ ] Verify persistence/replay/cleanup behavior as applicable; record measured latency, memory, lag, or cost where relevant.
- [ ] Document teardown, including persistent volumes/cloud resources; explain what teardown preserves or deletes.
- [ ] Explain why this technology was used on your resume and when a simpler alternative would work.

## Lab notebook

| Attempt/date | Versions/environment | Command or experiment | Expected | Observed | Evidence/next action |
| --- | --- | --- | --- | --- | --- |

Access status: not started. Record local, mocked, real-service, and account-blocked work distinctly. Consult the current official documentation before selecting commands or settings; linked entry points are not tested configurations.
