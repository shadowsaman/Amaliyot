package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"crud/models"

	"github.com/gin-gonic/gin"
)

// CreateSubscription godoc
// @ID create_subscription
// @Router /subscription [POST]
// @Summary Create Subscription
// @Description Create Subscription
// @Tags Subscription
// @Accept json
// @Produce json
// @Param Subscription body models.CreateSubscription true "CreateSubscriptionRequestBody"
// @Success 201 {object} models.Subscription "GetSubscriptionBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) CreateSubscription(c *gin.Context) {
	var Subscription models.CreateSubscription

	err := c.ShouldBindJSON(&Subscription)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Subscription().Create(context.Background(), &Subscription)
	if err != nil {
		log.Printf("error whiling Create: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling Create").Error())
		return
	}

	resp, err := h.storage.Subscription().GetByPKey(
		context.Background(),
		&models.SubscriptionPrimaryKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByIdSubscription godoc
// @ID get_by_id_subscription
// @Router /subscription/{id} [GET]
// @Summary Get By Id Subscription
// @Description Get By Id Subscription
// @Tags Subscription
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Subscription "GetSubscriptionBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetSubscriptionById(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.storage.Subscription().GetByPKey(
		context.Background(),
		&models.SubscriptionPrimaryKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetListSubscription godoc
// @ID get_list_subscription
// @Router /subscription [GET]
// @Summary Get List Subscription
// @Description Get List Subscription
// @Tags Subscription
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} models.GetListSubscriptionResponse "GetSubscriptionBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetSubscriptionList(c *gin.Context) {
	var (
		limit  int
		offset int
		err    error
	)

	limitStr := c.Query("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Printf("error whiling limit: %v\n", err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	offsetStr := c.Query("offset")
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Printf("error whiling limit: %v\n", err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	resp, err := h.storage.Subscription().GetList(
		context.Background(),
		&models.GetListSubscriptionRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)

	if err != nil {
		log.Printf("error whiling get list: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get list").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateSubscription godoc
// @ID update_Subscription
// @Router /Subscription/{id} [PUT]
// @Summary Update Subscription
// @Description Update Subscription
// @Tags Subscription
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Subscription body models.UpdateSubscriptionSwagger true "CreateSubscriptionRequestBody"
// @Success 200 {object} models.Subscription "GetSubscriptionsBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) UpdateSubscription(c *gin.Context) {

	var (
		Subscription models.UpdateSubscription
	)

	id := c.Param("id")

	if id == "" {
		log.Printf("error whiling update: %v\n", errors.New("required Subscription id").Error())
		c.JSON(http.StatusBadRequest, errors.New("required Subscription id").Error())
		return
	}

	err := c.ShouldBindJSON(&Subscription)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	Subscription.Id = id

	rowsAffected, err := h.storage.Subscription().Update(
		context.Background(),
		&Subscription,
	)

	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	if rowsAffected == 0 {
		log.Printf("error whiling update rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update rows affected").Error())
		return
	}

	resp, err := h.storage.Subscription().GetByPKey(
		context.Background(),
		&models.SubscriptionPrimaryKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteByIdSubscription godoc
// @ID delete_by_id_subscription
// @Router /subscription/{id} [DELETE]
// @Summary Delete By Id Subscription
// @Description Delete By Id Subscription
// @Tags Subscription
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Subscription "GetSubscriptionBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) DeleteSubscription(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		log.Printf("error whiling update: %v\n", errors.New("required Subscription id").Error())
		c.JSON(http.StatusBadRequest, errors.New("required Subscription id").Error())
		return
	}

	err := h.storage.Subscription().Delete(
		context.Background(),
		&models.SubscriptionPrimaryKey{
			Id: id,
		},
	)

	if err != nil {
		log.Printf("error whiling delete: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling delete").Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
