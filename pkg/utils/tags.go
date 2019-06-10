// Synse CLI
// Copyright (c) 2019 Vapor IO
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package utils

import (
	"fmt"
	"regexp"
	"strings"

	synse "github.com/vapor-ware/synse-server-grpc/go"
)

// NormalizeTags takes a slice specifying tags which may be comma-separated
// and produces a slice of tags where each element is an individual tag.
func NormalizeTags(tags []string) []string {
	var final []string
	for _, tag := range tags {
		final = append(final, strings.Split(tag, ",")...)
	}
	return final
}

// StringToTag converts a string representation of a Tag to its corresponding
// gRPC message.
func StringToTag(s string) (*synse.V3Tag, error) {
	tag := strings.TrimSpace(s)
	if strings.Contains(tag, " ") {
		return nil, fmt.Errorf("tag may not contain spaced: '%s'", s)
	}

	// This is the same check utilized in the SDK. We do not import the SDK
	// to use this utility as it is a large dependency for such small functionality.
	validTag := regexp.MustCompile(`(([^:/]+)/)?(([^:/]+):)?([^:/\s]+$)`)
	matches := validTag.FindStringSubmatch(tag)
	if len(matches) != 6 {
		return nil, fmt.Errorf("invalid tag string (match check): %s", tag)
	}
	namespace := matches[2]
	annotation := matches[4]
	label := matches[5]

	// Make sure that the original tag does not have a namespace delimiter
	// if no namespace was matched. This is indicative of a malformed tag which
	// the regex may not have choked on.
	if strings.Contains(tag, "/") && namespace == "" {
		return nil, fmt.Errorf("invalid tag string (namespace check): %s", tag)
	}

	// Make sure that the original tag does not have an annotation delimiter
	// if no annotation was matched. This is indicative of a malformed tag which
	// the regex may not have choked on.
	if strings.Contains(tag, ":") && annotation == "" {
		return nil, fmt.Errorf("invalid tag string (annotation check): %s", tag)
	}

	return &synse.V3Tag{
		Namespace:  namespace,
		Annotation: annotation,
		Label:      label,
	}, nil
}
