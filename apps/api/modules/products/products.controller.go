package products

import (
	"example/apps/api/domain/dto"
	"example/apps/api/domain/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const DEFAULT_PRODUCT_ROUTE = ""
const PRODUCT_ROUTE_ID = "/:id"

type ProductsController struct {
	service services.ProductsServices
}

func NewProductsController(service services.ProductsServices) ProductsController {
	return ProductsController{service: service}
}

func (p ProductsController) New(c *fiber.Ctx) error {
	var body dto.ProductDTO
	c.BodyParser(&body)
	product, err := p.service.CreateProduct(body)
	if err != nil {
		return c.Status(err.Code()).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"product": product,
	})
}

func (p ProductsController) ListAll(c *fiber.Ctx) error {
	products, err := p.service.GetAllProducts()

	id := c.Query("user_id", "")
	if len(id) > 0 {
		userId, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Couldn't get user_id",
			})
		}

		products, err = p.service.GetProductsByUserID(uint(userId))
	}

	if err != nil {
		return c.Status(err.Code()).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"products": products,
	})
}

func (p ProductsController) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id", "")
	productId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Couldn't get ID",
		})
	}

	var body dto.ProductDTO
	err = c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Couldn't parse body",
		})
	}

	body.ID = uint(productId)
	_, erro := p.service.SaveProduct(body)
	if erro != nil {
		return c.Status(erro.Code()).JSON(fiber.Map{
			"error": erro.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"product": body,
	})
}

func (p ProductsController) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id", "")
	productId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Couldn't get ID",
		})
	}

	_, erro := p.service.DeleteProduct(uint(productId))
	if erro != nil {
		return c.Status(erro.Code()).JSON(fiber.Map{
			"error": erro.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
