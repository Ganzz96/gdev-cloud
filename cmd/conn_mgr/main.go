package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/ganzz96/gdev-cloud/internal/common/configurer"
	"github.com/ganzz96/gdev-cloud/internal/conn_mgr"
	"github.com/ganzz96/gdev-cloud/internal/conn_mgr/config"
	"github.com/ganzz96/gdev-cloud/internal/conn_mgr/transport/websocket"
)

func main() {
	cli := cobra.Command{
		Use:   "conn_mgr",
		Short: "Connection Manager",
	}

	cli.AddCommand(runCommand())

	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run service",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfgPath, err := cmd.Flags().GetString("config")
			if err != nil {
				return errors.WithStack(err)
			}

			return initAndRun(cfgPath)
		},
	}

	cmd.Flags().StringP("config", "c", "", "")
	_ = cmd.MarkFlagRequired("config")

	return cmd
}

func initAndRun(cfgPath string) error {
	var cfg config.Config

	if err := configurer.Load(cfgPath, &cfg); err != nil {
		return errors.WithMessage(err, "failed to load config")
	}

	connectionManager := conn_mgr.New(nil)

	transportLayer := websocket.New(connectionManager)
	transportLayer.RegisterEndpoints("/connect")

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	return http.ListenAndServe(addr, nil)
}
