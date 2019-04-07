package main

import (
	. "github.com/belajar/config"
	. "github.com/belajar/dao"
)

var dao = MajalahDAO{}
var inputconfig = InputConfig{}
var network = Network{}

func init() {
	inputconfig.Read()
	dao.Database = inputconfig.Database
	dao.Server = inputconfig.Server
	network.Port = inputconfig.Port
}

func main() {
	dao.Koneksi()
	network.Route()
}
