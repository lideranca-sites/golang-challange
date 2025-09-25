package handlers

import (
	"errors"
	"example/apps/api/modules/auth/locals"
	"example/apps/api/modules/products/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HTTPError struct {
	Error string `json:"error" example:"Mensagem de erro aqui"`
}

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

type CreateProductDTO struct {
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required,gt=0"`
	Quantity int     `json:"quantity" validate:"required,gte=0"`
}

// CreateProduct godoc
// @Summary      Cria um produto
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product body CreateProductDTO true "Dados do Produto"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Security     ApiKeyAuth
// @Router       /products [post]
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	body := c.Locals("body").(*CreateProductDTO)
	userId := c.Locals(locals.UserIdLocal).(int)
	product, err := h.service.CreateProduct(body.Name, body.Price, body.Quantity, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPError{Error: "Failed to create product"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Product created successfully", "product": product})
}

// GetProducts godoc
// @Summary      Lista produtos
// @Tags         Products
// @Produce      json
// @Param        user_id query int false "Filtrar por ID do usuário"
// @Success      200 {object} map[string]interface{}
// @Failure      500  {object}  HTTPError
// @Router       /products [get]
func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	userId := c.Query("user_id")
	products, err := h.service.GetProducts(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPError{Error: "Failed to fetch products"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"products": products})
}

type UpdateProductDTO struct {
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required,gt=0"`
	Quantity int     `json:"quantity" validate:"required,gte=0"`
}

// UpdateProduct godoc
// @Summary      Atualiza um produto
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do Produto"
// @Param        product body UpdateProductDTO true "Dados para atualizar"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Security     ApiKeyAuth
// @Router       /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPError{Error: "Invalid product ID"})
	}
	body := c.Locals("body").(*UpdateProductDTO)
	product, err := h.service.UpdateProduct(id, body.Name, body.Price, body.Quantity)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(HTTPError{Error: "Product not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPError{Error: "Failed to update product"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product updated successfully", "product": product})
}

// DeleteProduct godoc
// @Summary      Deleta um produto
// @Tags         Products
// @Param        id   path      int  true  "ID do Produto"
// @Success      204
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Security     ApiKeyAuth
// @Router       /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPError{Error: "Invalid product ID"})
	}
	err = h.service.DeleteProduct(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(HTTPError{Error: "Product not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPError{Error: "Failed to delete product"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
