package service

import (
	"client-ccs/app/client"
	"fmt"
	"net/http"
	"time"
)

type requestCase int

const (
	headerAuthorization = "Authorization"

	getCurrenciesUri = "currencies"
	getRate          = "rate/from/USD/to/EUR"
	getRateByDay     = "rate/from/USD/to/EUR/date/2024-05-30"
	getRateByMonth   = "rate/from/USD/to/EUR/year/2024/month/5"
	postExchange     = "exchange"
)

const (
	caseCurrencies requestCase = iota
	caseRate
	caseByMonth
	caseByDay
	caseExchange
)

type LoadService struct {
	rps                 int
	client              *client.HttpClient
	baseUri             string
	authorizationHeader map[string]string
	headers             map[string]string
}

func NewLoadService(
	rps int,
	client *client.HttpClient,
	baseUri string,
	clientToken string,
) *LoadService {
	return &LoadService{
		rps:                 rps,
		client:              client,
		baseUri:             baseUri,
		authorizationHeader: createAuthorizationHeader(clientToken),
	}
}

func (ls *LoadService) StartLoadTest() {
	ls.headers = ls.client.BuildHeaders(ls.authorizationHeader)
	i := 1

	for {
		for j := 1; j <= ls.rps; j++ {
			switch reqCase := getCase(); reqCase {
			case caseCurrencies:
				go ls.requestCurrencies(i)
			case caseRate:
				go ls.requestRate(i)
			case caseByMonth:
				go ls.requestRateByMonth(i)
			case caseByDay:
				go ls.requestRateByDay(i)
			case caseExchange:
				go ls.requestExchange(i)
			}
		}
		i++
		time.Sleep(1 * time.Second)
	}

}

func getCase() requestCase {
	return caseCurrencies
	/*
		r := rand.IntN(13)

		if r == 1 {
			return caseCurrencies
		}

		if r > 1 && r <= 4 {
			return caseRate
		}

		if r > 4 && r <= 7 {
			return caseByMonth
		}

		if r > 7 && r <= 10 {
			return caseByDay
		}

		return caseExchange
	*/
}

func (ls *LoadService) requestCurrencies(i int) {
	startDoRequest := time.Now()

	code, resp, err := ls.client.Get(
		ls.baseUri+getCurrenciesUri,
		ls.headers,
	)

	fmt.Printf("get currency #%d took %v\n", i, time.Now().Sub(startDoRequest))

	parseResponse(code, resp, err)
}

func (ls *LoadService) requestRate(i int) {
	startDoRequest := time.Now()

	code, resp, err := ls.client.Get(
		ls.baseUri+getRate,
		ls.headers,
	)

	fmt.Printf("get rate #%d took %v\n", i, time.Now().Sub(startDoRequest))

	parseResponse(code, resp, err)
}

func (ls *LoadService) requestRateByDay(i int) {
	startDoRequest := time.Now()

	code, resp, err := ls.client.Get(
		ls.baseUri+getRateByDay,
		ls.headers,
	)

	fmt.Printf("get rate by day #%d took %v\n", i, time.Now().Sub(startDoRequest))

	parseResponse(code, resp, err)
}

func (ls *LoadService) requestRateByMonth(i int) {
	startDoRequest := time.Now()

	code, resp, err := ls.client.Get(
		ls.baseUri+getRateByMonth,
		ls.headers,
	)

	fmt.Printf("get rate by month #%d took %v\n", i, time.Now().Sub(startDoRequest))

	parseResponse(code, resp, err)
}

func (ls *LoadService) requestExchange(i int) {
	startDoRequest := time.Now()

	code, resp, err := ls.client.Post(
		ls.baseUri+postExchange,
		ls.headers,
		[]byte(`[{"from": "BNB","amount": 5,"to": "USD"},{"from": "BTC","amount": 0.56,"to": "EUR","rate": 15800}]`),
	)

	fmt.Printf("post exchange #%d took %v\n", i, time.Now().Sub(startDoRequest))

	parseResponse(code, resp, err)
}

func createAuthorizationHeader(clientToken string) map[string]string {
	authorizationHeader := make(map[string]string)
	authorizationHeader[headerAuthorization] = clientToken

	return authorizationHeader
}

func parseResponse(code int, resp []byte, err error) {
	if err != nil {
		println(err.Error())
		return
	}

	if code != http.StatusOK {
		println(string(resp))
		return
	}
}
