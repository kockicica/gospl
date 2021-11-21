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

const (
	serviceName        = "NBSProxy"
	serviceDisplayName = "NBS proxy"
	serviceDescription = "NBS proxy service"
)

//type program struct {
//	service  service.Service
//	server   *http.Server
//	listener net.Listener
//	exit     chan os.Signal
//	port     int
//}

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

//func (p *program) Start(s service.Service) error {
//	var err error
//
//	if serviceLogger == nil {
//		var err error
//		serviceLogger, err = s.Logger(nil)
//		if err != nil {
//			return err
//		}
//	}
//
//	_ = serviceLogger.Info("Starting service: " + s.String())
//
//	p.exit = make(chan os.Signal)
//	p.service = s
//	p.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", p.port))
//	if err != nil {
//		return err
//	}
//
//	http.DefaultServeMux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
//		writer.WriteHeader(http.StatusOK)
//		_, err := strings.NewReader("dadad").WriteTo(writer)
//		if err != nil {
//			return
//		}
//	})
//	p.server = &http.Server{Handler: http.DefaultServeMux}
//
//	go p.run()
//
//	_ = serviceLogger.Info("Started service: " + s.String() + fmt.Sprintf(" listening on port:%d", p.port))
//	return nil
//}
//
//func (p *program) Stop(s service.Service) error {
//	_ = serviceLogger.Info("Stopping service: " + s.String())
//	//err := p.server.Close()
//	err := p.server.Shutdown(context.Background())
//	if err != nil {
//		return err
//	}
//	close(p.exit)
//	_ = serviceLogger.Info("Stopped service: " + s.String())
//	return nil
//}
//
//func (p *program) run() {
//	signal.Notify(p.exit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
//
//	go func() {
//		err := p.server.Serve(p.listener)
//		if err == http.ErrServerClosed {
//			_ = serviceLogger.Error(err)
//		}
//	}()
//
//	<-p.exit
//}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run NBS web services proxy",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			return err
		}

		baseUrl := viper.GetString("url")
		username := viper.GetString("username")
		password := viper.GetString("password")
		licence := viper.GetString("licence")

		runner := &serverWrapper{}
		service, err := service.New(runner, &service.Config{
			Name:        serviceName,
			DisplayName: serviceDisplayName,
			Description: serviceDescription,
		})
		serviceLogger, err = service.Logger(nil)
		if err != nil {
			return err
		}
		runner.server = server.New(port, baseUrl, username, password, licence)

		if err != nil {
			return err
		}

		serviceLogger.Infof("Running server on port: %d", port)

		err = service.Run()
		if err != nil {
			return err
		}

		//runner := &program{
		//	port: port,
		//}
		//
		//service, err := service.New(runner, &service.Config{
		//	Name:        serviceName,
		//	DisplayName: serviceDisplayName,
		//	Description: serviceDescription,
		//})
		//
		//if err != nil {
		//	return err
		//}
		//
		//err = service.Run()
		//if err != nil {
		//	return err
		//}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP("port", "p", 31100, "Port server will listen at")

}
