package main

import (
	"github.com/nabinkatwal7/irlquest/routes"
	"github.com/nabinkatwal7/irlquest/utils"
)

func main() {
	utils.LoadEnv()
	utils.LoadDatabase()
	routes.ServeApplication()
}
