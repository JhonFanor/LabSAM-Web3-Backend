package middlewares

import (
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionMiddleware struct {
	permissionService services.PermissionService
	roleService       services.RoleService
}

func NewPermissionMiddleware(ps services.PermissionService, rs services.RoleService) *PermissionMiddleware {
	return &PermissionMiddleware{
		permissionService: ps,
		roleService:       rs,
	}
}

func (pm *PermissionMiddleware) RequirePermission(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsValue, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No se encontraron los claims del token"})
			return
		}

		claims, ok := claimsValue.(*security.Claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Formato inválido de claims"})
			return
		}

		role, err := pm.roleService.GetByName(claims.Role)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el rol"})
			return
		}

		permissions, err := pm.permissionService.GetAllPermissionsByUser(claims.UserID, role.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener permisos"})
			return
		}

		for _, p := range permissions {
			if p.Name == requiredPermission {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No tienes permiso para acceder a esta ruta"})
	}
}
