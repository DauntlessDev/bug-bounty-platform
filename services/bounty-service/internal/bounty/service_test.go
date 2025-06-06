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

func (m *MockRepository) GetBounties() ([]bounty.Bounty, error) {
	args := m.Called()
	return args.Get(0).([]bounty.Bounty), args.Error(1)
}

func (m *MockRepository) GetBountyByID(id string) (bounty.Bounty, error) {
	args := m.Called(id)
	return args.Get(0).(bounty.Bounty), args.Error(1)
}

func (m *MockRepository) CreateBounty(b *bounty.Bounty) error {
	args := m.Called(b)
	return args.Error(0)
}

func (m *MockRepository) UpdateBounty(b *bounty.Bounty) error {
	args := m.Called(b)
	return args.Error(0)
}

func TestService_GetBounties(t *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	expectedBounties := []bounty.Bounty{
		{ID: "1", Title: "Test Bounty 1"},
		{ID: "2", Title: "Test Bounty 2"},
	}

	mockRepo.On("GetBounties").Return(expectedBounties, nil).Once()

	bounties, err := service.GetBounties()

	assert.NoError(t, err)
	assert.Equal(t, expectedBounties, bounties)
	mockRepo.AssertExpectations(t)
}

func TestService_GetBountiesBy(t *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	expectedBounty := bounty.Bounty{ID: "1", Title: "Test Bounty 1"}
	bountyID := "1"

	mockRepo.On("GetBountyByID", bountyID).Return(expectedBounty, nil).Once()

	bounty, err := service.GetBountiesBy(bountyID)

	assert.NoError(t, err)
	assert.Equal(t, expectedBounty, bounty)
	mockRepo.AssertExpectations(t)
}

func TestService_GetBountiesBy_NotFound(t *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	bountyID := "non-existent-id"
	expectedError := errors.New("bounty not found")

	mockRepo.On("GetBountyByID", bountyID).Return(bounty.Bounty{}, expectedError).Once()

	_, err := service.GetBountiesBy(bountyID)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestService_CreateBounty(t *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	newBounty := &bounty.Bounty{ID: "3", Title: "New Bounty"}

	mockRepo.On("CreateBounty", newBounty).Return(nil).Once()

	err := service.CreateBounty(newBounty)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestService_CreateBounty_Error(t *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	newBounty := &bounty.Bounty{ID: "3", Title: "New Bounty"}
	expectedError := errors.New("failed to create bounty")

	mockRepo.On("CreateBounty", newBounty).Return(expectedError).Once()

	err := service.CreateBounty(newBounty)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestService_UpdateBounty(t *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	updatedBounty := &bounty.Bounty{ID: "1", Title: "Updated Bounty"}

	mockRepo.On("UpdateBounty", updatedBounty).Return(nil).Once()

	err := service.UpdateBounty(updatedBounty)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestService_UpdateBounty_Error(t *testing.T) {
	mockRepo := new(MockRepository)
	service := bounty.NewService(mockRepo)

	updatedBounty := &bounty.Bounty{ID: "1", Title: "Updated Bounty"}
	expectedError := errors.New("failed to update bounty")

	mockRepo.On("UpdateBounty", updatedBounty).Return(expectedError).Once()

	err := service.UpdateBounty(updatedBounty)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
