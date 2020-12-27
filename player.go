package cocwrapper

// CocPlayer is a player
type CocPlayer struct {
	Tag                string        `json:"tag"`
	Name               string        `json:"name"`
	THLevel            int           `json:"townHallLevel"`
	Level              int           `json:"expLevel"`
	Trophies           int           `json:"trophies"`
	TopTrophies        int           `json:"bestTrophies"`
	WarStars           int           `json:"warStars"`
	BHLevel            int           `json:"builderHallLevel"`
	VersusTrophies     int           `json:"versusTrophies"`
	BestVersusTrophies int           `json:"bestVersusTrophies"`
	Clan               cocPlayerClan `json:"clan"`
}

type cocPlayerClan struct {
	Tag   string    `json:"tag"`
	Name  string    `json:"name"`
	Level int       `json:"clanLevel"`
	Pic   cocBadges `json:"badgeUrls"`
}

type cocLeague struct {
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Pic  cocBadges `json:"iconUrls"`
}
