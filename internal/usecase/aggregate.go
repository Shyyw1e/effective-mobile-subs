package usecase

import (
	"fmt"
	"time"

	"github.com/Shyyw1e/effective-mobile-subs/internal/db"
	"github.com/Shyyw1e/effective-mobile-subs/pkg/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type TotalCostQuery struct {
	UserID      string
	ServiceName string
	From        string
	To          string
}

func parseMonthYear(s string) (time.Time, error) {
	return time.Parse("01-2006", s)
}

func CalculateTotalCost(database *gorm.DB, userID, service, from, to string) (uint64, error) {
	var subs []*db.Subscription
	
	query := database.Model(&db.Subscription{})

	if userID != "" {
		if _, err := uuid.Parse(userID); err != nil {
			logger.Log.Errorf("invalid user_id: %v", err)
			return 0, fmt.Errorf("invalid user_id: %v", err)
		}
		query = query.Where("user_id = ?", userID)
	}

	if service != "" {
		query = query.Where("service = ?", service)
	}

	if err := query.Find(&subs).Error; err != nil {
		logger.Log.Errorf("failed to find subscriptions: %v", err)
		return 0, err
	}
	
	var sum uint64
	parsedFrom, err := parseMonthYear(from)
	if err != nil {
		logger.Log.Errorf("failed to parse date: %v", err)
		return 0, err
	}
	parsedTo, err := parseMonthYear(to)
	if err != nil {
		logger.Log.Errorf("failed to parse date: %v", err)
		return 0, err
	}
	for _, el := range subs {
		parsedStart, err := parseMonthYear(el.Started_At)
		if err != nil {
			logger.Log.Errorf("failed to parse date: %v", err)
			return 0, err
		}
		var parsedEnd time.Time
		if el.Ends_At != "" {
    	parsedEnd, err = parseMonthYear(el.Ends_At)
    		if err != nil {
        		logger.Log.Errorf("failed to parse end date: %v", err)
        		return 0, err
    		}
		}
		if !parsedStart.After(parsedTo) && (parsedEnd.IsZero() || !parsedEnd.Before(parsedFrom)) {
			sum += el.Price
		}


	}
	return sum, nil

}