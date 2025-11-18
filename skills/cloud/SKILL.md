---
name: cloud-services
description: Master cloud computing fundamentals with AWS, Google Cloud, and Azure. Learn compute, storage, networking, databases, and building scalable cloud applications.
---

# Cloud Services

## Quick Start

```bash
# AWS CLI - Deploy application to EC2
aws ec2 run-instances --image-id ami-0c55b159cbfafe1f0 \
  --instance-type t2.micro --key-name my-key \
  --security-group-ids sg-12345678

# Create S3 bucket for storage
aws s3 mb s3://my-app-bucket

# Create RDS database
aws rds create-db-instance --db-instance-identifier mydb \
  --db-instance-class db.t2.micro \
  --engine postgres
```

## Core Cloud Concepts

### Compute Services
- Virtual machines (EC2, Compute Engine, Azure VMs)
- Managed containers (ECS, GKE, AKS)
- Serverless functions (Lambda, Cloud Functions)
- Managed Kubernetes

### Storage Solutions
- Object storage (S3, Cloud Storage, Blob Storage)
- Block storage (EBS, Persistent Disks)
- File storage (EFS, Cloud Filestore)
- Data warehousing (Redshift, BigQuery)

### Networking
- Virtual Private Cloud (VPC)
- Subnets and routing
- Load balancers
- Content delivery networks (CDN)
- VPN and interconnect

### Managed Databases
- Relational (RDS, Cloud SQL)
- NoSQL (DynamoDB, Firestore)
- Caching (ElastiCache, Memorystore)
- Message queues (SQS, Pub/Sub)

### Security & Identity
- Identity and access management
- Encryption at rest and in transit
- Secrets management
- Network security

## Multi-Cloud Comparison

- AWS (largest market share, most services)
- Google Cloud (strong in data/ML, good pricing)
- Azure (Microsoft ecosystem, enterprise)
- Multi-cloud strategy

## Cost Management

- Instance right-sizing
- Reserved instances
- Spot/preemptible instances
- Monitoring and optimization

## Essential Workflows

- Creating infrastructure
- Deploying applications
- Managing databases
- Monitoring and alerting
- Cost optimization

## Foundational Tools

- **AWS CLI** - Command-line management
- **Cloud Console** - Web UI
- **Terraform** - Infrastructure as code
- **CloudFormation** - AWS infrastructure templates
