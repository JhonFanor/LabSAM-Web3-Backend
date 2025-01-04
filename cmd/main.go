package main

import (
	"LABSAM-WEB3-BACKEND/config"
	"LABSAM-WEB3-BACKEND/internal/adapter/gormadapter"
	"LABSAM-WEB3-BACKEND/internal/api/routes"
	"LABSAM-WEB3-BACKEND/internal/constants"
	"LABSAM-WEB3-BACKEND/internal/logging"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func provideDB(appConfig *config.AppConfig, l logging.Logger) (*gorm.DB, *sql.DB, error) {
	var sslMode string
	if appConfig.Database.SSL {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Bogota",
		appConfig.Database.Host,
		appConfig.Database.User,
		appConfig.Database.Password,
		appConfig.Database.Database,
		appConfig.Database.Port,
		sslMode,
	)

	gormConfig := &gorm.Config{}

	logrusEntry := l.GetLogrusEntry()

	fileLogger := log.New(logrusEntry.Logger.Out, "\r\n", log.LstdFlags)

	if appConfig.LogLevel == constants.LogLevelDebug {
		gormConfig.Logger = logger.New(
			fileLogger,
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
			},
		)
	} else {
		gormConfig.Logger = logger.New(
			fileLogger,
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Silent,
			},
		)
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), gormConfig)

	if err != nil {
		l.LogError("Error connecting to database", err)
		return nil, nil, err
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		l.LogError("Error connecting to database", err)
		return nil, nil, err
	}
	return gormDB, sqlDB, nil
}

func main() {
	app := fx.New(
		fx.StartTimeout(time.Minute*10),
		fx.Provide(
			logging.NewLogger,
			config.GetConfig,
			provideDB,
			gormadapter.NewGormDBManager,
			gormadapter.NewGormQueryManager,
		),
		fx.Invoke(runServer),
	)

	app.Run()
}

func runServer(lc fx.Lifecycle, p routes.RegisterRoutesParams) {
	srv := routes.RegisterRoutes(p)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {

			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			p.Logger.LogInfo(fmt.Sprintf("Starting server on %s", srv.Addr))

			go func() {
				err := srv.Serve(ln)
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					p.Logger.LogError(err.Error())
				}
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			p.Logger.LogInfo("Stopping server...")
			return srv.Shutdown(ctx)
		},
	})
}
