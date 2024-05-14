package command_runner

import (
	"context"
	"os"
	"os/exec"

	"github.com/Artenso/command-runner/internal/model"
)

//go:generate mockgen -source=service.go -destination=./mocks/service_mock.go -package=mocks

// ICommandsRepository working with repository
type ICommandsRepository interface {
	AddCommand(ctx context.Context, command string) (int64, error)
	GetCommand(ctx context.Context, id int64) (*model.Command, error)
	ListCommand(ctx context.Context, limit, offset int64) ([]*model.Command, error)
	UpdateCommand(ctx context.Context, id int64, cmdInfo *model.CommandInfo) error
	StopCommand(ctx context.Context, id int64) (int64, error)
}

// ISystemCaller working with system calls
type ISystemCaller interface {
	StartCmd(cmd *exec.Cmd) error
	WaitCmd(cmd *exec.Cmd) error
	GetPid(cmd *exec.Cmd) int64
	IsProcessComplete(cmd *exec.Cmd) bool
	IsProcessKilled(cmd *exec.Cmd) bool
	FindProcess(pid int64) (*os.Process, error)
	CheckProcessExist(process *os.Process) error
	KillProcess(process *os.Process) error
}

// Service
type Service struct {
	commandsRepository ICommandsRepository
	systemCaller       ISystemCaller
}

// New creates new service
func New(commandsRepository ICommandsRepository, systemCaller ISystemCaller) *Service {
	return &Service{
		commandsRepository: commandsRepository,
		systemCaller:       systemCaller,
	}
}
