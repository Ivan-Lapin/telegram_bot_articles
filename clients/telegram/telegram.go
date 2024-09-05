package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"example.com/m/lib/errr"
)

const (
	getUpdatesMethod  = "getUpdates"
	SendMessageMethod = "sendMessage"
)

type Clients struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) Clients {
	return Clients{
		host:     host,
		basePath: NewBasePath(token),
		client:   http.Client{},
	}
}

func NewBasePath(token string) string {
	return "bot" + token
}

func (c *Clients) Update(offset int, limit int) (updates []Update, err error) {
	defer func() { err = errr.WrapIfErr("Can't to do request : ", err) }()

	q := url.Values{}
	q.Add(string(offset), strconv.Itoa(offset))
	q.Add(string(limit), strconv.Itoa(limit))

	// do request < - doRequest
	data, err := c.DoRequest("getUpdatesMethod", q)
	if err != nil {
		return nil, err
	}

	var res UpdateResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Clients) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chatID", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.DoRequest(SendMessageMethod, q)
	if err != nil {
		return errr.Wrap("Can't send message : ", err)
	}

	return nil
}

func (c *Clients) DoRequest(method string, query url.Values) (data []byte, err error) {

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
