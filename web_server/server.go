package main // Conoce todo lo de este paquete
import "net/http"

// Puerto del servidor
type Server struct {
	port   string
	router *Router
}

// Instanciar servidor y escuchar conexciones
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}

func (s *Server) Listen() error {
	// Punto de entrada de la aplicacion
	http.Handle("/", s.router)

	// Seguno parametro es un handler
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}

	return nil
}
