package application

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/photos-api/internal/account"
)

type AppManager struct {
	mu              sync.RWMutex
	accountManagers map[uuid.UUID]*account.Manager
}

func New() (*AppManager, error) {

	manager := &AppManager{
		accountManagers: make(map[uuid.UUID]*account.Manager),
	}

	defaultUserID, err := uuid.Parse("018f3a8b-1b32-729b-8f90-1234a5b6c7d8")
	if err != nil {
		return nil, err
	}

	_, err = manager.GetAccountManager(defaultUserID)
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *AppManager) GetAccountManager(userID uuid.UUID) (*account.Manager, error) {

	m.mu.Lock()
	defer m.mu.Unlock()

	if accountManager, exists := m.accountManagers[userID]; exists {
		return accountManager, nil
	}

	if userID == uuid.Nil {
		fmt.Println("account is nil")
		return nil, fmt.Errorf("account is nil")
	}

	accountManager, err := account.New(userID)
	if err != nil {
		return nil, err
	}
	m.accountManagers[userID] = accountManager
	return accountManager, err
}
