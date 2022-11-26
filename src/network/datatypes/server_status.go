package datatypes

import "encoding/json"

type ServerStatus struct {
	Description Description `json:"description"`
	Version     Version     `json:"version"`
	Players     Players     `json:"players"`
	Favicon     string      `json:"favicon"`
}

type Description struct {
	Text string `json:"text"`
}

type Version struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type Players struct {
	Max    int `json:"max"`
	Online int `json:"online"`
}

func GetDemoServerStatus(desc string, prot int) string {
	st := ServerStatus{
		Description: Description{desc},
		Version:     Version{"mango", prot},
		Players:     Players{69, -1},
		Favicon:     "data:image/png;base64,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",
	}

	res, err := json.Marshal(st)
	if err != nil {
		panic(err)
	}

	return string(res)
}