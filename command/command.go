package command

import (
	"encoding/json"
	"estemplate/file"
	"estemplate/report"
	"fmt"
	"os"
	"time"
)

type RenderTemplateCommand struct {
	TemplateURL string         `json:"templateUrl"`
	Data        map[string]any `json:"data"`
}

func NewRenderTemplateCommand(jsonData []byte) (*RenderTemplateCommand, error) {
	var cmd RenderTemplateCommand
	err := json.Unmarshal(jsonData, &cmd)
	if err != nil {
		return nil, err
	}
	return &cmd, nil
}

func Execute(cmd *RenderTemplateCommand) (string, func(), error) {
	filepath := fmt.Sprintf("downloaded_%d.odt", time.Now().Unix())

	err := file.Download(cmd.TemplateURL, filepath)
	if err != nil {
		return "", func() {}, err
	}

	pdf, err := report.Render(filepath, cmd.Data)
	if err != nil {
		return "", func() {}, err
	}

	return pdf, func() {
		report.CleanUp()
		os.Remove(filepath)
	}, err
}
