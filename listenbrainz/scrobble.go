package listenbrainz

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/jeselnik/grobble"
)

type listenResponse struct {
	Code   int    `json:"code"`
	Error  string `json:"error"`
	Status string `json:"status"`
}

type trackMetadata struct {
	ArtistName  string `json:"artist_name"`
	TrackName   string `json:"track_name"`
	ReleaseName string `json:"release_name"`
}

type trackPayload struct {
	ListenedAt    int           `json:"listened_at"`
	TrackMetadata trackMetadata `json:"track_metadata"`
}

type listenBrainzListen struct {
	ListenType string         `json:"listen_type"`
	Payload    []trackPayload `json:"payload"`
}

func (s *ListenBrainz) Scrobble(t grobble.Track) error {
	tr := []grobble.Track{}
	tr = append(tr, t)
	_, _, scr := s.BatchScrobble(tr)
	return scr
}

func (s *ListenBrainz) BatchScrobble(tracks []grobble.Track) ([]grobble.Track, []grobble.Track, error) {
	emptyTrackSlice := []grobble.Track{}

	if len(tracks) > grobble.BATCH_SCROBBLE_MAX {
		return emptyTrackSlice, tracks, errors.New("length of track slice is _, greater than the max of 50")
	}

	p := []trackPayload{}

	for _, t := range tracks {
		p = append(p, trackPayload{
			ListenedAt: t.Timestamp,
			TrackMetadata: trackMetadata{
				ArtistName:  t.Artist,
				TrackName:   t.Title,
				ReleaseName: t.Album,
			},
		})
	}

	listen := listenBrainzListen{
		ListenType: "import",
		Payload:    p,
	}

	jListen, err := json.Marshal(listen)
	if err != nil {
		return emptyTrackSlice, tracks, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", s.BaseURL+"submit-listens", bytes.NewBuffer(jListen))
	if err != nil {
		return emptyTrackSlice, tracks, err
	}
	req.Header.Set("Authorization", "Token "+s.Token)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return emptyTrackSlice, tracks, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return emptyTrackSlice, tracks, err
	}

	resBody := listenResponse{}
	err = json.Unmarshal(resData, &resBody)
	if err != nil {
		return emptyTrackSlice, tracks, err
	}

	if resBody.Status != "ok" {
		return emptyTrackSlice, tracks, errors.New("failed to scrobble with error: " + resBody.Error)
	}

	return tracks, emptyTrackSlice, nil
}

func (s *ListenBrainz) CapabilityBatchScrobble() bool {
	return true
}
