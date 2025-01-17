package gopwrstat

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"strings"
)

type Pwrstat struct {
	Status map[string]string `json:"status"`
}

func construct(content string) *Pwrstat {
	s := &Pwrstat{}
	s.Status = map[string]string{}

	lines := strings.Split(content, "\n")
	var statusArr []string
	for _, line := range lines {
		if len(line) > 0 {
			line = strings.Trim(line, "	")
			line = strings.Replace(line, ". ", ";", -1)
			line = strings.Replace(line, ".", "", -1)
			newline := strings.Split(line, ";")
			if len(newline) > 1 {
				statusArr = append(statusArr, newline...)
			}
		}
	}

	for i := 0; i < len(statusArr); i += 2 {
		s.Status[statusArr[i]] = statusArr[i+1]
	}

	return s
}

// A successful call returns err == nil.
func NewFromSystem() (*Pwrstat, error) {
	out, err := exec.Command("pwrstat", "-status").Output()
	if err != nil {
		return &Pwrstat{}, errors.New("pwrstat missing")
	}

	s := construct(string(out))

	return s, nil
}

func NewFromFile(path string) (*Pwrstat, error) {
	out, err := os.ReadFile(path)
	if err != nil {
		return &Pwrstat{}, err
	}

	s := construct(string(out))
	return s, nil
}

func (s *Pwrstat) JSON() string {
	out, _ := json.Marshal(s)

	return string(out)
}

func (s *Pwrstat) String() string {
	return s.JSON()
}
