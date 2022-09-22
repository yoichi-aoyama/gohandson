package main

import (
	"fmt"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()  // echo を利用する
	// GET リクエストでパスが `/` のとき第２引数の関数を実行する
	e.GET("/", func(c echo.Context) error {
		key := "count"
		conn, err := redis.Dial("tcp", "redis:6379")
		if err != nil {
			panic(err)
		}

		_, err = conn.Do("INCR", key)
		if err != nil {
			panic(err)
		}

		s, err := redis.String(conn.Do("GET", key))
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		return c.String(http.StatusOK, fmt.Sprintf("%s", s))
	})

	// 1323 ポートでリッスンを開始。 start がエラーを起こしたら Fatal を起こしてログに記録する
	e.Logger.Fatal(e.Start(":1323"))
}