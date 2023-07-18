package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rcherin/hackaton/database"
	"github.com/rcherin/hackaton/models"
)

func ListTrasactionsByAccountAndDate(c *fiber.Ctx) error {
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
	database.DB.Db.Unscoped().Where(&models.Transactions{Acc_id: c.Params("acc_id"), Status: c.Query("status")}).Where(dateTimeCondition, fromDateTime, toDateTime).Limit(pageSize).Offset(offset).Find(&transactions)

	return c.Status(200).JSON(transactions)
}

func ListTrasactionsById(c *fiber.Ctx) error {
	transactions := []models.Transactions{}
	database.DB.Db.Unscoped().Where(&models.Transactions{Tx_id: c.Params("tx_id")}).Find(&transactions)

	return c.Status(200).JSON(transactions)
}
