package command_runner

import (
	"context"

	"github.com/Artenso/command-runner/internal/logger"
	"github.com/Artenso/command-runner/internal/model"
)

// ListCommand gets commands with statuses and pids from repository
func (s *Service) ListCommand(ctx context.Context, limit, offset int64) ([]*model.Command, error) {
	commands, err := s.commandsRepository.ListCommand(ctx, limit, offset)
	if err != nil {
		logger.Errorf("failed to get commands list: %s", err.Error())
		return nil, err
	}

	return commands, nil
}
