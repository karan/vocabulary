package vocabulary

import (
  "encoding/json"
  "fmt"
  // "log"
  "io/ioutil"
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

// Makes an http GET request and returns the byte array contents
func makeReq(url string) ([]byte, error) {
  response, err := http.Get(url)
  if err != nil {
      fmt.Printf("%s", err)
      return nil, Error(fmt.Sprintf("%s", err))
  } else {
      defer response.Body.Close()
      contents, err := ioutil.ReadAll(response.Body)
      if err != nil {
          fmt.Printf("%s", err)
          return nil, Error(fmt.Sprintf("%s", err))
      }
      // fmt.Printf("%s\n", string(contents))
      return contents, nil
  }
}

// -----------------------------------------------------------------------------

// Config represents the configuration settings.
type Config struct {
  BigHugeLabsApiKey string  // API key from BigHugeLabs
  WordnikApiKey string  // API key from Wordnik
}

// -----------------------------------------------------------------------------

// Represents the part of speech of a word
type PartOfSpeech struct {
  POS  string // The part of speech for the word
  ExampleUsage  string  // An example usage for the word in POS
}

// Represents a word with all its information
type Word struct {
  Word  string  // The original word in the query
  Meanings []string // A list of meanings for this word
  Synonyms []string // A list of synonyms for this word
  Antonyms []string // A list of antonyms for this word
  PartOfSpeech []PartOfSpeech  // A list of part of speech for this word
  UsageExample []string  // A list of sentences showing usage example for the word
}

// -----------------------------------------------------------------------------

// Vocabulary object represents an instance of vocabulary
type Vocabulary struct {
  c  *Config
}

// Creates a new Word object collecting as much information as possible.
// Requires having all API keys.
func (v Vocabulary) Word(w string) (Word, error) {
  if w == "" {
    return Word{}, Error("word must be non-empty string")
  }

  meanings, err := v.Meanings(w)
  if err != nil {
    return Word{}, err
  }

  return Word{
    Word: w,
    Meanings: meanings,
  }, nil
}

// Returns a list of strings representing the meanings of the given word.
func (v Vocabulary) Meanings(w string) ([]string, error) {
  contents, err := makeReq(fmt.Sprintf(meaningApiUrl, w))
  if err != nil {
    return []string{}, err
  }

  var glosbe Glosbe
  err = json.Unmarshal(contents, &glosbe)

  if err != nil || glosbe.Result != "ok" {
    return []string{}, err
  }

  var result []string
  for _, gm := range glosbe.Tuc[0].Meanings {
    result = append(result, gm.Text)
  }

  return result, nil
}


// -----------------------------------------------------------------------------

// New Instantiates a new instance of Vocabulary with the passed config.
//
func New(c *Config) (Vocabulary, error) {
  v := Vocabulary{c: c}
  return v, nil
}
