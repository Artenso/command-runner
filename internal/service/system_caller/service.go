package system_caller

import (
	"os"
	"os/exec"
	"syscall"
)

type Service struct{}

// New creates new system call service
func New() *Service {
	return &Service{}
}

// StartCmd is wrapped exec.Cmd.Start()
func (s *Service) StartCmd(cmd *exec.Cmd) error {
	return cmd.Start()
}

// WaitCmd is wrapped exec.Cmd.Wait()
func (s *Service) WaitCmd(cmd *exec.Cmd) error {
	return cmd.Wait()
}

// GetPid gets command pid from exec.Cmd.Process.Pid
func (s *Service) GetPid(cmd *exec.Cmd) int64 {
	return int64(cmd.Process.Pid)
}

// IsProcessComplete checks process state
// returns true when exec.Cmd.ProcessState != nil
func (s *Service) IsProcessComplete(cmd *exec.Cmd) bool {
	return cmd.ProcessState != nil
}

// IsProcessKilled checks process state from exec.Cmd.ProcessState
// returns true when process killed by signal
func (s *Service) IsProcessKilled(cmd *exec.Cmd) bool {
	return cmd.ProcessState.String() == "signal: killed"
}

// FindProcess is wrapped os.FindProcess()
func (s *Service) FindProcess(pid int64) (*os.Process, error) {
	return os.FindProcess(int(pid))
}

// CheckProcessExist is wrapped os.Process.Signal(syscall.Signal(0))
func (s *Service) CheckProcessExist(process *os.Process) error {
	return process.Signal(syscall.Signal(0))
}

// KillProcess is wrapped os.Process.Kill()
func (s *Service) KillProcess(process *os.Process) error {
	return process.Kill()
}
