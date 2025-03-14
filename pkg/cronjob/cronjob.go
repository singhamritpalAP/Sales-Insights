package cronjob

import (
	cron "github.com/robfig/cron/v3"
	"log"
	"sales/internal/constants"
	"sales/internal/services"

	"gorm.io/gorm"
)

// SetupCronJob sets up the cron job for refreshing the database.
func SetupCronJob(db *gorm.DB) {
	c := cron.New()
	_, err := c.AddFunc(constants.CronTime, func() {
		err := services.RefreshDatabase(db)
		if err != nil {
			log.Println("Error refreshing database:", err)
		} else {
			log.Println("Database refreshed successfully via cron.")
		}
	})
	if err != nil {
		log.Println("Error adding cron function:", err)
	}

	c.Start()
}
