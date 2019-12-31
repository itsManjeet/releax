package confd

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

const ConfigDir = "/etc/conf.d"

type Config struct {
	Variable string
	Value    string
}

type Confd struct {
	File    string
	Configs []Config
}

// Load configurations from file = file address
func Load(file string) (*Confd, error) {

	filedata, err := ioutil.ReadFile(ConfigFor(file))
	if err != nil {
		return nil, err
	}

	var confd Confd
	confd.File = file
	confd.Configs = make([]Config, 0)

	for _, c := range strings.Split(string(filedata[:]), "\n") {
		c = strings.TrimSpace(c)
		if c[0] == '#' {
			continue
		} else if strings.Contains(c, "=") {
			linedata := strings.Split(c, "=")

			var lcf Config
			lcf.Variable = linedata[0]
			lcf.Value = strings.ReplaceAll(linedata[1], "\"", "")
			confd.Configs = append(confd.Configs, lcf)
		}
	}

	return &confd, nil
}

// Get value for 'v' variable
func (cd Confd) Get(v string) (string, error) {
	for _, a := range cd.Configs {
		if a.Variable == v {
			return a.Value, nil
		}
	}
	return "", fmt.Errorf("no variable found")
}

// Set value 'val' for variable 'v'
func (cd *Confd) Set(v, val string) error {
	for i, a := range cd.Configs {
		if a.Variable == v {
			cd.Configs[i].Value = val
			return nil
		}
	}
	return fmt.Errorf("no variable found")
}

// Append variable 'v' of value 'val'
func (cd *Confd) Append(v, val string) error {
	cd.Configs = append(cd.Configs, Config{
		Variable: v,
		Value:    val,
	})

	return nil
}

// Delete variable from configuration
func (cd *Confd) Delete(v string) error {
	for i, a := range cd.Configs {
		if a.Variable == v {
			cd.Configs = append(cd.Configs[:i], cd.Configs[i+1:]...)
		}
	}

	return nil
}

// Save Configuration file
func (cd *Confd) Save() error {
	d := make([]string, 0)

	for _, c := range cd.Configs {
		d = append(d, fmt.Sprintf("%s=\"%s\"", c.Variable, c.Value))
	}

	return ioutil.WriteFile(ConfigFor(cd.File),
		[]byte(strings.Join(d, "\n")), 0644,
	)
}

func ConfigFor(file string) string {
	return path.Join(ConfigDir, file)
}
