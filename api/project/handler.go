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

func NewHandlerProject() *HandlerProject {
	return &HandlerProject{}
}

func (hp *HandlerProject) upload(ctx *gin.Context) {
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
			fmt.Println(text[filename])
			ctx.JSON(http.StatusOK, text[filename])
		}
	}

}
