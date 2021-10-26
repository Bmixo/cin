package main

import (
	"github.com/Bmixo/cin"
	"github.com/Bmixo/cin/middleware/auth"
)

func main() {
	cin.Server(&auth.Auth{})
}
