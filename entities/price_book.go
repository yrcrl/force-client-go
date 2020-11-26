package force

const PriceBookObjectName = "Pricebook2"

type PriceBook struct {
	Active              bool   `json:"IsActive,omitempty"`
	IsStandardPriceBook bool   `json:"IsStandard,omitempty"`
	Name                string `json:"Name,omitempty"`
}
