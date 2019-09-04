package snow

import (
	"encoding/json"
)

type Store interface {
	GetList()
}

// The store obj has Content as we get the data from a file
type ServicenowStore struct {
	Content string
}

// Entire Incidents object
type Incidents struct {
	Name   string     `json:"Name"`
	Report []Incident `json:"Report"`
}

// Individual incident object
type Incident struct {
	Number      string `json:"number"`
	AssignedTo  string `json:"assigned_to"`
	Description string `json:"description"`
	State       string `json:"state"`
	Priority    string `json:"priority"`
	Severity    string `json:"severity"`
}

/*
Initializes the servicenow object with fileStore to read
*/
func Init(content string) (*ServicenowStore, error) {
	var snst ServicenowStore
	// for testing purpose, send file name
	if content != "" {
		snst.Content = content
	} else { // use default file
		snst.Content = defaultContent
	}
	// initialize session with serviceNow store
	return &snst, nil
}

/*
Add GetList method to ServicenowStore
Read the default content or passed content, extract json objects and map it to required struct
Return final json
*/
func (snst *ServicenowStore) GetList() (*[]byte, error) {

	var incidents Incidents

	// encode bytes to incidents struct object by mapping required fields
	err := json.Unmarshal([]byte(snst.Content), &incidents)

	if err != nil {
		return nil, err
	}

	// decode the json to string
	js, err := json.Marshal(&incidents)

	if err != nil {
		return nil, err
	}

	return &js, nil
}
