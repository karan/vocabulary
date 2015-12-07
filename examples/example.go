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
  c := &vocabulary.Config{BigHugeLabsApiKey: BigHugeLabsApiKey, WordnikApiKey: WordnikApiKey}

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

  fmt.Printf("word.Word = %s \n", word.Word)
  fmt.Printf("word.Meanings = %s \n", word.Meanings)
  fmt.Printf("word.Synonyms = %s \n", word.Synonyms)
  fmt.Printf("word.Antonyms = %s \n", word.Antonyms)
  fmt.Printf("word.PartOfSpeech = %s \n", word.PartOfSpeech)
  fmt.Printf("word.UsageExample = %s \n", word.UsageExample)

  // Get just the synonyms
  // synonyms, err := v.Synonyms("area")
  // if err != nil {
  //   log.Fatal(err)
  // }
  // for _, s := range synonyms {
  //   fmt.Println(s)
  // }
  //

  // Get just the antonyms
  // ants, err := v.Antonyms("love")
  // if err != nil {
  //   log.Fatal(err)
  // }
  // for _, a := range ants {
  //   fmt.Println(a)
  // }

  // Get just the part of speech
  // pos, err := v.PartOfSpeech("love")
  // if err != nil {
  //   log.Fatal(err)
  // }
  // for _, a := range pos {
  //   fmt.Println(a)
  // }

  // Can also use:
  //  v.Meanings(word)
  //  v.Antonyms(word)
  //  v.PartOfSpeech(word)
  //  v.UsageExample(word)

}
