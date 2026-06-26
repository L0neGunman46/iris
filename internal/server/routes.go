package server

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_url", r.CreateUrl)
	api.Get("/get_urls", r.GetUrls)
	api.Post("/update_short", r.UpdateShort)
	api.Get("/resolve_url", r.ReolveUrl)
	api.Delete("/delete_url", r.DeleteUrl)
}

func (r *Repository) CreateUrl(ctx *fiber.Ctx) error {
	return nil
}

func (r *Repository) GetUrls(ctx *fiber.Ctx) error {
	return nil
}

func (r *Repository) UpdateShort(ctx *fiber.Ctx) error {
	return nil
}

func (r *Repository) ReolveUrl(ctx *fiber.Ctx) error {
	return nil
}

func (r *Repository) DeleteUrl(ctx *fiber.Ctx) error {
	return nil
}
