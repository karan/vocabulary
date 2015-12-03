package main

// Simple example usage of
//  github.com/karan/vocabulary

import (
  "fmt"
  "log"

  "github.com/karan/vocabulary"
)


func main() {
  // Create a new vocabulary.Word object, and collects all possible information.
  word, err := vocabulary.Word("vuvuzela")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(word.Meanings)
  fmt.Println(word.Synonyms)
  fmt.Println(word.Antonyms)
  fmt.Println(word.PartOfSpeech)
  fmt.Println(word.UsageExample)

  // Get just the synonyms
  synonyms, err := vocabulary.Synonyms("area")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(synonyms)

  // Can also use:
  //  vocabulary.Meanings(word)
  //  vocabulary.Antonyms(word)
  //  vocabulary.PartOfSpeech(word)
  //  vocabulary.UsageExample(word)

}
