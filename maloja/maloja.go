package maloja

import (
	"github.com/jeselnik/grobble"
	"github.com/jeselnik/grobble/listenbrainz"
)

const (
	listenBrainzAPI = "/apis/listenbrainz/1/"
)

type Maloja struct {
	InstanceURL, APIKey string
	listenBrainz        *listenbrainz.ListenBrainz
}

func New(m Maloja) *Maloja {
	baseURL := m.InstanceURL + listenBrainzAPI
	lb := listenbrainz.New(listenbrainz.ListenBrainz{Token: m.APIKey, BaseURL: baseURL})
	return &Maloja{InstanceURL: m.InstanceURL, APIKey: m.APIKey, listenBrainz: lb}
}

func (m *Maloja) Login() error {
	return m.listenBrainz.Login()
}

func (m *Maloja) GetServiceName() string {
	return "Maloja server at " + m.listenBrainz.BaseURL
}

func (m *Maloja) Scrobble(t grobble.Track) error {
	return m.listenBrainz.Scrobble(t)
}

func (m *Maloja) BatchScrobble(t []grobble.Track) ([]grobble.Track, []grobble.Track, error) {
	return m.listenBrainz.BatchScrobble(t)
}

func (m *Maloja) CapabilityBatchScrobble() bool {
	return m.listenBrainz.CapabilityBatchScrobble()
}
