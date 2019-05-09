package utils

import "strings"

// NormalizeTags takes a slice specifying tags which may be comma-separated
// and produces a slice of tags where each element is an individual tag.
func NormalizeTags(tags []string) []string {
	var final []string
	for _, tag := range tags {
		final = append(final, strings.Split(tag, ",")...)
	}
	return final
}
