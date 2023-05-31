package utils

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func GetBinary(ctx *gin.Context) ([]byte, error) {
	bys, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}
	return bys, nil
}
