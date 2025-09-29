package help

import (
	"net/http"

	"github.com/mahdi-cpp/photos-api/mygin"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func SendError(c *mygin.Context, message string, code int) {
	c.JSON(http.StatusBadRequest, mygin.H{"message": message, "code": code})
}
