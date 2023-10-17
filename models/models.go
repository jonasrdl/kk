package models

import (
	"kk/config"
	"log"
)

type Indexcard struct {
	ID       int    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

var Models = []interface{}{
	&Indexcard{},
}

func AutoMigrate() {
	for _, model := range Models {
		if err := config.DB.AutoMigrate(model); err != nil {
			log.Fatalf("Auto-migration failed for model %T: %s", model, err)
		}
	}
}
