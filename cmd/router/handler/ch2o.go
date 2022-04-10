package handler

import (
	"env/cmd/router/resp"
	"env/internal/ch2o"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func addCh2o(c *gin.Context) {
	strVal := c.DefaultPostForm("ch2o", "-100")
	val, _ := strconv.ParseFloat(strVal, 64)
	fmt.Printf("CH2O: %.3f mg/m3", val/1000.0*1.23)
	ch2o.Append(val / 1000.0 * 1.23)
	c.JSON(200, resp.CommonResp{
		Success: true,
	})
}

func getCh2os(c *gin.Context) {
	startTime := c.DefaultQuery("startTime", time.Now().Format("2006-01-02")+" 00:00:00")
	endTime := c.DefaultQuery("endTime", time.Now().AddDate(0, 0, 1).Format("2006-01-02")+" 00:00:00")
	infos := ch2o.GetCh2oInfo(startTime, endTime)
	c.JSON(200, resp.CommonResp{
		Success: true,
		Result:  infos,
	})
}

func BuildCh2o(r *gin.Engine) {
	r.POST("/env/addCh2o", addCh2o)
	// r.POST("/fs/removeGroup", delRoot)
	r.GET("/env/getCh2os", getCh2os)

	// r.GET("/fs/rs", getResource)
	// r.GET("/fs/rs/enter", getChildrenResource)
}
