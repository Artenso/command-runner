package command_runner

import (
	"context"

	"github.com/Artenso/command-runner/internal/logger"
)

// AddCommand adds command to repository and runs it
func (s *Service) AddCommand(ctx context.Context, command string) (int64, error) {
	id, err := s.commandsRepository.AddCommand(ctx, command)
	if err != nil {
		logger.Errorf("failed to add command: %s", err.Error())
		return int64(0), err
	}

	go s.runCommand(id, command)

	return id, nil
}
