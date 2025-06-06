package controllers

import (
	"net/http"
	"rest-api/models"
	"rest-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type OrderController struct {
	service services.OrderService
}

func NewOrderController(s services.OrderService) *OrderController {
	return &OrderController{service: s}
}

func (c *OrderController) GetOrders(ctx *gin.Context) {
	userID := ctx.GetUint("userID")
	userRole := ctx.GetString("role")

	logrus.Infof("GetOrders called by userID: %d, role: %s", userID, userRole)

	orders, err := c.service.GetOrders(userID, userRole)
	if err != nil {
		logrus.Errorf("Error fetching orders: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (c *OrderController) UpdateOrder(ctx *gin.Context) {
	userID := ctx.GetUint("userID")
	userRole := ctx.GetString("role")

	orderIDStr := ctx.Param("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		logrus.Errorf("Invalid order Id format: %s, error: %v", orderIDStr, err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID format"})
		return
	}

	var updateRequest models.OrderUpdateRequest
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		logrus.Errorf("Error binding JSON for UpdateOrder: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload: " + err.Error()})
		return
	}

	logrus.Infof("UpdateOrder called for orderID: %d by userID: %d, role: %s", uint(orderID), userID, userRole)

	updatedOrder, err := c.service.UpdateOrderByID(uint(orderID), updateRequest, userID, userRole)
	if err != nil {
		logrus.Errorf("Error updating order %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order updated successfully",
		"order":   updatedOrder,
	})
}
