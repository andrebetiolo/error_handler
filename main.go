package error_handler

import(
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
func CheckDB(c *gin.Context, erro error){
	if erro != nil {
		c.String(500, "")
		c.Abort()
	}
}
