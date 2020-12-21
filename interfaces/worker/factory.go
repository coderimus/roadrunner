package worker

import (
	"context"
	"os/exec"
)

// Factory is responsible of wrapping given command into tasks WorkerProcess.
type Factory interface {
	// SpawnWorkerWithContext creates new WorkerProcess process based on given command with context.
	// Process must not be started.
	SpawnWorkerWithTimeout(context.Context, *exec.Cmd) (BaseProcess, error)
	// SpawnWorker creates new WorkerProcess process based on given command.
	// Process must not be started.
	SpawnWorker(*exec.Cmd) (BaseProcess, error)
	// Close the factory and underlying connections.
	Close() error
}