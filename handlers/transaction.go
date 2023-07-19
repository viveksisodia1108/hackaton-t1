package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viveksisodia1108/gogingorm/database"
	"github.com/viveksisodia1108/gogingorm/models"
)

func ListTrasactionsByAccountAndDate(c *gin.Context) {
	transactions := []models.Transactions{}
	currentTime := time.Now()

	toDateTime := c.Query("toDateTime")

	formattedCurrentDateTime := fmt.Sprintf("%d-%d-%d %d:%d:%d",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second())

	if toDateTime == "" {
		toDateTime = formattedCurrentDateTime
	}

	dateTimeCondition := "tx_ts >= ?  AND tx_ts <= ?"
	fromDateTime := c.Query("fromDateTime")
	if fromDateTime == "" {
		dateTimeCondition = "tx_ts <= ?  AND tx_ts <= ?"
		fromDateTime = formattedCurrentDateTime
	}

	pageSizeStr := c.Query("pageSize")
	pageSize := 50
	if pageSizeStr != "" {
		convPageSize, _ := strconv.Atoi(pageSizeStr)
		pageSize = convPageSize
	}
	pageStr := c.Query("page")
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * pageSize
	database.DB.Db.Unscoped().Where(&models.Transactions{Acc_id: c.Param("acc_id"), Status: c.Query("status")}).Where(dateTimeCondition, fromDateTime, toDateTime).Limit(pageSize).Offset(offset).Find(&transactions)

	c.IndentedJSON(http.StatusOK, transactions)
}

func ListTrasactionsById(c *gin.Context) {
	transactions := []models.Transactions{}
	database.DB.Db.Unscoped().Where(&models.Transactions{Tx_id: c.Param("tx_id")}).Find(&transactions)

	c.IndentedJSON(http.StatusOK, transactions)
}
