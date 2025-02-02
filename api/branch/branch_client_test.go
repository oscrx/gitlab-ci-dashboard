package branch

import (
	"fmt"
	"testing"

	"github.com/larscom/gitlab-ci-dashboard/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetBranchesWith1Page(t *testing.T) {
	const totalPages = 1
	client := NewBranchClient(mock.NewMockGitlabClient(totalPages, nil))

	branches := client.GetBranches(1)

	assert.Len(t, branches, 2)
	assert.Equal(t, "branch-1", branches[0].Name)
	assert.Equal(t, "branch-2", branches[1].Name)
}

func TestGetBranchesWith2Pages(t *testing.T) {
	const totalPages = 2
	client := NewBranchClient(mock.NewMockGitlabClient(totalPages, nil))

	branches := client.GetBranches(1)

	assert.Len(t, branches, 4)
	assert.Equal(t, "branch-1", branches[0].Name)
	assert.Equal(t, "branch-2", branches[1].Name)
	assert.Equal(t, "branch-3", branches[2].Name)
	assert.Equal(t, "branch-4", branches[3].Name)
}

func TestGetBranchesWithErrorEmptySlice(t *testing.T) {
	client := NewBranchClient(mock.NewMockGitlabClient(1, fmt.Errorf("ERROR")))

	branches := client.GetBranches(100)

	assert.Len(t, branches, 0)
}
