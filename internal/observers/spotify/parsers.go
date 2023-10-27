package spotify

import (
	"kingsley/dbus/internal/common"
	"kingsley/dbus/internal/helpers"
	"reflect"
	"strings"

	"github.com/godbus/dbus/v5"
)

const (
	playbackPlaying = "Playing"
	playbackPaused  = "Paused"
)

func parseVariant(variant []interface{}) map[string]dbus.Variant {
	return common.ParseVariant(common.NameHasOwner.String(), variant)
}

func parsePlaybackStatus(variant dbus.Variant) bool {
	status := variant.Value().(string)
	if status == playbackPaused {
		return false
	}
	return status == playbackPlaying
}

func parseMetadata(variant dbus.Variant) *Metadata {
	metadataMap := variant.Value().(map[string]dbus.Variant)
	metadata := new(Metadata)

	valueOf := reflect.ValueOf(metadata).Elem()
	typeOf := reflect.TypeOf(metadata).Elem()

	for key, val := range metadataMap {
		for i := 0; i < typeOf.NumField(); i++ {
			field := typeOf.Field(i)
			if field.Tag.Get("spotify") == key {
				if key == "xesam:artist" {
					// translate string array to comma separated string
					artists := strings.Join(val.Value().([]string), ", ")
					metadata.Artist = artists
				}
				field := valueOf.Field(i)
				field.Set(reflect.ValueOf(val.Value()))
			}
		}
	}

	return metadata
}

func parseCoverArt(s *Spotify, variant dbus.Variant) {
	url := variant.Value().(map[string]dbus.Variant)["mpris:artUrl"].String()
	if url != "" && url != s.Metadata.ArtURL {
		tmp := "/home/kingsley/.cache/eww/tmp.png"
		url = strings.Trim(url, "\"")
		err := helpers.DownloadCoverArt(url, tmp)
		helpers.LogError("Failed to download cover art.", err)
	}
}
