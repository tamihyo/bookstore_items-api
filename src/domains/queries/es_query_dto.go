package queries

type EsQuery struct {
	Equals []FieldValue `json:"value"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
