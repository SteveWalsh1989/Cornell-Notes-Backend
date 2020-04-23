package models

// reward : Sample struct for reward with type and score
type Reward struct {
	rewardType string `db:"type" json:"type"`
	score      int    `db:"score" json:"score"`
}
