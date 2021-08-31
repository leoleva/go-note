package main

import (
	"demoproject/src/container"
	"demoproject/src/router"
)

func main()  {

	ctr := container.Load()

	router.RegisterRoutes(ctr)
}
