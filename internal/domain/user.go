package domain

// DutyUser 值班人员信息
type DutyUser struct {
	Version SwitchVersion
	User
}

// User 用户信息， 域数据权限校验
type User struct {
	SsoUm     string
	UserName  string
	BizDomain string
}
