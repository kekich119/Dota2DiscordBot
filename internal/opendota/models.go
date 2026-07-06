package opendota

type Dotka struct {
	RankTier         int     `json:"rank_tier"`
	LeaderboardRank  int     `json:"leaderboard_rank"`
	ComputedMmr      float64 `json:"computed_mmr"`
	ComputedMmrTurbo float64 `json:"computed_mmr_turbo"`
	Profile          struct {
		AccountID      int    `json:"account_id"`
		Personaname    string `json:"personaname"`
		Name           string `json:"name"`
		Plus           bool   `json:"plus"`
		Cheese         int    `json:"cheese"`
		Steamid        string `json:"steamid"`
		Avatar         string `json:"avatar"`
		Avatarmedium   string `json:"avatarmedium"`
		Avatarfull     string `json:"avatarfull"`
		Profileurl     string `json:"profileurl"`
		LastLogin      string `json:"last_login"`
		Loccountrycode string `json:"loccountrycode"`
		IsContributor  bool   `json:"is_contributor"`
		IsSubscriber   bool   `json:"is_subscriber"`
	} `json:"profile"`
}

type AccountByPersonalName []struct {
	AccountID     int     `json:"account_id"`
	Avatarfull    string  `json:"avatarfull"`
	Personaname   string  `json:"personaname"`
	LastMatchTime string  `json:"last_match_time"`
	Similarity    float64 `json:"similarity"`
}
