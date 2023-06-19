package casbin

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"oasis/config"
	"oasis/pkg/log"
	"sync"
)

var (
	cachedEnforcer *casbin.CachedEnforcer
	once           sync.Once
)

func Casbin() *casbin.CachedEnforcer {
	once.Do(func() {
		db := config.DB
		a, err := gormadapter.NewAdapterByDB(db)
		if err != nil {
			log.Error(err.Error())
			return
		}

		rbac := config.RBAC
		m, err := model.NewModelFromString(rbac)
		if err != nil {
			log.Error(err.Error())
			return
		}

		cachedEnforcer, err = casbin.NewCachedEnforcer(m, a)
		if err != nil {
			log.Error(err.Error())
			return
		}

		cachedEnforcer.SetExpireTime(60 * 60)
		err = cachedEnforcer.LoadPolicy()
		if err != nil {
			log.Error(err.Error())
			return
		}
	})
	return cachedEnforcer
}

func GetUserPermissions(username string) ([][]string, error) {
	e := Casbin() // 获取Casbin enforcer实例
	if e == nil {
		return nil, errors.New("Casbin enforcer instance is nil")
	}

	// 使用GetFilteredPolicy获取该用户的所有策略
	// 0 是角色名（v0）在策略中的位置，username 是要匹配的角色名
	policies := e.GetFilteredPolicy(0, username)

	return policies, nil
}

func InitCasbin() {
	db := config.DB
	_, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Error("Casbin initialization error：" + err.Error())
		return
	}
}

func InitCasbinRule() error {

	db := config.DB
	_, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return err
	}

	table := []gormadapter.CasbinRule{
		// Menu
		{Ptype: "p", V0: "admin", V1: "/v1/menu", V2: "POST"},

		// User List
		{Ptype: "p", V0: "admin", V1: "/v1/user", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/v1/user", V2: "DELETE"},
		{Ptype: "p", V0: "admin", V1: "/v1/user", V2: "PATCH"},
		{Ptype: "p", V0: "admin", V1: "/v1/user/add", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/v1/user/list", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/v1/user/info", V2: "GET"},

		// User Role
		{Ptype: "p", V0: "admin", V1: "/v1/user/role", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/v1/user/role", V2: "DELETE"},
		{Ptype: "p", V0: "admin", V1: "/v1/user/role", V2: "PATCH"},
		{Ptype: "p", V0: "admin", V1: "/v1/user/role/add", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/v1/user/role/list", V2: "POST"},

		// User Group
		{Ptype: "p", V0: "admin", V1: "/v1/user/group", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/v1/user/group", V2: "DELETE"},
		{Ptype: "p", V0: "admin", V1: "/v1/user/group", V2: "PATCH"},
		{Ptype: "p", V0: "admin", V1: "/v1/user/group/add", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/v1/user/group/list", V2: "POST"},

		// Instance
		{Ptype: "p", V0: "admin", V1: "/v1/instance", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/v1/instance", V2: "DELETE"},
		{Ptype: "p", V0: "admin", V1: "/v1/instance", V2: "PATCH"},
		{Ptype: "p", V0: "admin", V1: "/v1/instance/list", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/v1/instance/add", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/v1/instance/ping", V2: "POST"},
	}

	var count int64
	db.Model(&gormadapter.CasbinRule{}).Count(&count)
	if count > 0 {
		return nil
	}

	if err := db.Create(&table).Error; err != nil {
		return err
	}

	return nil
}
