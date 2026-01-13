# Sea-TryGo 本地开发环境：可视化面板一览（先看这里）

下面这些组件提供了 Web UI / 可视化管理界面，建议优先收藏：

| 面板/系统 | 访问地址（宿主机） | 默认账号/密码 | 主要用途 |
|---|---|---|---|
| **Grafana** | http://localhost:33000 | `admin / admin` | 统一监控看板（Prometheus 数据源） |
| **Prometheus** | http://localhost:39090 | 无（默认无登录） | 指标查询、Targets 状态、告警规则（如有） |
| **Jaeger** | http://localhost:16686 | 无（默认无登录） | 分布式链路追踪查询 |
| **Kibana** | http://localhost:35601 | 无（当前 ES 关闭安全：`xpack.security.enabled=false`） | Elasticsearch 日志检索与可视化 |
| **Elasticsearch** | http://localhost:39200 | 无（当前关闭安全） | ES API（开发调试、索引管理） |
| **Kafka UI** | http://localhost:38080 | 无（默认无登录） | Kafka Topic / Consumer / Message 管理 |
| **Neo4j Browser** | http://localhost:37474 | `neo4j / Sea-TryGo` | 图数据库查询与管理（Cypher） |
| **Neodash** | http://localhost:35005 | 默认无登录（取决于 Neodash 自身配置） | Neo4j 数据看板（Dashboard） |
| **RedisInsight** | http://localhost:35540 | 首次进入通常需要创建账号（非 compose 固定账号） | Redis 数据浏览、慢查询、性能分析 |
| **MinIO Console** | http://localhost:39001 | `minioadmin / minioadmin` | 对象存储管理控制台 |
| **Flink Web UI** | http://localhost:38081 | 无（默认无登录） | Flink 集群与作业管理 |
| **cAdvisor** | http://localhost:38082 | 无（默认无登录） | 容器资源与性能监控（容器级） |

> 说明
> - **RedisInsight**：官方镜像一般是 Web 首次登录时自行创建账号/密码（存储在 `redisinsight_data` 卷里），因此不是 compose 里能固定写死的那种默认凭据。
> - **Neodash**：是否需要登录/默认密码与镜像版本/部署模式有关，你这份 compose 未配置认证参数，所以按“默认无登录/需看实际初始化页面”处理。

---

> 统一说明：本文所有地址均以 **宿主机访问** 为准（例如 `localhost:33000`）。容器间互访请使用服务名（如 `postgres:5432`、`kafka:9092`）。

---

# 1. 总体架构与组件分组

本 `docker-compose.yml` 主要覆盖以下能力域：

- **服务发现/配置中心**：etcd
- **存储层**：PostgreSQL、Redis、Neo4j、MinIO、Elasticsearch、Milvus（向量库）
- **消息队列**：Kafka（含 Kafka UI）
- **大数据/流处理**：Flink（JobManager/TaskManager）
- **可观测性**：Prometheus + Grafana + 各类 Exporter、Jaeger、node-exporter、cAdvisor、Kibana

默认网络被命名为：`Sea-TryGo`（`networks.default.name`）。

---

# 2. 端口总览（宿主机暴露）

> 表格仅列出对外暴露的端口；未显式声明 `ports` 的 exporter 默认仅在 Compose 网络内可访问。

| 组件 | 宿主机端口 → 容器端口 | 用途 |
|---|---|---|
| etcd | `32379 → 2379` | etcd Client API |
| postgres | `35432 → 5432` | PostgreSQL 连接端口 |
| redis | `36379 → 6379` | Redis 连接端口 |
| redisinsight | `35540 → 5540` | RedisInsight Web UI |
| neo4j | `37474 → 7474` | Neo4j Web UI（Browser） |
| neo4j | `37687 → 7687` | Neo4j Bolt 协议（驱动连接） |
| neo4j | `32004 → 2004` | Neo4j Prometheus metrics（需 Neo4j 配置启用） |
| neodash | `35005 → 5005` | Neodash Web UI |
| kafka | `39092 → 9092` | Kafka 对外端口（注意当前广告地址配置） |
| kafka-ui | `38080 → 8080` | Kafka UI |
| minio | `39000 → 9000` | MinIO S3 API |
| minio | `39001 → 9001` | MinIO Console |
| milvus | `19530 → 19530` | Milvus gRPC（SDK 连接） |
| milvus | `39091 → 9091` | Milvus HTTP（health/metrics 等） |
| elasticsearch | `39200 → 9200` | Elasticsearch HTTP API |
| kibana | `35601 → 5601` | Kibana Web UI |
| flink-jobmanager | `38081 → 8081` | Flink Web UI |
| prometheus | `39090 → 9090` | Prometheus Web UI/API |
| grafana | `33000 → 3000` | Grafana Web UI |
| jaeger | `16686 → 16686` | Jaeger Web UI |
| jaeger | `36831/udp → 6831/udp` | Agent Thrift Compact（旧客户端上报） |
| jaeger | `36832/udp → 6832/udp` | Agent Thrift Binary（旧客户端上报） |
| jaeger | `14268 → 14268` | Collector HTTP |
| jaeger | `14250 → 14250` | Collector gRPC |
| jaeger | `34317 → 4317` | OTLP gRPC（OpenTelemetry） |
| jaeger | `34318 → 4318` | OTLP HTTP（OpenTelemetry） |
| node-exporter | `39100 → 9100` | 主机指标 exporter |
| cadvisor | `38082 → 8080` | 容器指标 exporter/UI |

