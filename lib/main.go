package log2team

import (
	"bytes"
	"encoding/json"
	"html"
	"io"
	"io/ioutil"
	"net/http"
	//	"strings"
)

type PostDataT struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func Send(url, title string, reader io.Reader) error {
	textData, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	text := html.UnescapeString(string(textData))
	// text = strings.Replace(text, "\n", "<br />", -1)
	bin, err := json.Marshal(&PostDataT{
		Title: title,
		Text:  "<pre>\n" + string(text) + "\n</pre>",
	})
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(bin))
	if err != nil {
		return err
	}
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	return nil
}
