package utils

import (
	"encoding/base64"

	"github.com/google/uuid"
)

func GenerateRef() string {
	return base64.StdEncoding.EncodeToString([]byte(uuid.NewString()))[:16]
}
