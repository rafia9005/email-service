package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rafia9005/email-service/pkg/utils"
)

type EmailRequest struct {
	To      string `json:"to" validate:"required,email"`
	Subject string `json:"subject" validate:"required"`
	Body    string `json:"body" validate:"required"`
	IsHTML  bool   `json:"is_html"`
}

func SendEmail(c echo.Context) error {
	req := new(EmailRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err := utils.EmailSender(req.To, req.Subject, req.Body, req.IsHTML)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Email sent successfully"})
}
