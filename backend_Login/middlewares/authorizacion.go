package middlewares

import (
	"os"
	"github.com/golang-jwt/jwt"
	"github.com/gofiber/fiber/v2"
)


//toma el token en cadena de string
func AutenticacionToken(tokenString string) error{
	//checkear token
	//parsewith:Para analizar el token / 
	//func de validacion de token
	_, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{},error){
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil{
		return err
	}
	//Devuelve nil para indicar que el token es valido
	return nil
}

func validacion(c *fiber.Ctx) error{
	token :=
}



