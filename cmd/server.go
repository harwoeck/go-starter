package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"allaboutapps.dev/aw/go-starter/internal/api/router"
	"allaboutapps.dev/aw/go-starter/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	migrateFlag string = "migrate"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the server",
	Long: `Starts the stateless RESTful JSON server

Requires configuration through ENV and
and a fully migrated PostgreSQL database.`,
	Run: func(cmd *cobra.Command, args []string) {
		applyMigrations, err := cmd.Flags().GetBool(migrateFlag)
		if err != nil {
			fmt.Printf("Error while parsing flags: %v\n", err)
			os.Exit(1)
		}
		if applyMigrations {
			migrateCmdFunc(cmd, args)
		}

		runServer()
	},
}

func init() {
	serverCmd.Flags().BoolP(migrateFlag, "m", false, "Apply migrations before startup.")
	rootCmd.AddCommand(serverCmd)
}

func runServer() {
	config := config.DefaultServiceConfigFromEnv()

	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.SetGlobalLevel(config.Logger.Level)
	if config.Logger.PrettyPrintConsole {
		log.Logger = log.Output(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.TimeFormat = "15:04:05"
		}))
	}

	s := api.NewServer(config)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := s.InitDB(ctx); err != nil {
		cancel()
		log.Fatal().Err(err).Msg("Failed to initialize database")
	}
	cancel()

	if err := s.InitMailer(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize mailer")
	}

	if err := s.InitPush(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize push service")
	}

	router.Init(s)

	go func() {
		if err := s.Start(); err != nil {
			if err == http.ErrServerClosed {
				log.Info().Msg("Server closed")
			} else {
				log.Fatal().Err(err).Msg("Failed to start server")
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Failed to gracefully shut down server")
	}
}
