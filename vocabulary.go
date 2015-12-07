package vocabulary

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
)

// -----------------------------------------------------------------------------

const meaningApiUrl string = "https://glosbe.com/gapi/translate?from=en&dest=en&format=json&pretty=true&phrase=%s"
const synonymsApiUrl string = "https://glosbe.com/gapi/translate?from=en&dest=en&format=json&pretty=true&phrase=%s"
const antonymsApiUrl string = "http://words.bighugelabs.com/api/2/%s/%s/text"
const partOfSpeechApiUrl string = "http://api.wordnik.com/v4/word.json/%s/definitions?api_key=%s"
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

// Returns true if a is present in list
func stringInSlice(a string, list []string) bool {
  for _, b := range list {
    if b == a {
      return true
    }
  }
  return false
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
  POS  string     `json:"partOfSpeech"` // The part of speech for the word
  ExampleUsage  string   `json:"text"` // An example usage for the word in POS
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

  if v.c.BigHugeLabsApiKey == "" {
    return Word{}, Error("BigHugeLabsApiKey required.")
  }
  if v.c.WordnikApiKey == "" {
    return Word{}, Error("WordnikApiKey required.")
  }

  meanings, err := v.Meanings(w)
  if err != nil {
    return Word{}, err
  }

  synonyms, err := v.Synonyms(w)
  if err != nil {
    return Word{}, err
  }

  antonyms, err := v.Antonyms(w)
  if err != nil {
    return Word{}, err
  }

  pos, err := v.PartOfSpeech(w)
  if err != nil {
    return Word{}, err
  }

  ue, err := v.UsageExample(w)
  if err != nil {
    return Word{}, err
  }

  return Word{
    Word: w,
    Meanings: meanings,
    Synonyms: synonyms,
    Antonyms: antonyms,
    PartOfSpeech: pos,
    UsageExample: ue,
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

  var meanings GlosbeMeanings
  err = json.Unmarshal(glosbe.Tuc[0], &meanings)
  if err != nil {
    return []string{}, err
  }

  var result []string
  for _, gt := range meanings.Things {
    result = append(result, gt.Text)
  }

  return result, nil
}

// Returns a list of strings representing the synonyms of the given word.
func (v Vocabulary) Synonyms(w string) ([]string, error) {
  contents, err := makeReq(fmt.Sprintf(synonymsApiUrl, w))
  if err != nil {
    return []string{}, err
  }

  var glosbe Glosbe
  err = json.Unmarshal(contents, &glosbe)

  if err != nil || glosbe.Result != "ok" {
    return []string{}, err
  }

  var result []string
  for _, tuc_raw := range glosbe.Tuc[1:] {
    var gp GlosbePhrase
    err = json.Unmarshal(tuc_raw, &gp)
    if err != nil {
      return []string{}, err
    }
    result = append(result, gp.Thing.Text)
  }
  return result, nil
}

// Returns a list of strings representing the antonyms of the given word.
func (v Vocabulary) Antonyms(w string) ([]string, error) {
  contents, err := makeReq(fmt.Sprintf(antonymsApiUrl, v.c.BigHugeLabsApiKey, w))
  if err != nil {
    return []string{}, err
  }

  if string(contents) == "" {
    return []string{}, nil
  }

  var result []string
  lines := strings.Split(string(contents), "\n")
  for _, line := range lines {
    b := strings.Split(line, "|")
    if len(b) == 3 && b[1] == "ant" && !stringInSlice(b[2], result) {
      result = append(result, b[2])
    }
  }
  return result, nil
}

// Returns a list of PartOfSpeech structs representing the POS of the given word.
func (v Vocabulary) PartOfSpeech(w string) ([]PartOfSpeech, error) {
  contents, err := makeReq(fmt.Sprintf(partOfSpeechApiUrl, w, v.c.WordnikApiKey))
  if err != nil {
    return []PartOfSpeech{}, err
  }

  var result []PartOfSpeech
  err = json.Unmarshal(contents, &result)
  if err != nil {
    return []PartOfSpeech{}, err
  }
  return result, nil
}

// Returns a list of strings representing usage examples of the given word.
func (v Vocabulary) UsageExample(w string) ([]string, error) {
  contents, err := makeReq(fmt.Sprintf(usageExampleApiUrl, w))
  if err != nil {
    return []string{}, err
  }

  var resp UrbanDictResp
  err = json.Unmarshal(contents, &resp)
  if err != nil {
    return []string{}, err
  }

  var result []string
  for _, thing := range resp.Things {
    if thing.ThumbsUp > 2 * thing.ThumbsDown {
      text := strings.Replace(thing.Example, "\r", " ", -1)
      text = strings.Replace(thing.Example, "\n", " ", -1)
      result = append(result, text)
    }
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
