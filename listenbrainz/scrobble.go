package listenbrainz

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

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
	listen := listenBrainzListen{
		ListenType: "single",
		Payload: []trackPayload{
			{ListenedAt: int(time.Now().Unix()),
				TrackMetadata: trackMetadata{
					ArtistName:  t.Artist,
					TrackName:   t.Title,
					ReleaseName: t.Album,
				}},
		},
	}

	jListen, err := json.Marshal(listen)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", s.BaseURL+"submit-listens", bytes.NewBuffer(jListen))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Token "+s.Token)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	resBody := listenResponse{}
	err = json.Unmarshal(resData, &resBody)
	if err != nil {
		return err
	}

	if resBody.Status != "ok" {
		return errors.New("failed to scrobble with error: " + resBody.Error)
	}

	return nil
}

func (s *ListenBrainz) BatchScrobble([]grobble.Track) (int, int) {
	return 0, 0
}
