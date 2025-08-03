package category

import (
	"dompetin-api/pkg/errors"
	"dompetin-api/pkg/jwt"
	"dompetin-api/pkg/response"
	"dompetin-api/pkg/sanitize"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func sanitizeRequest(request *Request) {
	request.Name = sanitize.SanitizeStrict(request.Name)
}

func (h *Handler) Create(c *gin.Context) {
	user, exists := jwt.GetUserClaims(c)
	if !exists {
		response.Respond(c, http.StatusUnauthorized, false, "", nil, nil)
		return
	}

	userID := user.ID

	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.service.Create(request.Name, userID); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusCreated, true, MsgCategorySaved, nil, nil)
}

func (h *Handler) GetByUserID(c *gin.Context) {
	user, exists := jwt.GetUserClaims(c)
	if !exists {
		response.Respond(c, http.StatusUnauthorized, false, "", nil, nil)
		return
	}

	userID := user.ID
	categories, err := h.service.GetByUserID(userID)
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, MsgCategoryRetrieved, categories, nil)
}

func (h *Handler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Error(err)
		return
	}

	var request Request
	if err = c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeRequest(&request)

	if err = errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	if err = h.service.Update(id, request.Name); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, MsgCategoryUpdated, nil, nil)
}

func (h *Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Error(err)
		return
	}

	if err = h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, MsgCategoryDeleted, nil, nil)
}
