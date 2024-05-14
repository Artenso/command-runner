package command_runner

import (
	"context"

	"github.com/Artenso/command-runner/internal/converter"
	desc "github.com/Artenso/command-runner/pkg/command_runner"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCommand gets command with status, pid and output from repository
func (i *Implementation) GetCommand(ctx context.Context, req *desc.GetCommandRequest) (*desc.GetCommandResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	command, err := i.commandRunnerSrv.GetCommand(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.ToGetCommandResponse(command), nil
}
