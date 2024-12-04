package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// Se define la estructura Arg
type Args struct {
	Name string
}

type HelloService struct{}

// Método remoto para el servicio RPC
func (h *HelloService) SayHello(args *Args, reply *string) error {
	*reply = "Hello Word" + args.Name
	return nil
}

func main() {
	helloService := new(HelloService)

	// Registramos el servicio RPC
	err := rpc.Register(helloService)
	if err != nil {
		log.Fatal("Error al registrar el servicio RPC:", err)
	}

	// Creamos un listener TCP
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("Error al escuchar en el puerto:", err)
	}
	defer listener.Close()

	fmt.Println("Servidor RPC escuchando en puerto 3000 ...")

	// Aceptamos conexiones y las manejamos
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
