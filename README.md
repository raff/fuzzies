# fuzzies
A package for "fuzzy" string operations (string compare, index, contains)

fuzzies uses the "Levenshtein distance" between strings (see https://en.wikipedia.org/wiki/Levenshtein_distance)
to implement string comparison operations that allow for "mispelling".

Examples:

    package main

    import (
            "fmt"

            "github.com/raff/fuzzies"
    )

    func main() {
            fmt.Println(fuzzies.Contains("the quirk brown fox", "quick"))
            // true

            fmt.Println(fuzzies.Index("the quirk brown fox", "quick"))
            // 4

            fmt.Println(fuzzies.Compare("hello", "hullo"))
            // true

            fmt.Println(fuzzies.Compare("hello", "hull"))
            // true

            fmt.Println(fuzzies.Compare("hello", "halla"))
            // false

            fmt.Println(fuzzies.Compare("hello", "haloa"))
            // false
    }

