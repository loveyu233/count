package test

import (
	"count/model"
	"count/utils/pgsql"
	"encoding/json"
	"fmt"
	"testing"
)

func init() {
	dsn := "host=127.0.0.1 user=pgsql password=pgsql dbname=pgsql port=5432 sslmode=disable"

	pgsql.InitPgsql(dsn)
	model.InitModel()
}

func TestA(t *testing.T) {
	var datas []*model.Data
	m := make(map[string][]*model.Data)
	pgsql.DB.Model(&model.Data{}).Order("created_at desc").Find(&datas)
	for _, data := range datas {
		data.FormatDate = data.Model.CreatedAt.Format("2006-01-02")
		if _, ok := m[data.FormatDate]; ok {
			m[data.FormatDate] = append(m[data.FormatDate], data)
		} else {
			m[data.FormatDate] = []*model.Data{data}
		}
	}
	marshal, _ := json.Marshal(m)
	fmt.Println(string(marshal))
}

func TestBBB(t *testing.T) {
	data := model.Data{
		Content: "123",
		Size:    600,
	}
	data.Insert()
}
