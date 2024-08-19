package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/KozlovNikolai/test-task/internal/model"
	"github.com/KozlovNikolai/test-task/internal/store"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UserHandler is ...
type ProductHandler struct {
	logger *zap.Logger
	repoWR store.IRepository
	repoRO store.IRepository
}

// NewProductHandler is ...
func NewProductHandler(logger *zap.Logger, repoWR store.IRepository, repoRO store.IRepository) *ProductHandler {
	return &ProductHandler{
		logger: logger,
		repoWR: repoWR,
		repoRO: repoRO,
	}
}

// CreateProduct is ...
// CreateProductTags		godoc
// @Summary				Register user
// @Description			Save register data of user in Repo.
// @Param				product body model.AddProduct true "Create product"
// @Produce				application/json
// @Tags				product
// @Security     	BearerAuth
// @Success				200 {object} model.Product
// @failure				400 {string} err.Error()
// @failure				500 {string} err.Error()
// @Router				/product [post]
func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	var addProduct model.AddProduct
	// Заполняем структуру addProduct данными из запроса
	if err := c.ShouldBindJSON(&addProduct); err != nil {
		ph.logger.Error("Error binding JSON-addProduct", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var product model.Product
	// Заполняем структуру Product данными из addProduct
	product.Name = addProduct.Name
	product.Price = addProduct.Price
	product.ProviderID = addProduct.ProviderID
	product.Stock = addProduct.Stock

	// Валидация данных товара
	if err := product.Validate(); err != nil {
		ph.logger.Error("Error creating product", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Сохраняем товар в БД
	id, err := ph.repoWR.CreateProduct(context.TODO(), product)
	if err != nil {
		ph.logger.Error("Error creating product", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// загружаем из базы информацию о сохраненном пользователе:
	product, err = ph.repoRO.GetProductByID(context.TODO(), id)
	if err != nil {
		ph.logger.Error("Error load product from db after saving", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// GetProduct is ...
// GetProductTags 		godoc
// @Summary			Get product by product id.
// @Description		Return product with "id" number.
// @Param			product_id path int true "Product ID"
// @Tags			product
// @Success			200 {object} model.Product
// @failure			404 {string} err.Error()
// @Router			/product/{product_id} [get]
func (ph *ProductHandler) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := ph.repoRO.GetProductByID(context.TODO(), id)
	if err != nil {
		ph.logger.Error("Error getting product", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// GetProducts is ...
// GetProductsTags 		godoc
// @Summary			Get all products.
// @Description		Return products list.
// @Tags			product
// @Produce      json
// @Success			200 {object} []model.Product
// @failure			404 {string} err.Error()
// @Router			/products [get]
func (ph *ProductHandler) GetProducts(c *gin.Context) {

	users, err := ph.repoRO.GetAllProducts(context.TODO())
	if err != nil {
		ph.logger.Error("Error getting users", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	}
	c.JSON(http.StatusOK, users)
}
