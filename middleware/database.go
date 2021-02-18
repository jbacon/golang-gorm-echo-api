package middleware

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// DatabaseEchoContext struct
type DatabaseEchoContext struct {
	echo.Context
	DB *gorm.DB
}

// DatabaseMiddleware - exported
func DatabaseMiddleware(db *gorm.DB) func(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			customEchoContext := &DatabaseEchoContext{context, db}
			return handler(customEchoContext)
		}
	}
}
