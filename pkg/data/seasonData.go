package data

// seasons api responses
type SeasonResponse struct {
	Data  []SeasonData `json: "data"`
	Links interface{}  `json: "links,omitempty"`
	Meta  interface{}  `json: "meta,omitempty"`
}

type SeasonData struct {
	Type       string           `json: "type"`
	ID         string           `json: "id"`
	Attributes SeasonAttributes `json: "attributes"`
}

type SeasonAttributes struct {
	IsOffseason     bool `json: "isOffseason"`
	IsCurrentSeason bool `json: "isCurrentSeason"`
}

// playerSeason responses
type PlayerSeasonResponseData struct {
	Data  PlayerSeasonData `json: "data"`
	Links interface{}      `json: "links"`
	Meta  interface{}      `json: "meta"`
}

type PlayerSeasonData struct {
	Type          string                      `json: "type"`
	Attributes    PlayerSeasonAttributes      `json: "attributes"`
	Relationships map[string]RelationshipData `json: "relationships"`
	Links         interface{}                 `json: "links"`
	Meta          interface{}                 `json: "meta"`
}

type PlayerSeasonAttributes struct {
	GameModeStats map[string]GameModeStats `json: "gameModeStats"`
	BestRankPoint int                      `json: "bestRankPoint"`
}

type GameModeStats struct {
	Assits              int     `json: "assists"`
	Boosts              int     `json: "boosts"`
	DBNOs               int     `json: "dBNOs"`
	DailyKills          int     `json: "dailyKills"`
	DailyWins           int     `json: "dailyWins"`
	DamageDealt         float32 `json: "damageDealt"`
	Days                int     `json: "days"`
	HeadshotKills       int     `json: "headshotKills"`
	Heals               int     `json: "heals"`
	Kills               int     `json: "kills"`
	LongestKill         float32 `json: "longestKill"`
	LongestTimeSurvived float32 `json: "longestTimeSurvived"`
	Losses              int     `json: "losses"`
	MaxKillStreaks      int     `json; "maxKillStreaks"`
	MostSurvivalTime    float32 `json; "mostSurvivalTime"`
	RankPoints          float32 `json: "rankPoints"`
	RankPointsTitle     string  `json: "rankPointsTitle"`
	Revives             int     `json: "revives"`
	RideDistance        float32 `json: "rideDistance"`
	RoadKills           int     `json: "roadKills"`
	RoundMostKills      int     `json: "roundMostKills"`
	RoundsPlayed        int     `json: "roundsPlayed"`
	Suicides            int     `json: "suicides"`
	SwimDistance        float32 `json: "swimDistance"`
	TeamKills           int     `json: "teamKills"`
	TimeSurvived        float32 `json: "timeSurvived"`
	Top10s              int     `json: "top10s"`
	VehicleDestroys     int     `json: "vehicleDestroys"`
	WalkDistance        float32 `json: "walkDistance"`
	WeaponsAcquired     int     `json: "weaponsAcquired"`
	WeeklyKills         int     `json: "weeklyKills"`
	WeeklyWins          int     `json: "weeklyWins"`
	Wins                int     `json: "wins"`
}
