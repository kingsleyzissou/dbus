package spotify

import (
	"encoding/json"
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

type Spotify struct {
	Running  bool     `json:"running"`
	Playing  bool     `json:"playing"`
	Metadata Metadata `json:"metadata"`

	// we don't want these to be serialized
	signals   []common.Signal
	listeners []common.Listener
}

type Metadata struct {
	Artist  string   `json:"artist"`
	Artists []string `spotify:"xesam:artist"`
	Title   string   `spotify:"xesam:title" json:"title"`
	Album   string   `spotify:"xesam:album" json:"album"`
	Length  uint64   `spotify:"mpris:length" json:"length"`
	ArtURL  string   `spotify:"mpris:artUrl" json:"art"`
}

// Helpers
func (s Spotify) notify() {
	jsonSpotify, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	common.Notify("spotify", string(jsonSpotify))
}

func (s *Spotify) GetListeners() []common.Listener {
	return s.listeners
}

func (s *Spotify) GetSignals() []common.Signal {
	return s.signals
}
func (s *Spotify) Refresh(connection *dbus.Conn) {
	s.Running = isRunning(connection, callObjects[running])

	if s.Running {
		s.Playing = isPlaying(connection, callObjects[playing])
		s.Metadata = getMetadata(connection, callObjects[metadata])
	}

	s.notify()
}

func (s *Spotify) Initialise(connection *dbus.Conn) {
	s.signals = signals
	s.listeners = []common.Listener{
		s.isPlayingListener,
		s.metadataListener,
		s.coverArtListener,
		s.isRunningListener,
	}
	s.Refresh(connection)
}

func (s *Spotify) AttachSignals(attach func(common.Signal)) {
	for _, signal := range s.signals {
		attach(signal)
	}
}

func (s *Spotify) DetachSignals(detach func(common.Signal)) {
	for _, signal := range s.signals {
		detach(signal)
	}
}

func (s *Spotify) Broadcast(message *common.Message) {
	for _, listener := range s.listeners {
		listener(message)
	}
}

// Constructors
func AttachToBus(bus *common.Bus) {
	s := Spotify{}
	s.Initialise(bus.Connection)
	bus.RegisterObserver(&s)
}
