package types

import (
	"fmt"
	"time"
)

type FarmingStatusResponse struct {
	Balance                string `json:"balance"`
	ActiveFarmingStartedAt string `json:"activeFarmingStartedAt"`
	FarmingDurationInSec   int    `json:"farmingDurationInSec"`
	FarmingReward          int    `json:"farmingReward"`
}

type CheckTaskItem struct {
	Id             string      `json:"id"`
	Title          string      `json:"title"`
	Type           string      `json:"type"`
	Description    string      `json:"description"`
	Reward         int         `json:"reward"`
	TwitterAuthUrl interface{} `json:"twitterAuthUrl"`
	Url            string      `json:"url,omitempty"`
	ChatId         string      `json:"chatId,omitempty"`
	Submission     struct {
		Reward    int       `json:"reward"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"submission,omitempty"`
}

func (t CheckTaskItem) String() string {
	return fmt.Sprintf("%s:\n   %s (type: %s, reward: %d)", t.Id, t.Title, t.Type, t.Reward)
}
