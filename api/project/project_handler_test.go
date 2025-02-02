package project

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/larscom/gitlab-ci-dashboard/model"
	"github.com/stretchr/testify/assert"
)

type MockProjectService struct{}

func (s *MockProjectService) GetProjectsGroupedByStatus(groupId int) map[string][]model.Project {
	if groupId == 1 {
		return map[string][]model.Project{
			"success": {model.Project{Name: "project-1", LatestPipeline: &model.Pipeline{Id: 123}}},
		}
	}

	return make(map[string][]model.Project)
}

func TestHandleGetProjectsGroupedByStatus(t *testing.T) {
	app := fiber.New()

	app.Get("/:groupId", NewProjectHandler(&MockProjectService{}).HandleGetProjectsGroupedByStatus)

	resp, _ := app.Test(httptest.NewRequest("GET", "/1", nil), -1)
	body, _ := io.ReadAll(resp.Body)

	result := make(map[string][]model.Project)
	err := json.Unmarshal(body, &result)
	if err != nil {
		t.Fatal(err.Error())
	}

	success := result["success"]

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.Len(t, success, 1)
	assert.Equal(t, "project-1", success[0].Name)
	assert.Equal(t, 123, success[0].LatestPipeline.Id)
}

func TestHandleGetProjectsGroupedByStatusBadRequest(t *testing.T) {
	app := fiber.New()
	app.Get("/:groupId", NewProjectHandler(&MockProjectService{}).HandleGetProjectsGroupedByStatus)

	resp, _ := app.Test(httptest.NewRequest("GET", "/nan", nil), -1)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
