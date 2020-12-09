package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Add Account in Slice
// @Tags Device
// @version 1.0
// @produce application/json
// @param auth header string true "token"
// @param params body account true "params"
// @Success 200 {object} account "{"msg":"OK"}"
// @Failure 400 {string} json "{"msg":"fail"}"
// @Router /Device [post]
func InsertAccount(c *gin.Context) {
	var account account
	err := c.BindJSON(&account)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"msg": err.Error(),
		})
		return
	}
	a = append(a, account)
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

// @Summary Get Account by index
// @Tags Device
// @version 1.0
// @produce application/json
// @param auth header string true "token"
// @param index query string true "index"
// @Success 200 {object} account "{"msg":"OK"}"
// @Failure 400 {string} json "{"msg":"fail"}"
// @Router /Device [get]
func GetAccount(c *gin.Context) {
	indexStr := c.Request.URL.Query().Get("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, a[index])
}

// @Summary Delete Account by index
// @Tags Device
// @version 1.0
// @produce application/json
// @param auth header string true "token"
// @param index path string true "index"
// @Success 200 {string} json "{"msg":"OK"}"
// @Failure 400 {string} json "{"msg":"fail"}"
// @Router /Device [delete]
func DeleteAccount(c *gin.Context) {
	indexStr := c.Param("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"msg": err.Error(),
		})
		return
	}
	a = append(a[:index], a[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}
