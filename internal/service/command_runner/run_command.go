package command_runner

import (
	"context"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/Artenso/command-runner/internal/config"
	"github.com/Artenso/command-runner/internal/logger"
	"github.com/Artenso/command-runner/internal/model"
)

// runCommand runs command and updates info in repository
func (s *Service) runCommand(id int64, command string) {
	ctx := context.Background()

	wg := sync.WaitGroup{}

	cmdInfo := &model.CommandInfo{}

	output := &strings.Builder{}
	output.Grow(config.GetStrBuilderBaseCap())

	cmd := exec.CommandContext(ctx, "sh", "-c", command)

	cmd.Stdout = output
	cmd.Stderr = output

	err := s.systemCaller.StartCmd(cmd)
	if err != nil {
		cmdInfo.Status = model.StatusFailed
		cmdInfo.Output.String = err.Error()
		cmdInfo.Output.Valid = true
		s.updateCommandInfo(ctx, id, cmdInfo)

		logger.Errorf("failed to start command: %s", err)
		return
	}

	cmdInfo.Status = model.StatusInProgress
	cmdInfo.Pid.Int64 = int64(s.systemCaller.GetPid(cmd))
	cmdInfo.Pid.Valid = true
	s.updateCommandInfo(ctx, id, cmdInfo)

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.waitingCommandComplete(ctx, id, cmd, cmdInfo)
	}()

	s.readCommandOutput(ctx, cmd, id, output, cmdInfo)

	wg.Wait()
	if cmdInfo.Status == model.StatusInProgress {
		cmdInfo.Status = model.StatusDone
		cmdInfo.Output.String = output.String()
		cmdInfo.Output.Valid = true
		s.updateCommandInfo(ctx, id, cmdInfo)
	}

}

// readCommandOutput reads command output and updates info in repository
func (s *Service) readCommandOutput(ctx context.Context, cmd *exec.Cmd, id int64, output *strings.Builder, cmdInfo *model.CommandInfo) {
	for !s.systemCaller.IsProcessComplete(cmd) {
		time.Sleep(config.GetUpdatingFreq())
		cmdInfo.Output.String = output.String()
		cmdInfo.Output.Valid = true
		s.updateCommandInfo(ctx, id, cmdInfo)
	}
}

// updateCommandInfo updates command info in repository
func (s *Service) updateCommandInfo(ctx context.Context, id int64, cmdInfo *model.CommandInfo) {
	if err := s.commandsRepository.UpdateCommand(ctx, id, cmdInfo); err != nil {
		logger.Errorf("failed to update info in db: %s", err.Error())
		return
	}
}

// waitingCommandComplete waits command complete
func (s *Service) waitingCommandComplete(ctx context.Context, id int64, cmd *exec.Cmd, cmdInfo *model.CommandInfo) {
	err := s.systemCaller.WaitCmd(cmd)
	if err != nil {
		if s.systemCaller.IsProcessKilled(cmd) {
			cmdInfo.Status = model.StatusStopped
		} else {
			cmdInfo.Status = model.StatusFailed
			logger.Errorf("command running fialed: %s", err.Error())
		}
		s.updateCommandInfo(ctx, id, cmdInfo)
		return
	}
}
