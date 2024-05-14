package command_runner

import (
	"context"

	desc "github.com/Artenso/command-runner/pkg/command_runner"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// StopCommand kills running command
func (i *Implementation) StopCommand(ctx context.Context, req *desc.StopCommandRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	err := i.commandRunnerSrv.StopCommand(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
