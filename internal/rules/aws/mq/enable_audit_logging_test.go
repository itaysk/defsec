package mq

import (
	"testing"

	defsecTypes "github.com/aquasecurity/defsec/pkg/types"

	"github.com/aquasecurity/defsec/pkg/state"

	"github.com/aquasecurity/defsec/pkg/scan"

	"github.com/aquasecurity/defsec/pkg/providers/aws/mq"
	"github.com/stretchr/testify/assert"
)

func TestCheckEnableAuditLogging(t *testing.T) {
	tests := []struct {
		name     string
		input    mq.MQ
		expected bool
	}{
		{
			name: "AWS MQ Broker without audit logging",
			input: mq.MQ{
				Brokers: []mq.Broker{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Logging: mq.Logging{
							Metadata: defsecTypes.NewTestMetadata(),
							Audit:    defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "AWS MQ Broker with audit logging",
			input: mq.MQ{
				Brokers: []mq.Broker{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Logging: mq.Logging{
							Metadata: defsecTypes.NewTestMetadata(),
							Audit:    defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
						},
					},
				},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.AWS.MQ = test.input
			results := CheckEnableAuditLogging.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckEnableAuditLogging.Rule().LongID() {
					found = true
				}
			}
			if test.expected {
				assert.True(t, found, "Rule should have been found")
			} else {
				assert.False(t, found, "Rule should not have been found")
			}
		})
	}
}
