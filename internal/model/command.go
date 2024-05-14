package model

import "database/sql"

// Command - model of command with it's parameters
type Command struct {
	Id      int64       `db:"id"`
	Command string      `db:"command"`
	Info    CommandInfo `db:""`
}

// CommandInfo - model of command with it's parameters
type CommandInfo struct {
	Status string         `db:"status"`
	Pid    sql.NullInt64  `db:"pid"`
	Output sql.NullString `db:"output"`
}
