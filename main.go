package error_handler

import(
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

// Verifica se tem algum erro na operação
func Check(erro error, params ...string) {
	if erro != nil {
		log.Fatalln(erro, params)
	}
}

// Verifica se tem algum erro em uma operação feita em um banco de dados
func CheckDB(c *gin.Context, erro error) bool{
	valid := true
	if erro != nil {
		fmt.Println(erro)
		 c.AbortWithStatus(500)
		 valid = false
	}
	return valid
}
