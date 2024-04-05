package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()
	// loggear "Hola soy un log" usando la biblioteca log
	log.Println("Hola soy un log!")

	globals.ClientConfig = utils.IniciarConfiguracion("config.json")
	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuracion")
	}

	// loggeamos el valor de la config
	log.Printf("IP: %s\n", globals.ClientConfig.Ip)
	log.Printf("Message: %s\n", globals.ClientConfig.Mensaje)
	log.Printf("Port %d\n", globals.ClientConfig.Puerto)

	// enviar un mensaje al servidor con el valor de la config
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	// leer de la consola el mensaje
	valores := utils.LeerConsola()

	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, valores)
}
