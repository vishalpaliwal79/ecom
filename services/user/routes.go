package user

type Handler struct{

}

func NewHandler() *Handler{
	return &Handler{}
}

func(h *Handler) registerRoutes(router *mux.Router){
	router.HandleFunc("/login",h.HandleLogin).Methods("POST")
	router.HandleFunc("/register",h.HandleRegister).Methods("POST")
}
}
}

func(h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request){

}


func(h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request){

}
