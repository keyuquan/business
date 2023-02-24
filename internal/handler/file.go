package handler

import (
	"business/internal/model"
	"business/pkg/http"
	"fmt"
	"os"
)

func Upload(c *http.Context) {
	file, _ := c.FormFile("file")
	req := model.UploadReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		http.ResponseError(c, 400, err.Error())
		return
	}
	path := "./upload/"
	os.Mkdir(path, 0777)
	err = c.SaveUploadedFile(file, path+req.Name)
	if err != nil {
		http.ResponseError(c, 400, err.Error())
		return
	}
	http.Response(c, "success", nil)
}

func Download(c *http.Context) {
	filename, err := c.GetQuery("filename")
	if !err {
		http.ResponseError(c, 400, "没有此文件")
	}
	path := "./upload/"
	path += filename
	fmt.Println(path)
	c.File(path)
}
