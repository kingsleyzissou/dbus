package screencast

import (
	"kingsley/dbus/internal/common"
)

func (s *Screencast) isSharingListener(m *common.Message) {
	for _, variant := range m.Headers {
		if variant.Value() == screensharingOpened {
			s.Sharing = true
			s.notify()
			break
		}
		if variant.Value() == screensharingClosed {
			s.Sharing = false
			s.notify()
			break
		}
	}
}
