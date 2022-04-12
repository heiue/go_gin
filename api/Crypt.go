package api

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin/tools"
)

type Param struct {
	ResultKey string `form:"result_key"`
}

const key = "AxRwdPuI3ZWonhtB"
// 14219930231 => 8p9P8VS5AoMZzH4=

// EnCrypt 加密
func EnCrypt(ctx *gin.Context) {
	var res Param
	err := ctx.ShouldBindQuery(&res)
	if err != nil {
		panic("shouldBind error")
	}
	r := tools.Crypt([]byte(res.ResultKey), []byte(key))
	ctx.JSON(0, gin.H{
		"encrypt": r,
	})
}

// DeCrypt 解密
func DeCrypt(ctx *gin.Context) {
	var res Param
	err := ctx.ShouldBindQuery(&res)
	if err != nil {
		panic("shouldBind error")
	}
	resultBase, _ := base64.StdEncoding.DecodeString(res.ResultKey)
	fmt.Println(resultBase)
	r := tools.Crypt([]byte(resultBase), []byte(key))
	fmt.Println(base64.StdEncoding.EncodeToString(r))
	fmt.Println(string(r))
	ctx.JSON(0, gin.H{
		"decrypt": r,
	})
}
