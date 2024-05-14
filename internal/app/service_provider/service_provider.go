package service_provider

import (
	"context"

	cmdImpl "github.com/Artenso/command-runner/internal/api/command_runner"
	"github.com/Artenso/command-runner/internal/logger"
	cmdRepo "github.com/Artenso/command-runner/internal/repository/postgres"
	cmdService "github.com/Artenso/command-runner/internal/service/command_runner"
	systemCaller "github.com/Artenso/command-runner/internal/service/system_caller"
	"github.com/jackc/pgx/v5/pgxpool"
)

// serviceProvider di-container
type serviceProvider struct {
	dbConn         *pgxpool.Pool
	repository     *cmdRepo.Repository
	systemCaller   *systemCaller.Service
	service        *cmdService.Service
	implementation *cmdImpl.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) getDbConn(ctx context.Context) *pgxpool.Pool {
	if s.dbConn == nil {
		dbDSN := "postgres://postgres:postgres@postgres:5432/commands_storage"
		conn, err := pgxpool.New(ctx, dbDSN)
		if err != nil {
			logger.Fatalf("failed to init db connection: %s", err.Error())
		}

		s.dbConn = conn
	}

	return s.dbConn
}

func (s *serviceProvider) getRepository(ctx context.Context) *cmdRepo.Repository {
	if s.repository == nil {
		s.repository = cmdRepo.New(s.getDbConn(ctx))
	}
	return s.repository
}

func (s *serviceProvider) getSystemCaller() *systemCaller.Service {
	if s.systemCaller == nil {
		s.systemCaller = systemCaller.New()
	}
	return s.systemCaller
}

func (s *serviceProvider) getService(ctx context.Context) *cmdService.Service {
	if s.service == nil {
		s.service = cmdService.New(s.getRepository(ctx), s.getSystemCaller())
	}
	return s.service
}

func (s *serviceProvider) getCommandRunner(ctx context.Context) *cmdImpl.Implementation {
	if s.implementation == nil {
		s.implementation = cmdImpl.NewCommandRunner(s.getService(ctx))
	}
	return s.implementation
}
