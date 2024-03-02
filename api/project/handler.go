package project

import (
	"DRSP/common"
	"DRSP/internal/openai"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

//业务路由处理函数

type HandlerProject struct {
}

type RequestData struct {
	Question string `json:"question"`
}

type ResponseData struct {
	Reply string `json:"reply"`
}

func NewHandlerProject() *HandlerProject {
	return &HandlerProject{}
}

func (hp *HandlerProject) upload(ctx *gin.Context) {
	resp := &common.Result{}
	reply := ""
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Println(err.Error())
		return
	}
	if len(form.File) <= 0 {
		log.Println("文件为空！")
		ctx.JSON(http.StatusOK, resp.Fail(common.FileIsNull.Code, common.FileIsNull.Msg))
		return
	}
	text := make(map[string]string)
	for filePath, files := range form.File {
		for _, file := range files {
			filename := filepath.Base(file.Filename)
			absPath := fmt.Sprintf("./pkg/upload/%s%s", filePath, filename)
			if err := ctx.SaveUploadedFile(file, absPath); err != nil {
				log.Println("保存失败！")
				continue
			}
			text[filename], err = common.Ocr(absPath)
			if err != nil {
				log.Println(err.Error())
			}
			reply = fmt.Sprintf("%s%s:%s\n", reply, filename, text[filename])
			fmt.Println("reply:", reply)
		}
	}
	ctx.JSON(http.StatusOK, resp.Success(reply))
}

func (hp *HandlerProject) chat(ctx *gin.Context) {
	var requestData RequestData
	resp := &common.Result{}

	// 解析请求体中的JSON数据到requestData结构体中
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, resp.Fail(common.DataParseFail.Code, common.DataParseFail.Msg))
		return
	}
	fmt.Println(requestData.Question)
	reply := openai.GetGpt().Chat(requestData.Question)
	fmt.Println(reply)
	ctx.JSON(http.StatusOK, resp.Success(reply))
}
