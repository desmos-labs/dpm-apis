package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/desmos-labs/dpm-apis/caerus"
)

// Context contains all the data that can be useful while registering routes
type Context struct {
	Router *gin.Engine
	Caerus *caerus.Client
}
