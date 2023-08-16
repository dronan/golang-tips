package main

import (
	"encoding/json"
)

type CheckDetail struct {
	DefinitionID string                 `json:"definition_id"`
	Success      bool                   `json:"success"`
	Reason       string                 `json:"reason"`
	Specific     map[string]interface{} `json:"specific"`
}

type APMJson struct {
	Type        string                 `json:"type"`
	PerimeterID string                 `json:"perimeter_id"`
	Grade       string                 `json:"grade"`
	Score       int                    `json:"score"`
	ChecksCount int                    `json:"checks_count"`
	ChecksOK    int                    `json:"checks_ok"`
	ChecksKO    int                    `json:"checks_ko"`
	Ctime       string                 `json:"ctime"`
	Mtime       string                 `json:"mtime"`
	Specific    map[string]interface{} `json:"specific"`
	Details     []CheckDetail          `json:"details"`
}

func debugJsonResponse() APMJson {

	data := `[
		{
	"type": "apm",
	"perimeter_id": "",
	"grade": "A",
	"score": 100,
	"checks_count": 10,
	"checks_ok": 2,
	"checks_ko": 8,
	"ctime": "",
	"mtime": "",
	"specific": {},
	"details": [
		{
			"definition_id": "",
			"success": true,
			"reason": "",
			"specific": {}
		},
		{
			"definition_id": "",
			"success": false,
			"reason": "",
			"specific": {}
		}
	]
}
]`

	var result APMJson
	// get the json response from the API and unmarshal it to the result variable
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		panic(err)
	}

	return result
}
