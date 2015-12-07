package vocabulary

// -----------------------------------------------------------------------------

type GlosbeMeaning struct {
  Text  string  `json:"text"`
}

type GlosbeTuc struct {
  Meanings []GlosbeMeaning  `json:"meanings"`
}

type Glosbe struct {
  Result  string  `json:"result"`
  Tuc  []GlosbeTuc  `json:"tuc"`
}

// -----------------------------------------------------------------------------
