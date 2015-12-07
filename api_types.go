package vocabulary

import (
  "encoding/json"
)

// -----------------------------------------------------------------------------

type GlosbeThing struct {
  Text  string  `json:"text"`
}

type GlosbeMeanings struct {
  Things []GlosbeThing  `json:"meanings"`
}

type GlosbePhrase struct {
  Thing GlosbeThing  `json:"phrase"`
}

type Glosbe struct {
  Result  string  `json:"result"`
  Tuc     []json.RawMessage  `json:"tuc"`
}

// -----------------------------------------------------------------------------

type UrbanDictThing struct {
  Example  string  `json:"example"`
  ThumbsUp  int `json:"thumbs_up"`
  ThumbsDown  int `json:"thumbs_down"`
}

type UrbanDictResp struct {
  Things  []UrbanDictThing  `json:"list"`
}
