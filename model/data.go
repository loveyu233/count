package model

import (
	"count/utils/pgsql"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Data struct {
	gorm.Model
	Content    string  `gorm:"column:content" json:"content"`
	Size       float64 `gorm:"column:size" json:"size"`
	FormatDate string  `gorm:"-" json:"format_date"`
}

type DataTime struct {
	gorm.Model
	YearMonth string  `gorm:"column:year_month;uniqueIndex:data_time_year_month" json:"year_month"`
	Size      float64 `gorm:"column:size" json:"size"`
}

func (d *Data) Insert() error {
	d.ID = uint(time.Now().UnixMilli())
	pgsql.DB.Save(d)
	year, month, _ := time.Now().Date()
	yearMonth := fmt.Sprintf("%d-%d", year, month)
	pgsql.DB.Model(&DataTime{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "year_month"}},
		DoUpdates: clause.Assignments(map[string]any{"size": gorm.Expr("data_times.size+?", d.Size)}),
	}).Create(&DataTime{YearMonth: yearMonth, Size: d.Size})
	return nil
}
