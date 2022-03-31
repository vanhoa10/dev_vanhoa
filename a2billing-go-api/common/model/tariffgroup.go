package model

type TariffGroup struct {
	ID       int    `bun:"column:id;primary_key" json:"id"`               //
	Iduser   int    `bun:"column:iduser" json:"iduser"`                   //
	Loglevel string `bun:"column:tariffgroupname" json:"tariffgroupname"` //
}

// TableName sets the insert table name for this struct type
func (s *TariffGroup) TableName() string {
	return "cc_tariffgroup"
}
