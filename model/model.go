package model

type CDCData struct {
	Schema  Schema  `json:"schema"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Before      interface{} `json:"before"`
	After       interface{} `json:"after"`
	Source      Source      `json:"source"`
	Op          string      `json:"op"`
	TsMS        int64       `json:"ts_ms"`
	Transaction interface{} `json:"transaction"`
}

type Source struct {
	Version   string      `json:"version"`
	Connector string      `json:"connector"`
	Name      string      `json:"name"`
	TsMS      int64       `json:"ts_ms"`
	Snapshot  string      `json:"snapshot"`
	DB        string      `json:"db"`
	Sequence  string      `json:"sequence"`
	Schema    string      `json:"schema"`
	Table     string      `json:"table"`
	TxID      int64       `json:"txId"`
	Lsn       int64       `json:"lsn"`
	Xmin      interface{} `json:"xmin"`
}

type Schema struct {
	Type     string        `json:"type"`
	Fields   []SchemaField `json:"fields"`
	Optional bool          `json:"optional"`
	Name     string        `json:"name"`
	Version  int64         `json:"version"`
}

type SchemaField struct {
	Type     string       `json:"type"`
	Fields   []FieldField `json:"fields,omitempty"`
	Optional bool         `json:"optional"`
	Name     *string      `json:"name,omitempty"`
	Field    string       `json:"field"`
	Version  *int64       `json:"version,omitempty"`
}

type FieldField struct {
	Type       Type        `json:"type"`
	Optional   bool        `json:"optional"`
	Default    *Default    `json:"default"`
	Field      string      `json:"field"`
	Name       *string     `json:"name,omitempty"`
	Version    *int64      `json:"version,omitempty"`
	Parameters *Parameters `json:"parameters,omitempty"`
}

type Parameters struct {
	Allowed string `json:"allowed"`
}

type Type string

const (
	Int64  Type = "int64"
	String Type = "string"
)

type Default struct {
	Integer *int64
	String  *string
}
