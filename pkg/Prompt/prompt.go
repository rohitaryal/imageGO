// Package prompt holds prompt related things
package prompt

import (
	"fmt"

	Types "github.com/rohitaryal/imageGO/pkg/Types"
)

type Prompt struct {
	Seed            int
	Prompt          string
	NumberOfImages  int
	AspectRatio     Types.AspectRatio
	GenerationModel Types.Model
}

func (p *Prompt) String() string {
	return fmt.Sprintf(`{
            "userInput": {
                "candidatesCount": %d,
                "prompts": ["%s"],
                "seed": %d
            },
            "clientContext": {
                "sessionId": ";1757113025397",
                "tool": "IMAGE_FX"
            },
            "modelInput": {
                "modelNameType": "%s"
            },
            "aspectRatio": "%s"
        }`, p.NumberOfImages, p.Prompt, p.Seed, p.GenerationModel, p.AspectRatio)
}
