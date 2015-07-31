
package util

type Command interface {
  Name() string
  Matches() []Regex

  // For documentation purposes
  Description() string
  Format() string
  Examples() []string
  OtherDocs() []HelpTopic

  // Behavior
  Execute(IncomingSlackMessage)

}
