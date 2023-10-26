package spotify

import (
	"kingsley/dbus/internal/common"
)

func (s *Spotify) isRunningListener(m *common.Message) {
	variant := parseVariant(m.Variant)
	if val, ok := variant[running.String()]; ok {
		s.Running = val.Value().(bool)
		s.notify()
	}
}
func (s *Spotify) isPlayingListener(m *common.Message) {
	variant := parseVariant(m.Variant)
	if val, ok := variant[playing.String()]; ok {
		s.Playing = parsePlaybackStatus(val)
		s.notify()
	}
}

func (s *Spotify) metadataListener(m *common.Message) {
	variant := parseVariant(m.Variant)
	if val, ok := variant[metadata.String()]; ok {
		metadata := parseMetadata(val)
		s.Metadata = *metadata
		s.notify()
	}
}

func (s *Spotify) coverArtListener(m *common.Message) {
	variant := parseVariant(m.Variant)
	if val, ok := variant[metadata.String()]; ok {
		parseCoverArt(s, val)
	}
}
