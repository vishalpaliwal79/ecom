package api
import(
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
)


type APIserver struct{
	addr string
	db *sql.DB
}

func newAPIServer(addr sting,db *sql.DB)*APIserver{
return &APIserver{
	addr:addr,
	db:db,
}
}

func (s *APIserver) run() error{
router :=mux.NewRouter()
subrouter:=router.PathPrefix("api/v1").Subrouter()
userHandler:=user.NewHandler()
userHandler.RegisterRoutes(subrouter)
log.println("Listening on",s.addr) 
return http.ListenAndServe(s.addr, router)
}
