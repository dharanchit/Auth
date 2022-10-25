package initializers

import (
	"github.com/dharanchit/auth/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.UserGo{})
}
