# Resume technology labs

This track follows the LLD curriculum and covers the technologies and operational claims in your supplied resume. It is a practice plan, not installed infrastructure or completed lab solutions. Follow the numbered groups; run one small stack at a time. You can also do one lab per week alongside LLD.

## Breadth pass scope

The [study plan](../STUDY_PLAN.md) groups these **48 labs into 8 resume review bundles**. Use one representative demo per bundle and explain related technologies; all 48 full lab builds are optional later depth. The LLD track separately has 30 focus and 58 optional modules. Mark a full lab complete only when its own checks pass, not merely because its review bundle is done.

## Practice loop

1. Read the linked official setup reference and select compatible, pinned versions. Record OS, runtime, images, plugins, client libraries, ports, resources, and required accounts. Avoid copying old tutorial settings across image distributions.
2. Create the minimal runnable setup yourself: Compose/config, migrations/seed data, and a small client. Use Python, Go, or TypeScript where appropriate; Celery/Django/PyInstaller are Python-specific and frontend work uses TypeScript/JavaScript. Do not force three-language rewrites of infrastructure.
3. Verify the happy path from a fresh start. Then inject the listed failures, inspect logs/metrics/state, recover, and compare expected versus observed behavior.
4. Write exact start, verify, restart, backup/restore, and teardown commands plus troubleshooting notes in the lab folder. A running container alone is not completion.
5. Explain a five-minute architecture walkthrough, one ten-minute incident, and the tradeoff against an alternative. Repeat setup without following your own steps once.

## Completion evidence for every lab

- [ ] Reproducible setup and version/compatibility notes committed; secrets excluded.
- [ ] Small end-to-end demo with assertions or integration tests.
- [ ] One injected failure, diagnostic evidence, and recovery verified.
- [ ] Resource/cost observations and teardown verified.
- [ ] Explain configuration, guarantees, limits, and resume relevance without AI assistance.

Cloud labs use a separate sandbox, synthetic data, budget alerts, and documented teardown. Mocks demonstrate application logic, not real IAM, AWS delivery, signing trust, or RDS failover. Mark local-only and real-cloud passes separately. Windows packaging and Active Directory require an isolated Windows VM or appropriate test host. Never import customer data or company secrets.

Plan approximately 12–20+ weeks at 6–10 hours/week for a first pass; capstones and unfamiliar enterprise environments can take longer. Prioritize groups 01–03, then the technologies you will defend in your next interview. Account-gated exercises remain pending if you have no access.

## Ordered lab checklist

### 01. Local environment and languages

- [ ] [01. Docker and Docker Compose](01_local_environment/01_docker_compose/README.md)
- [ ] [02. Bash and Linux operations](01_local_environment/02_bash_linux/README.md)
- [ ] [03. Python, Go, TypeScript, JavaScript, C/C++ and SQL](01_local_environment/03_language_runtime/README.md)

### 02. Databases and cache operations

- [ ] [04. PostgreSQL setup and SQL](02_databases/01_postgresql/README.md)
- [ ] [05. SQLite and Redis-free KV cache](02_databases/02_sqlite_kv/README.md)
- [ ] [06. Redis setup and caching](02_databases/03_redis/README.md)
- [ ] [07. MongoDB setup and document modeling](02_databases/04_mongodb/README.md)
- [ ] [08. Database migrations and large backfills](02_databases/05_migrations/README.md)

### 03. Messaging, workers, and CDC

- [ ] [09. RabbitMQ setup and delivery](03_messaging_cdc/01_rabbitmq/README.md)
- [ ] [10. Celery, starvation, and memory](03_messaging_cdc/02_celery/README.md)
- [ ] [11. Kafka setup and consumer behavior](03_messaging_cdc/03_kafka/README.md)
- [ ] [12. Debezium PostgreSQL CDC](03_messaging_cdc/04_debezium/README.md)
- [ ] [13. Transactional outbox and CDC integration](03_messaging_cdc/05_outbox_cdc/README.md)
- [ ] [14. Kafka, RabbitMQ, Redis, Celery comparison](03_messaging_cdc/06_broker_comparison/README.md)

### 04. Backend frameworks and enterprise integrations

- [ ] [15. FastAPI](04_backend_connectors/01_fastapi/README.md)
- [ ] [16. Django and Django REST Framework](04_backend_connectors/02_django_drf/README.md)
- [ ] [17. Node.js and Express.js](04_backend_connectors/03_node_express/README.md)
- [ ] [18. GraphQL](04_backend_connectors/04_graphql/README.md)
- [ ] [19. OAuth and API connector lifecycle](04_backend_connectors/05_oauth_connectors/README.md)
- [ ] [20. DocuSign, Google, Okta, Notion, Airtable, Freshservice, ClickUp](04_backend_connectors/06_vendor_connectors/README.md)
- [ ] [21. Playwright browser integrations](04_backend_connectors/07_playwright/README.md)
- [ ] [22. Active Directory and restricted networks](04_backend_connectors/08_active_directory/README.md)
- [ ] [23. Large-scale streaming synchronization](04_backend_connectors/09_stream_sync/README.md)

### 05. On-prem packaging, updates, and signing

