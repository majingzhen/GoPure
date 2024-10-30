package model

// UserRole 用户角色关联模型
type UserRole struct {
	UserId int `json:"userId" gorm:"column:user_id;not null;comment:用户id;index:idx_user_id"`
	RoleId int `json:"roleId" gorm:"column:role_id;not null;comment:角色id;index:idx_role_id"`
}

// TableName 指定表名
func (UserRole) TableName() string {
	return "p_user_role"
}
