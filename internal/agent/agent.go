// Package agent runs on each Kubernetes node as a DaemonSet.
// It attaches eBPF probes and streams kernel telemetry to the control plane.
package agent

import "go.uber.org/zap"

// Agent represents a Carenel node agent instance.
type Agent struct {
	NodeName string
	Logger   *zap.Logger
}

// New creates a new Agent for the given node.
func New(nodeName string, logger *zap.Logger) *Agent {
	return &Agent{NodeName: nodeName, Logger: logger}
}

// Run starts the agent: attaches eBPF probes and begins streaming metrics.
func (a *Agent) Run() error {
	a.Logger.Info("carenel agent starting", zap.String("node", a.NodeName))
	// TODO: attach eBPF programs via cilium/ebpf
	// TODO: start sysctl state poller
	// TODO: open gRPC stream to control plane
	return nil
}
