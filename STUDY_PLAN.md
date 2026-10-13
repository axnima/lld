# Breadth-first study plan

This file controls our teaching order. [README.md](README.md) is the catalogue; each module README holds full requirements; code and revision comments live in its language workspace. Check sessions here as we finish them. A completed session is not a claim that every linked module or language pass is complete.

## Scope and pace

**15 LLD sessions + 8 resume review bundles = 23 planned units.** There are still **30 focus + 58 optional LLD modules** and **48 resume labs** behind them. The 50 LeetCode companions are alternate drills, not an additional to-do list.

A session combines one main implementation with short related experiments or design walkthroughs. Most sessions target roughly 2–4 focused hours; pipeline, API, and persistence sessions may need a second sitting. At one unit per study day, 23 days is the arithmetic minimum, not a deadline. Review, full builds, account access, and language ports add time. Today's language-toolkit work covers session 01's implementation; the recall check remains.

Use Go first. Later port a queue/pool, a service, and a cache to Python and TypeScript, adding runtime-specific exercises where needed. Keep critical validation, concurrency, and failure behavior. Stop after the session's outcome; save specialized variants for depth.

## LLD session checklist

### 01. Language fluency

- [ ] Session outcome demonstrated and explained.

Exercises: [Language toolkit](<01_language_and_design/01_language_toolkit_(focus)/README.md>).

**Do now:** Revise the notification exercise: structs, interfaces, copying, errors, injected clock/IDs, cleanup. Already implemented and tested in Go; do a short recall check.

**Move on when:** Explain the call path and fix a validation/error propagation bug without assistance.

### 02. Design essentials

- [ ] Session outcome demonstrated and explained.

Exercises: [Money, time, and identifiers](<01_language_and_design/03_money_time_identifiers_(focus)/README.md>), [SOLID and composition](<01_language_and_design/04_solid_composition_(focus)/README.md>), [Behavioral patterns](<01_language_and_design/05_behavioral_patterns_(focus)/README.md>).

**Do now:** Use one small pricing/state example for composition, Strategy, and State; briefly compare Observer/Command. Practice integer money and exact expiry boundaries.

**Move on when:** Choose a suitable interface/state model and explain rounding and time assumptions.

### 03. Synchronization

- [ ] Session outcome demonstrated and explained.

Exercises: [Thread-safe counter](<02_concurrency/01_thread_safe_counter_(focus)/README.md>), [Semaphore concurrency limiter](<02_concurrency/05_semaphore_limiter_(focus)/README.md>).

**Do now:** Counter with a mutex, then a semaphore concurrency limiter. Explain atomic versus multi-step updates, ordering handoff, and async versus parallel work.

**Move on when:** Show the exact final count and that in-flight work never exceeds K.

### 04. Queues and deadlocks

- [ ] Session outcome demonstrated and explained.

Exercises: [Bounded blocking queue](<02_concurrency/07_bounded_blocking_queue_(focus)/README.md>), [Deadlock and lock ordering](<02_concurrency/08_deadlock_lock_order_(focus)/README.md>).

**Do now:** Bounded producer-consumer queue, close/drain behavior, and two-account lock ordering. Briefly discuss condition predicates and reader/writer locks.

**Move on when:** Blocked workers exit on close; transfers cannot deadlock or partially apply.

### 05. Worker pool and DAG prerequisites

- [ ] Session outcome demonstrated and explained.

Exercises: [Fixed worker pool](<03_workers_and_pipelines/01_worker_pool_(focus)/README.md>), [Sequential topological sort](<03_workers_and_pipelines/03_topological_sort_(focus)/README.md>).

**Do now:** Build a fixed pool with bounded admission, results, failures, cancellation, and shutdown; implement sequential Kahn traversal as a short prerequisite.

**Move on when:** No worker leak; reject cycles and explain readiness versus completion.

### 06. Parallel dependencies and ordered pipelines

- [ ] Session outcome demonstrated and explained.

Exercises: [Parallel topological scheduler with semaphore](<03_workers_and_pipelines/04_parallel_topological_scheduler_(focus)/README.md>), [Image Transform Pipeline](<03_workers_and_pipelines/08_image_transform_pipeline_(focus)/README.md>).

**Do now:** Implement the semaphore DAG core. Reuse the pool to sketch/test an ordered image pipeline with backpressure; use small fake transforms.

**Move on when:** Dependents wait for completion; tasks run once; output order and memory bounds are explicit.

### 07. Caches

- [ ] Session outcome demonstrated and explained.

Exercises: [LRU cache](<04_caches_and_pools/01_lru_cache_(focus)/README.md>), [TTL cache and authentication tokens](<04_caches_and_pools/03_ttl_cache_(focus)/README.md>).

**Do now:** Implement LRU, then add TTL with a fake clock. Compare LFU and time-based lookup verbally.

**Move on when:** Correct eviction and exact expiry; discuss concurrency and stale cleanup.

### 08. Resource pools

- [ ] Session outcome demonstrated and explained.

Exercises: [Connection and resource pool](<04_caches_and_pools/12_resource_pool_(focus)/README.md>).

**Do now:** Implement acquire/release/close with a limit and timeout; connect it to semaphore ownership.

**Move on when:** Failed work returns resources; double return and shutdown behavior are defined.

### 09. Rate limits and deterministic policies

- [ ] Session outcome demonstrated and explained.

Exercises: [Rate Limiter](<05_backend_services/01_rate_limiter_(focus)/README.md>), [Feature flag](<05_backend_services/09_feature_flags_(focus)/README.md>).

**Do now:** Implement token bucket and fixed window; explain sliding-window tradeoffs. Add a small deterministic flag bucket function.

