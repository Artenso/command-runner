package command_runner

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Artenso/command-runner/internal/service/command_runner/mocks"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestService_StopCommand(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockICommandsRepository(ctrl)
		systemCaller := mocks.NewMockISystemCaller(ctrl)
		service := New(repository, systemCaller)

		ctx := context.Background()
		id := testCmd.Id
		var process *os.Process

		repoResponse := testCmd.Info.Pid.Int64

		repository.EXPECT().StopCommand(ctx, id).Return(repoResponse, nil)
		systemCaller.EXPECT().FindProcess(testCmd.Info.Pid.Int64).Return(process, nil)
		systemCaller.EXPECT().CheckProcessExist(process).Return(nil)
		systemCaller.EXPECT().KillProcess(process).Return(nil)

		err := service.StopCommand(ctx, id)

		require.NoError(t, err)
	})

	t.Run("Kill process err", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockICommandsRepository(ctrl)
		systemCaller := mocks.NewMockISystemCaller(ctrl)
		service := New(repository, systemCaller)

		ctx := context.Background()
		id := testCmd.Id
		killErr := errors.New("some error")

		var process *os.Process

		repoResponse := testCmd.Info.Pid.Int64

		repository.EXPECT().StopCommand(ctx, id).Return(repoResponse, nil)
		systemCaller.EXPECT().FindProcess(testCmd.Info.Pid.Int64).Return(process, nil)
		systemCaller.EXPECT().CheckProcessExist(process).Return(nil)
		systemCaller.EXPECT().KillProcess(process).Return(killErr)

		err := service.StopCommand(ctx, id)

		require.ErrorContains(t, err, "failed to kill process")
	})

	t.Run("Check process exist err", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockICommandsRepository(ctrl)
		systemCaller := mocks.NewMockISystemCaller(ctrl)
		service := New(repository, systemCaller)

		ctx := context.Background()
		id := testCmd.Id
		checkExistErr := errors.New("some error")

		var process *os.Process

		repoResponse := testCmd.Info.Pid.Int64

		repository.EXPECT().StopCommand(ctx, id).Return(repoResponse, nil)
		systemCaller.EXPECT().FindProcess(testCmd.Info.Pid.Int64).Return(process, nil)
		systemCaller.EXPECT().CheckProcessExist(process).Return(checkExistErr)

		err := service.StopCommand(ctx, id)

		require.ErrorContains(t, err, "failed to find process")
	})

	t.Run("Find process err", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockICommandsRepository(ctrl)
		systemCaller := mocks.NewMockISystemCaller(ctrl)
		service := New(repository, systemCaller)

		ctx := context.Background()
		id := testCmd.Id
		findErr := errors.New("some error")

		repoResponse := testCmd.Info.Pid.Int64

		repository.EXPECT().StopCommand(ctx, id).Return(repoResponse, nil)
		systemCaller.EXPECT().FindProcess(testCmd.Info.Pid.Int64).Return(nil, findErr)

		err := service.StopCommand(ctx, id)

		require.ErrorContains(t, err, "failed to find process")
	})

	t.Run("Not found err", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockICommandsRepository(ctrl)
		systemCaller := mocks.NewMockISystemCaller(ctrl)
		service := New(repository, systemCaller)

		ctx := context.Background()
		id := testCmd.Id

		var repoResponse int64

		repository.EXPECT().StopCommand(ctx, id).Return(repoResponse, pgx.ErrNoRows)

		err := service.StopCommand(ctx, id)

		require.ErrorContains(t, err, "bad id")
	})

	t.Run("Repo error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockICommandsRepository(ctrl)
		systemCaller := mocks.NewMockISystemCaller(ctrl)
		service := New(repository, systemCaller)

		ctx := context.Background()
		id := testCmd.Id

		repoErr := errors.New("Some error")
		var repoResponse int64

		repository.EXPECT().StopCommand(ctx, id).Return(repoResponse, repoErr)

		err := service.StopCommand(ctx, id)

		require.ErrorContains(t, err, "failed to get pid")
	})
}
