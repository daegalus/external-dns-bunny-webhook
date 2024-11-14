package bunny

type RecordType int

const (
	RecordTypeA RecordType = iota
	RecordTypeAAAA
	RecordTypeCNAME
	RecordTypeTXT
	RecordTypeMX
	RecordTypeRDR
	RecordTypeFlatten
	RecordTypePZ
	RecordTypeSRV
	RecordTypeCAA
	RecordTypePTR
	RecordTypeSCR
	RecordTypeNS
)

func (r RecordType) String() string {
	if r < RecordTypeA || r > RecordTypeNS {
		return "?"
	}

	return [...]string{"A", "AAAA", "CNAME", "TXT", "MX", "RDR", "FLATTEN", "PZ", "SRV", "CAA", "PTR", "SCR", "NS"}[r]
}

func RecordTypeFromString(s string) RecordType {
	switch s {
	case "A":
		return RecordTypeA
	case "AAAA":
		return RecordTypeAAAA
	case "CNAME":
		return RecordTypeCNAME
	case "TXT":
		return RecordTypeTXT
	case "MX":
		return RecordTypeMX
	case "RDR":
		return RecordTypeRDR
	case "FLATTEN":
		return RecordTypeFlatten
	case "PZ":
		return RecordTypePZ
	case "SRV":
		return RecordTypeSRV
	case "CAA":
		return RecordTypeCAA
	case "PTR":
		return RecordTypePTR
	case "SCR":
		return RecordTypeSCR
	case "NS":
		return RecordTypeNS
	}

	return RecordType(-1)
}

type Record struct {
	ID                    int64      `json:"Id"`
	Type                  RecordType `json:"Type"`
	TTLSeconds            int        `json:"Ttl"`
	Value                 string     `json:"Value"`
	Name                  string     `json:"Name"`
	Weight                int        `json:"Weight"`
	Priority              int        `json:"Priority"`
	Port                  int        `json:"Port"`
	Flags                 int        `json:"Flags"`
	Tag                   string     `json:"Tag"`
	Accelerated           bool       `json:"Accelerated"`
	AcceleratedPullZoneID int64      `json:"AcceleratedPullZoneId"`
	LinkName              string     `json:"LinkName"`
	Disabled              bool       `json:"Disabled"`
	Comment               string     `json:"Comment"`
}

type Zone struct {
	ID      int64     `json:"Id"`
	Domain  string    `json:"Domain"`
	Records []*Record `json:"Records"`
}
