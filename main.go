package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {

	// models.Personalidades = []models.Personalidade{
	// 	{Id: 1, Nome: "Albert Einstein", Historia: "Físico alemão"},
	// 	{Id: 2, Nome: "Isaac Newton", Historia: "Físico e matemático inglês"},
	// }
	database.Connect()
	fmt.Println("Iniciando o servidor Rest com Go")
	fmt.Println(`
	______     ______        ______     ______   __    
 /\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
 \ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
	\ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
	 \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	routes.HandleRequest()
}
