package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"oasis/config"
	"oasis/db/model"
	"oasis/pkg/log"
	"oasis/pkg/utils"
)

// 创建数据库连接并执行查询
// executeQuery creates a database connection and executes the query
//func executeQuery(db *gorm.DB, sql string) ([]map[string]any, error) {
//	// Run the query
//	records, err := queryRow(db, sql)
//	if err != nil {
//		return nil, err
//	}
//
//	// Return the records
//	return records, nil
//}

func HandleQuery(sql, instance, database string) ([]string, []map[string]any, error) {
	ins := model.Instance{
		Name: instance,
	}

	i, err := ins.GetInstanceByName()
	if err != nil {
		return nil, nil, err
	}

	passwordX, err := utils.DecryptWithAES(i.Password)
	if err != nil {
		log.Error(err.Error())
		return nil, nil, err
	}

	db, err := openInstance(i.User, passwordX, i.Host, i.Port, database)
	if err != nil {
		return nil, nil, err
	}

	// 执行查询
	columns, result, err := queryRow(db, sql)
	if err != nil {
		return nil, nil, err
	}

	return columns, result, nil
}

func openInstance(user, password, host, port, database string) (*gorm.DB, error) {
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

	maxOpen := config.NewOasisConfig().MySQL.MaxOpenConn
	maxIdle := config.NewOasisConfig().MySQL.MaxIdleConn
	DB.SetMaxOpenConns(maxOpen)
	DB.SetMaxIdleConns(maxIdle)

	if err := DB.Ping(); err != nil {
		return nil, err
	}

	return db, err

}

//func queryRow(db *gorm.DB, query string) ([]*orderedmap.OrderedMap[string, interface{}], error) {
//	rows, err := db.Raw(query).Rows()
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	columns, err := rows.Columns()
//	if err != nil {
//		return nil, err
//	}
//
//	values := make([]sql.RawBytes, len(columns))
//	scanArgs := make([]interface{}, len(values))
//	for i := range values {
//		scanArgs[i] = &values[i]
//	}
//
//	var result []*orderedmap.OrderedMap[string, interface{}]
//	for rows.Next() {
//		err = rows.Scan(scanArgs...)
//		if err != nil {
//			return nil, err
//		}
//		record := orderedmap.NewOrderedMap[string, interface{}]()
//		for i, col := range values {
//			if col == nil {
//				record.Set(columns[i], nil)
//			} else {
//				record.Set(columns[i], string(col))
//			}
//		}
//		result = append(result, record)
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}

// queryRow 负责具体执行SQL
func queryRow(db *gorm.DB, query string) ([]string, []map[string]any, error) {
	// 执行SQL
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, nil, err
	}

	// 关闭连接
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var result []map[string]interface{}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, nil, err
		}
		record := make(map[string]interface{})
		for i, col := range values {
			if col == nil {
				record[columns[i]] = nil
			} else {
				record[columns[i]] = string(col)
			}
		}
		result = append(result, record)
	}

	if err = rows.Err(); err != nil {
		return nil, nil, err
	}

	return columns, result, nil
}

// CheckMySQLAlive 检查MySQL 数据库是否存活
func CheckMySQLAlive(user, password, host, port, database string) (int, error) {
	db, err := openInstance(user, password, host, port, database)
	if err != nil {
		return 0, err
	}

	var result int
	err = db.Raw("SELECT 1").Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func GetAllDatabaseName(user, password, host, port, database string) ([]string, error) {
	db, err := openInstance(user, password, host, port, database)
	if err != nil {
		return nil, err
	}

	var result []string
	err = db.Raw("SHOW DATABASES").Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
