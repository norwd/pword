package dat

import (
	"regexp"
	"strings"
	"testing"
)

func TestAllCharacterClassesAreComplete(t *testing.T) {
	for name, class := range CharacterClasses {
		name := name
		class := class

		t.Run(name, func(t *testing.T) {
			regex, err := regexp.Compile("[[:" + name + ":]]")

			if err != nil {
				t.Fatalf("Could not compile regex: %q", err)
			}

			matches := regex.FindAllString(string(class), -1)

			if want, have := string(class), strings.Join(matches, ""); want != have {
				t.Errorf("Not all listed characters are matched by class: want %q but have %q", want, have)
			}
		})
	}
}
