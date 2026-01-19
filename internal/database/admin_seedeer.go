package database

import (
	"github.com/Simok666/ecommerce-app.git/internal/models"
	"github.com/Simok666/ecommerce-app.git/internal/services"
	"github.com/Simok666/ecommerce-app.git/logger"
)

func SeedAdmin() {
	var count int64
	var log = logger.InitLogger()

	adminEmail := "admin@gmail.com"
	adminPassword := "admin12345"
	adminName := "Super Admin"

	if adminEmail == "" || adminPassword == "" || adminName == "" {
		log.Info("Admin seeder skipped: env not set")
		return
	}

	DB.Model(&models.User{}).
		Where("email = ?", adminEmail).
		Count(&count)

	if count > 0 {
		log.Info("Admin already exists, skipping seeder")
		return
	}

	hashedPassword, err := services.HashPassword(adminPassword)
	if err != nil {
		log.Info("Failed to hash admin password:", err)
	}

	admin := models.User{
		Name:     adminName,
		Email:    adminEmail,
		Password: hashedPassword,
		Role:     "admin",
	}

	if err := DB.Create(&admin).Error; err != nil {
		log.Info("Failed to seed admin:", err)
	}

	log.Info("Admin user seeded successfully")

}
