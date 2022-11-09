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

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {

	_, exit := s.router.rules[path]
	if !exit {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
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
