package postgres

import (
	"context"
	"fmt"

	"github.com/Artenso/command-runner/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

// db table and columns
const (
	table = "commands"

	idCol      = "id"
	commandCol = "command"
	statusCol  = "status"
	pidCol     = "pid"
	outputCol  = "output"
)

// Repository with postgress connection
type Repository struct {
	dbConn *pgxpool.Pool
}

// New creates new repository object
func New(dbConn *pgxpool.Pool) *Repository {
	return &Repository{
		dbConn: dbConn,
	}
}

// AddCommand adds command to repository
func (r *Repository) AddCommand(ctx context.Context, command string) (int64, error) {
	builder := sq.Insert(table).
		Columns(commandCol, statusCol).
		Values(command, model.StatusNew).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return int64(0), fmt.Errorf("failed to build SQL query: %s", err.Error())
	}

	var id int64

	if err = r.dbConn.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return int64(0), err
	}

	return id, nil
}

// GetCommand gets command with status, pid and output from repository
func (r *Repository) GetCommand(ctx context.Context, id int64) (*model.Command, error) {
	builder := sq.Select("*").
		From(table).
		Where(sq.Eq{idCol: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %s", err.Error())
	}

	rows, err := r.dbConn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	cmd := &model.Command{}

	if err = pgxscan.ScanOne(cmd, rows); err != nil {
		return nil, err
	}

	return cmd, nil
}

// ListCommand gets commands with statuses and pids from repository
func (r *Repository) ListCommand(ctx context.Context, limit, offset int64) ([]*model.Command, error) {
	builder := sq.Select("*").
		From(table).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		GroupBy(idCol).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %s", err.Error())
	}

	rows, err := r.dbConn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var commands []*model.Command

	if err = pgxscan.ScanAll(&commands, rows); err != nil {
		return nil, err
	}

	return commands, nil
}

// UpdateCommand updates command info
func (r *Repository) UpdateCommand(ctx context.Context, id int64, cmdInfo *model.CommandInfo) error {
	builder := sq.Update(table).
		Where(sq.Eq{idCol: id}).
		PlaceholderFormat(sq.Dollar).
		Set(statusCol, cmdInfo.Status)

	if cmdInfo.Output.Valid {
		builder = builder.Set(outputCol, cmdInfo.Output.String)
	}
	if cmdInfo.Pid.Valid {
		builder = builder.Set(pidCol, cmdInfo.Pid.Int64)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %s", err.Error())
	}

	rows, err := r.dbConn.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

// StopCommand kills running command
func (r *Repository) StopCommand(ctx context.Context, id int64) (int64, error) {
	builder := sq.Select(pidCol).
		From(table).
		Where(sq.Eq{idCol: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return int64(0), fmt.Errorf("failed to build SQL query: %s", err.Error())
	}

	var pid int64

	if err = r.dbConn.QueryRow(ctx, query, args...).Scan(&pid); err != nil {
		return int64(0), err
	}

	return pid, nil
}
