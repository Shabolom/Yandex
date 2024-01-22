package migrate

import (
	"YandexPra/config"
	"YandexPra/iternal/domain"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gormigrate.v1"
)

// Migrate run migration for all entities and add constraints for them,
// создаем таблицы и закидываем в бл тут
func Migrate() {
	db := config.DB
	id, _ := uuid.NewV4()

	// создаем объект миграции данная строка всегда статична (всегда такая)
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			// id всех миграций кторые были проведены
			ID: id.String(),
			// переписываем так при создании таблицы изменяется только структура которую мы передаем
			Migrate: func(tx *gorm.DB) error {
				err := tx.AutoMigrate(&domain.Urls{}).Error
				if err != nil {
					return err
				}
				return nil
			},
			// это метод отмены миграции ни разу не использовал
			Rollback: func(tx *gorm.DB) error {
				err := tx.DropTable("urls").Error
				if err != nil {
					return err
				}
				return nil
			},
		},
	})

	err := m.Migrate()
	if err != nil {
		panic(err)
	}

	if err == nil {
		println("Migration did run successfully")
	} else {
		println("Could not migrate: %v", err)
	}
}
