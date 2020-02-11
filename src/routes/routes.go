package routes

func (server *ServiceServer) setRoutes() {
	erver.Router.HandleFunc("/channel", middleware.setJSONNormailDataAtMiddleware(channel.getChannels)).Methods("GET")
}

/*
func (server *ServiceServer) setRoutes() {

	server.Router.HandleFunc("/channel", middleware.setJSONNormailDataAtMiddleware(server.CreatePost)).Methods("GET")
}
*/
