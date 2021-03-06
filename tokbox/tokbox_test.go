package tokbox

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testKey    = "45822722"
	testSecret = "362f8bfbb5fff2f960c72ee4b798fa7029f9a601"
)

func TestTokboxCreateSession(t *testing.T) {
	tok := New(testKey, testSecret)
	session, err := tok.CreateSession()
	if err != nil {
		t.Errorf("expected nil err, but got %q", err)
	}
	if session.ID == "" {
		t.Errorf("expected non-empty id")
	}
	if session.ProjectID == "" {
		t.Errorf("expected non-empty project-id")
	}

}

func TestTokboxJwtToken(t *testing.T) {
	tok := New(testKey, testSecret)
	token1, err := tok.jwtToken()
	if err != nil {
		t.Errorf("expected nil error, got %q", err)
	}
	token2, err := tok.jwtToken()
	if err != nil {
		t.Errorf("expected nil error, got %q", err)
	}
	if token1 == token2 {
		t.Errorf("expected unique token, got %q, %q", token1, token2)
	}
}

func TestTokboxArchive(t *testing.T) {
	tok := New(testKey, testSecret)
	session, err := tok.CreateSession()
	require.NoError(t, err)
	archives, err := tok.Archives(session.ID)
	require.NoError(t, err)
	fmt.Println(archives)
}
