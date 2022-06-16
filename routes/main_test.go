package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTickerListHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	GetAllTickers(w, req)
	res := w.Result()
	defer res.Body.Close()

	wantResponseType := "application/json"

	gotResponseType := res.Header.Get("Content-Type")

	if wantResponseType != gotResponseType {
		t.Error("Wrong Response Type")
	}

	type TickerListResponseBody struct {
		Tickers []string
	}

	var bodyBytes []byte
	tlRB := &TickerListResponseBody{}

	res.Body.Read(bodyBytes)

	json.Unmarshal(bodyBytes, tlRB)

	notWantedTicketCount := 0
	gotTickerCount := len(tlRB.Tickers)

	if notWantedTicketCount == gotTickerCount {
		t.Errorf("ticker Count Empty")
	}

	// {
	//     "tickers" : ["AAPL", "MOON"]
	// }

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	//     t.Errorf("expected error to be nil got %v", err)
	// }
	// if string(data) != "sdgadgasdgasdgasdg" {
	//     t.Errorf("expected ABC got %v", string(data))
	// }
}