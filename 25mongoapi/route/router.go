package router

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("api/movies", controller.GetMovies).Methods("GET")
	router.HandleFunc("api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("api/movie/{id}", controller.DeleteMovie).Methods("PUT")
	router.HandleFunc("api/movie/deleteallmovie", controller.MarkAsWatched).Methods("PUT")

	return router
}
