package command_runner

import (
	"context"
	"errors"
	"testing"

	"github.com/Artenso/command-runner/internal/model"
	"github.com/Artenso/command-runner/internal/service/command_runner/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestService_ListCommand(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockICommandsRepository(ctrl)
		systemCaller := mocks.NewMockISystemCaller(ctrl)
		service := New(repository, systemCaller)

		ctx := context.Background()
		limit := int64(10)
		offst := int64(0)

		repoResponse := []*model.Command{
			testCmd,
			testCmd,
		}

		expected := repoResponse

		repository.EXPECT().ListCommand(ctx, limit, offst).Return(repoResponse, nil)

		actual, err := service.ListCommand(ctx, limit, offst)

		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockICommandsRepository(ctrl)
		systemCaller := mocks.NewMockISystemCaller(ctrl)
		service := New(repository, systemCaller)

		ctx := context.Background()
		limit := int64(10)
		offst := int64(0)

		repoResponse := []*model.Command{}

		expected := repoResponse

		repository.EXPECT().ListCommand(ctx, limit, offst).Return(repoResponse, nil)

		actual, err := service.ListCommand(ctx, limit, offst)

		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("Repo error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockICommandsRepository(ctrl)
		systemCaller := mocks.NewMockISystemCaller(ctrl)
		service := New(repository, systemCaller)

		ctx := context.Background()
		limit := int64(10)
		offst := int64(0)

		repoErr := errors.New("Some error")

		repository.EXPECT().ListCommand(ctx, limit, offst).Return(nil, repoErr)

		cmdList, err := service.ListCommand(ctx, limit, offst)

		expectedErr := repoErr

		require.Nil(t, cmdList)
		require.EqualError(t, err, expectedErr.Error())
	})
}
