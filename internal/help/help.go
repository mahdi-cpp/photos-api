package help

import (
	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/mygin"
)

// GetUserID از Gin context، user_id را به صورت string دریافت می‌کند.
func GetUserID(c *mygin.Context) (uuid.UUID, bool) {

	//// این تابع باید بعد از middleware احراز هویت استفاده شود
	//userID, exists := c.Read("X-User-ID")
	//if !exists {
	//	return "", false
	//}
	//
	//userIDStr, ok := userID.(string)
	//if !ok {
	//	return "", false
	//}
	//
	//return userIDStr, true

	//// این تابع باید بعد از middleware احراز هویت استفاده شود
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		return uuid.Nil, false
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		return [16]byte{}, false
	}

	return id, true
}

func StrPtr(str string) *string {
	return &str
}

func BoolPtr(bool bool) *bool {
	return &bool
}
