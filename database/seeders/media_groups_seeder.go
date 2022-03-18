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

    seed.Add("SeedMediaGroupsTable", func(db *gorm.DB) {

        mediaGroups  := factories.MakeMediaGroups(10)

        result := db.Table("media_groups").Create(&mediaGroups)

        if err := result.Error; err != nil {
            logger.LogIf(err)
            return
        }

        console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
    })
}