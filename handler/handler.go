package handler

import (
	"github.com/a-h/templ"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/playlist-pilot/components"
)

type Handler struct{}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{}
}

func (h *Handler) Index(c echo.Context) error {
	return c.JSON(200, "Hello fromr root!")
}

func (h *Handler) Component(c echo.Context) error {
	return render(c, components.Base())
}

func (h *Handler) SubComponent(c echo.Context) error {
	return render(c, components.SubBase())
}

// Shorthand to render a templ component using Echo context
func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
