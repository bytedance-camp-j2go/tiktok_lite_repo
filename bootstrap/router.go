package bootstrap

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/middleware"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
}
