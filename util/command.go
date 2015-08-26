
package util

type Command interface {
  Name() string
  Matches() []Regex

  // For documentation purposes
  Description() string
  Examples() []string
  OtherDocs() []HelpTopic

  // Behavior
  Execute(IncomingSlackMessage)

}
