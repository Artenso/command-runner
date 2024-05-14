package converter

import (
	"github.com/Artenso/command-runner/internal/model"
	desc "github.com/Artenso/command-runner/pkg/command_runner"
)

// ToGetCommandResponse converts internal command model into proto command message
func ToGetCommandResponse(command *model.Command) *desc.GetCommandResponse {
	response := &desc.GetCommandResponse{
		Command: command.Command,
		Status:  ToDescStatus(command.Info.Status),
	}
	if command.Info.Output.Valid {
		response.Output = command.Info.Output.String
	}
	if command.Info.Pid.Valid {
		response.Pid = command.Info.Pid.Int64
	}
	return response
}

// ToListCommandResponse converts internal commands list into proto commands list
func ToListCommandResponse(commands []*model.Command) *desc.ListCommandResponse {
	descCommands := make([]*desc.CommandInList, 0, len(commands))
	for _, command := range commands {
		descCommand := &desc.CommandInList{
			Command: command.Command,
			Status:  ToDescStatus(command.Info.Status),
		}
		if command.Info.Pid.Valid {
			descCommand.Pid = command.Info.Pid.Int64
		}
		descCommands = append(descCommands, descCommand)
	}
	return &desc.ListCommandResponse{
		Commands: descCommands,
	}
}
