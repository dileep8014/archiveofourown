package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/shyptr/archiveofourown/docs"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/internal/middleware"
	v1 "github.com/shyptr/archiveofourown/internal/routers/api/v1"
	"github.com/shyptr/archiveofourown/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBucket(limiter.BucketRule{
	Key:          "/login",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(
		middleware.Tracing(),
		middleware.AccessLog(),
		middleware.Tx(),
		middleware.Recovery(),
		middleware.RateLimiter(methodLimiters),
		middleware.ContextTimeout(time.Duration(global.AppSetting.ContextTimeout)*time.Second),
		middleware.Translations(),
	)

	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//auth
	r.GET("/auth", v1.NewAuth().Get)
	r.GET("/auth/:id", v1.NewAuth().Get)

	apiv1 := r.Group("/api/v1", middleware.Jwt())
	{
		v1.NewCategory().Router(apiv1)
		v1.NewArticle().Router(apiv1)
	}

	return r
}
