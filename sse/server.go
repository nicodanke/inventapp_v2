package sse

import (
	"fmt"
	"net"
	"net/http"

	"github.com/nicodanke/inventapp_v2/token"
	"github.com/nicodanke/inventapp_v2/utils"
	"github.com/rs/zerolog/log"
)

type Server struct {
	config       utils.Config
	tokenMaker   token.Maker
	handlerEvent *HandlerEvent
	mux          *http.ServeMux
	listener     net.Listener
}

func NewServer(config utils.Config, handlerEvent *HandlerEvent) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{tokenMaker: tokenMaker, config: config, handlerEvent: handlerEvent}

	server.setupRoutes()

	server.setupListener()

	return server, nil
}

func (server *Server) setupRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/events", server.handlerEvent.Handler)
	// mux.HandleFunc("/test1", HandlerTest1(server.handlerEvent))
	// mux.HandleFunc("/test2", HandlerTest2(server.handlerEvent))
	// mux.HandleFunc("/test3", HandlerTest3(server.handlerEvent))
	// mux.HandleFunc("/test4", HandlerTest4(server.handlerEvent))

	server.mux = mux
}

func (server *Server) setupListener() {
	listener, err := net.Listen("tcp", server.config.SSEAddress)
	if err != nil {
		LogError(err, "cannot create listener")
	}

	server.listener = listener
}

func (server *Server) Start(address string) error {
	log.Info().Msgf("SSE Server started at: %s", server.listener.Addr().String())
	return http.Serve(server.listener, server.mux)
}

// func HandlerTest1(notifier Notifier) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var data = map[string]any{}
// 		json.NewDecoder(r.Body).Decode(&data)
// 		notifier.BroadcastMessage(EventMessage{
// 			EventName: "greeting",
// 			Data:      data,
// 		})
// 	}
// }

// func HandlerTest2(notifier Notifier) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var data = map[string]any{}
// 		json.NewDecoder(r.Body).Decode(&data)
// 		notifier.BoadcastMessageToAccount(EventMessage{
// 			EventName: "greeting",
// 			Data:      data,
// 		}, 1)
// 	}
// }

// func HandlerTest3(notifier Notifier) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var data = map[string]any{}
// 		json.NewDecoder(r.Body).Decode(&data)
// 		notifier.BoadcastMessageToAccount(EventMessage{
// 			EventName: "greeting",
// 			Data:      data,
// 		}, 2)
// 	}
// }

// func HandlerTest4(notifier Notifier) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var data = map[string]any{}
// 		json.NewDecoder(r.Body).Decode(&data)
// 		notifier.SendMessageToUser(EventMessage{
// 			EventName: "greeting",
// 			Data:      data,
// 		}, 2)
// 	}
// }
