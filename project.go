package gopherstack

import (
	"net/url"
)

// List the available CloudStack projects
func (c CloudStackClient) ListProjects(name string) (string, error) {
	params := url.Values{}

	if name != "" {
		params.Set("name", name)
	}
	_, err := NewRequest(c, "listProjects", params)
	if err != nil {
		return "", err
	}

	// response.(ListProjectsResponse).Listprojectsresponse.Count
	// response.(ListProjectsResponse).Project[0]
	return "", err
}

type Project struct {
	Account     string        `json:"account"`
	Displaytext string        `json:"displaytext"`
	Domain      string        `json:"domain"`
	Domainid    string        `json:"domainid"`
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	State       string        `json:"state"`
	Tags        []interface{} `json:"tags"`
}

type ListProjectsResponse struct {
	Listprojectsresponse struct {
		Count   float64   `json:"count"`
		Project []Project `json:"project"`
	} `json:"listprojectsresponse"`
}