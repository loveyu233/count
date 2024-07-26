package handler

import (
	"count/model"
	"count/resp"
	"count/utils/pgsql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func Add() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type Req struct {
			Content string `json:"content"`
			Size    string `json:"size"`
		}
		var (
			req  = new(Req)
			data = new(model.Data)
			err  error
		)
		if err = c.BodyParser(req); err != nil {
			return resp.Resp400(c, "参数错误")
		}
		data.Content = req.Content
		data.Size, _ = strconv.ParseFloat(req.Size, 64)
		if err = data.Insert(); err != nil {
			return resp.Resp500(c, err.Error())
		}
		return resp.Resp200(c, data)
	}
}

func GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type Resp struct {
			List    map[string][]*model.Data `json:"list"`
			Current float64                  `json:"current"`
			Last    float64                  `json:"last"`
		}
		var datas []*model.Data
		var dataTime []*model.DataTime
		m := make(map[string][]*model.Data)
		pgsql.DB.Model(&model.Data{}).Order("created_at desc").Find(&datas)
		for _, data := range datas {
			data.FormatDate = data.Model.CreatedAt.UTC().Format("2006-01-02")
			if _, ok := m[data.FormatDate]; ok {
				m[data.FormatDate] = append(m[data.FormatDate], data)
			} else {
				m[data.FormatDate] = []*model.Data{data}
			}
		}
		year, month, _ := time.Now().Date()
		t1 := fmt.Sprintf("%d-%d", year, month)
		t2 := fmt.Sprintf("%d-%d", year, month-1)
		pgsql.DB.Model(&model.DataTime{}).Select("year_month", "size").Where("year_month = ? or year_month = ?", t1, t2).Find(&dataTime)
		respData := &Resp{
			List: m,
		}
		for _, d := range dataTime {
			if d.YearMonth == t1 {
				respData.Current = d.Size
			}
			if d.YearMonth == t2 {
				respData.Last = d.Size
			}
		}
		return resp.Resp200(c, respData)
	}
}
