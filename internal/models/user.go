package models

// TAdmin  商户后台会员记录
type TAdmin struct {
	ID                int64  `gorm:"column:id" json:"id"`
	Account           string `gorm:"column:account" json:"account"`                         // 后台用户名
	NickName          string `gorm:"column:nick_name" json:"nick_name"`                     //  昵称
	RoleId            int64  `gorm:"column:role_id" json:"role_id"`                         //  角色id
	LoginPassword     string `gorm:"column:login_password" json:"login_password"`           //  登录密码加密字符
	LoginPasswordSalt string `gorm:"column:login_password_salt" json:"login_password_salt"` //  密码言
	Status            int64  `gorm:"column:status" json:"status"`                           //  状态-1 停用 1启用
	Step2SecretKey    string `gorm:"column:step2_secret_key" json:"step2_secret_key"`       //  安全码秘钥
	LoginCount        int64  `gorm:"column:login_count" json:"login_count"`                 //  登录次数
	CreateIp          string `gorm:"column:create_ip" json:"create_ip"`                     //  创建ip
	CreateTime        int64  `gorm:"column:create_time" json:"create_time"`                 //  创建时间
	LastLoginIp       string `gorm:"column:last_login_ip" json:"last_login_ip"`             //  最后登陆ip
	LastLoginTime     int64  `gorm:"column:last_login_time" json:"last_login_time"`         //  最后登陆时间
	Auth              string `gorm:"column:auth" json:"auth"`                               //  目录权限
}
