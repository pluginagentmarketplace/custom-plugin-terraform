---
name: data-engineering
description: Master data engineering with ETL pipelines, data warehousing, Apache Spark, and building scalable data systems. Learn data modeling for analytics and big data tools.
---

# Data Engineering

## Quick Start

```python
# Python data pipeline with Pandas
import pandas as pd
from sqlalchemy import create_engine

# Extract data
df = pd.read_csv('raw_data.csv')

# Transform data
df['date'] = pd.to_datetime(df['date'])
df['amount'] = df['amount'].astype(float)
df = df.dropna()

# Load to database
engine = create_engine('postgresql://user:pass@localhost/warehouse')
df.to_sql('fact_sales', engine, if_exists='append')
```

## Core Competencies

### ETL Concepts
- Extract from various sources (APIs, databases, files)
- Transform data (cleaning, enrichment, aggregation)
- Load to target systems (warehouse, data lake)
- ELT approach (extract, load, transform)
- Real-time streaming alternatives

### Data Warehousing
- Kimball methodology
- Star schema (facts and dimensions)
- Snowflake schema for complex warehouses
- Slowly changing dimensions (SCD)
- Data marts and subject areas

### Tools & Frameworks
- **Apache Spark** - Distributed processing
- **Airflow** - Workflow orchestration
- **dbt** - Data transformation
- **Kafka** - Stream processing
- **Python** - Data processing scripts

### Big Data Technologies
- Hadoop ecosystem basics
- Distributed file systems (HDFS)
- Spark SQL and DataFrames
- RDD and lazy evaluation
- Partitioning and shuffling

### Data Quality
- Validation rules
- Anomaly detection
- Data profiling
- Quality metrics and KPIs
- Testing frameworks

### Cloud Data Platforms
- Snowflake - cloud-native warehouse
- BigQuery - serverless analytics
- Redshift - AWS data warehouse
- Synapse - Azure analytics
- Data lakes (S3, Azure Data Lake)

## Pipeline Patterns

- Batch processing (daily, hourly)
- Near real-time (event-driven)
- Stream processing (continuous)
- Lambda architecture
- Kappa architecture

## Performance Optimization

- Partition strategies
- Compression techniques
- Caching and materialized views
- Query optimization
- Cost optimization

## Real-World Projects

- Customer analytics pipeline
- Real-time monitoring system
- Data warehouse implementation
- Stream processing platform
- ML feature engineering
