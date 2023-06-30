package main

import (
	"fmt"

	"github.com/nehalshaquib/GoShellCommander/api"
)

func main() {
	// out, err := util.RunCommand(context.Background(), "docker", []string{"ps", "-a"})
	// if err != nil {
	// 	if strings.Contains(err.Error(), "not found") {
	// 		fmt.Println("not found: ", err)
	// 	}
	// 	fmt.Println("my error: ", err)
	// }
	// fmt.Println("my res: ", out)
	
	fmt.Println("starting shellCommander server...")
	server := api.NewServer()
	err := server.Run()
	if err != nil {
		fmt.Println("error in starting server: ", err)
	}

}
