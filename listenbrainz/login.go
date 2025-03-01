package listenbrainz

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type validateTokenResponse struct {
	Valid    bool   `json:"valid"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	UserName string `json:"user_name"`
}

func (s *ListenBrainz) Auth() error {
	/* User token authentication, no authentication procedure */
	return nil
}

func (s *ListenBrainz) Login() error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", s.BaseURL+"validate-token", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Token "+s.Token)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	tokenResObj := validateTokenResponse{}

	err = json.Unmarshal(resData, &tokenResObj)
	if err != nil {
		return err
	}

	if tokenResObj.Code != 200 {
		return errors.New("login failed with code : ")
	}

	if !tokenResObj.Valid {
		return errors.New("token invalid")
	}

	return nil
}
