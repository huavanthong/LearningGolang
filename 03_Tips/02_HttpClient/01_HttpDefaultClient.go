package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	hclog "github.com/hashicorp/go-hclog"
)

type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

// for multiple Cube values
type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

// define structure adapt value in vxml
type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}

// define internal value to getRates
func (e *ExchangeRates) getRates() error {
	// using http to get the exchange rate of currency from european central bank
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return nil
	}

	// handling error code from response using http package
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected error code 200 got %d", resp.StatusCode)
	}
	// any abnormal case, it need to be closed
	defer resp.Body.Close()

	// initialize a container to collect data
	md := &Cubes{}

	// create a decoder, and decode xml to get a raw data
	xml.NewDecoder(resp.Body).Decode(&md)

	// loop to parse all values to from the container of data
	// note: for each currency in CubeData
	for _, c := range md.CubeData {
		// parse rate value with 64 bit, from c.Rate
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}

		// assign data
		e.rates[c.Currency] = r
	}

	// because e.rates lack of the euro currency
	// and we set it to base currency.
	e.rates["EUR"] = 1

	return nil
}

func main() {

}
