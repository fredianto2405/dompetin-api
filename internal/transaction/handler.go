package transaction

import (
	"dompetin-api/pkg/errors"
	"dompetin-api/pkg/jwt"
	"dompetin-api/pkg/response"
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

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.service.Create(request, userID); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, MsgTransactionSaved, nil, nil)
}

func (h *Handler) History(c *gin.Context) {
	startDate := c.DefaultQuery("startDate", "")
	endDate := c.DefaultQuery("endDate", "")

	user, exists := jwt.GetUserClaims(c)
	if !exists {
		response.Respond(c, http.StatusUnauthorized, false, "", nil, nil)
		return
	}

	userID := user.ID

	details, err := h.service.GetByTransactionDateAndUserID(startDate, endDate, userID)
	if err != nil {
		c.Error(err)
		return
	}

	summary, err := h.service.SummaryByTransactionDateAndUserID(startDate, endDate, userID)
	if err != nil {
		c.Error(err)
		return
	}

	history := &HistoryResponse{
		Summary: summary,
		Details: details,
	}

	response.Respond(c, http.StatusOK, true, MsgTransactionRetrieved, history, nil)
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

	if err = errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	if err = h.service.Update(request, id); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, MsgTransactionUpdated, nil, nil)
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

	response.Respond(c, http.StatusOK, true, MsgTransactionDeleted, nil, nil)
}
