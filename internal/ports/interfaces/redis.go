package interfaces

import "time"

type Redis interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
}