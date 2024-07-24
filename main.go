/*
	Este exercício foi o mais desafiador, pois não sabia como retornar a latitude e longitude de uma unica vez.
	Gostei pois aprendi como utilizar alguns parâmetros do objeto do contexto como "Request.URL.Query().Get()"
	Evitei ao máximo me prender à tutoriais e GPT que não fossem para descobrir como coletar os dados do IP com o goip e outras opções como
	IPinfo.io e IP2Location
*/

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	goip "github.com/jpiontek/go-ip-api"
)

/*
Crio a função para fazer a chamada getIP utilizando o gin e goip para coletar a longitude e latitude.
Na função defino a variável givenIP para armazenar o IP informado na URL
Crio a instancia do goip e atribuo à variável client
Crio a variável result e chamo a função "GetLocationForIp()" passando a variável givenIP conforme exemplo da documentação https://pkg.go.dev/github.com/FairyTale5571/go-ip-api#section-readme
Testo para erro
Crio a variável location (aqui foi onde mais apanhei, pois não sabia como retornar os valores, nem me passou pela cabeça usar o map, e precisei recorrer ao Google)
Devolvo o Json para o status ok com o o location.
*/
func getIP(c *gin.Context) {

	givenIP := c.Request.URL.Query().Get("ip")

	client := goip.NewClient()
	result, err := client.GetLocationForIp(givenIP)
	if err != nil {
		fmt.Println("Parâmetro inválido: ", err)
	}

	location := map[string]float32{
		"latitude":  result.Lat,
		"longitude": result.Lon,
	}

	c.IndentedJSON(http.StatusOK, location)
}

func main() {
	router := gin.Default()
	router.GET("/", getIP)

	router.Run("localhost:8080")
}
