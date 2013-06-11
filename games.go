package twitch

import (
  "encoding/json"
)

// Represents a Game in the list of top Games from the Twitch API.
type TopGame struct {
  Channels int `json:"channels"`
  Viewers  int `json:"viewers"`
  Game     Game `json:"game"`
}

// Represents
type Game struct {
  Name        string `json:"name"`
  GameBoxes   Art `json:"box"`
  Logos       Art `json:"logo"`
  GiantbombId int `json:"giantbomb_id"`
  Id          int    `json:"_id"`
}

type Art struct {
  Small string `json:"small"`
  Medium string `json:"medium"`
  Large string `json:"large"`
  Template string `json:"template"`
}

func (client *Twitch) TopGames() (games []TopGame, total int, err error) {
  body, err := client.api("GET", "games/top")
  if err != nil {
    return
  }

  var g struct {
    Total int `json:"_total"`
    Games []TopGame `json:"top"`
  }

  err = json.Unmarshal(body, &g)
  if err != nil {
    return
  }

  games, total = g.Games, g.Total
  return
}
