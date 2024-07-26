package model

import "count/utils/pgsql"

func InitModel() {
	pgsql.DB.AutoMigrate(&Data{}, &DataTime{})
}
