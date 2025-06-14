package rest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	db "bounty-poc/internal/bountyservice/repo/postgres"
)

func TestCreateBountySuccess(t *testing.T) {
	// Mock the service function
	oldFn := createBounty
	createBounty = func(b *db.Bounty) (int, error) {
		return 99, nil
	}
	defer func() { createBounty = oldFn }()

	bountyJSON := `{"title":"Test","description":"desc","reward":"stars:2","category":"code","status":"open","deadline":"2024-01-01T00:00:00Z","creator_id":1,"assignee_id":2}`

	req := httptest.NewRequest(http.MethodPost, "/bounty", bytes.NewBufferString(bountyJSON))
	rec := httptest.NewRecorder()

	CreateBounty(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	body := rec.Body.String()
	if body != "99" {
		t.Fatalf("expected body 99, got %s", body)
	}
}
