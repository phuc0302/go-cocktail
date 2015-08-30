package oauth2

import "testing"

func TestGenerateToken(t *testing.T) {
	hash := GenerateToken()
	t.Error(hash)
}
