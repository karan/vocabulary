# Vocabulary

An English-language dictionary and thesaurus in a Go package.

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
  // Create a new vocabulary.Word object, and collects all information it can.
  word, err := vocabulary.Word("vuvuzela")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(word.Meanings)
  fmt.Println(word.Synonyms)
  fmt.Println(word.Antonyms)
  fmt.Println(word.PartOfSpeech)
  fmt.Println(word.UsageExample)

  // Or I just want the synonyms
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
```

## Tests

```go
$ go test
# TODO: test results here
```

## Bugs

Please use the [issue tracker](https://github.com/karan/vocabulary/issues) to submit any bugs or feature requests.

## License

MIT License © [Karan Goel](https://twitter.com/karangoel)
