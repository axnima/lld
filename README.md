# LLD practice roadmap — Python, Go, TypeScript

An ordered curriculum for object modeling, machine coding, concurrent algorithms, backend APIs, persistence, and debugging. These are practice specifications, not reference solutions. No finite list covers every interview variation; the variants below deliberately cover the recurring failure modes and design choices.

## How to use this repository

1. Follow the numbered phases and problem folders. Complete the core before its variants; extension-heavy exercises can span several sessions.
2. Pick one primary language for a first pass. Go is a useful first choice for the semaphore/DAG interview; Python or TypeScript is equally valid for service modeling. Then implement the same contract idiomatically in the other two languages.
3. Each folder has a prompt, ordered variants, acceptance checks, and `python/`, `go/`, `typescript/` workspaces. Empty workspace markers preserve the folders in Git. A master checkbox means the core, listed variants, language passes, and review are complete; partial progress lives in that problem's README.
4. First implement the core in memory unless persistence is explicitly required. Add a database or HTTP server only in the corresponding variant. Do not add distributed infrastructure to a local interview problem.
5. Keep APIs, invariants, and observable behavior consistent across languages; do not mechanically translate synchronization primitives.

## A realistic schedule

- Breadth pass: about 24 weeks at 8–12 hours/week, one primary language, roughly 3–5 cores/week depending on complexity. This is a planning estimate, not a promise to finish every extension.
- Weeks 1–3: phase 01 and early phase 02. Weeks 4–6: finish phase 02 and phase 03. Weeks 7–10: phase 04. Weeks 11–14: phase 05. Weeks 15–18: phase 06. Weeks 19–22: phase 07. Weeks 23–24: phase 08 and weak-area repeats.
- Depth and portability pass: budget another 12–24+ weeks for variants and the other languages. Completing every variation in all three languages is substantially more than a four-week plan.
- Interview in two weeks? Prioritize language/testing toolkit → semaphore limiter → blocking queue → worker pool → sequential topo → semaphore DAG → rate limiter → idempotency → ticket booking → one timed mock. Return to the full order afterward.

## Session protocol and completion gate

- 5–10 min: clarify scope, API, invariants, time semantics, failure behavior, and shared-state ownership.
- 30–60 min: implement the smallest working core without reading a solution. Use the supplied 45/60/75-minute limits for timed repetitions, not for all extensions combined.
- 15–25 min: test success, invalid input, exact boundaries, failures, and concurrency interleavings; fix the root cause.
- 10 min: explain complexity, lock scope, lifecycle, and tradeoffs. Record the bug you missed and repeat from scratch after 2, 7, and 21 days.
- For concurrency, coordinate tests with barriers/events/channels rather than arbitrary sleeps. Assert invariants (no duplicate execution, no oversell, bounded in-flight work), give tests deadlines, and verify workers stop. Stress tests support reasoning; they do not prove race freedom.
- For a repair exercise, first create a faulty local version or ask for a broken fixture without its solution. Preserve a failing test before fixing it; a fresh correct implementation alone is not debugging practice.

## Language-specific learning checklist

### Python

- [ ] Dataclasses, Protocol/ABC, exceptions, context managers, copying and mutable-default hazards.
- [ ] `threading.Lock`, `RLock`, `Condition`, `Event`, `Barrier`, `Semaphore`, `BoundedSemaphore`; `queue.Queue`.
- [ ] `concurrent.futures` executors, futures, exception propagation, and shutdown.
- [ ] `asyncio` tasks, locks, semaphores, queues, task groups, cancellation, and deadlines; keep blocking I/O off the event loop.
- [ ] Threads for I/O coordination; process workers for CPU parallelism in conventional GIL-enabled CPython. Do not rely on the GIL to protect multi-step invariants; optional free-threaded builds change runtime assumptions.
- [ ] `unittest` or existing pytest; temporary directories; `sqlite3` transactions; monotonic clocks versus UTC timestamps.

### Go

- [ ] Structs, small interfaces, pointer/value semantics, errors, `defer`, maps/slices and aliasing.
- [ ] Goroutines, buffered/unbuffered channels, `select`, `sync.Mutex`, `RWMutex`, `Cond`, `WaitGroup`, `Once`, atomics.
- [ ] Bounded channels as permit semaphores; context cancellation and deadlines; clear channel-close ownership.
- [ ] Race detector (`go test -race ./...` on supported setups), table-driven tests, leak/deadlock checks with deadlines.
- [ ] `net/http`, `httptest`, `database/sql`; a SQLite driver is an explicit dependency when doing SQLite work.

