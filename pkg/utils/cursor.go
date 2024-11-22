package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func EncodeCursorId(serialId int64) string {
	rand.Seed(time.Now().UnixNano())
	salt := rand.Int63n(1000000)

	data := fmt.Sprintf("%d-%d", serialId, salt)
	cursorId := base64.StdEncoding.EncodeToString([]byte(data))
	return cursorId
}

func DecodeCursorId(cursorId string) (int64, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(cursorId)
	if err != nil {
		return 0, err
	}

	// Split the decoded string
	parts := strings.Split(string(decodedBytes), "-")
	if len(parts) < 1 {
		return 0, fmt.Errorf("invalid cursorId format")
	}

	// Extract the serialId
	serialId, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, err
	}

	return serialId, nil
}

func GetNextCursor[T any](items []T, getSerialId func(item T) int64) string {
	// Check if the slice is empty
	if len(items) == 0 {
		return ""
	}

	// Get the last item and its SerialId
	lastItem := items[len(items)-1]
	serialId := getSerialId(lastItem)

	// Encode and return the cursor ID
	return EncodeCursorId(serialId)
}
