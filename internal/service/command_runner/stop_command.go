package command_runner

import (
	"context"
	"errors"
	"fmt"

	"github.com/Artenso/command-runner/internal/logger"
	"github.com/Artenso/command-runner/internal/model"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StopCommand kills running command
func (s *Service) StopCommand(ctx context.Context, id int64) error {
	pid, err := s.commandsRepository.StopCommand(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return status.Errorf(codes.InvalidArgument, "invalid request: bad id: %s", model.ErrNotFound)
		}
		logger.Errorf("failed to get pid: %s", err.Error())
		return fmt.Errorf("failed to get pid: %s", err.Error())
	}

	process, err := s.systemCaller.FindProcess(pid)
	if err != nil {
		logger.Errorf("failed to find process: %s", err.Error())
		return fmt.Errorf("failed to find process: %s", err.Error())
	}

	if err = s.systemCaller.CheckProcessExist(process); err != nil {
		logger.Errorf("failed to find process: %s", err.Error())
		return fmt.Errorf("failed to find process: %s", err.Error())
	}
	if err = s.systemCaller.KillProcess(process); err != nil {
		logger.Errorf("failed to kill process: %s", err.Error())
		return fmt.Errorf("failed to kill process: %s", err.Error())
	}

	return nil
}
