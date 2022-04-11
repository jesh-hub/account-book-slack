package abs

type Payment struct {
	Date               string `json:"date"`
	Name               string `json:"name"`
	Category           string `json:"category"`
	Method             string `json:"method"`
	Price              int    `json:"price"`
	MonthlyInstallment int    `json:"monthlyInstallment"`
}
