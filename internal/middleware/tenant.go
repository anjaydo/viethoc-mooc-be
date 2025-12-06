package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TenantIDKey is the context key for storing the Tenant ID
const TenantIDKey = "tenant_id"

// RequireTenant Middleware enforces that a Tenant ID is present in the headers/token.
// For the MVP, we assume the API Gateway or Auth Service has already validated the JWT
// and passed the Tenant ID as a header, OR we extract it from the JWT here.
//
// In this MVP "Wizard of Oz" phase, we might simulate it via a header for testing.
func RequireTenant() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Try to get from Header (e.g., from an upstream gateway or for testing)
		tenantID := c.GetHeader("X-Tenant-ID")

		// 2. TODO: In production, extract this from the Supabase JWT claims:
		// claims := jwt.ExtractClaims(c.Request)
		// tenantID = claims["app_metadata"].(map[string]interface{})["tenant_id"]

		if tenantID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Tenant ID"})
			c.Abort()
			return
		}

		// 3. Set into context for Controllers to use
		c.Set(TenantIDKey, tenantID)

		c.Next()
	}
}
