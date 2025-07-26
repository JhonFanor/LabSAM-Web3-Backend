package models

type DeniedPermissionUser struct {
	PermissionID uint `json:"permission_id"`
	UserID       uint `json:"user_id"`
}

func (DeniedPermissionUser) TableName() string {
	return "denied_permissions_user"
}