### TypeScript on Node.js

- [ ] Strict types, discriminated unions, interfaces, structural typing, reference aliasing, typed errors, runtime input validation.
- [ ] Promises, `async`/`await`, event-loop interleaving, `AbortController`, cancellation propagation, bounded async admission.
- [ ] Async semaphore/queue implementation with explicit waiter cleanup; a promise limiter coordinates async work but is not shared-memory parallelism.
- [ ] `worker_threads` for CPU parallelism; message ownership first, `SharedArrayBuffer`/`Atomics` as a separate advanced exercise. Do not block the main event loop with blocking waits.
- [ ] `node:test` or an existing runner, deterministic clock injection, HTTP tests, stream backpressure, number precision (`bigint`/validated safe integers for money).
- [ ] Choose and document a supported Node/TypeScript execution setup before coding; native TypeScript handling and SQLite availability depend on the selected runtime.

## Resources and source boundaries

Checked public sources on 2026-10-09. The 15 service prompts supplied in the conversation are labeled **supplied prompt**; their full official statements, hidden tests, and platform availability were not verified. Requirements here are original practice expansions, not claimed copies of official solutions. LeetCode language support and Premium access vary: implement locally when its judge cannot exercise your chosen runtime.

- [LeetCode concurrency collection](https://leetcode.com/problem-list/concurrency/) and [Print in Order](https://leetcode.com/problems/print-in-order/): coordination drills, not a full backend curriculum.
- [Python concurrent execution](https://docs.python.org/3/library/concurrency.html), [threading](https://docs.python.org/3/library/threading.html), [asyncio](https://docs.python.org/3/library/asyncio.html): primary language references.
- [Go concurrency tour](https://go.dev/tour/concurrency/1), [Go memory model](https://go.dev/ref/mem), [race detector](https://go.dev/doc/articles/race_detector): primitives and correctness.
- [Node worker threads](https://nodejs.org/api/worker_threads.html), [Node test runner](https://nodejs.org/api/test.html), [TypeScript handbook](https://www.typescriptlang.org/docs/handbook/intro.html): TypeScript/Node practice references.
- [SQLite transactions](https://www.sqlite.org/lang_transaction.html): persistence and repair exercises.

Use these topic references while implementing; this plan does not depend on Java playlists. Spend more time coding and testing than watching.

## Master checklist

### 01. Language, modeling, and test foundations

- [ ] 001. [Language toolkit](01_foundations/01_language_toolkit/README.md)
- [ ] 002. [Testing toolkit](01_foundations/02_testing_toolkit/README.md)
- [ ] 003. [Money, time, and identifiers](01_foundations/03_money_time_identifiers/README.md)
- [ ] 004. [SOLID and composition](01_foundations/04_solid_composition/README.md)
- [ ] 005. [Behavioral patterns](01_foundations/05_behavioral_patterns/README.md)
- [ ] 006. [Creational and structural patterns](01_foundations/06_creational_structural_patterns/README.md)
- [ ] 007. [Collections and iterators](01_foundations/07_collections_iterators/README.md)
- [ ] 008. [API contracts and validation](01_foundations/08_api_contracts/README.md)

### 02. Concurrency primitives and coordination

- [ ] 009. [Thread-safe counter](02_concurrency_primitives/01_thread_safe_counter/README.md)
- [ ] 010. [Ordered execution](02_concurrency_primitives/02_ordered_execution/README.md)
- [ ] 011. [Alternating workers](02_concurrency_primitives/03_alternating_workers/README.md)
- [ ] 012. [Barrier and molecule assembly](02_concurrency_primitives/04_barrier_h2o/README.md)
- [ ] 013. [Semaphore concurrency limiter](02_concurrency_primitives/05_semaphore_limiter/README.md)
- [ ] 014. [Producer-consumer coordination](02_concurrency_primitives/06_producer_consumer/README.md)
- [ ] 015. [Bounded blocking queue](02_concurrency_primitives/07_bounded_blocking_queue/README.md)
- [ ] 016. [Deadlock and lock ordering](02_concurrency_primitives/08_deadlock_lock_order/README.md)
- [ ] 017. [Readers-writers store](02_concurrency_primitives/09_readers_writers/README.md)
- [ ] 018. [Traffic-light intersection](02_concurrency_primitives/10_traffic_intersection/README.md)
- [ ] 019. [Cancellation and shutdown](02_concurrency_primitives/11_cancellation_shutdown/README.md)

### 03. Executors, dependency graphs, and pipelines

- [ ] 020. [Fixed worker pool](03_executors_graphs/01_worker_pool/README.md)
- [ ] 021. [Futures and completion](03_executors_graphs/02_futures_completion/README.md)
- [ ] 022. [Sequential topological sort](03_executors_graphs/03_topological_sort/README.md)
- [ ] 023. [Parallel topological scheduler with semaphore](03_executors_graphs/04_parallel_topological_scheduler/README.md)
- [ ] 024. [Dependency scheduler extensions](03_executors_graphs/05_dag_scheduler_advanced/README.md)
- [ ] 025. [Delayed and priority scheduler](03_executors_graphs/06_delayed_priority_scheduler/README.md)
- [ ] 026. [Concurrent web crawler](03_executors_graphs/07_web_crawler/README.md)
- [ ] 027. [Image Transform Pipeline](03_executors_graphs/08_image_transform_pipeline/README.md)
- [ ] 028. [Map/reduce and fan-out/fan-in](03_executors_graphs/09_map_reduce_pipeline/README.md)
- [ ] 029. [Batcher and singleflight](03_executors_graphs/10_batcher_singleflight/README.md)

### 04. Data structures and reusable backend components

- [ ] 030. [LRU cache](04_data_components/01_lru_cache/README.md)
- [ ] 031. [LFU cache](04_data_components/02_lfu_cache/README.md)
- [ ] 032. [TTL cache and authentication tokens](04_data_components/03_ttl_cache/README.md)
- [ ] 033. [Time-based and snapshot store](04_data_components/04_versioned_kv_store/README.md)
- [ ] 034. [Hash map and trie](04_data_components/05_hashmap_trie/README.md)
- [ ] 035. [Hit counters and streaming statistics](04_data_components/06_stream_counters/README.md)
- [ ] 036. [Calendar and interval reservations](04_data_components/07_calendar_intervals/README.md)
- [ ] 037. [Memory allocator simulation](04_data_components/08_memory_allocator/README.md)
- [ ] 038. [In-memory filesystem](04_data_components/09_filesystem/README.md)
- [ ] 039. [Text editor and undo/redo](04_data_components/10_editor_undo/README.md)
- [ ] 040. [Leaderboard and indexed rankings](04_data_components/11_leaderboard_rankings/README.md)
- [ ] 041. [Connection and resource pool](04_data_components/12_resource_pool/README.md)

### 05. Reliability, APIs, and backend services

- [ ] 042. [Rate Limiter](05_reliability_services/01_rate_limiter/README.md)
- [ ] 043. [Retry and backoff](05_reliability_services/02_retry_backoff/README.md)
- [ ] 044. [Circuit breaker](05_reliability_services/03_circuit_breaker/README.md)
- [ ] 045. [Idempotency store](05_reliability_services/04_idempotency_store/README.md)
- [ ] 046. [Optimistic concurrency control](05_reliability_services/05_optimistic_concurrency/README.md)
- [ ] 047. [Retry-Safe Issue Service](05_reliability_services/06_retry_safe_issue_service/README.md)
- [ ] 048. [URL Shortener Service](05_reliability_services/07_url_shortener/README.md)
- [ ] 049. [Comments & Replies](05_reliability_services/08_comments_replies/README.md)
- [ ] 050. [Feature flag](05_reliability_services/09_feature_flags/README.md)
- [ ] 051. [Notification service](05_reliability_services/10_notification_service/README.md)
- [ ] 052. [Logging system](05_reliability_services/11_logging_system/README.md)
- [ ] 053. [Publish/subscribe broker](05_reliability_services/12_pubsub_broker/README.md)
- [ ] 054. [In-memory message queue](05_reliability_services/13_message_queue/README.md)
- [ ] 055. [HTTP router and middleware](05_reliability_services/14_http_router_middleware/README.md)
- [ ] 056. [Authorization and audit trail](05_reliability_services/15_authorization_audit/README.md)

### 06. Classical object modeling and machine coding

- [ ] 057. [Parking lot](06_classic_lld/01_parking_lot/README.md)
- [ ] 058. [Elevator system](06_classic_lld/02_elevator/README.md)
- [ ] 059. [Vending machine](06_classic_lld/03_vending_machine/README.md)
- [ ] 060. [Expense sharing](06_classic_lld/04_splitwise/README.md)
- [ ] 061. [Chess and board games](06_classic_lld/05_chess_board_games/README.md)
- [ ] 062. [Snake and ladder](06_classic_lld/06_snake_ladder/README.md)
- [ ] 063. [Library management](06_classic_lld/07_library_management/README.md)
- [ ] 064. [ATM and bank accounts](06_classic_lld/08_atm_bank/README.md)
- [ ] 065. [Movie and event booking](06_classic_lld/09_ticket_booking/README.md)
- [ ] 066. [Hotel and room reservations](06_classic_lld/10_hotel_reservations/README.md)
- [ ] 067. [Food delivery and order workflow](06_classic_lld/11_food_delivery/README.md)
- [ ] 068. [Ride dispatch](06_classic_lld/12_ride_dispatch/README.md)
- [ ] 069. [Cart, coupons, and checkout](06_classic_lld/13_shopping_checkout/README.md)
- [ ] 070. [Auction and order book](06_classic_lld/14_auction_order_book/README.md)

### 07. Production-style debugging, persistence, and data repair

- [ ] 071. [Reservation Service Incident](07_debugging_persistence/01_reservation_incident/README.md)
- [ ] 072. [Loan Lifecycle & Balance Management](07_debugging_persistence/02_loan_lifecycle/README.md)
- [ ] 073. [Return Risk Assessment](07_debugging_persistence/03_return_risk/README.md)
- [ ] 074. [Comment Moderation](07_debugging_persistence/04_comment_moderation/README.md)
- [ ] 075. [Personalized Movie Recommendation](07_debugging_persistence/05_movie_recommendations/README.md)
- [ ] 076. [Keyword and User Blocking](07_debugging_persistence/06_keyword_user_blocking/README.md)
- [ ] 077. [Persistent Private Watchlists](07_debugging_persistence/07_private_watchlists/README.md)
- [ ] 078. [Password Reset](07_debugging_persistence/08_password_reset/README.md)
- [ ] 079. [File Deduplication](07_debugging_persistence/09_file_deduplication/README.md)
- [ ] 080. [Transactions and repositories](07_debugging_persistence/10_sqlite_transactions/README.md)
- [ ] 081. [Write-ahead log and crash recovery](07_debugging_persistence/11_wal_recovery/README.md)
- [ ] 082. [Reconciliation and repair tooling](07_debugging_persistence/12_repair_reconciliation/README.md)

### 08. Advanced integrations and timed capstones

- [ ] 083. [LLM API gateway](08_advanced_capstones/01_llm_gateway/README.md)
- [ ] 084. [Workflow engine](08_advanced_capstones/02_workflow_engine/README.md)
- [ ] 085. [Search and autocomplete service](08_advanced_capstones/03_search_autocomplete/README.md)
- [ ] 086. [Spreadsheet dependency engine](08_advanced_capstones/04_spreadsheet_dependencies/README.md)
- [ ] 087. [Multi-tenant job service](08_advanced_capstones/05_multitenant_job_service/README.md)
- [ ] 088. [Timed interview circuit](08_advanced_capstones/06_mock_interviews/README.md)

## LeetCode companion checklist

These are algorithm/data-structure drills to support the projects above, not replacements for their API, lifecycle, and failure tests. Links identify exercises; try them locally in all three languages if judge support differs.

- [ ] [1114. Print in Order](https://leetcode.com/problems/print-in-order/) — core locally, then relevant concurrency/service extension.
- [ ] [1115. Print FooBar Alternately](https://leetcode.com/problems/print-foobar-alternately/) — core locally, then relevant concurrency/service extension.
- [ ] [1116. Print Zero Even Odd](https://leetcode.com/problems/print-zero-even-odd/) — core locally, then relevant concurrency/service extension.
- [ ] [1195. Fizz Buzz Multithreaded](https://leetcode.com/problems/fizz-buzz-multithreaded/) — core locally, then relevant concurrency/service extension.
- [ ] [1117. Building H2O](https://leetcode.com/problems/building-h2o/) — core locally, then relevant concurrency/service extension.
- [ ] [1188. Design Bounded Blocking Queue](https://leetcode.com/problems/design-bounded-blocking-queue/) — core locally, then relevant concurrency/service extension.
- [ ] [1226. The Dining Philosophers](https://leetcode.com/problems/the-dining-philosophers/) — core locally, then relevant concurrency/service extension.
- [ ] [1279. Traffic Light Controlled Intersection](https://leetcode.com/problems/traffic-light-controlled-intersection/) — core locally, then relevant concurrency/service extension.
- [ ] [1242. Web Crawler Multithreaded](https://leetcode.com/problems/web-crawler-multithreaded/) — core locally, then relevant concurrency/service extension.
- [ ] [207. Course Schedule](https://leetcode.com/problems/course-schedule/) — core locally, then relevant concurrency/service extension.
- [ ] [210. Course Schedule II](https://leetcode.com/problems/course-schedule-ii/) — core locally, then relevant concurrency/service extension.
- [ ] [146. LRU Cache](https://leetcode.com/problems/lru-cache/) — core locally, then relevant concurrency/service extension.
- [ ] [460. LFU Cache](https://leetcode.com/problems/lfu-cache/) — core locally, then relevant concurrency/service extension.
- [ ] [981. Time Based Key-Value Store](https://leetcode.com/problems/time-based-key-value-store/) — core locally, then relevant concurrency/service extension.
- [ ] [1146. Snapshot Array](https://leetcode.com/problems/snapshot-array/) — core locally, then relevant concurrency/service extension.
- [ ] [1797. Design Authentication Manager](https://leetcode.com/problems/design-authentication-manager/) — core locally, then relevant concurrency/service extension.
- [ ] [359. Logger Rate Limiter](https://leetcode.com/problems/logger-rate-limiter/) — core locally, then relevant concurrency/service extension.
- [ ] [362. Design Hit Counter](https://leetcode.com/problems/design-hit-counter/) — core locally, then relevant concurrency/service extension.
- [ ] [933. Number of Recent Calls](https://leetcode.com/problems/number-of-recent-calls/) — core locally, then relevant concurrency/service extension.
- [ ] [705. Design HashSet](https://leetcode.com/problems/design-hashset/) — core locally, then relevant concurrency/service extension.
- [ ] [706. Design HashMap](https://leetcode.com/problems/design-hashmap/) — core locally, then relevant concurrency/service extension.
- [ ] [208. Implement Trie](https://leetcode.com/problems/implement-trie-prefix-tree/) — core locally, then relevant concurrency/service extension.
- [ ] [211. Design Add and Search Words Data Structure](https://leetcode.com/problems/design-add-and-search-words-data-structure/) — core locally, then relevant concurrency/service extension.
- [ ] [729. My Calendar I](https://leetcode.com/problems/my-calendar-i/) — core locally, then relevant concurrency/service extension.
- [ ] [731. My Calendar II](https://leetcode.com/problems/my-calendar-ii/) — core locally, then relevant concurrency/service extension.
- [ ] [732. My Calendar III](https://leetcode.com/problems/my-calendar-iii/) — core locally, then relevant concurrency/service extension.
- [ ] [588. Design In-Memory File System](https://leetcode.com/problems/design-in-memory-file-system/) — core locally, then relevant concurrency/service extension.
- [ ] [1166. Design File System](https://leetcode.com/problems/design-file-system/) — core locally, then relevant concurrency/service extension.
- [ ] [2502. Design Memory Allocator](https://leetcode.com/problems/design-memory-allocator/) — core locally, then relevant concurrency/service extension.
- [ ] [2296. Design a Text Editor](https://leetcode.com/problems/design-a-text-editor/) — core locally, then relevant concurrency/service extension.
- [ ] [1472. Design Browser History](https://leetcode.com/problems/design-browser-history/) — core locally, then relevant concurrency/service extension.
- [ ] [1244. Design A Leaderboard](https://leetcode.com/problems/design-a-leaderboard/) — core locally, then relevant concurrency/service extension.
- [ ] [2034. Stock Price Fluctuation](https://leetcode.com/problems/stock-price-fluctuation/) — core locally, then relevant concurrency/service extension.
- [ ] [2353. Design a Food Rating System](https://leetcode.com/problems/design-a-food-rating-system/) — core locally, then relevant concurrency/service extension.
- [ ] [3408. Design Task Manager](https://leetcode.com/problems/design-task-manager/) — core locally, then relevant concurrency/service extension.
- [ ] [1603. Design Parking System](https://leetcode.com/problems/design-parking-system/) — core locally, then relevant concurrency/service extension.
- [ ] [1845. Seat Reservation Manager](https://leetcode.com/problems/seat-reservation-manager/) — core locally, then relevant concurrency/service extension.
- [ ] [1912. Design Movie Rental System](https://leetcode.com/problems/design-movie-rental-system/) — core locally, then relevant concurrency/service extension.
- [ ] [2286. Booking Concert Tickets in Groups](https://leetcode.com/problems/booking-concert-tickets-in-groups/) — core locally, then relevant concurrency/service extension.
- [ ] [2043. Simple Bank System](https://leetcode.com/problems/simple-bank-system/) — core locally, then relevant concurrency/service extension.
- [ ] [2241. Design an ATM Machine](https://leetcode.com/problems/design-an-atm-machine/) — core locally, then relevant concurrency/service extension.
- [ ] [535. Encode and Decode TinyURL](https://leetcode.com/problems/encode-and-decode-tinyurl/) — core locally, then relevant concurrency/service extension.
- [ ] [355. Design Twitter](https://leetcode.com/problems/design-twitter/) — core locally, then relevant concurrency/service extension.
- [ ] [642. Design Search Autocomplete System](https://leetcode.com/problems/design-search-autocomplete-system/) — core locally, then relevant concurrency/service extension.
- [ ] [631. Design Excel Sum Formula](https://leetcode.com/problems/design-excel-sum-formula/) — core locally, then relevant concurrency/service extension.
- [ ] [348. Design Tic-Tac-Toe](https://leetcode.com/problems/design-tic-tac-toe/) — core locally, then relevant concurrency/service extension.
- [ ] [353. Design Snake Game](https://leetcode.com/problems/design-snake-game/) — core locally, then relevant concurrency/service extension.
- [ ] [622. Design Circular Queue](https://leetcode.com/problems/design-circular-queue/) — core locally, then relevant concurrency/service extension.
- [ ] [641. Design Circular Deque](https://leetcode.com/problems/design-circular-deque/) — core locally, then relevant concurrency/service extension.
- [ ] [295. Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/) — core locally, then relevant concurrency/service extension.

Total: **88 ordered problem folders**, **50 companion LeetCode drills**, and three language workspaces per problem. Original `rate_limiter/` notes are unchanged.

## 09. Resume technology setup and operations

Continue with the [resume technology labs](09_resume_technology_labs/README.md): **48 ordered hands-on labs** covering every named resume technology, setup, application integration, injected failures, recovery, and interview explanation. This is separate from the 88 LLD problems above.

- [ ] [Local environment and languages](09_resume_technology_labs/README.md#01-local-environment-and-languages)
- [ ] [Databases and cache operations](09_resume_technology_labs/README.md#02-databases-and-cache-operations)
- [ ] [Messaging, workers, and CDC](09_resume_technology_labs/README.md#03-messaging-workers-and-cdc)
- [ ] [Backend frameworks and enterprise integrations](09_resume_technology_labs/README.md#04-backend-frameworks-and-enterprise-integrations)
- [ ] [On-prem packaging, updates, and signing](09_resume_technology_labs/README.md#05-on-prem-packaging-updates-and-signing)
- [ ] [AWS, Terraform, and CI/CD](09_resume_technology_labs/README.md#06-aws-terraform-and-cicd)
- [ ] [Frontend, testing, and observability](09_resume_technology_labs/README.md#07-frontend-testing-and-observability)
- [ ] [AI workflows, RAG, and resume defense](09_resume_technology_labs/README.md#08-ai-workflows-rag-and-resume-defense)
