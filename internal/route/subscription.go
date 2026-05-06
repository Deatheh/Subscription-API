package route

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ashwingopalsamy/uuidcheck"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"subscription/internal/entities"
	utils "subscription/pkg"
)

type DPOSubscription struct {
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserUUID    string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

func (h *Handler) AddSubscription(ctx *gin.Context) {
	var dpo DPOSubscription
	err := ctx.BindJSON(&dpo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting request data: invalid request body"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error getting request data: invalid request body")
		return
	}
	if dpo.ServiceName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting request data: empty service name"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error getting request data: empty service name")
		return
	}
	if dpo.Price == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting request data: empty price"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error getting request data: empty price")
		return
	}
	if dpo.Price < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting request data: price is less than 0"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error getting request data: price is less than 0")
		return
	}
	if dpo.UserUUID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting request data: empty user uuid"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error getting request data: empty user uuid")
		return
	}
	if !uuidcheck.IsValidUUID(dpo.UserUUID) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting request data: not valid user uuid"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error getting request data: not valid user uuid")
		return
	}
	if dpo.StartDate == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting request data: empty start date"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error getting request data: empty start date")
		return
	}
	if !utils.IsValidMonthYear(dpo.StartDate) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting request data: wrong type start date"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error getting request data: wrong type start date")
		return
	}

	if dpo.EndDate != "" && !utils.IsValidMonthYear(dpo.EndDate) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting request data: wrong type end date"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error getting request data: wrong type end date")
		return
	}

	if dpo.EndDate != "" && !utils.IsValidMonthYearLength(dpo.StartDate, dpo.EndDate) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting request data: the end date is earlier than the start date"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error getting request data: the end date is earlier than the start date")
		return
	}

	var subscription entities.Subscription
	subscription.ServiceName = dpo.ServiceName
	subscription.Price = dpo.Price
	subscription.UserUUID = dpo.UserUUID
	subscription.StartDate, err = utils.MonthYearToDate(dpo.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting request data: wrong type start date"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error getting request data: wrong type start date")
		return
	}
	if dpo.EndDate != "" {
		subscription.EndDate, err = utils.MonthYearToDate(dpo.EndDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting request data: wrong type end date"})
			log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
				Msg("error getting request data: wrong type end date")
			return
		}
	}

	log.Debug().Msg("call h.services.Subscription.Add")
	outSubscription, err := h.services.Subscription.Add(&subscription)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error adding subscription: %s", err.Error())})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusInternalServerError).
			Msgf("error adding subscription: %s", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, outSubscription)
	log.Info().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusCreated).
		Msg("successful adding subscription")
}

func (h *Handler) GetAllSubscription(ctx *gin.Context) {
	log.Debug().Msg("call h.services.Subscription.GetAll")
	outSubscriptionSlice, err := h.services.Subscription.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error getting all subscriptions: %s", err.Error())})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusInternalServerError).
			Msgf("error getting all subscriptions: %s", err.Error())
		return
	}

	if len(*outSubscriptionSlice) == 0 {
		ctx.JSON(http.StatusNoContent, "")
		log.Info().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusOK).
			Msg("successful getting all subscriptions, but finding 0 rows")
		return
	}

	ctx.JSON(http.StatusOK, outSubscriptionSlice)
	log.Info().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusOK).
		Msg("successful getting all subscriptions")
}

func (h *Handler) GetSubscriptionById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing user data: empty id"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error parsing user data: empty id")
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing user data: count is not a number"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error parsing user data: count is not a number")
		return
	}
	if id < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing user data: invalid id"})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusBadRequest).
			Msg("error parsing user data: invalid id")
		return
	}

	log.Debug().Msg("call h.services.Subscription.GetById")
	outSubscription, err := h.services.Subscription.GetById(id)
	if err != nil {
		if err.Error() == "Subscription.GetById: sql: no rows in result set" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "error getting subscription by id: id not founded"})
			log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusNotFound).
				Msg("error getting subscription by id: id not founded")
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error getting subscription by id: %s", err.Error())})
		log.Error().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusInternalServerError).
			Msgf("error getting subscription by id: %s", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, outSubscription)
	log.Info().Str("method", ctx.Request.Method).Str("url", ctx.Request.URL.String()).Int("status", http.StatusOK).
		Msg("successful getting subscription by id")
}
