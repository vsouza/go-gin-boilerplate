package tests

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vsouza/go-gin-boilerplate/config"
	"github.com/vsouza/go-gin-boilerplate/controllers"
)

func Test(t *testing.T) { Testing(t) }

var _ = Suite(&UserSuite{})

type UserSuite struct {
	config *viper.Viper
	router *gin.Engine
}

func (s *UserSuite) SetUpTest(c *C) {
	config.Init("test")
	s.config = config.GetConfig()
	s.router = SetupRouter()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	health := new(controllers.HealthController)
	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.Retrieve)
		}
	}
	return router
}

func TestMain(m *testing.M) {
	SetupRouter()
}
