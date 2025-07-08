package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"warehouse/routes"
	"warehouse/utils"
)

func main() {
	// TODO: look into:
	// https://developer.nvidia.com/gpugems/gpugems3/part-vi-gpu-computing/chapter-37-efficient-random-number-generation-and-application

	// Not needed since go 1.5:
	// https://www.soroushjp.com/2015/02/07/go-concurrency-is-not-parallelism-real-world-lessons-with-monte-carlo-simulations/
	//runtime.GOMAXPROCS(runtime.NumCPU())

	if !utils.IsRunningInDocker() {
		err := godotenv.Load("backend.env")
		if err != nil {
			fmt.Println("Error loading .env file")
		}
	}

	// Create server
	e := echo.New()

	// Middleware
	if utils.GetEnvBool("IS_DEVELOPMENT", true) {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
		fmt.Println("Running in development mode")
	} else {
		// TODO: when thinking of deploying to prod, change
		//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//	AllowOrigins: []string{""},
		//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		//}))
		fmt.Println("Running in production mode")
	}

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Inserting into an db is to slow and takes to much space to be viable
	// TODO: use an simulation seed instead to save the simulation and to re run it on the fly

	// Define routes
	e.POST("/restapi/sim/run", routes.RunSimulation)

	// Start server
	e.Logger.Fatal(e.Start(utils.GetEnvString("PORT", "0.0.0.0:4000")))
}
