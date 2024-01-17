package types

type typeApi string
type nameFile string
type route string

const (
	Rest  typeApi = "rest"
	Graph typeApi = "graph"
)

const (
	Handlers nameFile = "handlers.go"
	Routes   nameFile = "routes.go"
	Server   nameFile = "server.go"
	Main     nameFile = "main.go"
)

const (
	RouteServer route = "server/"
)
