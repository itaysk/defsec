package specs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSpecs(t *testing.T) {
	tests := []struct {
		name         string
		specName     string
		wantSpecPath string
	}{
		{name: "nsa spec", specName: "nsa", wantSpecPath: "./compliance/nsa-1.0.yaml"},
		{name: "awscis1.2", specName: "awscis1.2", wantSpecPath: "./compliance/aws-cis-1.2.yaml"},
		{name: "awscis1.4", specName: "awscis1.4", wantSpecPath: "./compliance/aws-cis-1.4.yaml"},
		{name: "awscis1.2 by filepath", specName: "@./compliance/aws-cis-1.2.yaml", wantSpecPath: "./compliance/aws-cis-1.2.yaml"},
		{name: "bogus spec", specName: "foobarbaz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantSpecPath != "" {
				wantSpecData, err := os.ReadFile(tt.wantSpecPath)
				assert.NoError(t, err)
				gotSpecData := GetSpec(tt.specName)
				assert.Equal(t, gotSpecData, string(wantSpecData))
			} else {
				assert.Empty(t, GetSpec(tt.specName), tt.name)
			}
		})
	}
}
