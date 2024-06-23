package postgres

import (
	"time"
	"context"

	"bounty-poc/internal/database/postgres"
	"bounty-poc/pkg/models"
)

type Bounty models.Bounty
type BountyService models.BountyService

func (bs *Bounty) CreateBounty() (int, error) {
	now := time.Now()
	ctx := context.Background()
	err := postgres.DbClient.QueryRow(ctx,`INSERT INTO bounty (title, description, reward, category, status, createdat, updatedat, deadline, creatorid, assigneeid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;`,
	bs.Title,
	bs.Description,
	bs.Reward,
        bs.Category,
        bs.Status,
        now, // Assuming 'now' is used for both createdat and updatedat
        now,
        bs.Deadline,
        bs.CreatorID,
        bs.AssigneeID,
    ).Scan(&bs.ID)
	if err != nil {
		return -1, nil
	}
	return bs.ID, nil
}

func (bs *Bounty) GetBountyByID() (Bounty, error) {
	return Bounty{}, nil
}
