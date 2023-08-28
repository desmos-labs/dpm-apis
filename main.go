package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/caerus/analytics"
	"github.com/desmos-labs/caerus/runner"
	"github.com/desmos-labs/caerus/utils"
	"github.com/desmos-labs/desmos/v6/app"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/desmos-labs/dpm-apis/caerus"
	"github.com/desmos-labs/dpm-apis/logging"
	"github.com/desmos-labs/dpm-apis/routes"
	linksroutes "github.com/desmos-labs/dpm-apis/routes/links"
)

func main() {
	// Setup Cosmos-related stuff
	app.SetupConfig(sdk.GetConfig())

	// Build the clients
	caerusClient := caerus.NewClientFromEnvVariables()

	// Setup the CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")

	// Build the Gin server
	router := gin.New()
	router.Use(logging.ZeroLog(), gin.Recovery(), cors.New(corsConfig))

	// Build the routes context
	ctx := routes.Context{
		Router: router,
		Caerus: caerusClient,
	}

	// Register the routes
	linksroutes.RegisterWithContext(ctx)

	// Build the HTTP server to be able to shut it down if needed
	runningAddress := utils.GetEnvOr(runner.EnvServerAddress, "0.0.0.0")
	runningPort := utils.GetEnvOr(runner.EnvServerPort, "3000")
	httpServer := &http.Server{
		Addr:              fmt.Sprintf("%s:%s", runningAddress, runningPort),
		Handler:           router,
		ReadHeaderTimeout: time.Minute,
		ReadTimeout:       time.Minute,
		WriteTimeout:      time.Minute,
	}

	// Listen for and trap any OS signal to gracefully shutdown and exit
	go trapSignal(httpServer)

	// Start the HTTP server
	// Block main process (signal capture will call WaitGroup's Done)
	log.Info().Str("address", httpServer.Addr).Msg("Starting API server")
	err := httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

// trapSignal traps the stops signals to gracefully shut down the server
func trapSignal(httpServer *http.Server) {
	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	// Kill (no param) default send syscall.SIGTERM
	// Kill -2 is syscall.SIGINT
	// Kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Debug().Msg("Shutting down API server")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("API server forces to shutdown")
	}

	log.Info().Msg("API server shutdown")

	// Perform the cleanup of other things
	analytics.Stop()
}
