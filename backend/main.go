// https://gin-gonic.com/docs/quickstart/
package main

import (
	"github.com/gin-gonic/gin"

	"backend/infrastructure"
	"backend/interface/handlers"
	"backend/router"
	"backend/usecase"
	"backend/utils"
)

func main() {
	// infrastructure
	memoRepository := infrastructure.NewMemoRepository(utils.NewDB())
	// usecase
	memoUsecase := usecase.NewMemoUsecase(memoRepository)
	// interface
	memoHandler := handlers.NewMemoHandler(memoUsecase)

	r := gin.Default()

	router.NewRouter(r, memoHandler)

	r.Run(":8080")
}
