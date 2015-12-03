package main

// Simple example usage of
//  github.com/karan/vocabulary

import (
  "fmt"
  "log"

  "github.com/karan/vocabulary"
)


func main() {
  // Set the API keys
  // Some functions require API keys. Refer to docs.
  // If API keys are not required, simple set empty strings as config:
  c := &vocabulary.Config{BigHugeLabsApiKey: "", WordnikApiKey:""}

  // Instantiate a Vocabulary object with your config
  v, err := vocabulary.New(c)
  if err != nil {
    log.Fatal(err)
  }

  // Create a new vocabulary.Word object, and collects all possible information.
  word, err := v.Word("vuvuzela")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(word.Meanings)
  fmt.Println(word.Synonyms)
  fmt.Println(word.Antonyms)
  fmt.Println(word.PartOfSpeech)
  fmt.Println(word.UsageExample)

  // Get just the synonyms
  synonyms, err := v.Synonyms("area")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(synonyms)

  // Can also use:
  //  v.Meanings(word)
  //  v.Antonyms(word)
  //  v.PartOfSpeech(word)
  //  v.UsageExample(word)

}
