# 06. Redis setup and caching

[Track and completion rules](../../README.md) · [Official reference](https://redis.io/docs/latest/operate/oss_and_stack/install/install-stack/docker/)

## Goal

Start Redis with persistence and use it for cache-aside reads.

## Ordered practice checklist

- [ ] Strings/hashes/sets/sorted sets
- [ ] TTL and exact expiry
- [ ] Eviction under maxmemory
- [ ] RDB versus AOF and restart test
- [ ] Atomic Lua check/update
- [ ] Pipeline versus transaction
- [ ] Cache stampede
- [ ] Invalidation
- [ ] ACL/auth
- [ ] Inspect slowlog and memory
- [ ] Compare Pub/Sub and Streams delivery
- [ ] Lock ownership and lease expiry

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
