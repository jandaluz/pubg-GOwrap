package data

// generic responses
type ResponseData struct {
	Data  []interface{} `json: "data"`
	Links interface{}   `json: "links"`
	Meta  interface{}   `json: "meta"`
}

type PubgData struct {
	Type          string      `json: "type"`
	ID            string      `json: "id"`
	PubgInfo      interface{} `json: "attributes,omitempty"`
	Relationships interface{} `json: "relationships,omitempty"`
	Links         interface{} `json: "links,omitempty"`
}

// player api responses
type PlayerResponse struct {
	Data  []PlayerData `json: "data"`
	Links interface{}  `json: "links"`
	Meta  interface{}  `json: "meta"`
}

type PlayerData struct {
	Type          string           `json: "type"`
	ID            string           `json: "id"`
	Attributes    PlayerAttributes `json: "attributes,"`
	Relationships interface{}      `json: "relationships"`
	Links         interface{}      `json: "links"`
}

type PlayerAttributes struct {
	Name         string `json: "name"`
	ShardID      string `json: "shardId"`
	PatchVersion string `json: "patchVersion,omitempty"`
}

type RelationshipData struct {
	Data PubgData `json: "data"`
}
