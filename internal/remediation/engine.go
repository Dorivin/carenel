// Package remediation correlates kernel telemetry with a knowledge base
// of best-practice sysctl configurations and produces risk-scored suggestions.
package remediation

// RiskLevel represents the severity of a misconfiguration.
type RiskLevel string

const (
	RiskHigh   RiskLevel = "HIGH"
	RiskMedium RiskLevel = "MEDIUM"
	RiskLow    RiskLevel = "LOW"
)

// Finding represents a detected kernel parameter issue.
type Finding struct {
	Key            string
	CurrentVal     string
	RecommendedVal string
	Risk           RiskLevel
	Reason         string
	Docs           string
}

// Rule defines a single kernel parameter best-practice check.
type Rule struct {
	Key      string
	Evaluate func(current string) *Finding
}

// Engine holds the knowledge base and produces findings.
type Engine struct {
	rules []Rule
}

// New creates a remediation Engine with built-in rules.
func New() *Engine {
	return &Engine{rules: defaultRules()}
}

// Analyze runs all rules against the provided sysctl snapshot and returns findings.
func (e *Engine) Analyze(snapshot map[string]string) []Finding {
	var findings []Finding
	for _, rule := range e.rules {
		val, ok := snapshot[rule.Key]
		if !ok {
			continue
		}
		if f := rule.Evaluate(val); f != nil {
			findings = append(findings, *f)
		}
	}
	return findings
}

func defaultRules() []Rule {
	// TODO: expand with full knowledge base
	return []Rule{
		{
			Key: "net.core.somaxconn",
			Evaluate: func(v string) *Finding {
				if v == "128" {
					return &Finding{
						Key:            "net.core.somaxconn",
						CurrentVal:     v,
						RecommendedVal: "65535",
						Risk:           RiskHigh,
						Reason:         "Default value severely limits TCP listen backlog under high concurrency workloads.",
						Docs:           "https://www.kernel.org/doc/html/latest/networking/ip-sysctl.html",
					}
				}
				return nil
			},
		},
		{
			Key: "vm.swappiness",
			Evaluate: func(v string) *Finding {
				if v == "60" {
					return &Finding{
						Key:            "vm.swappiness",
						CurrentVal:     v,
						RecommendedVal: "10",
						Risk:           RiskMedium,
						Reason:         "High swappiness causes excessive swap usage on workload nodes, increasing latency.",
						Docs:           "https://www.kernel.org/doc/html/latest/admin-guide/sysctl/vm.html",
					}
				}
				return nil
			},
		},
	}
}
