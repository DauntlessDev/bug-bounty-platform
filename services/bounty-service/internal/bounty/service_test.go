package bounty_test

import (
	"errors"
	"testing"

	"github.com/DauntlessDev/bug-bounty-platform/services/bounty-service/internal/bounty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mockRepo *MockRepository) GetBounties() ([]bounty.Bounty, error) {
	args := mockRepo.Called()
	return args.Get(0).([]bounty.Bounty), args.Error(1)
}

func (mockRepo *MockRepository) GetBountyByID(bountyID string) (bounty.Bounty, error) {
	arguments := mockRepo.Called(bountyID)
	return arguments.Get(0).(bounty.Bounty), arguments.Error(1)
}

func (mockRepo *MockRepository) CreateBounty(bountyItem *bounty.Bounty) error {
	arguments := mockRepo.Called(bountyItem)
	return arguments.Error(0)
}

func (mockRepo *MockRepository) UpdateBounty(bountyItem *bounty.Bounty) error {
	arguments := mockRepo.Called(bountyItem)
	return arguments.Error(0)
}

func TestService_GetBounties(test *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	expectedBounties := []bounty.Bounty{
		{ID: "1", Title: "Test Bounty 1"},
		{ID: "2", Title: "Test Bounty 2"},
	}

	mockRepo.On("GetBounties").Return(expectedBounties, nil).Once()

	bounties, err := service.GetBounties()

	assert.NoError(test, err)
	assert.Equal(test, expectedBounties, bounties)
	mockRepo.AssertExpectations(test)
}

func TestService_GetBountiesBy(test *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	expectedBounty := bounty.Bounty{ID: "1", Title: "Test Bounty 1"}
	bountyID := "1"

	mockRepo.On("GetBountyByID", bountyID).Return(expectedBounty, nil).Once()

	bounty, err := service.GetBountiesBy(bountyID)

	assert.NoError(test, err)
	assert.Equal(test, expectedBounty, bounty)
	mockRepo.AssertExpectations(test)
}

func TestService_GetBountiesBy_NotFound(test *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	bountyID := "non-existent-id"
	expectedError := errors.New("bounty not found")

	mockRepo.On("GetBountyByID", bountyID).Return(bounty.Bounty{}, expectedError).Once()

	_, err := service.GetBountiesBy(bountyID)

	assert.Error(test, err)
	assert.Equal(test, expectedError, err)
	mockRepo.AssertExpectations(test)
}

func TestService_CreateBounty(test *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	newBounty := &bounty.Bounty{ID: "3", Title: "New Bounty"}

	mockRepo.On("CreateBounty", newBounty).Return(nil).Once()

	err := service.CreateBounty(newBounty)

	assert.NoError(test, err)
	mockRepo.AssertExpectations(test)
}

func TestService_CreateBounty_Error(test *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	newBounty := &bounty.Bounty{ID: "3", Title: "New Bounty"}
	expectedError := errors.New("failed to create bounty")

	mockRepo.On("CreateBounty", newBounty).Return(expectedError).Once()

	err := service.CreateBounty(newBounty)

	assert.Error(test, err)
	assert.Equal(test, expectedError, err)
	mockRepo.AssertExpectations(test)
}

func TestService_UpdateBounty(test *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	updatedBounty := &bounty.Bounty{ID: "1", Title: "Updated Bounty"}

	mockRepo.On("UpdateBounty", updatedBounty).Return(nil).Once()

	err := service.UpdateBounty(updatedBounty)

	assert.NoError(test, err)
	mockRepo.AssertExpectations(test)
}

func TestService_UpdateBounty_Error(test *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	updatedBounty := &bounty.Bounty{ID: "1", Title: "Updated Bounty"}
	expectedError := errors.New("failed to update bounty")

	mockRepo.On("UpdateBounty", updatedBounty).Return(expectedError).Once()

	err := service.UpdateBounty(updatedBounty)

	assert.Error(test, err)
	assert.Equal(test, expectedError, err)
	mockRepo.AssertExpectations(test)
}
