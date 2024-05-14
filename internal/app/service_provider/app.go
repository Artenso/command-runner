package service_provider

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/Artenso/command-runner/internal/config"
	"github.com/Artenso/command-runner/internal/logger"
	desc "github.com/Artenso/command-runner/pkg/command_runner"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	pathToConfig = "config.json"
)

// App is the application struct
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

// NewApp creates new App
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run starts app
func (a *App) Run(ctx context.Context) error {
	group, groupCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		list, err := net.Listen("tcp", config.GetGrpcPort())
		if err != nil {
			return fmt.Errorf("failed to mapping port: %s", err.Error())
		}

		if err := a.grpcServer.Serve(list); err != nil {
			return fmt.Errorf("failed to server: %s", err.Error())
		}

		return nil
	})

	group.Go(func() error {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()} // nolint: staticcheck

		err := desc.RegisterCommandRunnerHandlerFromEndpoint(groupCtx, mux, config.GetGrpcPort(), opts)
		if err != nil {
			return err
		}

		return http.ListenAndServe(config.GetHttpPort(), mux)
	})

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}

// initDeps initialize dependencies
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initServiceProvider,
		a.initGrpcServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	config.Read(pathToConfig)
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	newLogger, err := logger.New(config.GetLogFile())
	if err != nil {
		log.Fatalf("logger settingup error: %s", err.Error())
	}

	logger.SetLogger(newLogger)

	err = logger.SetLogLevel(config.GetLogLevel())
	if err != nil {
		log.Fatalf("logger settingup error: %s", err.Error())
	}

	logger.Sync()
	return nil
}

// initServiceProvider initialize serviceProvider
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

// initGrpcServer initialize gRPC server
func (a *App) initGrpcServer(ctx context.Context) error {
	s := grpc.NewServer()
	desc.RegisterCommandRunnerServer(s, a.serviceProvider.getCommandRunner(ctx))

	reflection.Register(s)

	a.grpcServer = s
	return nil
}
