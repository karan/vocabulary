package vocabulary_test

import (
  "testing"
  "reflect"

  "github.com/karan/vocabulary"
  e "github.com/karan/vocabulary/examples"
)

// -----------------------------------------------------------------------------

func TestMeanings(t *testing.T) {
  c := &vocabulary.Config{BigHugeLabsApiKey: "", WordnikApiKey: ""}
  v, err := vocabulary.New(c)
  if err != nil {
    t.Errorf("Test failed: %s", err)
  }

  expected := []string{"A plastic blowing horn, typically 65 cm long, that produces a loud and monotone note."}
  actual, _ := v.Meanings("vuvuzela")
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}

func TestMeaningsInvalidWord(t *testing.T) {
  c := &vocabulary.Config{BigHugeLabsApiKey: "", WordnikApiKey: ""}
  v, err := vocabulary.New(c)
  if err != nil {
    t.Errorf("Test failed: %s", err)
  }

  expected := []string{}
  actual, _ := v.Meanings("asfbhsdjfhdsj")
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}

// -----------------------------------------------------------------------------

func TestSynonyms(t *testing.T) {
  c := &vocabulary.Config{BigHugeLabsApiKey: "", WordnikApiKey: ""}
  v, err := vocabulary.New(c)
  if err != nil {
    t.Errorf("Test failed: %s", err)
  }

  expected := []string{"lepatata"}
  actual, _ := v.Synonyms("vuvuzela")
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}

func TestSynonymsInvalidWord(t *testing.T) {
  c := &vocabulary.Config{BigHugeLabsApiKey: "", WordnikApiKey: ""}
  v, err := vocabulary.New(c)
  if err != nil {
    t.Errorf("Test failed: %s", err)
  }

  expected := []string{}
  actual, _ := v.Synonyms("asfbhsdjfhdsj")
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}

// -----------------------------------------------------------------------------

func TestAntonyms(t *testing.T) {
  c := &vocabulary.Config{BigHugeLabsApiKey: e.BigHugeLabsApiKey, WordnikApiKey: ""}
  v, err := vocabulary.New(c)
  if err != nil {
    t.Errorf("Test failed: %s", err)
  }

  expected := []string{"hate"}
  actual, _ := v.Antonyms("love")
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}

func TestAntonymsInvalidWord(t *testing.T) {
  c := &vocabulary.Config{BigHugeLabsApiKey: e.BigHugeLabsApiKey, WordnikApiKey: ""}
  v, err := vocabulary.New(c)
  if err != nil {
    t.Errorf("Test failed: %s", err)
  }

  expected := []string{}
  actual, _ := v.Antonyms("asfbhsdjfhdsj")
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}

// -----------------------------------------------------------------------------

func TestPartOfSpeech(t *testing.T) {
  c := &vocabulary.Config{BigHugeLabsApiKey: "", WordnikApiKey: e.WordnikApiKey}
  v, err := vocabulary.New(c)
  if err != nil {
    t.Errorf("Test failed: %s", err)
  }

  expected := vocabulary.PartOfSpeech{"adverb", "With speed; in a rapid manner."}
  actual, _ := v.PartOfSpeech("rapidly")
  if !reflect.DeepEqual(expected, actual[0]) {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}

func TestPartOfSpeechInvalidWord(t *testing.T) {
  c := &vocabulary.Config{BigHugeLabsApiKey: "", WordnikApiKey: e.WordnikApiKey}
  v, err := vocabulary.New(c)
  if err != nil {
    t.Errorf("Test failed: %s", err)
  }

  expected := []vocabulary.PartOfSpeech{}
  actual, _ := v.PartOfSpeech("asfbhsdjfhdsj")
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}

// -----------------------------------------------------------------------------

func TestUsageExample(t *testing.T) {
  c := &vocabulary.Config{BigHugeLabsApiKey: "", WordnikApiKey: ""}
  v, err := vocabulary.New(c)
  if err != nil {
    t.Errorf("Test failed: %s", err)
  }

  expected := []string{"I went to the to of the hillock to look around."}
  actual, _ := v.UsageExample("hillock")
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}

func TestUsageExampleInvalidWord(t *testing.T) {
  c := &vocabulary.Config{BigHugeLabsApiKey: "", WordnikApiKey: ""}
  v, err := vocabulary.New(c)
  if err != nil {
    t.Errorf("Test failed: %s", err)
  }

  expected := []string{}
  actual, _ := v.UsageExample("asfbhsdjfhdsj")
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}
