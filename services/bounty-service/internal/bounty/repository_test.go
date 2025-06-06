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

func (mockQueries *MockQueries) GetBounties(ctx context.Context) ([]db.Bounty, error) {
	arguments := mockQueries.Called(ctx)
	return arguments.Get(0).([]db.Bounty), arguments.Error(1)
}

func (mockQueries *MockQueries) GetBountyByID(ctx context.Context, bountyID uuid.UUID) (db.Bounty, error) {
	arguments := mockQueries.Called(ctx, bountyID)
	return arguments.Get(0).(db.Bounty), arguments.Error(1)
}

func (mockQueries *MockQueries) CreateBounty(ctx context.Context, createBountyParams db.CreateBountyParams) error {
	arguments := mockQueries.Called(ctx, createBountyParams)
	return arguments.Error(0)
}

func (mockQueries *MockQueries) UpdateBounty(ctx context.Context, updateBountyParams db.UpdateBountyParams) error {
	arguments := mockQueries.Called(ctx, updateBountyParams)
	return arguments.Error(0)
}

func TestGetBounties_Success(test *testing.T) {
	mockQueries := new(MockQueries)
	dbBounties := []db.Bounty{{ID: uuid.New(), Title: "Sample"}}
	mockQueries.On("GetBounties", mock.Anything).Return(dbBounties, nil)

	repository := bounty.NewDBRepository(mockQueries)
	result, err := repository.GetBounties()

	assert.NoError(test, err)
	assert.Len(test, result, 1)
	assert.Equal(test, "Sample", result[0].Title)
	mockQueries.AssertExpectations(test)
}

func TestGetBountyByID_InvalidUUID(test *testing.T) {
	repository := bounty.NewDBRepository(nil)
	_, err := repository.GetBountyByID("invalid-uuid")
	assert.Error(test, err)
}

func TestGetBountyByID_Success(test *testing.T) {
	mockQueries := new(MockQueries)
	bountyID := uuid.New()
	dbItem := db.Bounty{ID: bountyID, Title: "Test"}
	mockQueries.On("GetBountyByID", mock.Anything, bountyID).Return(dbItem, nil)

	repository := bounty.NewDBRepository(mockQueries)
	result, err := repository.GetBountyByID(bountyID.String())

	assert.NoError(test, err)
	assert.Equal(test, "Test", result.Title)
	mockQueries.AssertExpectations(test)
}

func TestCreateBounty_ErrorFromToDBParams(test *testing.T) {
	repository := bounty.NewDBRepository(nil)
	bountyItem := &bounty.Bounty{}
	// TODO: This will fail if toDBParams returns an error on invalid input
	assert.Error(test, repository.CreateBounty(bountyItem))
}

func TestUpdateBounty_Success(test *testing.T) {
	mockQueries := new(MockQueries)
	bountyItem := &bounty.Bounty{ID: uuid.New().String(), Title: "Updated"}
	params, _ := bounty.ToDBUpdateParams(*bountyItem)
	mockQueries.On("UpdateBounty", mock.Anything, params).Return(nil)

	repository := bounty.NewDBRepository(mockQueries)
	err := repository.UpdateBounty(bountyItem)
	assert.NoError(test, err)
	mockQueries.AssertExpectations(test)
}
