package kamel_snake

import (
	"strings"
	"unicode"
)

func CamelToSnake(s string) string {
	input_runes := []rune(s);
	input_runes[0] = unicode.ToLower(input_runes[0]);
	sb := strings.Builder{};
	for _, c := range input_runes {
		if unicode.IsUpper(c) {
			sb.WriteRune('_');
		}
		sb.WriteRune(c);
	}
	return strings.ToLower(sb.String());
}

func SnakeToCamel(s string) string {
	words := strings.Split(s, "_");
	fixed := make([]string, len(words));
	for i, w := range words {
		fixed[i] = strings.Title(w);
	}
	return strings.Join(fixed, "");
}
