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

		if (*metrics)["roi5Y"] != nil {
			metricData.ROI5Y = (*metrics)["roi5Y"].(float64)
		}
		if (*metrics)["revenueGrowth5Y"] != nil {
			metricData.RevenueGrowth5Y = (*metrics)["revenueGrowth5Y"].(float64)
		}
		if (*metrics)["epsGrowth5Y"] != nil {
			metricData.EPSGrowth5Y = (*metrics)["epsGrowth5Y"].(float64)
		}
		// Not being rendered, load it anyway
		if (*metrics)["peNormalizedAnnual"] != nil {
			metricData.PENormalizedAnnual = (*metrics)["peNormalizedAnnual"].(float64)
		}
		// Not being rendered, load it anyway
		if (*metrics)["peExclExtraTTM"] != nil {
			metricData.PEExclExtraTTM = (*metrics)["peExclExtraTTM"].(float64)
		}
		if (*metrics)["bookValueShareGrowth5Y"] != nil {
			metricData.BookValueGrowth5Y = (*metrics)["bookValueShareGrowth5Y"].(float64)
		}
		if (*metrics)["revenueShareGrowth5Y"] != nil {
			metricData.RevenueShareGrowth5Y = (*metrics)["revenueShareGrowth5Y"].(float64)
		}
		if (*metrics)["longTermDebt/equityAnnual"] != nil {
			metricData.LongTermDebtPerequityAnnual = (*metrics)["longTermDebt/equityAnnual"].(float64)
		}
		if (*metrics)["totalDebt/totalEquityAnnual"] != nil {
			metricData.TotalDebtPerTotalEquityAnnual = (*metrics)["totalDebt/totalEquityAnnual"].(float64)

		}
		if (*metrics)["focfCagr5Y"] != nil {
			metricData.FOCFCagr5Y = (*metrics)["focfCagr5Y"].(float64)

		}

		return metricData, nil
	}
}
