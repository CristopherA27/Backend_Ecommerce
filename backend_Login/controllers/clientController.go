package controllers

import (
	db "github.com/CristopherA27/backend_login/tree/master/Config"
	"github.com/CristopherA27/backend_login/tree/master/models"
	"github.com/gofiber/fiber/v2"
)

// Status 200: Exito
func getClientes(c *fiber.Ctx) error {
	var clientes []models.Cliente

	db.DB.Select("id, rut, nombre, apellido, correo, direccion, celular").Find(&clientes)
	//Devolver respuesta HTTP al cliente
	return c.Status(200).JSON(fiber.Map{
		"success": true, //true para indicar que la solicitud fue un exito
		"message": "Success",
		"data":    clientes, //datos enviados
	})
}

func getCliente(c *fiber.Ctx) error {
	clienteID := c.Params("clienteID") //el clienteID es el texto despues de mi /
	var cliente models.Cliente
	db.DB.Select("id,rut,nombre,apellido,correo,direccion,celular").Where("id= ?", clienteID).First(&cliente)

	if cliente.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Cliente no encontrado"})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    cliente,
	})

}
