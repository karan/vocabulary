# Vocabulary

[![](http://i.imgur.com/RJSJHlA.png)](https://xkcd.com/1443/)

An English-language dictionary and thesaurus in a Go package by [Karan Goel](https://twitter.com/karangoel).

[![Build Status](https://drone.io/github.com/karan/vocabulary/status.png)](https://drone.io/github.com/karan/vocabulary/latest)

Golang port of [the Python counterpart](https://github.com/prodicus/vocabulary/).

For a given word, `Vocabulary` will give you:

* **Meaning**
* **Synonyms**
* **Antonyms**
* **Part of speech**: whether the word is a noun, interjection or an adverb et el
* **Usage example**: a quick example on how to use the word in a sentence

## Features

* Written in idiomatic Go
* No external dependencies
* So easy, a five-year-old can use it
* Works on Mac, Linux and Windows.

## Installation

```bash
$ go get -u github.com/karan/vocabulary
```

## Usage

### Get API keys

1. [Big Huge Thesaurus](http://words.bighugelabs.com/getkey.php)
  * Required for `Antonyms`
2. [Wordnik](http://developer.wordnik.com/)
  * Required for `PartOfSpeech`

Calling `vocabulary.Word()` with any word as a string will return a `vocabulary.Word` type object that has all possible information about it.

Or if you just want selective information, you can call individual functions passing in a word (`vocabulary.Meanings("hallucination")`).

```go
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
  //  v.UsageExample(word)

}

```

## Tests

Create `examples/api_keys.go` with your API keys:

```go
package main

const (
  BigHugeLabsApiKey = "xxxx"
  WordnikApiKey     = "xxxx"
)

```

Then, to run the tests, use this command:

```bash
$ go test
PASS
```

## Bugs

Please use the [issue tracker](https://github.com/karan/vocabulary/issues) to submit any bugs or feature requests.

## License

MIT License © [Karan Goel](https://twitter.com/karangoel)
