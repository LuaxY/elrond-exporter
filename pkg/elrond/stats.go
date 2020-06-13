package elrond

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (n *Node) Statistics() (*StatisticsInfo, error) {
	resp, err := http.Get(n.url + "/node/statistics")

	if err != nil {
		return nil, errors.Wrap(err, "http get /node/statistics")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var statisticsInfo StatisticsInfo

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return nil, errors.Wrapf(err, "bad http code (%s), unable to read body", resp.Status)
		}

		return nil, errors.Wrapf(err, "bad http code (%s): %s", resp.Status, string(body))
	}

	j := json.NewDecoder(resp.Body)

	if err := j.Decode(&statisticsInfo); err != nil {
		return nil, errors.Wrap(err, "parse json status")
	}

	return &statisticsInfo, nil
}
