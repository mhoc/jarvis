
package util

import (
  "regexp"
  "strings"
)

type HelpGenerator struct {
  CommandName string
  Description string
  RegexMatches []*regexp.Regexp
  Format string
  Examples []string
  OtherTopics []HelpGeneratorTopic
}

// A topic is any other pieces of documentation we want to include with
// this command.
type HelpGeneratorTopic struct {
  Name string
  Body string
}

// The order in which help text is generated is always:
// CommandName/Description
// Matches On
// Format
// Examples
// Other topics
func (h HelpGenerator) Generate() string {
  h.Description = "  " + strings.Replace(h.Description, "\n", "\n  ", -1)
  help := "```\n"
  help += h.CommandName + "\n"
  help += h.Description + "\n\n"
  help += "matches on\n"
  for _, match := range h.RegexMatches {
    help += "  " + match.String() + "\n"
  }
  help += "\nformat\n  " + h.Format + "\n\n"
  help += "examples\n"
  for _, ex := range h.Examples {
    help += "  " + ex + "\n"
  }
  for _, topic := range h.OtherTopics {
    help += "\n" + topic.Name + "\n"
    help += "  " + strings.Replace(topic.Body, "\n", "\n  ", -1) + "\n"
  }
  help += "```"
  return help
}
