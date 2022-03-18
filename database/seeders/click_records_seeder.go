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

    seed.Add("SeedClickRecordsTable", func(db *gorm.DB) {

        clickRecords  := factories.MakeClickRecords(10)

        result := db.Table("click_records").Create(&clickRecords)

        if err := result.Error; err != nil {
            logger.LogIf(err)
            return
        }

        console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
    })
}