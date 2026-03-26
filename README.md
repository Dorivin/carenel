# Carenel 🫀

> **Care for the Kernel** — Linux kernel observability and remediation with Kubernetes-native node intelligence.

[![CI](https://github.com/Dorivin/carenel/actions/workflows/ci.yaml/badge.svg)](https://github.com/Dorivin/carenel/actions)
[![Go Version](https://img.shields.io/badge/go-1.22-00ADD8?logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-AGPL--3.0-green)](LICENSE)

---

## What is Carenel?

Carenel is the first tool that lets you **see** and **act** on your Linux kernel configuration — with:

- **eBPF-powered observability** — deep kernel tracing without patching
- **Node drift detection** — sysctl heatmap across your entire Kubernetes cluster  
- **Intelligent remediation** — correlates anomalies with actionable parameter fixes
- **GitOps-native** — every change is a Git commit, works with FluxCD & ArgoCD
- **Dry-run simulation** — preview impact before applying anything
- **One-click rollback** — full audit trail, instant revert

---

## Architecture

```
┌─────────────────────────────────────────┐
│          Linux Kernel (eBPF)            │
└────────────────┬────────────────────────┘
                 ↓
┌─────────────────────────────────────────┐
│     Carenel Node Agent (DaemonSet)      │  ← reads /proc/sys, attaches eBPF probes
└────────────────┬────────────────────────┘
                 ↓
┌─────────────────────────────────────────┐
│   Control Plane + Remediation Engine    │  ← Go / gRPC / Kubernetes CRDs
└────────────────┬────────────────────────┘
                 ↓
┌─────────────────────────────────────────┐
│         GitOps Bridge                   │  ← FluxCD / ArgoCD / Git
└────────────────┬────────────────────────┘
                 ↓
┌─────────────────────────────────────────┐
│           Carenel UI                    │  ← Web dashboard
└─────────────────────────────────────────┘
```

---

## Quick Start

```bash
# Install
go install github.com/Dorivin/carenel/cmd/carenel@latest

# Scan a node
carenel scan --node node-01

# Preview a remediation plan (dry-run)
carenel apply --dry-run

# Show kernel parameter drift across all cluster nodes
carenel diff
```

---

## Project Structure

```
carenel/
├── cmd/carenel/          # CLI entrypoint
├── internal/
│   ├── agent/            # Node agent (DaemonSet)
│   ├── ebpf/             # eBPF program management
│   ├── sysctl/           # /proc/sys read/write
│   ├── remediation/      # Knowledge base & rule engine
│   ├── gitops/           # Git commit generation
│   └── api/              # gRPC / REST API
├── pkg/
│   ├── kernel/           # Kernel version & capability detection
│   └── k8s/              # Kubernetes client helpers
├── deploy/k8s/           # DaemonSet, RBAC, Namespace manifests
├── web/                  # Landing page & dashboard
└── docs/                 # Architecture & contribution docs
```

---

## Deploy to Kubernetes

```bash
kubectl apply -f deploy/k8s/daemonset.yaml
```

> ⚠️ The agent requires `privileged: true` for eBPF probe attachment.

---

## Roadmap

- [ ] eBPF probe attachment via `cilium/ebpf`
- [ ] sysctl snapshot & drift detection
- [ ] Remediation rule knowledge base
- [ ] GitOps commit bridge (FluxCD / ArgoCD)
- [ ] Web dashboard
- [ ] Workload-aware tuning profiles (Kafka, vLLM, Postgres...)
- [ ] Dry-run simulation engine
- [ ] Slack / PagerDuty alerting

---

## Contributing

PRs and issues are very welcome. See [docs/CONTRIBUTING.md](docs/CONTRIBUTING.md).

---

## License

Apache 2.0 — see [LICENSE](LICENSE).
