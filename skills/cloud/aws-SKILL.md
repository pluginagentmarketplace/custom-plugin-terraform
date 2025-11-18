---
name: aws-mastery
description: Master AWS cloud platform with 200+ services. Learn EC2, S3, Lambda, RDS, VPC, and building enterprise-scale applications on AWS.
---

# AWS Mastery

## Quick Start

```hcl
# Terraform - AWS infrastructure as code
provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "web" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"

  tags = {
    Name = "web-server"
  }
}

resource "aws_s3_bucket" "app_bucket" {
  bucket = "my-app-bucket"
}
```

## Core AWS Services

### Compute
- **EC2** - Virtual machines with fine-grained control
- **Lambda** - Serverless functions, event-driven
- **ECS/EKS** - Container orchestration
- **Elastic Beanstalk** - Managed application platform
- **App Runner** - Managed container runner

### Storage
- **S3** - Object storage, highly scalable
- **EBS** - Block storage for EC2
- **EFS** - Shared file system
- **Glacier** - Long-term archival

### Databases
- **RDS** - Managed relational databases (PostgreSQL, MySQL, etc.)
- **DynamoDB** - NoSQL managed database
- **Aurora** - High-performance managed database
- **ElastiCache** - In-memory caching

### Networking
- **VPC** - Virtual private cloud, network isolation
- **Security Groups** - Stateful firewalls
- **Network ACLs** - Subnet-level firewalls
- **CloudFront** - Global CDN
- **Route 53** - DNS and health checks
- **Elastic Load Balancing** - ALB, NLB, CLB

### Application Services
- **SQS** - Message queue service
- **SNS** - Publish-subscribe messaging
- **Kinesis** - Real-time data streaming
- **EventBridge** - Event routing and processing

### Monitoring & Management
- **CloudWatch** - Metrics, logs, alarms
- **X-Ray** - Distributed tracing
- **CloudTrail** - API logging and compliance
- **Systems Manager** - Parameter store, patch management

## Specializations

### Architecture Patterns
- High availability and disaster recovery
- Multi-region deployments
- Serverless architecture
- Microservices on AWS

### Security & Compliance
- IAM policies and roles
- Encryption strategies
- Network security
- Compliance frameworks

### Cost Optimization
- Reserved instances
- Spot instances
- Storage optimization
- Compute optimization

## Real-World Scenarios

- Hosting web applications
- Building APIs
- Data pipelines
- Real-time analytics
- Machine learning inference

## Essential Tools

- **AWS Management Console** - Web interface
- **AWS CLI** - Command-line interface
- **AWS SDKs** - Programmatic access
- **Terraform** - Infrastructure as code
- **AWS CloudFormation** - Native IaC
