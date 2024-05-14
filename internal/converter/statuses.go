package converter

import (
	"github.com/Artenso/command-runner/internal/model"
	desc "github.com/Artenso/command-runner/pkg/command_runner"
)

// ToDescStatus converts internal status into proto status
func ToDescStatus(status string) desc.Status {
	if value, found := desc.Status_value[status]; found {
		return desc.Status(value)
	}

	return desc.Status_UNSPECIFIED
}

// ToInternalStatus converts proto status into internal status
func ToInternalStatus(descStatus desc.Status) string {
	if value, found := desc.Status_name[int32(descStatus)]; found {
		return value
	}

	return model.StatusUnspecified
}
