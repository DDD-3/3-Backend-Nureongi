package server

func (server *ServiceServer) setRoutes() {
	server.router.HandleFunc("/channel", middleware.setJSONNormailDataAtMiddleware(channel.getChannels)).Methods("GET")
}
