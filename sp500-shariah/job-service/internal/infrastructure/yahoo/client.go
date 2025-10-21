package yahoo

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type YahooClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewYahooClient() *YahooClient {
	return &YahooClient{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://query1.finance.yahoo.com/v8/finance/chart",
	}
}

type YahooResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Symbol   string `json:"symbol"`
				Currency string `json:"currency"`
			} `json:"meta"`
			Timestamp  []int64 `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Close []float64 `json:"close"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"chart"`
}

func (c *YahooClient) FetchStockData(symbol string) (*YahooResponse, error) {
	url := fmt.Sprintf("%s/%s?interval=1d&range=1mo", c.baseURL, symbol)

	maxRetries := 5
	for attempt := 1; attempt <= maxRetries; attempt++ {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to request Yahoo API: %w", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			var data YahooResponse
			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				return nil, fmt.Errorf("failed to decode Yahoo response: %w", err)
			}
			return &data, nil
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			delay := time.Duration(rand.Intn(5)+5) * time.Second
			fmt.Printf("⚠️  Too many requests for %s — retrying in %v (attempt %d/%d)\n", symbol, delay, attempt, maxRetries)
			time.Sleep(delay)
			continue
		}

		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Yahoo API error: %s", string(body))
	}

	return nil, fmt.Errorf("max retries exceeded for %s", symbol)
}
