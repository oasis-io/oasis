package v1

import (
	"github.com/gin-gonic/gin"
	"oasis/app/response"
	"oasis/config"
)

type User struct {
	Host                  string `gorm:"column:Host"`
	User                  string `gorm:"column:User"`
	SelectPriv            string `gorm:"column:Select_priv"`
	InsertPriv            string `gorm:"column:Insert_priv"`
	UpdatePriv            string `gorm:"column:Update_priv"`
	DeletePriv            string `gorm:"column:Delete_priv"`
	CreatePriv            string `gorm:"column:Create_priv"`
	DropPriv              string `gorm:"column:Drop_priv"`
	ReloadPriv            string `gorm:"column:Reload_priv"`
	ShutdownPriv          string `gorm:"column:Shutdown_priv"`
	ProcessPriv           string `gorm:"column:Process_priv"`
	FilePriv              string `gorm:"column:File_priv"`
	GrantPriv             string `gorm:"column:Grant_priv"`
	ReferencesPriv        string `gorm:"column:References_priv"`
	IndexPriv             string `gorm:"column:Index_priv"`
	AlterPriv             string `gorm:"column:Alter_priv"`
	ShowDbPriv            string `gorm:"column:Show_db_priv"`
	SuperPriv             string `gorm:"column:Super_priv"`
	CreateTmpTablePriv    string `gorm:"column:Create_tmp_table_priv"`
	LockTablesPriv        string `gorm:"column:Lock_tables_priv"`
	ExecutePriv           string `gorm:"column:Execute_priv"`
	ReplSlavePriv         string `gorm:"column:Repl_slave_priv"`
	ReplClientPriv        string `gorm:"column:Repl_client_priv"`
	CreateViewPriv        string `gorm:"column:Create_view_priv"`
	ShowViewPriv          string `gorm:"column:Show_view_priv"`
	CreateRoutinePriv     string `gorm:"column:Create_routine_priv"`
	AlterRoutinePriv      string `gorm:"column:Alter_routine_priv"`
	CreateUserPriv        string `gorm:"column:Create_user_priv"`
	EventPriv             string `gorm:"column:Event_priv"`
	TriggerPriv           string `gorm:"column:Trigger_priv"`
	CreateRolePriv        string `gorm:"column:Create_role_priv"`
	DropRolePriv          string `gorm:"column:Drop_role_priv"`
	PasswordExpiredPriv   string `gorm:"column:Password_expired_priv"`
	ManageGrantsPriv      string `gorm:"column:Manage_grants_priv"`
	AuthenticationString  string `gorm:"column:authentication_string"`
	PasswordLastChanged   string `gorm:"column:password_last_changed"`
	PasswordLifetime      string `gorm:"column:password_lifetime"`
	AccountLocked         string `gorm:"column:account_locked"`
	CreateTablesPriv      string `gorm:"column:Create_tables_priv"`
	SslType               string `gorm:"column:ssl_type"`
	SslCipher             string `gorm:"column:ssl_cipher"`
	X509Issuer            string `gorm:"column:x509_issuer"`
	X509Subject           string `gorm:"column:x509_subject"`
	MaxQuestions          int    `gorm:"column:max_questions"`
	MaxUpdates            int    `gorm:"column:max_updates"`
	MaxConnections        int    `gorm:"column:max_connections"`
	MaxUserConnections    int    `gorm:"column:max_user_connections"`
	Plugin                string `gorm:"column:plugin"`
	PasswordExpired       string `gorm:"column:password_expired"`
	PasswordLastChangedAt string `gorm:"column:password_last_changed_at"`
	PasswordLifeTime      string `gorm:"column:password_lifetime"`
}

func GetSqlQueryData(c *gin.Context) {
	//api := model.Api{}
	//allApi, err := api.GetAllApi()
	//if err != nil {
	//	log.Error("获取菜单API失败!", zap.Error(err))
	//	response.Error(c, err.Error())
	//	return
	//}
	var users []User
	db := config.DB

	result := db.Raw("SELECT * FROM mysql.user").Scan(&users)
	if result.Error != nil {
		response.Error(c, result.Error.Error())
		return
	}

	response.SendSuccessData(c, "获取用户成功", users)
}
