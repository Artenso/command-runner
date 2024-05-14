package command_runner

import (
	"context"
	"errors"

	"github.com/Artenso/command-runner/internal/logger"
	"github.com/Artenso/command-runner/internal/model"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCommand gets command with status, pid and output from repository
func (s *Service) GetCommand(ctx context.Context, id int64) (*model.Command, error) {
	command, err := s.commandsRepository.GetCommand(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid request: bad id: %s", model.ErrNotFound)
		}
		logger.Errorf("failed to get command: %s", err.Error())
		return nil, err
	}

	return command, nil
}
