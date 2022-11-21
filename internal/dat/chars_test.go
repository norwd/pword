package dat

import (
	"regexp"
	"strings"
	"testing"
)

func TestAllCharacterClassesAreProvided(t *testing.T) {
	for name, class := range CharacterClasses {
		name := name
		class := class

		t.Run(name, func(t *testing.T) {
			if len(class) == 0 {
				t.Errorf("Empty character class %q", name)
			}
		})
	}
}

func TestAllCharacterClassesAreComplete(t *testing.T) {
	for name, class := range CharacterClasses {
		name := name
		class := string(class)

		t.Run(name, func(t *testing.T) {
			regex, err := regexp.Compile("[[:" + name + ":]]")

			if err != nil {
				t.Fatalf("Could not compile regex: %q", err)
			}

			matches := regex.FindAllString(class, -1)

			if have := strings.Join(matches, ""); class != have {
				t.Errorf("Not all listed characters are matched by class: want %q but have %q", class, have)
			}
		})
	}
}
