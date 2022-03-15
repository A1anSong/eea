package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetPageParam(c *gin.Context) (page, perPage int) {
	pageStr := c.Query("page")
	perPageStr := c.Query("perPage")
	n, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		logrus.Errorf("GetPageParam page: %s error: %s", pageStr, err.Error())
		n = 1
	}
	page = int(n)
	n, err = strconv.ParseInt(perPageStr, 10, 64)
	if err != nil {
		logrus.Errorf("GetPageParam perPage: %s error: %s", perPageStr, err.Error())
		n = 10
		return
	}
	perPage = int(n)
	return
}
