package requests

type PermissionRequest struct {
	UserID uint `json:"user_id" validate:"required"`
	RoleID uint `json:"role_id" validate:"required"`
}
type PermissionUserRequest struct {
	PermissionID uint `json:"permission_id" validate:"required"`
	UserID       uint `json:"user_id" validate:"required"`
}
