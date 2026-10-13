# 13. Transactional outbox and CDC integration

[Track and completion rules](../../README.md) · [Official reference](https://debezium.io/documentation/reference/stable/transformations/outbox-event-router.html)

## Goal

Write business state and an outbox row in one PostgreSQL transaction; stream events to Kafka and a materialized view.

## Ordered practice checklist

- [ ] Outbox event identity
- [ ] Ordering per aggregate
- [ ] Debezium routing
- [ ] Duplicate replay
- [ ] Tombstones/deletes
- [ ] Consumer crash
- [ ] Poison event quarantine
- [ ] Redis projection rebuild
- [ ] Compare polling outbox
- [ ] Reconcile source and projection

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
