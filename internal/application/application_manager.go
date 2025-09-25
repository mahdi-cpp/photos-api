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

	//defaultUserID, err := uuid.Parse("01997cba-6dab-7636-a1f8-2c03174c7b6e")
	//if err != nil {
	//	return nil, err
	//}

	//accountManager, err := manager.GetAccountManager(defaultUserID)
	//if err != nil {
	//	return nil, err
	//}

	//with := &photo.SearchOptions{
	//	Sort:      "createdAt",
	//	SortOrder: "desc",
	//}
	//all, err := accountManager.ReadAll(with)
	//if err != nil {
	//	return nil, err
	//}
	//
	//for _, a := range all {
	//	fmt.Println(a.CreatedAt, a.FileInfo.MimeType)
	//}
	//
	//fmt.Println(len(all))

	return manager, nil
}

func (m *AppManager) GetAccountManager(userID uuid.UUID) (*account.Manager, error) {

	m.mu.Lock()
	defer m.mu.Unlock()

	if userID == uuid.Nil {
		fmt.Println("account is nil")
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
