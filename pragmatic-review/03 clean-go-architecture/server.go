package main

import (
	"clean-go-architecture/controller"
	router "clean-go-architecture/http"
	"fmt"
	"net/http"
)

var (
	postController controller.PostController = controller.NewPostController()
	httpRouter     router.Router             = router.NewChiRouter()
	// httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
