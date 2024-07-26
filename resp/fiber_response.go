package resp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	MSG200 = "请求成功"
	MSG202 = "请求成功, 请稍后..."
	MSG400 = "请求参数错误"
	MSG401 = "登录已过期, 请重新登录"
	MSG403 = "请求权限不足"
	MSG404 = "请求资源未找到"
	MSG429 = "请求过于频繁, 请稍后再试"
	MSG500 = "服务器开小差了, 请稍后再试"
	MSG501 = "功能开发中, 尽情期待"
)

type Error struct {
	Status uint32
	Msg    string
	Err    error
	Data   any
}

func NewError(status uint32, msg string, err error, data any) Error {
	return Error{
		Status: status,
		Msg:    msg,
		Err:    err,
		Data:   data,
	}
}

func (e Error) Error() string {
	if e.Msg != "" {
		return e.Msg
	}

	switch e.Status {
	case 200:
		return MSG200
	case 202:
		return MSG202
	case 400:
		return MSG400
	case 401:
		return MSG401
	case 403:
		return MSG403
	case 404:
		return MSG404
	case 429:
		return MSG429
	case 500:
		return MSG500
	case 501:
		return MSG501
	}

	return e.Err.Error()
}

func handleEmptyMsg(status uint32, msg string) string {
	if msg == "" {
		switch status {
		case 200:
			msg = MSG200
		case 202:
			msg = MSG202
		case 400:
			msg = MSG400
		case 401:
			msg = MSG401
		case 403:
			msg = MSG403
		case 404:
			msg = MSG404
		case 429:
			msg = MSG429
		case 500:
			msg = MSG500
		case 501:
			msg = MSG501
		}
	}

	return msg
}

func Resp(c *fiber.Ctx, status uint32, msg string, err string, data any) error {
	msg = handleEmptyMsg(status, msg)

	c.Set("SONAR-STATUS", strconv.Itoa(int(status)))

	if data == nil {
		return c.JSON(fiber.Map{"status": status, "msg": msg, "err": err})
	}

	return c.JSON(fiber.Map{"status": status, "msg": msg, "err": err, "data": data})
}

func RespError(c *fiber.Ctx, err error) error {
	if err == nil {
		return Resp(c, 500, MSG500, "response with nil error", nil)
	}

	var re = &Error{}
	if errors.As(err, re) {
		if re.Err == nil {
			return Resp(c, re.Status, re.Msg, re.Msg, re.Data)
		}

		return Resp(c, re.Status, re.Msg, re.Err.Error(), re.Data)
	}

	return Resp(c, 500, MSG500, err.Error(), nil)
}

func Resp200(c *fiber.Ctx, data any, msgs ...string) error {
	msg := MSG200

	if len(msgs) > 0 && msgs[0] != "" {
		msg = fmt.Sprintf("%s: %s", msg, strings.Join(msgs, "; "))
	}

	return Resp(c, 200, msg, "", data)
}

func Resp202(c *fiber.Ctx, data any, msgs ...string) error {
	msg := MSG202

	if len(msgs) > 0 && msgs[0] != "" {
		msg = fmt.Sprintf("%s: %s", msg, strings.Join(msgs, "; "))
	}

	return Resp(c, 202, msg, "", data)
}

func Resp400(c *fiber.Ctx, data any, msgs ...string) error {
	msg := MSG400
	err := ""

	if len(msgs) > 0 && msgs[0] != "" {
		msg = fmt.Sprintf("%s: %s", msg, strings.Join(msgs, "; "))
		err = msg
	}

	return Resp(c, 400, msg, err, data)
}

func Resp401(c *fiber.Ctx, data any, msgs ...string) error {
	msg := MSG401
	err := ""

	if len(msgs) > 0 && msgs[0] != "" {
		msg = fmt.Sprintf("%s: %s", msg, strings.Join(msgs, "; "))
		err = msg
	}

	return Resp(c, 401, msg, err, data)
}

func Resp403(c *fiber.Ctx, data any, msgs ...string) error {
	msg := MSG403
	err := ""

	if len(msgs) > 0 && msgs[0] != "" {
		msg = fmt.Sprintf("%s: %s", msg, strings.Join(msgs, "; "))
		err = msg
	}

	return Resp(c, 403, msg, err, data)
}

func Resp429(c *fiber.Ctx, data any, msgs ...string) error {
	msg := MSG429
	err := ""

	if len(msgs) > 0 && msgs[0] != "" {
		msg = fmt.Sprintf("%s: %s", msg, strings.Join(msgs, "; "))
		err = ""
	}

	return Resp(c, 429, msg, err, data)
}

func Resp500(c *fiber.Ctx, data any, msgs ...string) error {
	msg := MSG500
	err := ""

	if len(msgs) > 0 && msgs[0] != "" {
		msg = fmt.Sprintf("%s: %s", msg, strings.Join(msgs, "; "))
		err = msg
	}

	return Resp(c, 500, msg, err, data)
}
