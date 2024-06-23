package models

import "time"

type Bounty struct {
	ID          int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Reward      string    `json:"reward"` // stars:2 or coins: 1 or currency:20.00
	Category    string    `json:"category"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Deadline    time.Time `json:"deadline"`
	CreatorID   int       `json:"creator_id"`
	AssigneeID  int       `json:"assignee_id"`
}

type BountyService interface {
	CreateBounty() (int, error)
	GetBountyByID(bountyID int) (Bounty, error)
}
