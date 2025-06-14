package rest

import (
	"encoding/json"

	"net/http"
	"strconv"
	"time"

	db "bounty-poc/internal/bountyservice/repo/postgres"
	"bounty-poc/internal/bountyservice/service"
)

type BountyRequest struct {
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

var createBounty = service.CreateBounty

func CreateBounty(w http.ResponseWriter, r *http.Request) {

	bounty := &db.Bounty{}
	// Decode the request body directly into the bounty struct
	err := json.NewDecoder(r.Body).Decode(bounty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdID, err := createBounty(bounty)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Write([]byte(strconv.Itoa(createdID)))
}
