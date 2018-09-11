package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/al45tair/docker-machine-driver-linode/pkg/drivers/linode"
)

func main() {
	plugin.RegisterDriver(linode.NewDriver("", ""))
}
