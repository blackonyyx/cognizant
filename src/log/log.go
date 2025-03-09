package log

import (
	"log"

	"github.com/gin-gonic/gin"
)


// normal
func CtxInfo(ctx *gin.Context, msg ...string) {
	path := ctx.FullPath()
	log.Println("[INFO]: ", path, msg)
}

// 204
func CtxWarning (ctx *gin.Context, msg ...string) {
	path := ctx.FullPath()
	log.Println("[WARN]: ", path, msg)
}

// All else
func CtxError (ctx *gin.Context, errorCode int, msg ...string) {
	path := ctx.FullPath()
	log.Println("[ERROR]", path, "error:", errorCode, msg)
}

// normal
func Info(msg ...string) {
	log.Println("[INFO]: ", msg)
}

// 204
func Warning (msg ...string) {
	log.Println("[WARN]: ", msg)
}

// All else
func Error (errorCode int, msg ...string) {
	log.Println("[ERROR]","error:", errorCode, msg)
}