package model

type CallerId struct {
	ID        int64  `bun:"column:id;primary_key" json:"-"`    //
	Cid       string `bun:"column:cid" json:"cid"`             //
	IDCcCard  int64  `bun:"column:id_cc_card" json:"-"`        //
	Activated string `bun:"column:activated" json:"activated"` //
}

// TableName sets the insert table name for this struct type
func (c *CallerId) TableName() string {
	return "cc_callerid"
}
