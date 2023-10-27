package main

import (
	"flag"
	"os"
	"sync"

	"kingsley/dbus/internal/common"

	// system bus observers
	"kingsley/dbus/internal/observers/bluetooth"
	"kingsley/dbus/internal/observers/lid"
	"kingsley/dbus/internal/observers/network"
	"kingsley/dbus/internal/observers/power"

	// session bus observers
	"kingsley/dbus/internal/observers/dnd"
	"kingsley/dbus/internal/observers/screencast"
	"kingsley/dbus/internal/observers/spotify"
)

func newSystemBus() common.Bus {
	bus := common.NewSystemBus()
	bluetooth.AttachToBus(&bus)
	lid.AttachToBus(&bus)
	network.AttachToBus(&bus)
	power.AttachToBus(&bus)
	return bus
}

func newSessionBus() common.Bus {
	bus := common.NewSessionBus()
	spotify.AttachToBus(&bus)
	screencast.AttachToBus(&bus)
	dnd.AttachToBus(&bus)
	return bus
}

func listen(wg *sync.WaitGroup, callback func()) {
	defer wg.Done()
	callback()
}

func main() {
	r := flag.Bool("refresh", false, "refresh dbus sessions")
	flag.Parse()

	sessionBus := newSessionBus()
	systemBus := newSystemBus()

	if *r {
		sessionBus.RefreshObservers()
		systemBus.RefreshObservers()
		os.Exit(0)
	}

	var wg sync.WaitGroup

	wg.Add(2) // should probably do 1 at a time
	go listen(&wg, sessionBus.Listen)
	go listen(&wg, systemBus.Listen)

	wg.Wait()
}
