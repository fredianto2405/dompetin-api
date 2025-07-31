package auth

import (
	"dompetin-api/internal/user"
	"dompetin-api/pkg/errors"
	"dompetin-api/pkg/jwt"
	"dompetin-api/pkg/response"
	"dompetin-api/pkg/sanitize"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service     *Service
	userService *user.Service
}

func NewHandler(service *Service, userService *user.Service) *Handler {
	return &Handler{
		service:     service,
		userService: userService,
	}
}

func sanitizeRegisterRequest(request *RegisterRequest) {
	request.Email = sanitize.SanitizeStrict(request.Email)
}

func (h *Handler) Register(c *gin.Context) {
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeRegisterRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	emailExists, err := h.userService.IsEmailExist(request.Email)
	if err != nil {
		c.Error(err)
		return
	}

	if emailExists {
		response.Respond(c, http.StatusBadRequest, false, MsgEmailExists, nil, nil)
		return
	}

	if err = h.userService.Create(request.Email, request.Password); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, MsgRegisterSuccess, nil, nil)
}

func sanitizeLoginRequest(request *LoginRequest) {
	request.Email = sanitize.SanitizeStrict(request.Email)
}

func (h *Handler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	sanitizeLoginRequest(&request)

	if err := errors.Validate.Struct(request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	user, err := h.service.Login(&request)
	if err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	token, err := jwt.GenerateJWT(user.ID, user.Email)
	if err != nil {
		response.Respond(c, http.StatusInternalServerError, false, MsgFailedGenerateJWT, nil, nil)
		return
	}

	data := &LoginResponse{AccessToken: token}
	response.Respond(c, http.StatusOK, true, MsgLoginSuccess, data, nil)
}
