package service

import (
	db "bounty-poc/internal/bountyservice/repo/postgres"
)

// CreateBounty persists a new bounty using the repository implementation.
// The previous implementation used a package level interface variable which
// could lead to unexpected behaviour when accessed concurrently. Instead we
// directly call the repository method on the provided bounty instance.
func CreateBounty(b *db.Bounty) (int, error) {
	return b.CreateBounty()
}
