package cmd

import (
	"wait4it/MongoDbChecker"
	"wait4it/MySQLChecker"
	"wait4it/PostgreSQLChecker"
	"wait4it/RedisChecker"
	"wait4it/TcpChecker"
	"wait4it/httpChecker"
)

var cm = map[string]interface{}{
	"tcp":      &TcpChecker.Tcp{},
	"mysql":    &MySQLChecker.MySQLConnection{},
	"postgres": &PostgreSQLChecker.PostgresSQLConnection{},
	"http":     &httpChecker.HttpCheck{},
	"mongo":    &MongoDbChecker.MongoDbConnection{},
	"redis":    &RedisChecker.RedisConnection{},
}
