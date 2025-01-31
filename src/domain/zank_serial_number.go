package domain

type OrderMarkCodesSerialNumber struct {
	ID               int64  `db:"id" json:"id" toml:"id" yaml:"id"`
	IDOrderMarkCodes int64  `db:"id_order_mark_codes" json:"id_order_mark_codes,omitempty" toml:"id_order_mark_codes" yaml:"id_order_mark_codes,omitempty"`
	Gtin             string `db:"gtin" json:"gtin,omitempty" toml:"gtin" yaml:"gtin,omitempty"`
	SerialNumber     string `db:"serial_number" json:"serial_number,omitempty" toml:"serial_number" yaml:"serial_number,omitempty"`
	Code             string `db:"code" json:"code,omitempty" toml:"code" yaml:"code,omitempty"`
	BlockID          string `db:"block_id" json:"block_id,omitempty" toml:"block_id" yaml:"block_id,omitempty"`
	Status           string `db:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
}
