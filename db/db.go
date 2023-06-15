package db

import (
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"oasis/config"
	"oasis/db/model"
	"oasis/pkg/log"
	"sort"
	"strings"
)

func OpenOasis() (*gorm.DB, error) {
	return openOasis()
}

func openOasis() (*gorm.DB, error) {
	config := config.NewConfig()
	user := config.MySQL.User
	password := config.MySQL.Password
	host := config.MySQL.Host
	port := config.MySQL.Port
	database := config.MySQL.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)

	mysqlConfig := mysql.Config{
		DriverName:                "mysql",
		DSN:                       dsn,
		DefaultStringSize:         255,
		SkipInitializeWithVersion: false,
		DisableDatetimePrecision:  true,
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	DB, err := db.DB()
	if err != nil {
		return nil, err
	}

	maxOpen := config.MySQL.MaxOpenConn
	maxIdle := config.MySQL.MaxIdleConn
	DB.SetMaxOpenConns(maxOpen)
	DB.SetMaxIdleConns(maxIdle)

	if err := DB.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

func OpenInstance() (*gorm.DB, error) {
	return openInstance()
}

func openInstance() (*gorm.DB, error) {
	return nil, nil
}

// Login 登陆验证
func Login(username, password string) (*model.User, error) {
	db := config.DB

	user := model.User{}

	// 根据用户名查询用户记录
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Error("query user error", zap.Error(err))
		return nil, err
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Error("password verification failed", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

func GetMenuTree() ([]model.Menu, error) {
	var menus []model.Menu
	// 获取所有的菜单项
	db := config.DB
	if err := db.Find(&menus).Error; err != nil {
		return nil, err
	}

	// 构建菜单树
	return buildMenuTree(menus, "0"), nil
}

func buildMenuTree(menus []model.Menu, parentId string) []model.Menu {
	var result []model.Menu
	for _, menu := range menus {
		// 将 menu.ID 转换为 string 来比较
		if menu.ParentID == parentId {
			children := buildMenuTree(menus, fmt.Sprintf("%d", menu.Sort))
			menu.Children = children
			result = append(result, menu)
		}
	}

	// 对二级菜单进行排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].Sort < result[j].Sort
	})

	return result
}

func GetMenuTreeMap() ([]model.Menu, error) {
	var menus []model.Menu

	// 获取所有的菜单项
	db := config.DB
	if err := db.Find(&menus).Error; err != nil {
		return nil, err
	}

	// 将菜单按父级ID分类到一个map中
	menuMap := make(map[string][]model.Menu)
	for _, menu := range menus {
		menuMap[menu.ParentID] = append(menuMap[menu.ParentID], menu)
	}

	// 构建菜单树
	return buildMenuTreeWithMap(menuMap, "0"), nil
}

func buildMenuTreeWithMap(menuMap map[string][]model.Menu, parentId string) []model.Menu {
	var result []model.Menu

	for _, menu := range menuMap[parentId] {
		children := buildMenuTreeWithMap(menuMap, fmt.Sprintf("%d", menu.Sort))
		menu.Children = children
		result = append(result, menu)
	}

	// 对子菜单进行排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].Sort < result[j].Sort
	})

	return result
}

// GetMenuTreeMapForRole 通过角色名查询该角色拥有的菜单
func GetMenuTreeMapForRole(roleName string) ([]model.Menu, error) {
	var role model.UserRole
	var menus []model.Menu

	// 获取角色
	if err := config.DB.Where("name = ?", roleName).First(&role).Error; err != nil {
		return nil, err
	}

	// 获取角色关联的菜单ID
	var relation []model.RoleMenuRelation
	if err := config.DB.Where("role_id = ?", role.ID).Find(&relation).Error; err != nil {
		return nil, err
	}

	// 获取这些菜单的信息
	menuIDs := make([]uint, len(relation))
	for i, r := range relation {
		menuIDs[i] = r.MenuID
	}
	if err := config.DB.Where("id IN ?", menuIDs).Find(&menus).Error; err != nil {
		return nil, err
	}

	// 构建菜单树
	menuMap := make(map[string][]model.Menu)
	for _, menu := range menus {
		menuMap[menu.ParentID] = append(menuMap[menu.ParentID], menu)
	}

	return buildMenuTreeWithMap(menuMap, "0"), nil
}

func AddApiPermissions(roleName string, apis []model.Api) error {
	var permissions []gormadapter.CasbinRule

	for _, api := range apis {
		if api.Path != "" && api.Method != "" {
			permission := gormadapter.CasbinRule{
				Ptype: "p",
				V0:    strings.ToUpper(roleName),
				V1:    api.Path,
				V2:    api.Method,
			}
			permissions = append(permissions, permission)
		}
	}

	db := config.DB

	for _, permission := range permissions {
		result := db.Create(&permission)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func GetApisByRole(roleName string) ([]model.Api, error) {
	// Query Role ID
	var role model.UserRole
	if err := config.DB.Where("name = ?", roleName).First(&role).Error; err != nil {
		return nil, err
	}

	// Query related permissions
	var permissions []gormadapter.CasbinRule
	if err := config.DB.Where("v0 = ?", strings.ToUpper(role.Name)).Find(&permissions).Error; err != nil {
		return nil, err
	}

	// Query all APIs
	apiModel := &model.Api{}
	allApis, err := apiModel.GetAllApi()
	if err != nil {
		return nil, err
	}

	// Convert to API format
	var apis []model.Api
	for _, api := range allApis {
		for _, permission := range permissions {
			// If the API path and method matches with permission, append it to the list
			if api.Path == permission.V1 && api.Method == permission.V2 {
				apis = append(apis, api)
			}
		}
	}

	return apis, nil
}
