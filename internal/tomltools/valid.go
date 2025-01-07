package tomltools

import "github.com/osamikoyo/sea/internal/tomltools/valider"

func isValid(text string) int {
	return valider.IsValid(text)
}
