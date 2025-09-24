package help

import (
	"net/http"

	"github.com/mahdi-cpp/iris-tools/mygin"
)

const (
	ErrorUserID         = "User ID is not a string"
	ErrorInvalidRequest = "Invalid request"
)

// AbortWithError یک پاسخ JSON خطا را ارسال و درخواست را Abort می‌کند.
func AbortWithError(c *mygin.Context, status int, message string) {
	c.JSON(status, mygin.H{"error": message})
	c.Abort()
}

// AbortWithUserIDInvalid یک پاسخ خطا برای زمانی که user_id نامعتبر است، ارسال می‌کند.
func AbortWithUserIDInvalid(c *mygin.Context) {
	c.JSON(http.StatusInternalServerError, mygin.H{"error": ErrorUserID})
	c.Abort() // This stops the next handler from running
}

func AbortWithRequestInvalid(c *mygin.Context) {
	c.JSON(http.StatusInternalServerError, mygin.H{"error": ErrorInvalidRequest})
	c.Abort()
}
