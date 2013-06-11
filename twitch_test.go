package twitch

import "testing"

func TestNew(t *testing.T) {
	clientId := "foo"
	client := New(clientId)

	if client.clientId != clientId {
		t.Errorf("ClientId failed to initialize properly")
	}
}
