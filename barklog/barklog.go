package barklog

import (
	"encoding/json"
	"time"
)

// Struct representing a log in Bark
type BarkLog struct {
	Id          int64           `db:"id"`
	LogTime     time.Time       `db:"log_time"`
	LogLevel    int             `db:"log_level"`
	ServiceName string          `db:"service_name"`
	Code        string          `db:"code"`
	Message     string          `db:"msg"`
	MoreData    json.RawMessage `db:"more_data"`
}
