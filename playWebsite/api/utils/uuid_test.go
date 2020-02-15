package utils

import (
	"testing"
)

func TestNewUUID(t *testing.T) {
	strUUID, err := NewUUID()
	if err != nil {
		t.Errorf("Error ocurred in NewUUID %v", err)
	}
	t.Logf("UUID is %s", strUUID)
}