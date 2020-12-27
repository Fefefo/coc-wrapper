package cocwrapper

// CocClans is an array of clans
type CocClans struct {
	Clan []CocClan `json:"items"`
}

// CocClan is a clan
type CocClan struct {
	Tag              string       `json:"tag"`
	Name             string       `json:"name"`
	Type             string       `json:"type"`
	Level            int          `json:"clanLevel"`
	RequiredTrophies int          `json:"requiredTrophies"`
	WarWins          int          `json:"warWins"`
	WarWinStreak     int          `json:"warWinStreak"`
	Members          int          `json:"members"`
	Location         cocLocation  `json:"location"`
	Pic              cocBadges    `json:"badgeUrls"`
	WarLeague        cocWarLeague `json:"warLeague"`
	Label            []cocLabel   `json:"labels"`
}

type cocLocation struct {
	Name        string `json:"name"`
	ID          int    `json:"id"`
	IsCountry   bool   `json:"isCountry"`
	CountryCode string `json:"countryCode"`
}

type cocWarLeague struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type cocLabel struct {
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Pic  cocBadges `json:"iconUrls"`
}

type cocBadges struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}
