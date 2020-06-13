package elrond

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
)

func Discovery() ([]*Node, error) {
	var nodes []*Node
	matches, err := filepath.Glob("/etc/systemd/system/elrond-node-*.service")

	if err != nil {
		return nil, errors.Wrap(err, "node discovery")
	}

	for _, match := range matches {
		file, err := ioutil.ReadFile(match)

		if err != nil {
			return nil, errors.Wrap(err, "read service node file")
		}

		re := regexp.MustCompile(`(?m)-rest-api-interface ([A-Za-z0-9-.]+):(\d{1,6})`)
		m := re.FindAllStringSubmatch(string(file), -1)

		if len(m) == 0 {
			log.Println("warn: no rest api interface found in " + match + " (no match)")
			continue
		}

		node, err := NewNode(fmt.Sprintf("http://%s:%s", m[0][1], m[0][2]))

		if err != nil {
			return nil, errors.Wrapf(err, "init node %s:%s", m[0][1], m[0][2])
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}
