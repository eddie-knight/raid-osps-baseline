package armory

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RestData struct {
	owner string
	repo  string
	Repo  RepoData
}

type RepoData struct {
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

func (r *RestData) loadData() error {
	r.owner = GlobalConfig.GetString("owner")
	r.repo = GlobalConfig.GetString("repo")

	if r.Repo.Name == "" {
		r.Repo.getData(r.owner, r.repo)
	}
	return nil
}

func (r *RepoData) getData(owner, repo string) error {
	endpoint := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	responseData, err := makeGetRequest(endpoint, false)
	if err != nil {
		return err
	}
	return json.Unmarshal(responseData, r)
}

func makeGetRequest(endpoint string, authRequired bool) (body []byte, err error) {
	GlobalConfig.Logger.Trace(fmt.Sprintf("GET %s", endpoint))
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	if authRequired && Authenticated {
		request.Header.Set("Authorization", "Bearer "+GlobalConfig.GetString("token"))
	} else if authRequired && !Authenticated {
		err = fmt.Errorf("auth required but not authenticated")
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		err = fmt.Errorf("error making http call: %s", err.Error())
		return nil, err
	}
	return io.ReadAll(response.Body)
}