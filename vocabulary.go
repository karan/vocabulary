package vocabulary

import (
  "fmt"
  "log"
  "net/http"
)

// -----------------------------------------------------------------------------

const meaningApiUrl string = "https://glosbe.com/gapi/translate?from=en&dest=en&format=json&pretty=true&phrase=%s"
const synonymsApiUrl string = "https://glosbe.com/gapi/translate?from=en&dest=en&format=json&pretty=true&phrase=%s"
const antonymsApiUrl string = "http://words.bighugelabs.com/api/2/%s/%s/json"
const partOfSpeechApiUrl string = "http://api.wordnik.com/v4/word.json/{word}/{action}?api_key=%s"
const usageExampleApiUrl string = "http://api.urbandictionary.com/v0/define?term=%s"

// -----------------------------------------------------------------------------

// Error represents an error.
type Error string

// Error implements the built-in error interface.
func (e Error) Error() string {
  return string(e)
}

// -----------------------------------------------------------------------------

// Config represents the configuration settings.
type Config struct {
  BigHugeLabsApiKey string  // API key from BigHugeLabs
  WordnikApiKey string  // API key from Wordnik
}

// -----------------------------------------------------------------------------

// Vocabulary object represents an instance of vocabulary
type Vocabulary struct {
  c  *Config
}

// -----------------------------------------------------------------------------

// New Instantiates a new instance of Vocabulary with the passed config.
//
func New(c *Config) (Vocabulary, error) {
  v := Vocabulary{c: c}
  return v, nil
}
