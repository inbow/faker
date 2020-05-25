package server

import (
	"net/http"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type (
	Data struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		DisplayURL  string `json:"displayURL"`
		Redirect    string `json:"redirect"`
		CPC         string `json:"cpc"`
		Image       string `json:"image"`
	}

	VolvoResponse struct {
		Data []Data `json:"data,omitempty"`

		Errors []string `json:"errors,omitempty"`
	}
)

func (s *Server) Volvo(w http.ResponseWriter, r *http.Request) {
	response := s.NewResponse()
	price, delay, skip := s.RequestValues(r.URL.Query())

	defer func() {
		time.Sleep(time.Duration(delay) * time.Millisecond)

		w.Header().Set("Content-Type", "application/javascript")
		w.WriteHeader(response.StatusCode)

		if response.StatusCode != http.StatusNoContent && len(response.Body) > 0 {
			_, _ = w.Write(response.Body)
		}
	}()

	if skip {
		response.StatusCode = http.StatusNoContent
		return
	}

	handlerResponse := &VolvoResponse{
		Data: []Data{
			{
				Title:       "Yukon Gold Casino",
				Description: "125 Chances to Win Jackpot. One spin can change your life!",
				DisplayURL:  "www.yukongoldcasino.eu",
				Redirect:    "https://rsalchk.com/performance/bdv_rd.dbm?enparms2=qwerty",
				CPC:         strconv.FormatFloat(price, 'g', 6, 64),
				Image:       "//bvadimgs.scdn7.secure.raxcdn.com/bidvertiser/images/ad_image/9/96226.png",
			},
		},
	}

	data, err := jsoniter.Marshal(handlerResponse)
	if err != nil {
		response.StatusCode = http.StatusBadGateway
		response.Body = []byte(err.Error())
	}

	response.StatusCode = http.StatusOK
	response.Body = data
}
