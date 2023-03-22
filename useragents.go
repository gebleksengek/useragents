package useragents

import (
	_ "embed"
	"encoding/json"
	"math/rand"
)

type UserAgents struct {
	Chrome  []string `json:"chrome"`
	Edge    []string `json:"edge"`
	Firefox []string `json:"firefox"`
	Safari  []string `json:"safari"`
}

//go:embed useragents.json
var useragents []byte

// List of available user-agent
var List = UserAgents{}

func init() {
	err := json.Unmarshal(useragents, &List)
	if err != nil {
		panic(err)
	}
}

func random(list []string) string {
	if len(list) <= 0 {
		return ""
	}

	return list[rand.Intn(len(list))]
}

// Generate a random user-agent
func Random() string {
	switch rand.Intn(4) {
	case 0:
		return Chrome()
	case 1:
		return Edge()
	case 2:
		return Firefox()
	default:
		return Safari()
	}
}

// Generate a random user-agent with the latest version
func RandomLatest() string {
	switch rand.Intn(4) {
	case 0:
		return ChromeLatest()
	case 1:
		return EdgeLatest()
	case 2:
		return FirefoxLatest()
	default:
		return SafariLatest()
	}
}

// Generate a chrome based user-agent
func Chrome() string {
	return random(List.Chrome)
}

// Generate a chrome based user-agent with the latest version
func ChromeLatest() string {
	if len(List.Chrome) <= 0 {
		return ""
	}

	return List.Chrome[0]
}

// Generate a edge based user-agent
func Edge() string {
	return random(List.Edge)
}

// Generate a edge based user-agent with the latest version
func EdgeLatest() string {
	if len(List.Edge) <= 0 {
		return ""
	}

	return List.Edge[0]
}

// Generate a firefox based user-agent
func Firefox() string {
	return random(List.Firefox)
}

// Generate a firefox based user-agent with the latest version
func FirefoxLatest() string {
	if len(List.Firefox) <= 0 {
		return ""
	}

	return List.Firefox[0]
}

// Generate a safari based user-agent
func Safari() string {
	return random(List.Safari)
}

// Generate a safari based user-agent with the latest version
func SafariLatest() string {
	if len(List.Safari) <= 0 {
		return ""
	}

	return List.Safari[0]
}
