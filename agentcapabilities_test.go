package gocbcore

import "testing"

func TestNoClusterCapabilities(t *testing.T) {
	cfg := loadConfigFromFile(t, "testdata/full_25.json")
	capabilities := globalAgent.buildClusterCapabilities(cfg)
	if capabilities != 0 {
		t.Fatalf("Expected no capabilities to be returned but was %v", capabilities)
	}
}

func TestClusterCapabilitiesEnhancedPreparedStatements(t *testing.T) {
	cfg := loadConfigFromFile(t, "testdata/full_65.json")
	globalAgent.updateClusterCapabilities(cfg)

	if !globalAgent.SupportsClusterCapability(ClusterCapabilityEnhancedPreparedStatements) {
		t.Fatalf("Expected agent to support enhanced prepared statements")
	}
}