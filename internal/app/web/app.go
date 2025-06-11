package web

//
//import (
//	"awesomeProject/internal/clients/rest"
//	"awesomeProject/internal/domain"
//	"log/slog"
//	"net/http"
//	"strconv"
//)
//
//type App struct {
//	port    int
//	request *rest.Requester
//	log     *slog.Logger
//}
//
//func New(port int, use domain.ResponderUseCase, log *slog.Logger) *App {
//
//	request := rest.New(use)
//
//	return &App{port: port, request: request, log: log}
//}
//
//func (a *App) MustRun() {
//	const op = "webApp.Run"
//
//	addr := createAddress(a.port)
//
//	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))
//
//	http.HandleFunc("/", a.request.ResponseHandler)
//
//	log.Info("starting WebServer", slog.String("address:", addr))
//
//	err := http.ListenAndServe(addr, nil)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func createAddress(port int) string {
//	return "localhost:" + strconv.Itoa(port)
//}
