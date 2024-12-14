# protobq

![CI](https://github.com/averak/protobq/workflows/CI/badge.svg)

`protobq` is a tool designed to simplify and streamline schema management for materialized views in BigQuery.
Instead of managing both the base tables and materialized views separately, developers only need to define the schema of the materialized view.
Based on this schema, `protobq` automatically constructs and maintains the corresponding base table.

This approach ensures consistency between the materialized view and its source data, allowing developers to focus on high-level data modeling without worrying about the complexities of table creation or maintenance.

Key features:

- **Idempotent Schema Management**: Define your materialized view schema declaratively, and let `protobq` handle updates and changes seamlessly.
- **Base Table Automation**: Automatically create and manage the base table from your materialized view schema.
- **BigQuery-Native Optimization**: Leverage BigQuery’s best practices, such as partitioning, clustering, and incremental refreshes, directly through schema definitions.
- **Protocol Buffers Integration**: Use Protocol Buffers to define your schemas, enabling compatibility, extensibility, and multi-language support.

## Philosophy

### Why Protocol Buffers?

- **Schema-First Approach**  
  BigQuery’s schema-driven nature aligns with protobuf, enabling structured and type-safe schema definitions.

- **Versioning and Evolution**  
  Protobuf supports backward and forward compatibility, simplifying schema updates and ensuring long-term maintainability.

- **Seamless BigQuery Integration**  
  BigQuery types map directly to protobuf types (`STRING` → `string`, etc.), ensuring consistency and reducing conversion complexity.

- **Readable and Extensible**  
  Protobuf schemas are both human-readable and machine-readable, aiding collaboration, automation, and extensibility.

### Why Materialized View First?

- **Simplified Architecture**  
  Consolidating data into a unified base table simplifies data pipelines and downstream processes.

- **Query Optimization**  
  Materialized views allow flexible clustering and partitioning, improving query performance for diverse use cases.

- **Cost and Performance Benefits**  
  Precomputed results lower query costs and significantly improve performance for repetitive workloads.

- **Consistency and Reusability**  
  A single base table ensures data integrity and facilitates schema reuse across multiple materialized views.
