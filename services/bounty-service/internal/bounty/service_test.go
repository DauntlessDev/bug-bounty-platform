package bounty

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mockRepo *MockRepository) GetBounties(ctx context.Context) ([]Bounty, error) {
	args := mockRepo.Called(ctx)
	return args.Get(0).([]Bounty), args.Error(1)
}

func (mockRepo *MockRepository) GetBountyByID(ctx context.Context, bountyID string) (Bounty, error) {
	arguments := mockRepo.Called(ctx, bountyID)
	return arguments.Get(0).(Bounty), arguments.Error(1)
}

func (mockRepo *MockRepository) CreateBounty(ctx context.Context, bountyItem *Bounty) error {
	arguments := mockRepo.Called(ctx, bountyItem)
	return arguments.Error(0)
}

func (mockRepo *MockRepository) UpdateBounty(ctx context.Context, bountyItem *Bounty) error {
	arguments := mockRepo.Called(ctx, bountyItem)
	return arguments.Error(0)
}

func TestService_GetBounties(test *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	expectedBounties := []Bounty{
		{ID: "1", Title: "Test Bounty 1"},
		{ID: "2", Title: "Test Bounty 2"},
	}

	mockRepo.On("GetBounties", mock.Anything).Return(expectedBounties, nil).Once()

	bounties, err := service.GetBounties(context.Background())

	assert.NoError(test, err)
	assert.Equal(test, expectedBounties, bounties)
	mockRepo.AssertExpectations(test)
}

func TestService_GetBountiesBy(test *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	expectedBounty := Bounty{ID: "1", Title: "Test Bounty 1"}
	bountyID := "1"

	mockRepo.On("GetBountyByID", mock.Anything, bountyID).Return(expectedBounty, nil).Once()

	bounty, err := service.GetBountiesBy(context.Background(), bountyID)

	assert.NoError(test, err)
	assert.Equal(test, expectedBounty, bounty)
	mockRepo.AssertExpectations(test)
}

func TestService_GetBountiesBy_NotFound(test *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	bountyID := "non-existent-id"
	expectedError := errors.New("bounty not found")

	mockRepo.On("GetBountyByID", mock.Anything, bountyID).Return(Bounty{}, expectedError).Once()

	_, err := service.GetBountiesBy(context.Background(), bountyID)

	assert.Error(test, err)
	assert.Equal(test, expectedError, err)
	mockRepo.AssertExpectations(test)
}

func TestService_CreateBounty(test *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	newBounty := &Bounty{ID: "3", Title: "New Bounty"}

	mockRepo.On("CreateBounty", mock.Anything, newBounty).Return(nil).Once()

	err := service.CreateBounty(context.Background(), newBounty)

	assert.NoError(test, err)
	mockRepo.AssertExpectations(test)
}

func TestService_CreateBounty_Error(test *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	newBounty := &Bounty{ID: "3", Title: "New Bounty"}
	expectedError := errors.New("failed to create bounty")

	mockRepo.On("CreateBounty", mock.Anything, newBounty).Return(expectedError).Once()

	err := service.CreateBounty(context.Background(), newBounty)

	assert.Error(test, err)
	assert.Equal(test, expectedError, err)
	mockRepo.AssertExpectations(test)
}

func TestService_UpdateBounty(test *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	updatedBounty := &Bounty{ID: "1", Title: "Updated Bounty"}

	mockRepo.On("UpdateBounty", mock.Anything, updatedBounty).Return(nil).Once()

	err := service.UpdateBounty(context.Background(), updatedBounty)

	assert.NoError(test, err)
	mockRepo.AssertExpectations(test)
}

func TestService_UpdateBounty_Error(test *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	updatedBounty := &Bounty{ID: "1", Title: "Updated Bounty"}
	expectedError := errors.New("failed to update bounty")

	mockRepo.On("UpdateBounty", mock.Anything, updatedBounty).Return(expectedError).Once()

	err := service.UpdateBounty(context.Background(), updatedBounty)

	assert.Error(test, err)
	assert.Equal(test, expectedError, err)
	mockRepo.AssertExpectations(test)
}
