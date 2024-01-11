package project

import (
	"DRSP/common"
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
	reply := ""
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Println(err.Error())
		return
	}
	if len(form.File) <= 0 {
		log.Println("文件为空！")
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
			text[filename] = common.Ocr(absPath)
			reply = fmt.Sprintf("%s%s:%s\n", reply, filename, text[filename])
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"reply": reply,
	})
}

func (hp *HandlerProject) chat(ctx *gin.Context) {
	var requestData RequestData
	//var responseData ResponseData

	// 解析请求体中的JSON数据到requestData结构体中
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(requestData.Question)
	//responseData.Reply = openai.GetGpt().Chat(requestData.Question)
	//fmt.Println(responseData.Reply)
	//ctx.JSON(http.StatusOK, responseData)
	ctx.JSON(http.StatusOK, gin.H{
		"reply": "返回响应",
	})
}
