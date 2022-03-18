package seeders

import (
	"adcenter/database/factories"
	"adcenter/pkg/console"
	"adcenter/pkg/logger"
	"adcenter/pkg/seed"
	"fmt"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedAdvertisingsTable", func(db *gorm.DB) {

		advertisings := factories.MakeAdvertisings(10)

		result := db.Table("advertisings").Create(&advertisings)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
