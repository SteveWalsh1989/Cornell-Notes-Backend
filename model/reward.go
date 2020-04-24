package models

// Reward : Sample struct for reward with type and score
type Reward struct {
	RewardType string `db:"type" json:"type"`
	Score      int    `db:"score" json:"score"`
	Added      int    `db:"added" json:"added"`
}
