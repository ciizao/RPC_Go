package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// Establecemos conexión con el servidor RPC en el puerto 3000
	client, err := rpc.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("Error de conexión:", err)
		return
	}
	defer client.Close()

	// Creamos los argumentos que enviaremos al servidor
	args := struct {
		Name string
	}{
		Name: " + RPC and Go",
	}
	var reply string

	// Realizamos la llamada al método remoto "HelloService.SayHello"
	err = client.Call("HelloService.SayHello", args, &reply)
	if err != nil {
		fmt.Println("Error al llamar al método remoto:", err)
		return
	}

	// Mostramos la respuesta recibida del servidor
	fmt.Println(reply)
}
