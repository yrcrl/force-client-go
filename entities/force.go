package force

type ID string
type Text string
type Phone string
type Date string     // yyyy-MM-dd
type DateTime string // yyyy-MM-ddTHH:mm:ss.000+0000
type TextArea string
type LongTextArea string
type Checkbox bool
type Email string
type Fax string
type Picklist string
type Number int64
type AutoIncrement string
type URL string
type Decimal float64

//type Address string  // Address Compound Fields
type Location string // Geolocation Compound Field
type Currency string

type BaseRecordSet struct {
	TotalSize int    `json:"totalSize"`
	Done      bool   `json:"done"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type Attr struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
