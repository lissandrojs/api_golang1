package main

import (
	"api_golang1/restfull"
)

func main() {
	r := restfull.InitialSetupRouter()
	r.Run(":8080")
}
