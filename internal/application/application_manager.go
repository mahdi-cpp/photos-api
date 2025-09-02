package application

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type AppManager struct {
	mu                  sync.RWMutex
	Users               map[string]string
	userManagers        map[string]*UserManager // Maps user IDs to their UserManager
	rdb                 *redis.Client
	ctx                 context.Context
	cancel              context.CancelFunc
	wg                  sync.WaitGroup
	initialUserList     chan struct{} // Add this channel
	initialListReceived bool          // Add this flag
}

func NewApplicationManager() (*AppManager, error) {

	ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	manager := &AppManager{
		userManagers: make(map[string]*UserManager),
		Users:        make(map[string]string),
		rdb: redis.NewClient(&redis.Options{
			Addr: "localhost:50001",
			DB:   0,
		}),
		ctx:             ctx,
		cancel:          cancel, // Store the cancel function
		initialUserList: make(chan struct{}),
	}

	// Verify Redis connection with timeout
	ctxPing, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := manager.rdb.Ping(ctxPing).Err(); err != nil {
		return nil, fmt.Errorf("redis connection failed: %w", err)
	}

	//var err error

	//// Initialize account manager
	//.AccountManager, err = account.NewClientManager()
	//if err != nil {
	//	cancel()
	//	return nil, fmt.Errorf("failed to initialize account manager: %w", err)
	//}
	//
	//// Register callback
	//manager.AccountManager.Register(manager.accountCallback)

	// Start account operations in a separate goroutine
	manager.wg.Add(1)
	//go manager.initAccountOperations()
	go manager.runMainSubscription()

	return manager, nil
}

func (m *AppManager) runMainSubscription() {
	pubsub := m.rdb.Subscribe(m.ctx, "upload")
	defer func(pubsub *redis.PubSub) {
		err := pubsub.Close()
		if err != nil {
		}
	}(pubsub)

	// Wait for subscription confirmation
	if _, err := pubsub.Receive(m.ctx); err != nil {
		log.Printf("Failed to receive subscription confirmation: %v", err)
		return
	}

	// Signal that subscription is ready
	//close(m.subReady)

	ch := pubsub.Channel()
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				log.Printf("Main subscription channel closed")
				return
			}
			m.handleMessage(msg)
		case <-m.ctx.Done():
			log.Printf("Main subscription exiting due to context cancellation")
			return
		}
	}
}

func (m *AppManager) handleMessage(msg *redis.Message) {
	switch msg.Channel {
	case "upload":

	default:
		log.Printf("Received message on %s: %s", msg.Channel, msg.Payload)
	}
}

//func (manager *AppManager) fetchUsers() {
//
//	//// Alternative using PHCollection directly
//	//type UserPHCollection struct {
//	//	collection []*person_test.PHCollection[account.User] `json:"collections"`
//	//}
//
//	ac := account.NewAccountManager()
//	Users, err := ac.GetAll()
//	if err != nil {
//		return
//	}
//
//	for _, user := range *Users {
//		fmt.Printf("User ID: %d\n", user.ID)
//		fmt.Printf("Username: %s\n", user.Username)
//		fmt.Printf("Name: %s %s\n", user.FirstName, user.LastName)
//		manager.Users[user.ID] = &user
//	}
//}

func (m *AppManager) StartUpgrade() error {
	return nil
}

func (m *AppManager) GetUserManager(c *gin.Context, userID string) (*UserManager, error) {

	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if userStorage already exists for this user
	//if userManager, exists := m.userManagers[userID]; exists {
	//	return userManager, nil
	//}

	createManager, err := m.CreateManager(userID)
	if err != nil {
		return nil, err
	}

	return createManager, nil
}

func (m *AppManager) CreateManager(userID string) (*UserManager, error) {

	//var user = m.Users[userID]
	if userID == "" {
		fmt.Println("user is nil")
		return nil, fmt.Errorf("user not found")
	}

	userManager, err := NewUserManager(userID)
	if err != nil {
		return nil, err
	}

	m.userManagers[userID] = userManager

	return userManager, err
}

func (m *AppManager) RemoveStorageForUser(userID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if storage, exists := m.userManagers[userID]; exists {
		// Cancel any background operations
		storage.cancelMaintenance()
		// Remove from map
		delete(m.userManagers, userID)
	}
}
