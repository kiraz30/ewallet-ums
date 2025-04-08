package main

import (
	"ewallet-framework/cmd"
	helpers "ewallet-framework/helpers"
)

func main() {
	//load config
	helpers.SetupConfig()

	//load log
	helpers.SetupLogger()

	//load database
	// helpers.SetupMySQL()

	//run grpc
	go cmd.ServeGRPC()

	//run http
	cmd.ServeHTTP()

}
