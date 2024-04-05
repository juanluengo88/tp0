package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()
	log.Println("hola soy un log")
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")
	//utils.EnviarMensaje(globals.ClientConfig.Ip,globals.ClientConfig.Puerto,globals.ClientConfig.Mensaje)
	valores:=utils.LeerConsola()
	utils.GenerarYEnviarPaquete(valores)
	
}
