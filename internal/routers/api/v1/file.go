package v1

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	fileTypeMap = map[string]string{
		"ffd8ffe000104a464946": "jpg",
		"89504e470d0a1a0a0000": "png",
	}
)

type File struct {
}

func NewFile() File {
	return File{}
}

// @Summary 上传图片
// @Tags 文件
// @Description 上传图片
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "文件" image
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/upload [post]
func (f File) Upload(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	// 获取文件
	file, err := ctx.FormFile("file")
	if err != nil {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(err.Error()))
		return
	}
	// 校验文件信息
	openFile, err := file.Open()
	if err != nil {
		res.ToErrorResponse(errcode.ErrorFileUpload.WithError(err))
		return
	}
	defer openFile.Close()
	src, err := ioutil.ReadAll(openFile)
	if err != nil {
		res.ToErrorResponse(errcode.ErrorFileUpload.WithError(err))
		return
	}
	if GetFileType(src) == "" {
		res.ToErrorResponse(errcode.ErrorFileUpload.WithDetails("文件类型错误，只接受png/jpg文件"))
		return
	}
	if file.Size/1024/1024 > 2 {
		res.ToErrorResponse(errcode.ErrorFileUpload.WithDetails("图片必须小于2M"))
		return
	}
	// 存储文件
	// md5 hash取名
	fileName := fmt.Sprintf("%x", md5.Sum(src))
	saveFile, err := os.Create("./storage/image/" + fileName)
	if err != nil {
		res.ToErrorResponse(errcode.ErrorFileUpload.WithError(err))
		return
	}
	defer saveFile.Close()
	n, err := saveFile.Write(src)
	if err != nil {
		res.ToErrorResponse(errcode.ErrorFileUpload.WithError(err))
		return
	}
	if n != len(src) {
		res.ToErrorResponse(errcode.ErrorFileUpload.WithError(err))
		return
	}

	res.ToResponse(global.ServerSetting.BaseUrl + "/image/" + fileName)
}

// 获取前面结果字节的二进制
func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	for _, v := range src {
		sub := v & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
}

// 用文件前面几个字节来判断
// fSrc: 文件字节流（就用前面几个字节）
func GetFileType(fSrc []byte) string {
	var fileType string
	fileCode := bytesToHexString(fSrc)

	for k, v := range fileTypeMap {
		if strings.HasPrefix(fileCode, strings.ToLower(k)) ||
			strings.HasPrefix(k, strings.ToLower(fileCode)) {
			fileType = v
		}
	}

	return fileType
}
