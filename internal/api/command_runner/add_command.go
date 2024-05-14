package command_runner

import (
	"context"

	desc "github.com/Artenso/command-runner/pkg/command_runner"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddCommand adds command to repository and runs it
func (i *Implementation) AddCommand(ctx context.Context, req *desc.AddCommandRequest) (*desc.AddCommandResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	id, err := i.commandRunnerSrv.AddCommand(ctx, req.GetCommand())
	if err != nil {
		return nil, err
	}

	return &desc.AddCommandResponse{Id: id}, nil
}
