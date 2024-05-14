package command_runner

import (
	"context"

	"github.com/Artenso/command-runner/internal/converter"
	desc "github.com/Artenso/command-runner/pkg/command_runner"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListCommand gets commands with statuses and pids from repository
func (i *Implementation) ListCommand(ctx context.Context, req *desc.ListCommandRequest) (*desc.ListCommandResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	command, err := i.commandRunnerSrv.ListCommand(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		return nil, err
	}

	return converter.ToListCommandResponse(command), nil
}
