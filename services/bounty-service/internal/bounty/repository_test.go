package bounty_test

import (
	"context"
	"testing"

	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/bounty"
	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/db"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQueries struct {
	mock.Mock
}

func (m *MockQueries) GetBounties(ctx context.Context) ([]db.Bounty, error) {
	args := m.Called(ctx)
	return args.Get(0).([]db.Bounty), args.Error(1)
}

func (m *MockQueries) GetBountyByID(ctx context.Context, id uuid.UUID) (db.Bounty, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(db.Bounty), args.Error(1)
}

func (m *MockQueries) CreateBounty(ctx context.Context, arg db.CreateBountyParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockQueries) UpdateBounty(ctx context.Context, arg db.UpdateBountyParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func TestGetBounties_Success(t *testing.T) {
	mockQ := new(MockQueries)
	dbBounties := []db.Bounty{{ID: uuid.New(), Title: "Sample"}}
	mockQ.On("GetBounties", mock.Anything).Return(dbBounties, nil)

	repo := bounty.NewDBRepository(mockQ)
	result, err := repo.GetBounties()

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Sample", result[0].Title)
	mockQ.AssertExpectations(t)
}

func TestGetBountyByID_InvalidUUID(t *testing.T) {
	repo := bounty.NewDBRepository(nil)
	_, err := repo.GetBountyByID("invalid-uuid")
	assert.Error(t, err)
}

func TestGetBountyByID_Success(t *testing.T) {
	mockQ := new(MockQueries)
	id := uuid.New()
	dbItem := db.Bounty{ID: id, Title: "Test"}
	mockQ.On("GetBountyByID", mock.Anything, id).Return(dbItem, nil)

	repo := bounty.NewDBRepository(mockQ)
	result, err := repo.GetBountyByID(id.String())

	assert.NoError(t, err)
	assert.Equal(t, "Test", result.Title)
	mockQ.AssertExpectations(t)
}

func TestCreateBounty_ErrorFromToDBParams(t *testing.T) {
	repo := bounty.NewDBRepository(nil)
	b := &bounty.Bounty{}
	// TODO: This will fail if toDBParams returns an error on invalid input
	assert.Error(t, repo.CreateBounty(b))
}

func TestUpdateBounty_Success(t *testing.T) {
	mockQ := new(MockQueries)
	b := &bounty.Bounty{ID: uuid.New().String(), Title: "Updated"}
	params, _ := bounty.ToDBUpdateParams(*b)
	mockQ.On("UpdateBounty", mock.Anything, params).Return(nil)

	repo := bounty.NewDBRepository(mockQ)
	err := repo.UpdateBounty(b)
	assert.NoError(t, err)
	mockQ.AssertExpectations(t)
}
