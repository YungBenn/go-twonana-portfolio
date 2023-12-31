package handler

import (
	"net/http"
	"time"

	"github.com/YungBenn/go-twonana-portfolio/internal/admin"
	"github.com/YungBenn/go-twonana-portfolio/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type AuthHandler struct {
	s     admin.AdminService
	store *session.Store
}

func NewAuthHandler(s admin.AdminService, store *session.Store) *AuthHandler {
	return &AuthHandler{s, store}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var input admin.AdminRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewResponse(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
	}

	ctx := c.Context()
	data, err := h.s.CreateAdmin(ctx, input)
	if err != nil {
		return c.Status(int(err.Code)).JSON(response.NewResponse(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
	}

	c.Status(http.StatusCreated).JSON(response.NewResponse(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		data,
	))

	return nil
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input admin.AdminRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewResponse(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
	}

	ctx := c.Context()
	data, err := h.s.Login(ctx, input)
	if err != nil {
		return c.Status(int(err.Code)).JSON(response.NewResponse(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
	}

	sess, errSess := h.store.Get(c)
	if errSess != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			errSess,
		))
	}

	sess.SetExpiry(1 * time.Hour)
	sess.Set("username", data.Username)

	if err := sess.Save(); err != nil {
		panic(err)
	}

	userData := admin.LoginResponse{
		ID:       data.ID.Hex(),
		Username: data.Username,
	}

	c.Status(http.StatusOK).JSON(response.NewResponse(
		http.StatusOK,
		"Login Success",
		userData,
	))

	return nil
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	sess, err := h.store.Get(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err,
		))
	}

	sess.Delete("username")

	if err := sess.Destroy(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewResponse(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err,
		))
	}

	c.Status(http.StatusOK).JSON(response.NewResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"Logout Success",
	))

	return nil
}
