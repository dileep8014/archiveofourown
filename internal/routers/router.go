package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/shyptr/archiveofourown/docs"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/internal/middleware"
	v1 "github.com/shyptr/archiveofourown/internal/routers/api/v1"
	"github.com/shyptr/archiveofourown/internal/routers/api/v1/admin"
	"github.com/shyptr/archiveofourown/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBucket(limiter.BucketRule{
	Key:          "/api/v1/register",
	FillInterval: time.Second,
	Capacity:     60,
	Quantum:      20,
}, limiter.BucketRule{
	Key:          "/api/v1/currentUser",
	FillInterval: time.Second,
	Capacity:     60,
	Quantum:      20,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(
		middleware.CORS(),
		middleware.Tracing(),
		middleware.AccessLog(),
		middleware.Recovery(),
		middleware.RateLimiter(methodLimiters),
		middleware.ContextTimeout(time.Duration(global.AppSetting.ContextTimeout)*time.Second),
	)
	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// file resource
	r.GET("/image/*any", func(ctx *gin.Context) {
		http.FileServer(http.Dir("./storage")).ServeHTTP(ctx.Writer, ctx.Request)
	})
	// common
	apiv1 := r.Group("/api/v1", middleware.Jwt())
	{
		v1.NewUser().Router(apiv1)
		v1.NewCalendar().Router(apiv1)
		v1.NewCategory().Router(apiv1)
		v1.NewChapter().Router(apiv1)
		v1.NewCollege().Router(apiv1)
		v1.NewComment().Router(apiv1)
		v1.NewMessage().Router(apiv1)
		v1.NewNews().Router(apiv1)

		// image upload
		apiv1.POST("/upload", v1.NewFile().Upload)
	}

	// admin
	adminApiv1 := r.Group("/api/v1/admin", middleware.Jwt(), middleware.Root())
	{
		admin.NewCategory().Router(adminApiv1)
		admin.NewChapter().Router(adminApiv1)
		admin.NewNews().Router(adminApiv1)
	}

	return r
}
