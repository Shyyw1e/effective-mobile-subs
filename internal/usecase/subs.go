package usecase

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/Shyyw1e/effective-mobile-subs/internal/db"
	"github.com/Shyyw1e/effective-mobile-subs/pkg/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type SubscriptionReq struct {
	UserID		string			`json:"user_id"`
	Service		string			`json:"service"`
	Price 		uint64			`json:"price"`

	StartedAt	string			`json:"started_at"`
	EndsAt		string			`json:"ends_at,omitempty"`

}
func CreateSub(database	*gorm.DB, req SubscriptionReq) (*db.Subscription, error) {
	user_id, err := uuid.Parse(req.UserID)
	if err != nil {
		logger.Log.Errorf("invalid user_id: %v", err)
		return nil, err
	}
	
	if req.Price <= 0 {
		logger.Log.Errorf("invalid price: %v", req.Price)
		return nil, fmt.Errorf("invalid price: %v", req.Price)
	}

	if !isValidMonthYear(req.StartedAt) {
		logger.Log.Errorf("invalid sub start date: %v", req.StartedAt)
		return nil, fmt.Errorf("invalid sub start date: %v", req.StartedAt)
	}

	if req.EndsAt != "" {
		if !isValidMonthYear(req.EndsAt) {
			logger.Log.Errorf("invalid sub end date: %v", req.EndsAt)
			return nil, fmt.Errorf("invalid sub end date: %v", req.EndsAt)
		}
	}

	newSub := db.Subscription {
		UserID: user_id,
		Price: req.Price,
		Service: req.Service,
		Started_At: req.StartedAt,
		Ends_At: req.EndsAt,
	}

	if err := database.Create(&newSub).Error; err != nil {
		logger.Log.Errorf("Failed to create subscription: %v", err)
		return nil, err
	}

	logger.Log.Infof("Subscription created: user=%s, service=%s, price=%d", req.UserID, req.Service, req.Price)
	return &newSub, nil
}

func ListSubs(database *gorm.DB, userID string, service string) ([]*db.Subscription, error) {
	var subs []*db.Subscription
	
	query := database.Model(&db.Subscription{})

	if userID != "" {
		if _, err := uuid.Parse(userID); err != nil {
			logger.Log.Errorf("invalid user_id: %v", err)
			return nil, fmt.Errorf("invalid user_id: %v", err)
		}
		query = query.Where("user_id = ?", userID)
	}

	if service != "" {
		query = query.Where("service = ?", service)
	}

	if err := query.Find(&subs).Error; err != nil {
		logger.Log.Errorf("failed to find subscriptions: %v", err)
		return nil, err
	}

	return subs, nil
}

func GetSub(database *gorm.DB, id string) (*db.Subscription, error) {
	var sub db.Subscription
	if err := database.First(&sub, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &sub, nil
}

func UpdateSub(database *gorm.DB, id string, req SubscriptionReq) error {
	var sub db.Subscription
	if err := database.First(&sub, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	sub.Service = req.Service
	sub.Price = req.Price
	sub.Started_At = req.StartedAt
	sub.Ends_At = req.EndsAt

	return database.Save(&sub).Error
}

func PatchSub(database *gorm.DB, id string, patch map[string]interface{}) error {
	result := database.Model(&db.Subscription{}).Where("id = ?", id).Updates(patch)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func DeleteSub(database *gorm.DB, id string) error {
	result := database.Delete(&db.Subscription{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}


func isValidMonthYear(s string) bool {
	re := regexp.MustCompile(`^(0[1-9]|1[0-2])-\d{4}$`)
	return re.MatchString(s)
}