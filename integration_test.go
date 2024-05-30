//go:build integration

package notify_test

import (
	"testing"

	"github.com/karapetianash/notifier"
)

func TestSend(t *testing.T) {
	n := notify.New("test title", "test msg", notify.SeverityNormal)

	err := n.Send()

	if err != nil {
		t.Error(err)
	}
}
