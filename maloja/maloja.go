package maloja

import (
	"github.com/jeselnik/grobble"
	"github.com/jeselnik/grobble/listenbrainz"
)

const (
	listenBrainzAPI = "/apis/listenbrainz/1/"
)

type Params struct {
	InstanceURL, APIKey string
}

type Maloja struct {
	instanceURL  string
	listenBrainz *listenbrainz.ListenBrainz
}

func New(p Params) *Maloja {
	baseURL := p.InstanceURL + listenBrainzAPI
	lb := listenbrainz.New(listenbrainz.Params{Token: p.APIKey, BaseURL: baseURL})
	return &Maloja{instanceURL: p.InstanceURL, listenBrainz: lb}
}

func (m *Maloja) Login() error {
	return m.listenBrainz.Login()
}

func (m *Maloja) GetServiceName() string {
	return "Maloja server at " + m.instanceURL
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
