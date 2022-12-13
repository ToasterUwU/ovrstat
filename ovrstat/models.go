package ovrstat

// PlayerStats holds all stats on a specified Overwatch player
type PlayerStats struct {
	Icon             string                     `json:"icon"`
	Name             string                     `json:"name"`
	Level            int                        `json:"level"`
	LevelIcon        string                     `json:"levelIcon"`
	Endorsement      int                        `json:"endorsement"`
	EndorsementIcon  string                     `json:"endorsementIcon"`
	Prestige         int                        `json:"prestige"`
	PrestigeIcon     string                     `json:"prestigeIcon"`
	Ratings          []Rating                   `json:"ratings"`
	GamesWon         int                        `json:"gamesWon"`
	QuickPlayStats   QuickPlayStatsCollection   `json:"quickPlayStats"`
	CompetitiveStats CompetitiveStatsCollection `json:"competitiveStats"`
	Private          bool                       `json:"private"`
}

type Rating struct {
	Group    string `json:"group"`
	Level    int    `json:"level"`
	Role     string `json:"role"`
	RoleIcon string `json:"roleIcon"`
	RankIcon string `json:"rankIcon"`
}

type StatsCollection struct {
	TopHeroes   map[string]*TopHeroStats `json:"topHeroes"`
	CareerStats map[string]*CareerStats  `json:"CareerStats"`
}

type CompetitiveStatsCollection struct {
	Season *int `json:"season"`
	StatsCollection
}

type QuickPlayStatsCollection struct {
	StatsCollection
}

// TopHeroStats holds basic stats for each hero
type TopHeroStats struct {
	TimePlayed          string  `json:"timePlayed"`
	GamesWon            int     `json:"gamesWon"`
	WinPercentage       int     `json:"winPercentage"`
	WeaponAccuracy      int     `json:"weaponAccuracy"`
	EliminationsPerLife float64 `json:"eliminationsPerLife"`
	MultiKillBest       int     `json:"multiKillBest"`
	ObjectiveKills      float64 `json:"objectiveKills"`
}

// CareerStats holds very detailed stats for each hero
type CareerStats struct {
	Assists       map[string]interface{} `json:"assists"`
	Average       map[string]interface{} `json:"average"`
	Best          map[string]interface{} `json:"best"`
	Combat        map[string]interface{} `json:"combat"`
	Deaths        map[string]interface{} `json:"deaths"`
	HeroSpecific  map[string]interface{} `json:"heroSpecific"`
	Game          map[string]interface{} `json:"game"`
	MatchAwards   map[string]interface{} `json:"matchAwards"`
	Miscellaneous map[string]interface{} `json:"miscellaneous"`
}

// Platform represents a response from the search-by-name api request
type Platform struct {
	BattleTag string `json:"battleTag"`
	Portrait  string `json:"portrait"`
	Frame     string `json:"frame"`
	IsPublic  bool   `json:"isPublic"`
}
