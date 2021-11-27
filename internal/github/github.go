package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tunaitis/contributor-map/internal/util"
	"strings"
)

type contributor struct {
	Login string
	Location string
	Contributions int
}

func GetContributors(accessToken string, name string, page int, useCache bool) ([]contributor, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/contributors?per_page=100&page=%d", name, page)

	resp, err := util.HttpRequest(accessToken, "GET", url, nil, useCache)
	if err != nil {
		return nil, err
	}

	var contributors []contributor
	err = json.Unmarshal(resp, &contributors)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("repository not found: %s", name))
	}

	if len(contributors) == 100 {
		x, err := GetContributors(accessToken, name, page+1, useCache)
		if err != nil {
			return nil, err
		}

		for _, c := range x {
			contributors = append(contributors, c)
		}
	}

	return contributors, nil
}

func prepareLocationsQuery(contributors []contributor) []byte {
	var sb strings.Builder
	sb.WriteString("{")
	for i, c := range contributors {
		if strings.Contains(c.Login, "[") {
			continue
		}
		sb.WriteString(fmt.Sprintf(`user%d:user(login: "%s"){location login}`, i, c.Login))
	}
	sb.WriteString("}")

	jsonData := map[string]string{
		"query": sb.String(),
	}
	jsonValue, _ := json.Marshal(jsonData)

	return jsonValue
}

func GetLocations(accessToken string, contributors []contributor, useCache bool) ([]contributor, error) {
	payload := prepareLocationsQuery(contributors)

	body, err := util.HttpRequest(accessToken, "POST",
		"https://api.github.com/graphql",
		payload, useCache)

	var resp map[string]map[string]map[string]string
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	for _, k := range resp["data"] {
		loc := k["location"]
		uid := k["login"]
		if loc == "" {
			continue
		}

		for i := range contributors {
			if contributors[i].Login == uid {
				contributors[i].Location = loc
			}
		}
	}

	return contributors, nil
}
