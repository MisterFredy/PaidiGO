package main

import (
	. "github.com/belajar/config"
	. "github.com/belajar/dao"
)

var core = Core{}
var inputconfig = InputConfig{}
var network = Network{}

func init() {
	inputconfig.Read()
	core.Database = inputconfig.Database
	core.Server = inputconfig.Server
	network.Port = inputconfig.Port
}

func main() {
	core.Koneksi()
	network.Route()
}
