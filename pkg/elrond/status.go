package elrond

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (n *Node) Status() (*StatusInfo, error) {
	resp, err := http.Get(n.url + "/node/status")

	if err != nil {
		return nil, errors.Wrap(err, "http get /node/status")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var statusInfo StatusInfo

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return nil, errors.Wrapf(err, "bad http code (%s), unable to read body", resp.Status)
		}

		return nil, errors.Wrapf(err, "bad http code (%s): %s", resp.Status, string(body))
	}

	j := json.NewDecoder(resp.Body)

	if err := j.Decode(&statusInfo); err != nil {
		return nil, errors.Wrap(err, "parse json status")
	}

	return &statusInfo, nil
}
