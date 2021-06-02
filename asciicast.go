package asciicast

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type Duration float64
type Env struct {
	Term  string `json:"TERM"`
	Shell string `json:"SHELL"`
}

type Asciicast struct {
	Version   int     `json:"version"`
	Width     int     `json:"width"`
	Height    int     `json:"height"`
	Timestamp int64   `json:"timestamp"`
	Title     string  `json:"title"`
	Stdout    []Frame `json:"-"`
	Env       *Env    `json:"env"`
}

func (a *Asciicast) Save(path string) error {
	by, err := json.Marshal(a)
	buf := bytes.NewBuffer(by)
	if err != nil {
		return err
	}
	buf.WriteString("\n")
	for _, f := range a.Stdout {
		buf.WriteString(f.String())
		buf.WriteString("\n")
	}
	err = ioutil.WriteFile(path, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func NewAsciicast(width, height int, timestamp int64, title string,
	frames []Frame, env map[string]string) *Asciicast {
	return &Asciicast{
		Version:   2,
		Width:     width,
		Height:    height,
		Timestamp: timestamp,
		Title:     title,
		Env:       &Env{Term: env["TERM"], Shell: env["SHELL"]},
		Stdout:    frames,
	}
}
