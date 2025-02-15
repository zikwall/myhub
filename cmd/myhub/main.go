package main

import (
	"context"
	"net"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"

	"github.com/zikwall/myhub/config"
	"github.com/zikwall/myhub/internal/application/container"
	"github.com/zikwall/myhub/internal/infrastructure/http/myhub"
	"github.com/zikwall/myhub/pkg/fiberext"
	"github.com/zikwall/myhub/pkg/log"
	"github.com/zikwall/myhub/pkg/signal"
)

func main() {
	application := cli.App{
		Name: "muHub",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config-file",
				Required: true,
				Usage:    "YAML config filepath",
				EnvVars:  []string{"CONFIG_FILE"},
				FilePath: "/srv/secret/config_file",
			},
			&cli.StringFlag{
				Name:     "bind-address",
				Usage:    "IP и порт сервера, например: 0.0.0.0:3023",
				Required: false,
				Value:    "0.0.0.0:3023",
				EnvVars:  []string{"BIND_ADDRESS"},
			},
			&cli.StringFlag{
				Name:     "bind-socket",
				Usage:    "Путь к Unix сокет файлу",
				Required: false,
				Value:    "/tmp/myhub.sock",
				EnvVars:  []string{"BIND_SOCKET"},
			},
			&cli.IntFlag{
				Name:     "listener",
				Usage:    "Unix socket or TCP",
				Required: false,
				Value:    1,
				EnvVars:  []string{"LISTENER"},
			},
		},
		Action: Main,
		After: func(_ *cli.Context) error {
			log.Info("stopped")
			return nil
		},
	}

	if err := application.Run(os.Args); err != nil {
		log.Error(err)
	}
}

// nolint:funlen // it's ok
func Main(ctx *cli.Context) error {
	appContext, cancel := context.WithCancel(ctx.Context)
	defer func() {
		cancel()
		log.Info("app context is canceled, myHub is down!")
	}()

	cfg, err := config.New(ctx.String("config-file"))
	if err != nil {
		return err
	}

	contrib, err := container.New(appContext, cfg)
	if err != nil {
		return err
	}

	defer func() {
		contrib.Shutdown(func(err error) {
			log.Warning(err)
		})
		contrib.Stacktrace()
	}()

	await, stop := signal.Notifier(func() {
		log.Info("start shutdown process..")
	})

	go func() {
		app := fiber.New(fiber.Config{
			ServerHeader: "myHub Server",
			Prefork:      cfg.Prefork,
			ErrorHandler: fiberext.ErrorHandler,
		})

		myhub.New(nil).MountRoutes(app)

		var ln net.Listener
		if ln, err = signal.Listener(
			contrib.Context(),
			ctx.Int("listener"),
			ctx.String("bind-socket"),
			ctx.String("bind-address"),
		); err != nil {
			stop(err)
			return
		}
		if err = app.Listener(ln); err != nil {
			stop(err)
		}
	}()

	log.Info("myHub: service is launched")
	return await()
}
