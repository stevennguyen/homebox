package repo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Group_Create(t *testing.T) {
	g, err := tRepos.Groups.GroupCreate(context.Background(), "test")

	assert.NoError(t, err)
	assert.Equal(t, "test", g.Name)

	// Get by ID
	foundGroup, err := tRepos.Groups.GroupByID(context.Background(), g.ID)
	assert.NoError(t, err)
	assert.Equal(t, g.ID, foundGroup.ID)
}

func Test_Group_Update(t *testing.T) {
	g, err := tRepos.Groups.GroupCreate(context.Background(), "test")
	assert.NoError(t, err)

	g, err = tRepos.Groups.GroupUpdate(context.Background(), g.ID, GroupUpdate{
		Name:     "test2",
		Currency: "eur",
	})
	assert.NoError(t, err)
	assert.Equal(t, "test2", g.Name)
	assert.Equal(t, "EUR", g.Currency)
}

func Test_Group_GroupStatistics(t *testing.T) {
	useItems(t, 20)
	useLabels(t, 20)

	stats, err := tRepos.Groups.GroupStatistics(context.Background(), tGroup.ID)

	assert.NoError(t, err)
	assert.Equal(t, 20, stats.TotalItems)
	assert.Equal(t, 20, stats.TotalLabels)
	assert.Equal(t, 1, stats.TotalUsers)
	assert.Equal(t, 1, stats.TotalLocations)
}
