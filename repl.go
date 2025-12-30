// by convention, we name our package the same as the directory
package msf_pokedex

import "strings"

//"fmt"
//"strings"
//"io"
//"net/http"
//"bytes"
//"encoding/json"

func cleanInput(text string) []string {
	var sliceStrings []string

	if len(text) > 0 {
		//cleanedString := strings.TrimSpace(text)
		//cleanedString = strings.ToLower(cleanedString)
		//itemsToAdd := strings.Split(cleanedString, " ")
		//sliceStrings = append(sliceStrings, itemsToAdd...)
		sliceStrings = strings.Fields(text)
	}

	return sliceStrings
}
