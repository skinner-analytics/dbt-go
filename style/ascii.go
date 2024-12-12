package style

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

// png to ascii
var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(str string) string {
	return ansiRegex.ReplaceAllString(str, "")
}

func CenterText(text string, width int) string {
	stripped := stripANSI(text)
	if len(stripped) >= width {
		return text
	}
	spaces := width - len(stripped)
	left := spaces / 2
	return strings.Repeat(" ", left) + text
}

func GetASCIIArt() string {
	filePath := ".img/logo.png"
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{26, 13}      // Adjust as needed
	flags.FontColor = [3]int{255, 165, 0} // RGB for orange

	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		return fmt.Sprintf("Error generating ASCII art: %v", err)
	}
	return asciiArt
}
