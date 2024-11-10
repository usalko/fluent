// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"strings"
	"unicode"
)

func ToSnakeCaseSpecialFromCamel(input_string string) string {
	if len(input_string) == 0 {
		return ""
	}
	var result strings.Builder
	for i, r := range input_string {
		// Special case - a first letter
		if i == 0 {
			result.WriteRune(unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			result.WriteRune('_')
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
