package service

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// QueryInt ...数値型でクエリを取得します。デフォルト値は0で、クエリが存在しない場合はエラーがnil以外を返します。
func QueryInt(c *gin.Context, key string) (int, error) {
	if s, ok := c.GetQuery(key); ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		return n, nil
	}
	return 0, errors.New("query not found")
}

// QueryString ...文字列型でクエリを取得します。クエリが存在しない場合はエラーがnil以外を返します。
func QueryString(c *gin.Context, key string) (string, error) {
	if s, ok := c.GetQuery(key); ok {
		return s, nil
	}
	return "", errors.New("query not found")
}
