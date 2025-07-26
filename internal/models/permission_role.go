package models

type PermissionRole struct {
	PermissionID uint `gorm:"primary_key" json:"permission_id"`
	RoleID       uint `gorm:"primary_key" json:"role_id"`
}

func (PermissionRole) TableName() string {
	return "permission_rol"
}
