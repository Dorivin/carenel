// Package sysctl reads and writes Linux kernel parameters via /proc/sys.
// All writes are journaled for audit and GitOps commit generation.
package sysctl

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const procSysBase = "/proc/sys"

// Get reads the current value of a kernel parameter (e.g. "net.core.somaxconn").
func Get(key string) (string, error) {
	path := keyToPath(key)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("sysctl get %s: %w", key, err)
	}
	return strings.TrimSpace(string(data)), nil
}

// Set writes a kernel parameter value.
// NOTE: In production this goes through the GitOps bridge, not a direct write.
func Set(key, value string) error {
	path := keyToPath(key)
	return os.WriteFile(path, []byte(value+"\n"), 0644)
}

// keyToPath converts dot-notation key to /proc/sys path.
// e.g. "net.core.somaxconn" → "/proc/sys/net/core/somaxconn"
func keyToPath(key string) string {
	parts := strings.ReplaceAll(key, ".", "/")
	return filepath.Join(procSysBase, parts)
}
