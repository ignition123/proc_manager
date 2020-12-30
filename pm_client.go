package main

import(
	"lib"
	"server"
)

func main(){
	
	defer lib.Handlepanic()

	// loading json config file to inmemory
	server.LoadConfig()

	// connecting server socket to publish status of the process
	server.ConnectTCP()

	// running process
	server.ExecProcesses()
}