**Move on when:** Atomic admission, boundary checks, stable assignments; no need for every policy.

### 10. Reliable service API and events

- [ ] Session outcome demonstrated and explained.

Exercises: [Retry-Safe Issue Service](<05_backend_services/06_retry_safe_issue_service_(focus)/README.md>), [Publish/subscribe broker](<05_backend_services/12_pubsub_broker_(focus)/README.md>), [In-memory message queue](<05_backend_services/13_message_queue_(focus)/README.md>).

**Do now:** Build the issue-service core with authorization, idempotency and version checks. Trace queue ack/redelivery and pub/sub fan-out using small in-memory examples.

**Move on when:** Retries do not duplicate effects, updates reject stale versions, delivery semantics are explained.

### 11. State machines and dispatch

- [ ] Session outcome demonstrated and explained.

Exercises: [Vending machine](<06_object_modeling/03_vending_machine_(focus)/README.md>), [Elevator system](<06_object_modeling/02_elevator_(focus)/README.md>).

**Do now:** Implement vending-machine transitions. Sketch elevator states and a simple dispatch policy as a second application.

**Move on when:** Reject illegal transitions and explain failure/refund handling and starvation.

### 12. Money and reservations

- [ ] Session outcome demonstrated and explained.

Exercises: [Expense sharing](<06_object_modeling/04_splitwise_(focus)/README.md>), [Movie and event booking](<06_object_modeling/09_ticket_booking_(focus)/README.md>).

**Do now:** Implement an exact-money split and a seat hold/confirm core. Discuss hotel/library variations instead of rebuilding them.

**Move on when:** Balances reconcile; last-seat contention and hold expiry cannot oversell.

### 13. Debugging state corruption

- [ ] Session outcome demonstrated and explained.

Exercises: [Reservation Service Incident](<07_debugging_and_data/01_reservation_incident_(focus)/README.md>), [Loan Lifecycle & Balance Management](<07_debugging_and_data/02_loan_lifecycle_(focus)/README.md>).

**Do now:** Fix one broken reservation fixture; use a short loan/transfer example to check money conservation and repeat-safe transitions.

**Move on when:** A failing reproducer passes after the root-cause fix; explain safe repair preconditions.

### 14. Persistence, authorization and files

- [ ] Session outcome demonstrated and explained.

Exercises: [Persistent Private Watchlists](<07_debugging_and_data/07_private_watchlists_(focus)/README.md>), [Password Reset](<07_debugging_and_data/08_password_reset_(focus)/README.md>), [File Deduplication](<07_debugging_and_data/09_file_deduplication_(focus)/README.md>).

**Do now:** Use SQLite watchlists for ownership/transaction basics; sketch atomic reset-code consumption. Implement the size/hash/byte-compare dedup core with bounded reads.

**Move on when:** Explain isolation, single-use expiry, collision safety, and partial-failure behavior.

### 15. Timed interview and gap review

- [ ] Session outcome demonstrated and explained.

Exercises: [Timed interview circuit](<08_mocks_and_capstones/06_mock_interviews_(focus)/README.md>).

**Do now:** Pick one unseen 60–90 minute variant from the completed families. Review weak concepts and schedule targeted repeats.

**Move on when:** Working scoped solution, important edge checks, clear tradeoffs, and a written gap list.

## Resume review bundles

Use [the resume lab catalogue](09_resume_labs/README.md) for detailed references. In each bundle, explain every named technology, then run one representative demo focused on your weakest claim. Reuse the same application. Full standalone implementations of all 48 labs are later depth work. An unavailable cloud/Windows account is recorded as pending, never passed by a mock.

- [ ] **16. Runtime and packaging:** Docker/Compose, Bash, language/runtime setup, PyInstaller, Windows lifecycle, checksum/signature distinction and rollback. Demo one packaged client's safe restart/update.
- [ ] **17. Databases and caching:** PostgreSQL, SQLite, Redis, MongoDB, migrations. Demo one transaction plus cache expiry; compare document modeling and backup behavior.
- [ ] **18. Messaging and CDC:** RabbitMQ, Celery, Kafka, Debezium/outbox. Demo a CDC stream or ack/redelivery path first; explain both architectures and record the other setup as a follow-up.
- [ ] **19. Backend and connectors:** FastAPI, Django/DRF, Node/Express, GraphQL, OAuth, vendor APIs, Playwright, Active Directory, streaming sync. Demo one authenticated resumable connector; compare framework-specific behavior and directory access requirements.
- [ ] **20. Cloud and delivery:** Lambda, SNS, S3, Signer, EC2, Secrets Manager, RDS, Terraform, GitHub Actions, Jenkins/Jira/GoLive. Draw the deployment/IAM flow and exercise one sandbox delivery path; real service checks remain explicit.
- [ ] **21. Frontend, quality and observability:** React/Svelte, Postman, Sentry, Prometheus/Grafana, OpenAPI, Makefile. Instrument one API and display it in one UI; inspect a failed request.
- [ ] **22. Project and AI workflows:** Bloom filter mathematics/hashing/benchmarks, RAG chunking, fuzzy/entity matching, Parallel Search; compare Codex/Claude Code/Copilot/Devin using one bounded task and independent checks.
- [ ] **23. Resume defense:** Explain each claim, your contribution, denominators/baselines, one incident and one tradeoff; prioritize any unproven setup skill for a deeper lab.

## Later depth

Choose `_(optional)` folders based on interview gaps or interest. Focus folders also contain optional extensions. Complete the full module checklists here only when pursuing that deeper goal. No projects or solutions have been deleted.
