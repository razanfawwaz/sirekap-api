package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/razanfawwaz/sirekap-api/application"
	"github.com/razanfawwaz/sirekap-api/config"
	"net/http"
	"strconv"
)

func GetAllData(c *gin.Context) {
	// limit & offset from query params is int
	limit, offset := c.Query("limit"), c.Query("offset")
	// convert limit & offset to int
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	// if limit & offset is not provided, return error
	if limit == "" || offset == "" {
		response := config.BaseResponse{
			Data:    nil,
			Status:  http.StatusBadRequest,
			Message: "Limit is required",
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data, err := application.GetAllData(limitInt, offsetInt)
	if err != nil {
		response := config.BaseResponse{
			Data:    nil,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// return data as JSON from base response
	response := config.BaseResponse{
		Data:    data,
		Status:  http.StatusOK,
		Message: "Success",
	}

	c.JSON(http.StatusOK, response)
}
