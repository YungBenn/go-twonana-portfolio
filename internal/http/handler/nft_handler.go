package handler

import (
	"net/http"

	"github.com/YungBenn/go-twonana-portfolio/internal/nft"
	"github.com/YungBenn/go-twonana-portfolio/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type NftHandler struct {
	s nft.NftService
}

func NewNftHandler(s nft.NftService) *NftHandler {
	return &NftHandler{s}
}

//	CreateNft		godoc
//	@Summary		Create a new Nft
//	@Tags			Nfts
//	@Accept 		multipart/form-data
//	@Produce 		json
//	@Param 			ImageUrl formData nft.CreateNftRequest true "Add Nft"
//	@Success 		200 {object} nft.Nft
//	@Router 		/api/v1/nft [post]
func (h *NftHandler) CreateNft(c *fiber.Ctx) error {
	var input nft.CreateNftRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
	}

	file, _ := c.FormFile("image_url")
	input.ImageUrl = file

	ctx := c.Context()
	data, err := h.s.CreateNft(ctx, input)
	if err != nil {
		return c.Status(int(err.Code)).JSON(response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
	}

	c.Status(http.StatusCreated).JSON(response.Response(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		data,
	))

	return nil
}

//	GetAllNft		godoc
//	@Summary		List nfts
//	@Tags			Nfts
//	@Accept 		json
//	@Produce 		json
//	@Success 		200 {object} nft.Nft
//	@Router 		/api/v1/nft [get]
func (h *NftHandler) GetAllNft(c *fiber.Ctx) error {
	ctx := c.Context()
	data, err := h.s.FindAllNft(ctx)
	if err != nil {
		return c.Status(int(err.Code)).JSON(response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
	}

	c.Status(http.StatusOK).JSON(response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))

	return nil
}

//	GetOneNft		godoc
//	@Summary		Detail nft
//	@Tags			Nfts
//	@Accept 		json
//	@Produce 		json
//	@Param			id path int true "Nft ID" Format(string)
//	@Success 		200 {object} nft.Nft
//	@Router 		/api/v1/nft/{id} [get]
func (h *NftHandler) GetOneNft(c *fiber.Ctx) error {
	id := c.Params("id")

	ctx := c.Context()
	data, err := h.s.FindOneNft(ctx, id)
	if err != nil {
		return c.Status(int(err.Code)).JSON(response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
	}

	c.Status(http.StatusOK).JSON(response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))

	return nil
}

//	UpdateNft		godoc
//	@Summary		Update Nft
//	@Tags			Nfts
//	@Accept 		multipart/form-data
//	@Produce 		json
//	@Param			id path int true "Nft ID" Format(string)
//	@Param 			nft body nft.UpdateNftRequest true "Update Nft"
//	@Success 		200 {object} nft.Nft
//	@Router 		/api/v1/nft/{id} [put]
func (h *NftHandler) UpdateNft(c *fiber.Ctx) error {
	id := c.Params("id")

	var input nft.UpdateNftRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
	}

	file, _ := c.FormFile("image_url")
	input.ImageUrl = file

	ctx := c.Context()
	data, err := h.s.UpdateNft(ctx, id, input)
	if err != nil {
		return c.Status(int(err.Code)).JSON(response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
	}

	c.Status(http.StatusOK).JSON(response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))

	return nil
}

//	DeleteNft		godoc
//	@Summary		Update Nft
//	@Tags			Nfts
//	@Accept 		json
//	@Produce 		json
//	@Param			id path int true "Nft ID" Format(string)
//	@Success 		200 {object} response.ApiResponse
//	@Router 		/api/v1/nft/{id} [delete]
func (h *NftHandler) DeleteNft(c *fiber.Ctx) error {
	id := c.Params("id")

	ctx := c.Context()
	err := h.s.DeleteNft(ctx, id)
	if err != nil {
		return c.Status(int(err.Code)).JSON(response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
	}

	c.Status(http.StatusOK).JSON(response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"Deleted",
	))

	return nil
}

func (h *NftHandler) GetAllCategory(c *fiber.Ctx) error {
	ctx := c.Context()
	data, err := h.s.FindAllCategory(ctx)
	if err != nil {
		return c.Status(int(err.Code)).JSON(response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
	}

	c.Status(http.StatusOK).JSON(response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))

	return nil
}

func (h *NftHandler) GetAllNftByCategory(c *fiber.Ctx) error {
	category := c.Query("category")

	ctx := c.Context()
	data, err := h.s.FindNftByCategory(ctx, category)
	if err != nil {
		return c.Status(int(err.Code)).JSON(response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
	}

	c.Status(http.StatusOK).JSON(response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))

	return nil
}

func (h *NftHandler) GetNftByTitle(c *fiber.Ctx) error {
	title := c.Query("title")

	ctx := c.Context()
	data, err := h.s.FindNftByTitle(ctx, title)
	if err != nil {
		return c.Status(int(err.Code)).JSON(response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
	}

	c.Status(http.StatusOK).JSON(response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))

	return nil
}