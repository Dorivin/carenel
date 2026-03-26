// Package gitops generates Git commits for every sysctl change,
// enabling a full audit trail and FluxCD/ArgoCD integration.
package gitops

import "fmt"

// Change represents a single kernel parameter change to be committed.
type Change struct {
	Node   string
	Key    string
	OldVal string
	NewVal string
	Author string
	Reason string
}

// CommitMessage returns a standardised commit message for a kernel parameter change.
func (c Change) CommitMessage() string {
	return fmt.Sprintf("carenel(%s): %s %s → %s\n\nReason: %s\nAuthor: %s",
		c.Node, c.Key, c.OldVal, c.NewVal, c.Reason, c.Author)
}

// Bridge generates and pushes commits for kernel parameter changes.
type Bridge struct {
	RepoURL string
	Branch  string
}

// New creates a GitOps bridge targeting the given repo and branch.
func New(repoURL, branch string) *Bridge {
	return &Bridge{RepoURL: repoURL, Branch: branch}
}

// Commit records a kernel parameter change as a Git commit.
func (b *Bridge) Commit(c Change) error {
	// TODO: clone/pull repo, write sysctl manifest, git commit + push
	fmt.Println(c.CommitMessage())
	return nil
}