---

# 3. 组件说明（作用 + 端口 + 关键配置）

## 3.1 etcd（服务发现 & 配置中心）
- **作用**：提供分布式 KV 存储，常用作服务发现/配置中心/一致性协调（本编排中 Milvus 依赖它）。
- **端口**
    - `32379`：对外暴露 etcd Client API（容器内为 `2379`）。
- **关键配置**
    - 数据目录：`./volumes/etcd:/etcd`
    - 自动压缩与保留策略：通过 `ETCD_AUTO_COMPACTION_*` 等环境变量配置。
    - 健康检查：`etcdctl endpoint health`

---

## 3.2 PostgreSQL（关系型数据库）
- **作用**：业务主数据存储、事务支持、SQL 查询。
- **端口**
    - `35432`：宿主机访问 PostgreSQL（容器内为 `5432`）。
- **默认账号（来自 compose 环境变量）**
    - `POSTGRES_USER=admin`
    - `POSTGRES_PASSWORD=Sea-TryGo`
    - `POSTGRES_DB=first_db`

---

## 3.3 Redis（缓存 & KV 存储）
- **作用**：缓存、分布式锁、队列、计数器等。
- **端口**
    - `36379`：宿主机访问 Redis（容器内为 `6379`）。
- **关键配置**
    - 启用 AOF：`redis-server --appendonly yes`

---

## 3.4 RedisInsight（Redis 可视化管理）
- **作用**：Redis 数据浏览、查询、慢查询/性能分析、连接管理。
- **端口**
    - `35540`：Web UI（容器内为 `5540`）。
- **依赖**
    - `depends_on: redis`
- **数据卷**
    - `redisinsight_data:/data`

---

## 3.5 Neo4j（图数据库）
- **作用**：图关系建模与查询（Cypher），适用于关系探索、知识图谱、路径分析等。
- **端口**
    - `37474`：Web UI（容器内 `7474`）
    - `37687`：Bolt 协议（容器内 `7687`，供驱动连接）
    - `32004`：Prometheus metrics（容器内 `2004`，**需要 Neo4j 配置启用**）
- **关键配置**
    - 认证：`NEO4J_AUTH=neo4j/Sea-TryGo`
    - 挂载目录：
        - 数据：`./data/neo4j/data:/data`
        - 日志：`./data/neo4j/logs:/logs`
        - 导入：`./data/neo4j/import:/var/lib/neo4j/import`
        - 插件：`./data/neo4j/plugins:/plugins`
        - 配置：`./data/neo4j/conf:/var/lib/neo4j/conf`

---

## 3.6 Neodash（Neo4j 看板）
- **作用**：基于 Neo4j 的 Dashboard 工具，适合给图数据做可视化看板。
- **端口**
    - `35005`：Web UI（容器内 `5005`）。
- **依赖**
    - `depends_on: neo4j`

---

## 3.7 Kafka（消息队列）
- **作用**：事件流/消息总线，支持发布订阅、消费者组等。
- **端口**
    - `39092`：宿主机访问 Kafka（容器内 `9092`）。
- **关键配置提示（非常重要）**
    - 当前设置：`KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092`
    - 这意味着**容器内**客户端用 `kafka:9092` 没问题；但**宿主机**客户端直接连 `localhost:39092` 时，可能会因为 broker “广告地址”返回 `kafka:9092` 而无法解析（取决于你的客户端环境/DNS）。
    - 如果你希望宿主机客户端也稳定连接，通常需要为宿主机额外配置一个 `ADVERTISED_LISTENERS`（例如 `PLAINTEXT_HOST://localhost:39092`）并做 listeners 映射。

---

## 3.8 Kafka UI（Kafka 可视化管理）
- **作用**：Topic/Partition/Consumer Group 管理、消息浏览、动态配置等。
- **端口**
    - `38080`：Web UI（容器内 `8080`）。
- **依赖**
    - `depends_on: kafka`
- **关键配置**
    - 连接集群：`kafka:9092`

---

## 3.9 MinIO（对象存储）
- **作用**：S3 兼容对象存储（Milvus 依赖它作为对象存储后端）。
- **端口**
    - `39000`：S3 API（容器内 `9000`）
    - `39001`：Console（容器内 `9001`）
- **默认凭据（来自 compose 环境变量）**
    - `MINIO_ACCESS_KEY=minioadmin`
    - `MINIO_SECRET_KEY=minioadmin`
- **数据卷**
    - `./volumes/minio:/minio_data`
- **健康检查**
    - `http://localhost:9000/minio/health/live`

---

## 3.10 Milvus（向量数据库，Standalone）
- **作用**：向量检索/相似度搜索（RAG、Embedding 检索、ANN）。
- **端口**
    - `19530`：Milvus gRPC 主端口（SDK/客户端连接）
    - `39091`：`9091`（healthz 等 HTTP）
- **依赖**
    - `depends
