package ListenAndServer

import (
	"fmt"
	"net/http"
	"os"
	"todo_list/conf"
	"todo_list/router"
)

func OpenSever() {
	r := router.InitRouter()
	addr := fmt.Sprintf("%s:%s", conf.Address, conf.ServerPort)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error Occurred When ListenAndServer", err)
		os.Exit(-1)
	}
}
