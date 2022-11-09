package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", handlerRoot)
	server.Handle("/api", HandleHome)

	server.Listen()
}
