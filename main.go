package main

import (
	"github.com/krau/shisoimg/cmd"
	"github.com/krau/shisoimg/dao"
)

func main() {
	dao.Init()
	cmd.Execute()
}
