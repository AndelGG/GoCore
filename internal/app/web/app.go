package web

import (
	rest "awesomeProject/internal/clients/http"
	"awesomeProject/internal/usecases"
	"fmt"
	"net/http"
	"strconv"
)

type App struct {
	addr    string
	request *rest.Requester
}

func New(port int, use *usecases.ResponderUseCase) *App {

	addr := "localhost:" + strconv.Itoa(port)

	request := rest.New(use)

	return &App{addr: addr, request: request}
}

func (a *App) MustRun() {
	http.HandleFunc("/", a.request.ResponseHandler)
	fmt.Println("Server is running on " + a.addr)
	err := http.ListenAndServe(a.addr, nil)
	if err != nil {
		panic(err)
	}
}
