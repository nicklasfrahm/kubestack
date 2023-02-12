package nxos

import (
	"bufio"
	"regexp"
	"strings"
)

var (
	regexCommentOrVersionInfo = regexp.MustCompile(`(!|version).*\n`)
)

// Section is a section of an NX-OS configuration snippet.
// TODO: This is a very naive implementation. It does not
// support nested sections.
type Section struct {
	// Header the name of the section.
	Header string
	// Lines are the configuration lines of the section.
	Lines []string
}

// Config is the parsed representation of an NX-OS configuration snippet.
type Config struct {
	// Sections are the sections of the configuration snippet.
	Sections []Section
}

// Parse parses an NX-OS configuration snippet.
func Parse(raw string) (*Config, error) {
	// Strip comments, version info and excessive lines.
	raw = strings.TrimSpace(regexCommentOrVersionInfo.ReplaceAllString(raw, ""))

	// Use a scanner to iterate over the lines.
	scanner := bufio.NewScanner(strings.NewReader(raw))

	config := &Config{
		Sections: make([]Section, 0),
	}
	var section *Section
	for scanner.Scan() {
		line := scanner.Text()

		// This marks the start of a new section.
		if !strings.HasPrefix(line, " ") {
			// Create a new section and update the pointer to the current section.
			config.Sections = append(config.Sections, Section{
				Header: line,
			})
			section = &config.Sections[len(config.Sections)-1]
			continue
		}

		// Remove leading whitespace and skip empty lines.
		line = strings.TrimSpace(line)
		if line != "" {
			section.Lines = append(section.Lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}
