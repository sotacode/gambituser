package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sotacode/gambituser/models"
	"github.com/sotacode/gambituser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("comienza registro")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close() //Al terminar la ejecucion de la funcion que cierre la conexion

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" +
		sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Registro completado exitosamente.")
	return nil
}
