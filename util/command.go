// Struct definitions for commands, subcommands, and documentation structs

package util

// I define a command as being a set of potentially multiple divergent functionalities
// which can adequately share a single piece of documentation.
type Command interface {
  Name() string
  Description() string
  Examples() []string
  OtherDocs() []HelpTopic
  SubCommands() []SubCommand
}

// Commands are formed by combining one or more subcommands, each of which
// defines a regex sequence and an execution function through which the
// incoming message and the matching regex sequence is passed.
type SubCommand struct {
  Pattern Regex
  Exec func(IncomingSlackMessage, Regex)
}

func NewSubCommand(pattern string, do func(IncomingSlackMessage, Regex)) SubCommand {
  return SubCommand{
    Pattern: NewRegex(pattern),
    Exec: do,
  }
}

// A topic is any other pieces of documentation we want to include with
// this command.
type HelpTopic struct {
  Name string
  Body string
}
