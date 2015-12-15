// Contains color code information for printing colorized logs

package log

import (
	"fmt"
)

const (
	FORMAT      string = "\033[%vm%v\033[0;00m"
	DEFAULT     string = "0;00"
	BLACK       string = "0;30"
	RED         string = "0;31"
	GREEN       string = "0;32"
	BROWN       string = "0;33"
	BLUE        string = "0;34"
	PURPLE      string = "0;35"
	CYAN        string = "0;36"
	BOLD_GRAY   string = "1;30"
	BOLD_RED    string = "1;31"
	BOLD_GREEN  string = "1;32"
	YELLOW      string = "1;33"
	BOLD_BLUE   string = "1;34"
	BOLD_PURPLE string = "1;35"
	BOLD_CYAN   string = "1;36"
)

func FormatColor(msg string, color string) string {
	return fmt.Sprintf(FORMAT, color, msg)
}
