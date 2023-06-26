package v1

import (
	"github.com/gin-gonic/gin"
	"oasis/app/response"
	"oasis/db"
	"oasis/pkg/log"
)

type Query struct {
	Sql      string `json:"sql"`
	Database string `json:"database"`
	Instance string `json:"instance"`
}

type Result struct {
	Columns []string                 `json:"columns"`
	Data    []map[string]interface{} `json:"data"`
}

func QueryData(c *gin.Context) {
	var req Query

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	columns, result, err := db.HandleQuery(req.Sql, req.Instance, req.Database)
	if err != nil {
		log.Error(err.Error())
		response.Error(c, "查询错误")
		return
	}

	// 创建一个新的结构体以包含列的名称和对应的数据
	data := &Result{
		Columns: columns,
		Data:    result,
	}

	response.SendSuccessData(c, "返回数据查询结果", data)
}
