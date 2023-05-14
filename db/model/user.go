package model

// All database tables do not use foreign keys
// The suffix of the associated table of the two tables is unified with relation

type User struct {
	Model
	Username   string       `json:"username" gorm:"column:username;index:uk_name,unique;not null;"`
	Password   string       `json:"password"  gorm:"column:password;not null;"`
	Email      string       `json:"email" gorm:"column:email;"`
	Phone      string       `json:"phone" gorm:"column:phone;"`
	IsEnable   bool         `json:"is_enable" gorm:"column:is_enable;type:tinyint(1);default:0;comment:0:enable,1:disabled"`
	Roles      []*UserRole  `gorm:"many2many:user_role_relation;"`
	UserGroups []*UserGroup `gorm:"many2many:user_group_relation"`
}
