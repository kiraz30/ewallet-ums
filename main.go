package main

import (
	"ewallet-ums/cmd"
	helpers "ewallet-ums/helpers"
	"fmt"
)

func main() {
	//load config
	helpers.SetupConfig()

	// Debug: Pastikan env sudah terisi dengan benar
	fmt.Println("Env map in main:", helpers.Env)

	// Pastikan jwtSecret terisi dengan benar
	jwtSecret := []byte(helpers.GetEnv("APP_SECRET", ""))
	fmt.Println("JWT Secret main:", jwtSecret)

	//load log
	helpers.SetupLogger()

	//load database
	helpers.SetupMySQL()

	//run grpc
	go cmd.ServeGRPC()

	//run http
	cmd.ServeHTTP()

}
