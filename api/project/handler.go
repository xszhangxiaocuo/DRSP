package project

import (
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

	for filePath, files := range form.File {
		for _, file := range files {
			filename := filepath.Base(file.Filename)
			absPath := fmt.Sprintf("pkg/upload/%s%s", filePath, filename)
			log.Println("保存路径：", absPath)
			if err := ctx.SaveUploadedFile(file, absPath); err != nil {
				log.Println("保存失败！")
				continue
			}
		}
	}
	ctx.JSON(http.StatusOK, "ok")
}
