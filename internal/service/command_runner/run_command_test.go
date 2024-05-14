package command_runner

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/Artenso/command-runner/internal/model"
	"github.com/Artenso/command-runner/internal/service/command_runner/mocks"
	"go.uber.org/mock/gomock"
)

var cmdInfo1 = model.CommandInfo{
	Status: model.StatusInProgress,
	Pid:    testCmd.Info.Pid,
	Output: sql.NullString{
		String: "",
		Valid:  false,
	},
}

func TestService_runCommand(t *testing.T) {
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
		command := testCmd.Command

		cmdInfo2 := cmdInfo1
		cmdInfo2.Output.Valid = true
		cmdInfo2.Status = model.StatusDone

		systemCaller.EXPECT().StartCmd(gomock.Any()).Return(nil)
		systemCaller.EXPECT().GetPid(gomock.Any()).Return(testCmd.Info.Pid.Int64)
		repository.EXPECT().UpdateCommand(ctx, id, &cmdInfo1).Return(nil)
		systemCaller.EXPECT().WaitCmd(gomock.Any()).Return(nil)
		systemCaller.EXPECT().IsProcessComplete(gomock.Any()).Return(true)
		repository.EXPECT().UpdateCommand(ctx, id, &cmdInfo2).Return(nil)

		service.runCommand(id, command)

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
		command := testCmd.Command
		repoErr := errors.New("some error")

		cmdInfo2 := cmdInfo1
		cmdInfo2.Output.Valid = true
		cmdInfo2.Status = model.StatusDone

		systemCaller.EXPECT().StartCmd(gomock.Any()).Return(nil)
		systemCaller.EXPECT().GetPid(gomock.Any()).Return(testCmd.Info.Pid.Int64)
		repository.EXPECT().UpdateCommand(ctx, id, &cmdInfo1).Return(repoErr)
		systemCaller.EXPECT().WaitCmd(gomock.Any()).Return(nil)
		systemCaller.EXPECT().IsProcessComplete(gomock.Any()).Return(true)
		repository.EXPECT().UpdateCommand(ctx, id, &cmdInfo2).Return(repoErr)

		service.runCommand(id, command)

	})

	t.Run("Wait error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockICommandsRepository(ctrl)
		systemCaller := mocks.NewMockISystemCaller(ctrl)
		service := New(repository, systemCaller)

		ctx := context.Background()

		id := testCmd.Id
		command := testCmd.Command
		waitErr := errors.New("some wait error")

		cmdInfo2 := cmdInfo1
		cmdInfo2.Status = model.StatusFailed

		systemCaller.EXPECT().StartCmd(gomock.Any()).Return(nil)
		systemCaller.EXPECT().GetPid(gomock.Any()).Return(testCmd.Info.Pid.Int64)
		systemCaller.EXPECT().WaitCmd(gomock.Any()).Return(waitErr)
		systemCaller.EXPECT().IsProcessKilled(gomock.Any()).Return(false)
		systemCaller.EXPECT().IsProcessComplete(gomock.Any()).Return(true)
		repository.EXPECT().UpdateCommand(ctx, id, &cmdInfo1).Return(nil)
		repository.EXPECT().UpdateCommand(ctx, id, &cmdInfo2).Return(nil)

		service.runCommand(id, command)

	})

	t.Run("Kill pocess", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockICommandsRepository(ctrl)
		systemCaller := mocks.NewMockISystemCaller(ctrl)
		service := New(repository, systemCaller)

		ctx := context.Background()

		id := testCmd.Id
		command := testCmd.Command
		waitErr := errors.New("some wait error")

		cmdInfo2 := cmdInfo1
		cmdInfo2.Status = model.StatusStopped

		systemCaller.EXPECT().StartCmd(gomock.Any()).Return(nil)
		systemCaller.EXPECT().GetPid(gomock.Any()).Return(testCmd.Info.Pid.Int64)
		systemCaller.EXPECT().WaitCmd(gomock.Any()).Return(waitErr)
		systemCaller.EXPECT().IsProcessKilled(gomock.Any()).Return(true)
		systemCaller.EXPECT().IsProcessComplete(gomock.Any()).Return(true)
		repository.EXPECT().UpdateCommand(ctx, id, &cmdInfo1).Return(nil)
		repository.EXPECT().UpdateCommand(ctx, id, &cmdInfo2).Return(nil)

		service.runCommand(id, command)

	})

	t.Run("Start error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockICommandsRepository(ctrl)
		systemCaller := mocks.NewMockISystemCaller(ctrl)
		service := New(repository, systemCaller)

		ctx := context.Background()

		id := testCmd.Id
		command := testCmd.Command
		startErr := errors.New("some start error")

		cmdInfo := model.CommandInfo{
			Status: model.StatusFailed,
			Output: sql.NullString{
				String: startErr.Error(),
				Valid:  true,
			},
		}
		systemCaller.EXPECT().StartCmd(gomock.Any()).Return(startErr)
		repository.EXPECT().UpdateCommand(ctx, id, &cmdInfo).Return(nil)

		service.runCommand(id, command)

	})
}
