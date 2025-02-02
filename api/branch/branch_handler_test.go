package branch

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/larscom/gitlab-ci-dashboard/model"
	"github.com/stretchr/testify/assert"
)

type MockBranchService struct{}

func (s *MockBranchService) GetBranchesWithLatestPipeline(projectId int) []model.Branch {
	if projectId == 1 {
		return []model.Branch{{Name: "branch-1"}}
	}
	return make([]model.Branch, 0)
}

func TestHandleGetBranchesWithLatestPipelineByProjectId(t *testing.T) {
	app := fiber.New()
	app.Get("/:projectId", NewBranchHandler(&MockBranchService{}).HandleGetBranchesWithLatestPipeline)

	resp, _ := app.Test(httptest.NewRequest("GET", "/1", nil), -1)
	body, _ := io.ReadAll(resp.Body)

	result := make([]model.Branch, 0)
	err := json.Unmarshal(body, &result)
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.Len(t, result, 1)
	assert.Equal(t, result[0].Name, "branch-1")
}

func TestHandleGetBranchesWithLatestPipelineByProjectIdNoMatch(t *testing.T) {
	app := fiber.New()
	app.Get("/:projectId", NewBranchHandler(&MockBranchService{}).HandleGetBranchesWithLatestPipeline)

	resp, _ := app.Test(httptest.NewRequest("GET", "/123", nil), -1)
	body, _ := io.ReadAll(resp.Body)

	result := make([]model.Branch, 0)
	err := json.Unmarshal(body, &result)
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.Len(t, result, 0)
}

func TestHandleGetBranchesWithLatestPipelineBadRequest(t *testing.T) {
	app := fiber.New()
	app.Get("/:projectId", NewBranchHandler(&MockBranchService{}).HandleGetBranchesWithLatestPipeline)

	resp, _ := app.Test(httptest.NewRequest("GET", "/nan", nil), -1)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
