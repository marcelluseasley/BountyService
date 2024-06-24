package service

import (
	db "bounty-poc/internal/bountyservice/repo/postgres"
	//"bounty-poc/pkg/models"
)

var bs db.BountyService
//var bounty db.Bounty


func CreateBounty(b *db.Bounty) (int, error) {
	bs = b 
	id, err := bs.CreateBounty()
	if err != nil {
		return id, err
	}
	return id, nil
}
