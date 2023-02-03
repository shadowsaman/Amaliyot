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

// CreateUserSubscription godoc
// @ID create_user_subscription
// @Router /user-subscription [POST]
// @Summary Create UserSubscription
// @Description Create UserSubscription
// @Tags UserSubscription
// @Accept json
// @Produce json
// @Param UserSubscription body models.CreateUserSubscription true "CreateUserSubscriptionRequestBody"
// @Success 201 {object} models.UserSubscription "GetUserSubscriptionBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) CreateUserSubscription(c *gin.Context) {
	var UserSubscription models.CreateUserSubscription

	err := c.ShouldBindJSON(&UserSubscription)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.UserSubscription().Create(context.Background(), &UserSubscription)
	if err != nil {
		log.Printf("error whiling Create: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling Create").Error())
		return
	}

	resp, err := h.storage.UserSubscription().GetByPKey(
		context.Background(),
		&models.UserSubscriptionPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByIdUserSubscription godoc
// @ID get_by_id_user_subscription
// @Router /user-subscription/{id} [GET]
// @Summary Get By Id UserSubscription
// @Description Get By Id UserSubscription
// @Tags UserSubscription
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.UserSubscription "GetUserSubscriptionBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetUserSubscriptionById(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.storage.UserSubscription().GetByPKey(
		context.Background(),
		&models.UserSubscriptionPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetByUserIdUserSubscription godoc
// @ID get_by_user_id_subscription
// @Router /user-id-subscription/{id} [GET]
// @Summary Get By User Id UserSubscription
// @Description Get By User Id UserSubscription
// @Tags UserSubscription
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.GetSubscriptionByUserId "GetUserSubscriptionBodyByUserId"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetUserSubscriptionByUserId(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.storage.UserSubscription().GetSubscriptionsByUserId(
		context.Background(),
		&models.UserIdSubscription{UserId: id},
	)

	if err != nil {
		log.Printf("error whiling user subscription: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling user-subscription-get-by-user-id").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetListUserSubscription godoc
// @Security ApiKeyAuth
// @ID get_list_user_subscription
// @Router /user-subscription [GET]
// @Summary Get List UserSubscription
// @Description Get List UserSubscription
// @Tags UserSubscription
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} models.GetListUserSubscriptionResponse "GetUserSubscriptionBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetUserSubscriptionList(c *gin.Context) {
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

	resp, err := h.storage.UserSubscription().GetList(
		context.Background(),
		&models.GetListUserSubscriptionRequest{
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

// HasAccess godoc
// @ID has_access
// @Router /has-access/{id} [GET]
// @Summary Has Access UserSubscription
// @Description Has Access UserSubscription
// @Tags UserSubscription
// @Accept json
// @Produce json
// @Param id path bool true "id"
// @Success 200 {object} string "HasAcces"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) HasAccess(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.storage.UserSubscription().HasAccess(
		context.Background(),
		&models.UserIdSubscription{UserId: id},
	)

	if err != nil {
		log.Printf("error whiling user subscription: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling user-subscription-get-by-user-id").Error())
		return
	}

	if resp == "INACTIVE" || resp == "FREE_TRIAL_EXPIRED" {
		c.JSON(http.StatusOK, "false")
	}

	c.JSON(http.StatusOK, "true")
}

// MakeActive godoc
// @ID make_active
// @Router /make-active/{id} [GET]
// @Summary Make Active UserSubscription
// @Description Make Active UserSubscription
// @Tags UserSubscription
// @Accept json
// @Produce json
// @Param id path bool true "id"
// @Success 200 {object} models.UserHistory "GetUserHistoryBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) MakeActive(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.storage.UserSubscription().MakeSubscriptionActive(
		context.Background(),
		&models.UserSubscriptionId{UserSubscriptionId: id},
	)

	if err != nil {
		log.Printf("error whiling user history active subscription: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling user-subscription-active").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// // UpdateUserSubscription godoc
// // @ID update_UserSubscription
// // @Router /UserSubscription/{id} [PUT]
// // @Summary Update UserSubscription
// // @Description Update UserSubscription
// // @Tags UserSubscription
// // @Accept json
// // @Produce json
// // @Param id path string true "id"
// // @Param UserSubscription body models.UpdateUserSubscriptionSwagger true "CreateUserSubscriptionRequestBody"
// // @Success 200 {object} models.UserSubscription "GetUserSubscriptionsBody"
// // @Response 400 {object} string "Invalid Argument"
// // @Failure 500 {object} string "Server Error"
// func (h *HandlerV1) UpdateUserSubscription(c *gin.Context) {

// 	var (
// 		UserSubscription models.UpdateUserSubscription
// 	)

// 	id := c.Param("id")

// 	if id == "" {
// 		log.Printf("error whiling update: %v\n", errors.New("required UserSubscription id").Error())
// 		c.JSON(http.StatusBadRequest, errors.New("required UserSubscription id").Error())
// 		return
// 	}

// 	err := c.ShouldBindJSON(&UserSubscription)
// 	if err != nil {
// 		log.Printf("error whiling update: %v\n", err)
// 		c.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	UserSubscription.Id = id

// 	rowsAffected, err := h.storage.UserSubscription().Update(
// 		context.Background(),
// 		&UserSubscription,
// 	)

// 	if err != nil {
// 		log.Printf("error whiling update: %v", err)
// 		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
// 		return
// 	}

// 	if rowsAffected == 0 {
// 		log.Printf("error whiling update rows affected: %v", err)
// 		c.JSON(http.StatusInternalServerError, errors.New("error whiling update rows affected").Error())
// 		return
// 	}

// 	resp, err := h.storage.UserSubscription().GetByPKey(
// 		context.Background(),
// 		&models.UserSubscriptionPrimarKey{Id: id},
// 	)

// 	if err != nil {
// 		log.Printf("error whiling GetByPKey: %v\n", err)
// 		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, resp)
// }

// // DeleteByIdUserSubscription godoc
// // @ID delete_by_id_UserSubscription
// // @Router /UserSubscription/{id} [DELETE]
// // @Summary Delete By Id UserSubscription
// // @Description Delete By Id UserSubscription
// // @Tags UserSubscription
// // @Accept json
// // @Produce json
// // @Param id path string true "id"
// // @Success 200 {object} models.UserSubscription "GetUserSubscriptionBody"
// // @Response 400 {object} string "Invalid Argument"
// // @Failure 500 {object} string "Server Error"
// func (h *HandlerV1) DeleteUserSubscription(c *gin.Context) {

// 	id := c.Param("id")
// 	if id == "" {
// 		log.Printf("error whiling update: %v\n", errors.New("required UserSubscription id").Error())
// 		c.JSON(http.StatusBadRequest, errors.New("required UserSubscription id").Error())
// 		return
// 	}

// 	err := h.storage.UserSubscription().Delete(
// 		context.Background(),
// 		&models.UserSubscriptionPrimarKey{
// 			Id: id,
// 		},
// 	)

// 	if err != nil {
// 		log.Printf("error whiling delete: %v", err)
// 		c.JSON(http.StatusInternalServerError, errors.New("error whiling delete").Error())
// 		return
// 	}

// 	c.JSON(http.StatusNoContent, nil)
// }
