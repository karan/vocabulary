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
