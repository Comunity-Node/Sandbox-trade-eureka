package main

import (
    "sandbox-invest/routes"
)

func main() {
    r := routes.SetupRouter()
    r.Run(":8080")
}
