package runtime

import (
	"os"
	"sync/atomic"
	"time"
)

// Snapshot captures process-local state for health checks.
type Snapshot struct {
	Status      string    `json:"status"`
	Service     string    `json:"service"`
	Version     string    `json:"version"`
	PID         int       `json:"pid"`
	StartedAt   time.Time `json:"started_at"`
	Uptime      string    `json:"uptime"`
	CallCount   uint64    `json:"healthcheck_call_count"`
	CheckedAt   time.Time `json:"checked_at"`
	GoVersion   string    `json:"go_version"`
}

var (
	startedAt   = time.Now()
	callCount   uint64
	serviceName = "edc-mcp"
	version     = "0.1.0"
)

// Health returns a fresh snapshot and bumps the healthcheck counter.
func Health() Snapshot {
	n := atomic.AddUint64(&callCount, 1)
	now := time.Now()

	return Snapshot{
		Status:      "ok",
		Service:     serviceName,
		Version:     version,
		PID:         os.Getpid(),
		StartedAt:   startedAt.UTC(),
		Uptime:      now.Sub(startedAt).Round(time.Millisecond).String(),
		CallCount:   n,
		CheckedAt:   now.UTC(),
		GoVersion:   getGoVersion(),
	}
}