package api

import (
	"context"
	"log"

	"github.com/ansufw/gomono/internal/dto"
	"github.com/ansufw/gomono/internal/infrastructure/database/pg"
	"github.com/ansufw/gomono/internal/infrastructure/database/sqlc"
	"github.com/ansufw/gomono/internal/pkg"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	query *sqlc.Queries
	ctx   context.Context
}

func NewHandler(db *pg.PG) *Handler {
	return &Handler{
		query: sqlc.New(db.DB),
	}
}

// @Summary      List users
// @Description  Get a list of users with pagination
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        X-MyHeader      header    string    true    "MyHeader must be set for valid response"
// @Param        X-API-VERSION   header    string    true    "API version eg.: 1.0"
// @Param        limit           query     int       false   "Limit the number of users returned"
// @Param        offset          query     int       false   "Offset for pagination"
// @Success      200             {object}  pkg.Payload{message=dto.UserListResponse}
// @Failure      500             {object}  pkg.Payload
// @Router       /api/users [get]
func (h *Handler) ListUsers(c *fiber.Ctx) error {
	// Get pagination parameters with defaults
	limit := 10
	offset := 0

	// Parse query parameters if provided
	if l := c.QueryInt("limit"); l > 0 {
		limit = l
	}
	if o := c.QueryInt("offset"); o >= 0 {
		offset = o
	}

	dbUsers, err := h.query.ListUsers(c.Context(), sqlc.ListUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		log.Printf("error listing users: %v", err)
		payload := pkg.Payload{
			Error:   true,
			Message: "Failed to retrieve users",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(payload)
	}

	log.Printf("retrieved %d users", len(dbUsers))
	users := dto.MapUsersToResponse(dbUsers)
	payload := pkg.Payload{
		Error:   false,
		Message: users,
	}

	return c.JSON(payload)
}
