<div align="center">

<!-- Animated Typing Banner -->
<img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=600&size=28&duration=3000&pause=1000&color=623CE4&center=true&vCenter=true&multiline=true&repeat=true&width=600&height=100&lines=Terraform+Development+Assistant;8+Agents+%7C+12+Skills;Claude+Code+Plugin" alt="Terraform Assistant" />

<br/>

<!-- Badge Row 1: Status Badges -->
[![Version](https://img.shields.io/badge/Version-2.1.0-blue?style=for-the-badge)](https://github.com/pluginagentmarketplace/custom-plugin-terraform/releases)
[![License](https://img.shields.io/badge/License-Custom-yellow?style=for-the-badge)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Production-brightgreen?style=for-the-badge)](#)
[![SASMP](https://img.shields.io/badge/SASMP-v1.3.0-blueviolet?style=for-the-badge)](#)

<!-- Badge Row 2: Content Badges -->
[![Agents](https://img.shields.io/badge/Agents-8-orange?style=flat-square&logo=robot)](#-agents)
[![Skills](https://img.shields.io/badge/Skills-12-purple?style=flat-square&logo=lightning)](#-skills)
[![Commands](https://img.shields.io/badge/Commands-4-green?style=flat-square&logo=terminal)](#-commands)
[![Terraform](https://img.shields.io/badge/Terraform-1.5+-623CE4?style=flat-square&logo=terraform)](#)

<br/>

<!-- Quick CTA Row -->
[**Install Now**](#-quick-start) · [**Explore Agents**](#-agents) · [**Documentation**](#-skills) · [**Star this repo**](https://github.com/pluginagentmarketplace/custom-plugin-terraform)

---

### What is this?

> **Terraform Development Assistant** is a Claude Code plugin with **8 agents** and **12 skills** for comprehensive Infrastructure as Code development with Terraform.

</div>

---

## Table of Contents

<details>
<summary>Click to expand</summary>

- [Quick Start](#-quick-start)
- [Features](#-features)
- [Agents](#-agents)
- [Skills](#-skills)
- [Commands](#-commands)
- [Contributing](#-contributing)
- [License](#-license)

</details>

---

## Quick Start

### Prerequisites

- Claude Code CLI v2.0.27+
- Active Claude subscription

### Installation

<details open>
<summary><strong>Option 1: From Marketplace (Recommended)</strong></summary>

```bash
# Step 1: Add the marketplace
/plugin add marketplace pluginagentmarketplace/custom-plugin-terraform

# Step 2: Install the plugin
/plugin install custom-plugin-terraform@pluginagentmarketplace-terraform

# Step 3: Restart Claude Code
# Close and reopen your terminal/IDE
```

</details>

<details>
<summary><strong>Option 2: Local Installation</strong></summary>

```bash
# Clone the repository
git clone https://github.com/pluginagentmarketplace/custom-plugin-terraform.git
cd custom-plugin-terraform

# Load locally
/plugin load .

# Restart Claude Code
```

</details>

### Verify Installation

After restart, you should see these agents:

```
custom-plugin-terraform:tf-fundamentals
custom-plugin-terraform:tf-providers
custom-plugin-terraform:tf-resources
custom-plugin-terraform:tf-state
custom-plugin-terraform:tf-modules
custom-plugin-terraform:tf-workspaces
custom-plugin-terraform:tf-cloud
custom-plugin-terraform:tf-advanced
```

---

## Features

- **8 Specialized Agents** covering all Terraform/IaC topics
- **12 Golden Format Skills** with assets, scripts, and references
- **4 Commands** for common Terraform operations
- **SASMP v1.3.0 Compliant** agent-skill architecture
- **Multi-Cloud Support** for AWS, Azure, GCP, and more

---

## Agents

| Agent | Description | Primary Skill |
|-------|-------------|---------------|
| `tf-fundamentals` | HCL syntax and basics | tf-basics |
| `tf-providers` | Multi-cloud provider configuration | tf-providers |
| `tf-resources` | Resource management and lifecycle | tf-resources |
| `tf-state` | State backends and locking | tf-state |
| `tf-modules` | Reusable module development | tf-modules |
| `tf-workspaces` | Environment management | tf-variables |
| `tf-cloud` | Terraform Cloud integration | tf-cloud |
| `tf-advanced` | Functions and patterns | tf-functions |

---

## Skills

### Core Skills
| Skill | Description |
|-------|-------------|
| `tf-basics` | Terraform fundamentals |
| `tf-hcl` | HCL syntax and expressions |

### Configuration Skills
| Skill | Description |
|-------|-------------|
| `tf-providers` | Provider configuration |
| `tf-resources` | Resource blocks |
| `tf-variables` | Input variables |
| `tf-outputs` | Output values |

### State Skills
| Skill | Description |
|-------|-------------|
| `tf-state` | State management |
| `tf-modules` | Module development |

### Advanced Skills
| Skill | Description |
|-------|-------------|
| `tf-functions` | Built-in functions |
| `tf-provisioners` | Provisioners and connections |
| `tf-cloud` | Terraform Cloud |
| `tf-best-practices` | Best practices |

---

## Commands

| Command | Description |
|---------|-------------|
| `/tf-check` | Validate Terraform configuration |
| `/tf-init` | Initialize Terraform directory |
| `/tf-plan` | Preview infrastructure changes |
| `/tf-apply` | Apply infrastructure changes |

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines.

---

## License

Custom License - See [LICENSE](LICENSE) for details.

Copyright (c) 2025 Dr. Umit Kacar & Muhsin Elcicek

---

<div align="center">
  <sub>Built with ULTRATHINK by Claude Code Plugin Team</sub>
</div>
