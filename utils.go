package dribbble

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// writeIfNotEmpty writes formatted key-value pairs to a StringBuilder if the value is not empty.
func writeIfNotEmpty(sb *strings.Builder, key, value string) {
	if value != "" {
		sb.WriteString(fmt.Sprintf("%s: %s\n", color.HiBlackString(key), value))
	}
}

func writeTitleString(sb *strings.Builder, key string) {
	sb.WriteString(fmt.Sprintf("--%s--\n", color.HiBlackString(key)))
}

// formatTags formats a slice of tags into a comma-separated string.
func formatTags(tags []string) string {
	if len(tags) > 0 {
		return strings.Join(tags, ", ")
	}
	return ""
}

// ------------------------------------------------------------------------