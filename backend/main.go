package main

import "ecommerce/cmd"

func main() {
	// check .env
	// cnf := config.GetConfig()
	// fmt.Println(cnf.Version)
	// fmt.Println(cnf.ServiceName)
	// fmt.Println(cnf.HttpPort)
	cmd.Serve()
}
