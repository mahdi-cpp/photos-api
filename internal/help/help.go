package help

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/goccy/go-json"
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

func MakeRequestBody(method, endpoint string, body interface{}) (*http.Response, error) {

	// Build URL with query parameters
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}

	// Marshal body if provided
	var bodyReader io.Reader

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshaling body: %w", err)
	}
	bodyReader = bytes.NewReader(jsonData)

	fmt.Println(u.String())
	fmt.Println("")

	// create request
	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("user_id", "01997cba-6dab-7636-a1f8-2c03174c7b6e")

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	return resp, nil
}
