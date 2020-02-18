package inputParser

import (
	"flag"
	"wait4it/model"
)

func GetInput() model.CheckContext {
	ct := flag.String("type", "tcp", "define the type of check, currently [tcp,mysql] are supported")
	timeout := flag.Int("t", 30, "Timeout, amount of time wait4it waits for the port in seconds")
	host := flag.String("h", "127.0.0.1", "IP of the host you want to test against")
	port := flag.Int("p", 0, "Port")
	username := flag.String("u", "", "Username of the service")
	password := flag.String("P", "", "Password of the service")
	databaseName := flag.String("n", "", "Name of the database")
	sslMode := flag.String("ssl", "disable", "Enable or Disable ssl mode (for some database or services)")
	httpCode := flag.Int("status-code", 200, "Status code to be expected from http call")
	httpText := flag.String("http-text", "", "Text to check inside http response")
	flag.Parse()

	c := model.CheckContext{
		Config: model.ConfigurationContext{
			CheckType: *ct,
			Timeout:   *timeout,
		},
		Host:         *host,
		Port:         *port,
		Username:     *username,
		Password:     *password,
		DatabaseName: *databaseName,
		DBConf: model.DatabaseSpecificConf{
			SSLMode: *sslMode,
		},
		HttpConf: model.HttpSpecificConf{
			StatusCode: *httpCode,
			Text:       *httpText,
		},
	}
	return c
}