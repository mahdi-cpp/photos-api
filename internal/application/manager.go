package application

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/photos-api/internal/account"
)

type AppManager struct {
	mu              sync.RWMutex
	accountManagers map[uuid.UUID]*account.Manager //key is userID
}

func New() (*AppManager, error) {

	manager := &AppManager{
		accountManagers: make(map[uuid.UUID]*account.Manager),
	}

	return manager, nil
}

func (m *AppManager) GetAccountManager(userID uuid.UUID) (*account.Manager, error) {

	m.mu.Lock()
	defer m.mu.Unlock()

	if userID == uuid.Nil {
		return nil, fmt.Errorf("account is nil")
	}

	if accountManager, exists := m.accountManagers[userID]; exists {
		return accountManager, nil
	}

	accountManager, err := account.New(userID)
	if err != nil {
		return nil, err
	}
	m.accountManagers[userID] = accountManager
	return accountManager, err
}
