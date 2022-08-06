package quote

import (
	"context"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

type ValueData struct {
	ROI5Y                         float64
	RevenueGrowth5Y               float64
	EPSGrowth5Y                   float64
	PENormalizedAnnual            float64
	PEExclExtraTTM                float64
	BookValueGrowth5Y             float64
	RevenueShareGrowth5Y          float64
	LongTermDebtPerequityAnnual   float64
	TotalDebtPerTotalEquityAnnual float64
	FOCFCagr5Y                    float64
}

func GetValueFunc(
	finnhubClient *finnhub.DefaultApiService,
) func(symbol string) (*ValueData, error) {
	return func(symbol string) (*ValueData, error) {
		basicFinancials, _, err := finnhubClient.CompanyBasicFinancials(context.Background()).Symbol(symbol).Metric("all").Execute()
		if err != nil {
			return nil, ErrCompanyBasicFinancialsNotFound
		}

		metrics, ok := basicFinancials.GetMetricOk()
		if !ok {
			return nil, ErrMetricsNotFound
		}

		metricData := &ValueData{}

		metricData.ROI5Y = (*metrics)["roi5Y"].(float64)
		metricData.RevenueGrowth5Y = (*metrics)["revenueGrowth5Y"].(float64)
		metricData.EPSGrowth5Y = (*metrics)["epsGrowth5Y"].(float64)
		metricData.PENormalizedAnnual = (*metrics)["peNormalizedAnnual"].(float64)
		metricData.PEExclExtraTTM = (*metrics)["peExclExtraTTM"].(float64)
		metricData.BookValueGrowth5Y = (*metrics)["bookValueShareGrowth5Y"].(float64)
		metricData.RevenueShareGrowth5Y = (*metrics)["revenueShareGrowth5Y"].(float64)
		metricData.LongTermDebtPerequityAnnual = (*metrics)["longTermDebt/equityAnnual"].(float64)
		metricData.TotalDebtPerTotalEquityAnnual = (*metrics)["totalDebt/totalEquityAnnual"].(float64)
		metricData.FOCFCagr5Y = (*metrics)["focfCagr5Y"].(float64)

		return metricData, nil
	}
}
