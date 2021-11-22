/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gospl/server"

	"github.com/kardianos/service"
	_ "github.com/kardianos/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serviceLogger service.Logger
var createdService service.Service
var serverPort int

const (
	serviceName        = "NBSProxy"
	serviceDisplayName = "NBS proxy"
	serviceDescription = "NBS proxy service"
)

type serverWrapper struct {
	service service.Service
	server  *server.Server
	exit    chan os.Signal
}

func (w *serverWrapper) Start(s service.Service) error {
	var err error

	w.exit = make(chan os.Signal)
	w.service = s

	if err != nil {
		return err
	}
	_ = serviceLogger.Info("Starting service:" + s.String())

	go w.run()

	_ = serviceLogger.Info("Started service:" + s.String())

	return nil
}

func (w *serverWrapper) run() {

	signal.Notify(w.exit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		err := w.server.Start()
		if err != http.ErrServerClosed {
			_ = serviceLogger.Error(err)
		} else {
			_ = serviceLogger.Info("Server closed")
		}
	}()

	<-w.exit
}

func (w *serverWrapper) Stop(s service.Service) error {
	_ = serviceLogger.Info("Stopping service: " + s.String())
	err := w.server.Stop()
	if err != nil {
		return err
	}

	close(w.exit)
	_ = serviceLogger.Info("Service: " + s.String() + " stopped")

	return nil
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run NBS web services proxy server",
	Long:  ``,
	Args:  cobra.NoArgs,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		serverPort, err = cmd.Flags().GetInt("port")
		if err != nil {
			return err
		}

		baseUrl := viper.GetString("url")
		username := viper.GetString("username")
		password := viper.GetString("password")
		licence := viper.GetString("licence")

		runner := &serverWrapper{}

		runner.server = server.New(serverPort, baseUrl, username, password, licence)

		createdService, err = service.New(runner, &service.Config{
			Name:        serviceName,
			DisplayName: serviceDisplayName,
			Description: serviceDescription,
			Arguments: []string{
				"--url", baseUrl,
				"--username", username,
				"--password", password,
				"--licence", licence,
				"serve",
				"--port", fmt.Sprintf("%d", serverPort),
			},
		})
		if err != nil {
			return err
		}
		serviceLogger, err = createdService.Logger(nil)
		if err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		serviceLogger.Infof("Running server on port: %d", serverPort)

		err := createdService.Run()
		if err != nil {
			return err
		}

		return nil
	},
}

var serviceInstallCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i"},
	Short:   "Install server as system service",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Installing service")
		user := cmd.Flag("user").Value.String()
		if service.Platform() == "windows-service" {
			user = "Nt Authority\\Network service"
		}
		fmt.Println("User:", user)
		err := service.Control(createdService, "install")
		if err != nil {
			return err
		}
		fmt.Println("Service installed")
		return nil
	},
}

var serviceUninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Aliases: []string{"u"},
	Short:   "Uninstall server as system service",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Uninstalling service")
		err := service.Control(createdService, "uninstall")
		if err != nil {
			return err
		}
		fmt.Println("Service uninstalled")
		return nil
	},
}

var serviceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server system service",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Starting service")
		err := service.Control(createdService, "start")
		if err != nil {
			return err
		}
		fmt.Println("Service started")
		return nil
	},
}

var serviceStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop server system service",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Stopping service")
		err := service.Control(createdService, "stop")
		if err != nil {
			return err
		}
		fmt.Println("Service stopped")
		return nil
	},
}

var serviceRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart server system service",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Restarting service")
		err := service.Control(createdService, "restart")
		if err != nil {
			return err
		}
		fmt.Println("Service restarted")
		return nil
	},
}

var serviceStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get server system service status",
	RunE: func(cmd *cobra.Command, args []string) error {
		status, err := createdService.Status()
		if err != nil {
			return err
		}
		switch status {
		case service.StatusRunning:
			fmt.Println("Service is running")
		case service.StatusStopped:
			fmt.Println("Service is stopped")
		case service.StatusUnknown:
			fmt.Println("Service status is unknown, or service is not installed")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.AddCommand(serviceInstallCmd)
	serveCmd.AddCommand(serviceUninstallCmd)
	serveCmd.AddCommand(serviceStartCmd)
	serveCmd.AddCommand(serviceStopCmd)
	serveCmd.AddCommand(serviceRestartCmd)
	serveCmd.AddCommand(serviceStatusCmd)

	serveCmd.PersistentFlags().IntP("port", "p", 31100, "Port server will listen at")

	serviceInstallCmd.Flags().String("user", "", "User service will run as")

}
