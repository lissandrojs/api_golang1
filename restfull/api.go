package restfull

import (
	"api_golang1/database"
	"api_golang1/models"

	"github.com/gin-gonic/gin"
)

func InitialSetupRouter() *gin.Engine {
	router := gin.Default()

	// db, err := database.InitDB()
	// if err != nil {
	// 	panic("Error ao conectar ao banco de dados")
	// }

	router.GET("/products", GetProduts)

	return router
}

func GetProduts(context *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	context.JSON(200, products)
}

func CreateProduct(context *gin.Context) {
	var product models.Product
	context.BindJSON(&product)
	database.DB.Create(&product)
	context.JSON(201, product)
}

func GetProdutById(context *gin.Context) {
	id := context.Params.ByName("id")
	var product models.Product
	database.DB.First(&product, id)
	context.JSON(200, product)
}

func UpdateProduct(context *gin.Context) {
	id := context.Params.ByName("id")
	var product models.Product
	database.DB.First(&product, id)
	context.BindJSON(&product)
	database.DB.Save(&product)
	context.JSON(200, product)
}

func DeleteProduct(context *gin.Context) {
	id := context.Params.ByName("id")
	var product models.Product
	database.DB.Delete(&product, id)
	context.JSON(200, gin.H{"message": "Produto deletado"})
}
