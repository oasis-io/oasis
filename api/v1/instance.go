package v1

import (
	"github.com/gin-gonic/gin"
	"oasis/app/response"
	"oasis/db"
	"oasis/db/model"
	"oasis/pkg/log"
	"oasis/pkg/utils"
)

type instanceResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	DBType string `json:"db_type"`
	Host   string `json:"host"`
	Port   string `json:"port"`
	User   string `json:"user"`
}

func GetInstanceList(c *gin.Context) {
	var req PageInfo
	var insRes []model.Instance

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	ins := model.Instance{}
	insList, count, err := ins.GetInstanceList(req.PageSize, req.CurrentPage)
	if err != nil {
		log.Error(err.Error())
		response.Error(c, err.Error())
		return
	}

	if len(insList) <= 0 {
		log.Error("Failed to get instance")
		response.Error(c, "Failed to get  instance")
		return
	}

	for _, v := range insList {
		insRes = append(
			insRes, model.Instance{
				Name:   v.Name,
				DBType: v.DBType,
				Host:   v.Host,
				Port:   v.Port,
				User:   v.User,
			})
	}

	response.SendSuccessData(c, "获取实例列表成功", PageResponse{
		Data:        insRes,
		Total:       count,
		PageSize:    req.PageSize,
		CurrentPage: req.CurrentPage,
	})
}

func GetInstance(c *gin.Context) {
	var req model.Instance
	var insRes []instanceResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	ins := model.Instance{
		Name: req.Name,
	}

	instance, err := ins.GetInstanceByName()
	if err != nil {
		log.Error(err.Error())
		response.Error(c, err.Error())
		return
	}

	insRes = append(insRes, instanceResponse{
		ID:     instance.ID,
		Name:   instance.Name,
		DBType: instance.DBType,
		Host:   instance.Host,
		Port:   instance.Port,
		User:   instance.User,
	})

	response.SendSuccessData(c, "获取示例成功", PageResponse{
		Data: insRes,
	})
}

func CreateInstance(c *gin.Context) {
	var req model.Instance

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	ins := model.Instance{
		Name:     req.Name,
		DBType:   req.DBType,
		Host:     req.Host,
		Port:     req.Port,
		User:     req.User,
		Password: req.Password,
	}

	err := ins.CreateInstance()
	if err != nil {
		log.Error(err.Error())
		response.Error(c, err.Error())
		return
	}

	response.Success(c)
}

func UpdateInstance(c *gin.Context) {
	var req model.Instance

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	ins := model.Instance{
		Model: model.Model{
			ID: req.ID,
		},
		Name:     req.Name,
		DBType:   req.DBType,
		Host:     req.Host,
		Port:     req.Port,
		User:     req.User,
		Password: req.Password,
	}

	if err := ins.UpdateUserGroupByID(); err != nil {
		log.Error("database update error：" + err.Error())
		response.Error(c, "database update instance error")
		return
	}
}

func UpdateInstancePassword(c *gin.Context) {
	var req model.Instance
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	ins := model.Instance{
		Name:     req.Name,
		Password: req.Password,
	}

	if err := ins.UpdateInstancePassword(); err != nil {
		log.Error("database update error：" + err.Error())
		response.Error(c, "database update error")
		return
	}

	response.Success(c)
}

func DeleteInstance(c *gin.Context) {
	var req model.Instance

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	if req.Name == "" {
		response.Error(c, "name field is empty")
		return
	}

	ins := model.Instance{
		Name: req.Name,
	}

	err := ins.DeleteInstance()
	if err != nil {
		log.Error(err.Error())
		response.Error(c, err.Error())
		return
	}

	response.Success(c)
}

func PingInstance(c *gin.Context) {
	var req model.Instance

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	ins := model.Instance{
		Name: req.Name,
	}

	i, err := ins.GetInstanceByName()
	if err != nil {
		log.Error(err.Error())
		response.Error(c, err.Error())
		return
	}

	user := i.User
	port := i.Port
	host := i.Host
	database := "mysql"

	password, err := utils.DecryptWithAES(i.Password)
	if err != nil {
		log.Error(err.Error())
		response.Error(c, "Failed to decrypt database password")
		return
	}

	isAlive, err := db.CheckMySQLAlive(user, password, host, port, database)
	if err != nil {
		log.Error("The database checks whether: " + err.Error())
		response.Error(c, "The database checks whether")
		return
	}

	response.SendSuccessData(c, "is alive", isAlive)
}

func GetInstanceName(c *gin.Context) {
	var insRes []instanceResponse

	ins := model.Instance{}
	insList, err := ins.GetAllInstance()
	if err != nil {
		log.Error(err.Error())
		response.Error(c, err.Error())
		return
	}

	for _, v := range insList {
		insRes = append(
			insRes, instanceResponse{
				Name: v.Name,
			})
	}

	response.SendSuccessData(c, "获取实例名称成功", insRes)
}

func GetInstanceDatabase(c *gin.Context) {
	var req model.Instance

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("parameter binding errors: " + err.Error())
		response.Error(c, "parameter binding errors")
		return
	}

	ins := model.Instance{
		Name: req.Name,
	}

	i, err := ins.GetInstanceByName()
	if err != nil {
		log.Error("The database checks whether: " + err.Error())
		response.Error(c, "The database checks whether")
		return
	}

	passwordX, err := utils.DecryptWithAES(i.Password)
	if err != nil {
		response.Error(c, "password error")
		log.Error(err.Error())
		return
	}

	database := "mysql"

	databaseName, err := db.GetAllDatabaseName(i.User, passwordX, i.Host, i.Port, database)
	if err != nil {
		log.Error("The database checks whether: " + err.Error())
		response.Error(c, "The database checks whether")
		return
	}

	response.SendSuccessData(c, "all database name", databaseName)
}
