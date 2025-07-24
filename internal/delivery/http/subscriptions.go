package http

import (
	"encoding/json"
	"net/http"

	"github.com/Shyyw1e/effective-mobile-subs/internal/usecase"
	"github.com/Shyyw1e/effective-mobile-subs/pkg/logger"
	"gorm.io/gorm"
)

// @Summary Подсчёт суммы подписок за период
// @Tags subscriptions
// @Produce json
// @Param user_id query string false "ID пользователя (UUID)"
// @Param service_name query string false "Название сервиса"
// @Param from query string true "Начало периода (MM-YYYY)"
// @Param to query string true "Конец периода (MM-YYYY)"
// @Success 200 {object} map[string]uint64
// @Failure 400 {string} string "Невалидные параметры"
// @Failure 500 {string} string "Внутренняя ошибка"
// @Router /subscriptions/total-cost [get]
func TotalCostHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	query := r.URL.Query()

	userID := query.Get("user_id")
	service := query.Get("service_name")
	from := query.Get("from")
	to := query.Get("to")

	if from == "" || to == "" {
		http.Error(w, "params 'from' and 'to' are required", http.StatusBadRequest)
		return
	}

	sum, err := usecase.CalculateTotalCost(db, userID, service, from, to)
	if err != nil {
		logger.Log.Errorf("failed to calculate total cost: %v", err)
		http.Error(w, "failed to calculate total cost", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]uint64{
		"total_cost": sum,
	})
}
