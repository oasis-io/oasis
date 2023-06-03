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
		log.Error("获取用户错误", zap.Error(err))
		return nil, err
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Error("密码校验失败", zap.Error(err))
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

func GetCasbinRulesByRole(roleName string) ([]gormadapter.CasbinRule, error) {
	var rules []gormadapter.CasbinRule

	db := config.DB
	err := db.Where("v0 = ?", roleName).Find(&rules).Error
	if err != nil {
		return nil, err
	}

	return rules, nil
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
