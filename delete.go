package cloud

import (
	"fmt"
	"io"
	"net/http"
)

type DeleteData struct {
	Url           string
	Params        map[string]string
	Headers       map[string]string
	Authorization string
	ContentType   string
	Accept        string
}

func Delete(data GetData) ([]byte, int, error) {

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", data.Url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	if data.Authorization != "" {
		req.Header.Add("Authorization", data.Authorization)
	}

	if data.ContentType != "" {
		req.Header.Add("Content-Type", data.ContentType)
	}

	if data.Accept != "" {
		req.Header.Add("Accept", data.Accept)
	}

	for key, element := range data.Headers {
		req.Header.Add(key, element)
	}

	q := req.URL.Query()
	for key, element := range data.Params {
		q.Add(key, element)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, 0, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	return bodyBytes, resp.StatusCode, err
}