- [ ] [24. PyInstaller and multithreaded clients](05_packaging_delivery/01_pyinstaller/README.md)
- [ ] [25. Checksum, backup, rollback, restart](05_packaging_delivery/02_safe_updater/README.md)
- [ ] [26. Code signing and AWS Signer](05_packaging_delivery/03_code_signing/README.md)

### 06. AWS, Terraform, and CI/CD

- [ ] [27. AWS IAM and Secrets Manager](06_cloud_iac_ci/01_aws_identity/README.md)
- [ ] [28. AWS S3](06_cloud_iac_ci/02_s3/README.md)
- [ ] [29. AWS Lambda and SNS](06_cloud_iac_ci/03_lambda_sns/README.md)
- [ ] [30. AWS EC2 and RDS](06_cloud_iac_ci/04_ec2_rds/README.md)
- [ ] [31. Terraform](06_cloud_iac_ci/05_terraform/README.md)
- [ ] [32. GitHub Actions](06_cloud_iac_ci/06_github_actions/README.md)
- [ ] [33. Jira, Apwide GoLive, and Jenkins](06_cloud_iac_ci/07_jira_jenkins_golive/README.md)

### 07. Frontend, testing, and observability

- [ ] [34. React.js and Svelte](07_frontend_testing_observability/01_react_svelte/README.md)
- [ ] [35. Postman and HTTP contract testing](07_frontend_testing_observability/02_postman/README.md)
- [ ] [36. Sentry](07_frontend_testing_observability/03_sentry/README.md)
- [ ] [37. Prometheus and Grafana](07_frontend_testing_observability/04_prometheus_grafana/README.md)
- [ ] [38. OpenAPI, Makefile, tests, and benchmarks](07_frontend_testing_observability/05_openapi_make_benchmarks/README.md)
- [ ] [39. Bloom filter project defense](07_frontend_testing_observability/06_bloom_service/README.md)

### 08. AI workflows, RAG, and resume defense

- [ ] [40. RAG parsing, chunking, and fuzzy/entity matching](08_ai_and_capstones/01_rag_entity_resolution/README.md)
- [ ] [41. Parallel Search API and entity validation](08_ai_and_capstones/02_parallel_search/README.md)
- [ ] [42. Codex practice](08_ai_and_capstones/03_codex/README.md)
- [ ] [43. Claude Code practice](08_ai_and_capstones/04_claude_code/README.md)
- [ ] [44. GitHub Copilot practice](08_ai_and_capstones/05_github_copilot/README.md)
- [ ] [45. Devin practice](08_ai_and_capstones/06_devin/README.md)
- [ ] [46. Connector event pipeline capstone](08_ai_and_capstones/07_event_pipeline_capstone/README.md)
- [ ] [47. Enterprise on-prem deployment capstone](08_ai_and_capstones/08_onprem_capstone/README.md)
- [ ] [48. Resume claim evidence and mock interviews](08_ai_and_capstones/09_resume_defense/README.md)

## Resume coverage map

| Resume area | Where to practice |
| --- | --- |
| Python, Go, TypeScript, JavaScript, C/C++, SQL, Bash | Group 01 plus the existing LLD language passes |
| PostgreSQL, MongoDB, SQLite, Redis, migrations | Group 02 |
| RabbitMQ, Celery, Kafka, Debezium | Group 03 |
| FastAPI, Django, DRF, Node.js, Express.js, GraphQL | Group 04 |
| OAuth, APIs, DocuSign, Google, Okta, Notion, Airtable, Freshservice, ClickUp | Group 04 connector labs |
| Playwright, Active Directory, large sync streams | Group 04 |
| PyInstaller, multi-threaded clients, safe updates, code signing | Group 05 |
| Lambda, SNS, S3, Signer, EC2, Secrets Manager, RDS | Groups 05–06 |
| Terraform, GitHub Actions, Jira, Apwide GoLive, Jenkins | Group 06 |
| React, Svelte, Postman, Sentry, Prometheus, Grafana | Group 07 |
| Bloom filters, MurmurHash3, double hashing, OpenAPI, Makefile, benchmarks | Group 07 |
| RAG, parsing/chunking, fuzzy/entity matching, Parallel Search | Group 08 |
| Codex, Claude Code, GitHub Copilot, Devin | Group 08, one separate lab per tool |
| Enterprise onboarding, deployment, operational metrics | Group 08 capstones and resume defense |

## Source and verification notes

Official entry points for the core infrastructure were checked on 2026-10-13: [Kafka quickstart](https://kafka.apache.org/quickstart/), [Debezium PostgreSQL connector](https://debezium.io/documentation/reference/stable/connectors/postgresql.html), [RabbitMQ tutorials](https://www.rabbitmq.com/tutorials), [Redis Docker setup](https://redis.io/docs/latest/operate/oss_and_stack/install/install-stack/docker/), [PostgreSQL getting started](https://www.postgresql.org/docs/current/tutorial-start.html). The Celery, PyInstaller, AWS Signer, and Codex official references informed the corresponding exercises. Other official links are learning entry points; every page, account capability, and version combination has not been independently verified.

No services were provisioned and no cloud/account actions were performed while writing this plan. Do not mark a setup exercise complete until its actual setup and failure-recovery checks work in your recorded environment.
