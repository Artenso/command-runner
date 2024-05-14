package command_runner

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/Artenso/command-runner/internal/model"
	"github.com/Artenso/command-runner/internal/service/command_runner/mocks"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

var testCmd = &model.Command{
	Id:      1,
	Command: "some command",
	Info: model.CommandInfo{
		Status: model.StatusDone,
		Pid: sql.NullInt64{
			Int64: 11111,
			Valid: true,
		},
		Output: sql.NullString{
			String: "some output",
			Valid:  true,
		},
	},
}

func TestService_GetCommand(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockICommandsRepository(ctrl)
		systemCaller := mocks.NewMockISystemCaller(ctrl)
		service := New(repository, systemCaller)

		id := testCmd.Id

		ctx := context.Background()

		repoResponse := testCmd

		expected := repoResponse

		repository.EXPECT().GetCommand(ctx, id).Return(repoResponse, nil)

		actual, err := service.GetCommand(ctx, id)

		require.NoError(t, err)
		require.Equal(t, expected, actual)
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

		repository.EXPECT().GetCommand(ctx, id).Return(nil, pgx.ErrNoRows)

		command, err := service.GetCommand(ctx, id)

		require.Nil(t, command)
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

		repository.EXPECT().GetCommand(ctx, id).Return(nil, repoErr)

		expectedErr := repoErr
		command, err := service.GetCommand(ctx, id)

		require.Nil(t, command)
		require.EqualError(t, err, expectedErr.Error())
	})
}
