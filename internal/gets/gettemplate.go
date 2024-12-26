package gets

import (
	"context"
	"github.com/BurntSushi/toml"
	"gitlab.com/osamikoyo/sea/internal/tomltools"
	"io/ioutil"
	"net/http"
	"time"
)

func GetTemplateFromUrl(url string) (tomltools.TEMP, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	resp, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return tomltools.TEMP{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tomltools.TEMP{}, err
	}

	var template tomltools.TEMP

	if _, err = toml.Decode(string(body), &template); err != nil {
		return tomltools.TEMP{}, err
	}

	return template, nil
}
