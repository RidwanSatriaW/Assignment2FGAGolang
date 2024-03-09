package controllers

import (
	"assignment2/database"
	"assignment2/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	db := database.GetDB()

	var requestData struct {
		CustomerName string        `json:"customerName"`
		Items        []models.Item `json:"items"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		CustomerName: requestData.CustomerName,
	}

	if err := db.Debug().Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i := range requestData.Items {
		requestData.Items[i].OrderID = order.ID
	}

	if err := db.Create(&requestData.Items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}

func GetOrders(c *gin.Context) {
	db := database.GetDB()

	var Orders []models.Order
	err := db.Debug().Preload("Items").Find(&Orders).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": Orders,
	})
}

func UpdateOrder(c *gin.Context) {
    var requestData struct {
        CustomerName string        `json:"customerName"`
        Items        []models.Item `json:"items"`
        OrderedAt    time.Time     `json:"orderedAt"`
    }

    if err := c.ShouldBindJSON(&requestData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := database.GetDB()

    orderID := c.Param("id")

    var order models.Order
    if err := db.Where("id = ?", orderID).First(&order).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    for _, itemReq := range requestData.Items {
        var item models.Item
        if err := db.Where("order_id = ? AND id = ?", orderID, itemReq.ID).First(&item).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
            return
        }
        // Lakukan pembaruan item
        item.ItemCode = itemReq.ItemCode
        item.Description = itemReq.Description
        item.Quantity = itemReq.Quantity
        if err := db.Save(&item).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
            return
        }
    }

    // Update informasi order jika diperlukan
    order.CustomerName = requestData.CustomerName
    order.OrderedAt = requestData.OrderedAt
    if err := db.Save(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

func DeleteOrder(c *gin.Context) {
	db := database.GetDB()
	OrderID := c.Param("id")

	// Hapus item yang terkait dengan pesanan
	if err := db.Where("id = ?", OrderID).Delete(&models.Item{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hapus pesanan itu sendiri
	if err := db.Where("id = ?", OrderID).Delete(&models.Order{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order deleted successfully",
	})
}
