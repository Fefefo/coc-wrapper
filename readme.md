# Golang bindings for the Clash of Clans API

A scraping wrapper for Gophers which provides a layer of abstraction for the [Clash of clans API](https://developer.clashofclans.com/#/)


## Installation

`go get github.com/Fefefo/coc-wrapper`


## Project using coc-wrapper

-  [fefegobot](https://github.com/Fefefo/fefegobot) Multifunctional telegram bot written in GO


## Structs

```go
// This is the array of clans
type CocClans struct {
	Clan []CocClan
}

// This is the structure of a clan
type CocClan struct {
	Tag              string
	Name             string
	Type             string
	Level            int
	RequiredTrophies int
	WarWins          int
	WarWinStreak     int
	Members          int
	Location         cocLocation
	Pic              cocBadges
	WarLeague        cocWarLeague
	Label            []cocLabel
}

// This is the structure of a player
type CocPlayer struct {
	Tag                string
	Name               string
	THLevel            int
	Level              int
	Trophies           int
	TopTrophies        int
	WarStars           int
	BHLevel            int
	VersusTrophies     int
	BestVersusTrophies int
	Clan               cocPlayerClan
}

// The struct for require clans
type ReqClans struct {
	Name      string
	MinLevel  int
	MaxLevel  int
	Limit     int
	MinPoints int
}
```
