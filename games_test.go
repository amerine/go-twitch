package twitch

import (
	"testing"
)

func TestListGames(t *testing.T) {
	c := New("foo")

	_, _, err := c.TopGames()
	if err != nil {
		t.Errorf("Error getting Top Games")
	}
}
