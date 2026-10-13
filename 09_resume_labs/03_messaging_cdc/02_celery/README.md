# 10. Celery, starvation, and memory

[Track and completion rules](../../README.md) · [Official reference](https://docs.celeryq.dev/en/stable/userguide/optimizing.html)

## Goal

Connect Python Celery workers to RabbitMQ and reproduce long jobs delaying short jobs.

## Ordered practice checklist

- [ ] Separate short/long queues
- [ ] Routing
- [ ] Concurrency and prefetch experiments
- [ ] Early/late ack tradeoff
- [ ] Worker crash redelivery
- [ ] Idempotent tasks
- [ ] Time limits
- [ ] Retries and jitter
- [ ] Result expiration
- [ ] Periodic task overlap
- [ ] Bounded chunk processing
- [ ] Measure queue wait and RSS before/after

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
