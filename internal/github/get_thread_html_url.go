package github

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
	"time"

	"github.com/umekikazuya/gh-notify/internal/app"
)

func GetTreadHTMLURL(rawURL string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(
		ctx,
		"gh",
		"api",
		rawURL,
	)
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	var data targetResource
	err = json.Unmarshal(stdOut.Bytes(), &data)
	if err != nil {
		return "", err
	}
	return data.URL, nil
}

type targetResource struct {
	URL string `json:"html_url"`
}

var _ app.GetTreadHTMLURLFn = GetTreadHTMLURL
