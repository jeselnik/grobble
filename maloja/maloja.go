package maloja

import (
	"errors"

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
	instanceURL, apiKey string
	listenBrainz        *listenbrainz.ListenBrainz
}

func New(p Params) (*Maloja, error) {
	malojaInstance := &Maloja{}

	if p.InstanceURL == "" {
		return malojaInstance, errors.New("an instance URL must be set")
	}

	if p.APIKey == "" {
		return malojaInstance, errors.New("an API key must be set")
	}

	baseURL := p.InstanceURL + listenBrainzAPI

	lb := listenbrainz.New(listenbrainz.Params{Token: p.APIKey, BaseURL: baseURL})
	malojaInstance.listenBrainz = lb

	malojaInstance.instanceURL = p.InstanceURL
	malojaInstance.apiKey = p.APIKey
	return malojaInstance, nil
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
