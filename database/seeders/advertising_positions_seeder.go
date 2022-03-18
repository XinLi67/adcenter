package seeders

import (
    "fmt"
    "adcenter/database/factories"
    "adcenter/pkg/console"
    "adcenter/pkg/logger"
    "adcenter/pkg/seed"

    "gorm.io/gorm"
)

func init() {

    seed.Add("SeedAdvertisingPositionsTable", func(db *gorm.DB) {

        advertisingPositions  := factories.MakeAdvertisingPositions(10)

        result := db.Table("advertising_positions").Create(&advertisingPositions)

        if err := result.Error; err != nil {
            logger.LogIf(err)
            return
        }

        console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
    })
}