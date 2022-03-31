package model

import (
	"database/sql"
	"time"
)

type Card struct {
	ID                 int64          `bun:"column:id;primary_key" json:"id"`                       //
	Creationdate       time.Time      `bun:"column:creationdate" json:"creationdate"`               //
	Firstusedate       time.Time      `bun:"column:firstusedate" json:"firstusedate"`               //
	Expirationdate     time.Time      `bun:"column:expirationdate" json:"expirationdate"`           //
	Enableexpire       sql.NullInt64  `bun:"column:enableexpire" json:"enableexpire"`               //
	Expiredays         sql.NullInt64  `bun:"column:expiredays" json:"expiredays"`                   //
	Username           string         `bun:"column:username" json:"username"`                       //
	Useralias          string         `bun:"column:useralias" json:"useralias"`                     //
	Uipass             string         `bun:"column:uipass" json:"uipass"`                           //
	Credit             float64        `bun:"column:credit" json:"credit"`                           //
	Tariff             sql.NullInt64  `bun:"column:tariff" json:"tariff"`                           //
	IDDidgroup         sql.NullInt64  `bun:"column:id_didgroup" json:"id_didgroup"`                 //
	Activated          string         `bun:"column:activated" json:"activated"`                     //
	Status             int            `bun:"column:status" json:"status"`                           //
	Lastname           string         `bun:"column:lastname" json:"lastname"`                       //
	Firstname          string         `bun:"column:firstname" json:"firstname"`                     //
	Address            string         `bun:"column:address" json:"address"`                         //
	City               string         `bun:"column:city" json:"city"`                               //
	State              string         `bun:"column:state" json:"state"`                             //
	Country            string         `bun:"column:country" json:"country"`                         //
	Zipcode            string         `bun:"column:zipcode" json:"zipcode"`                         //
	Phone              string         `bun:"column:phone" json:"phone"`                             //
	Email              string         `bun:"column:email" json:"email"`                             //
	Fax                string         `bun:"column:fax" json:"fax"`                                 //
	Inuse              sql.NullInt64  `bun:"column:inuse" json:"inuse"`                             //
	Simultaccess       sql.NullInt64  `bun:"column:simultaccess" json:"simultaccess"`               //
	Currency           sql.NullString `bun:"column:currency" json:"currency"`                       //
	Lastuse            time.Time      `bun:"column:lastuse" json:"lastuse"`                         //
	Nbused             sql.NullInt64  `bun:"column:nbused" json:"nbused"`                           //
	Typepaid           sql.NullInt64  `bun:"column:typepaid" json:"typepaid"`                       //
	Creditlimit        sql.NullInt64  `bun:"column:creditlimit" json:"creditlimit"`                 //
	Voipcall           sql.NullInt64  `bun:"column:voipcall" json:"voipcall"`                       //
	SipBuddy           sql.NullInt64  `bun:"column:sip_buddy" json:"sip_buddy"`                     //
	IaxBuddy           sql.NullInt64  `bun:"column:iax_buddy" json:"iax_buddy"`                     //
	Language           sql.NullString `bun:"column:language" json:"language"`                       //
	Redial             string         `bun:"column:redial" json:"redial"`                           //
	Runservice         sql.NullInt64  `bun:"column:runservice" json:"runservice"`                   //
	Nbservice          sql.NullInt64  `bun:"column:nbservice" json:"nbservice"`                     //
	IDCampaign         sql.NullInt64  `bun:"column:id_campaign" json:"id_campaign"`                 //
	NumTrialsDone      sql.NullInt64  `bun:"column:num_trials_done" json:"num_trials_done"`         //
	Vat                float32        `bun:"column:vat" json:"vat"`                                 //
	Servicelastrun     time.Time      `bun:"column:servicelastrun" json:"servicelastrun"`           //
	Initialbalance     float64        `bun:"column:initialbalance" json:"initialbalance"`           //
	Invoiceday         sql.NullInt64  `bun:"column:invoiceday" json:"invoiceday"`                   //
	Autorefill         sql.NullInt64  `bun:"column:autorefill" json:"autorefill"`                   //
	Loginkey           string         `bun:"column:loginkey" json:"loginkey"`                       //
	MacAddr            string         `bun:"column:mac_addr" json:"mac_addr"`                       //
	IDTimezone         sql.NullInt64  `bun:"column:id_timezone" json:"id_timezone"`                 //
	Tag                string         `bun:"column:tag" json:"tag"`                                 //
	VoicemailPermitted int            `bun:"column:voicemail_permitted" json:"voicemail_permitted"` //
	VoicemailActivated int            `bun:"column:voicemail_activated" json:"voicemail_activated"` //
	LastNotification   time.Time      `bun:"column:last_notification" json:"last_notification"`     //
	EmailNotification  string         `bun:"column:email_notification" json:"email_notification"`   //
	NotifyEmail        int            `bun:"column:notify_email" json:"notify_email"`               //
	CreditNotification int            `bun:"column:credit_notification" json:"credit_notification"` //
	IDGroup            int            `bun:"column:id_group" json:"id_group"`                       //
	CompanyName        string         `bun:"column:company_name" json:"company_name"`               //
	CompanyWebsite     string         `bun:"column:company_website" json:"company_website"`         //
	VatRn              sql.NullString `bun:"column:vat_rn" json:"vat_rn"`                           //
	Traffic            sql.NullInt64  `bun:"column:traffic" json:"traffic"`                         //
	TrafficTarget      string         `bun:"column:traffic_target" json:"traffic_target"`           //
	Discount           float64        `bun:"column:discount" json:"discount"`                       //
	Restriction        int            `bun:"column:restriction" json:"restriction"`                 //
	IDSeria            sql.NullInt64  `bun:"column:id_seria" json:"id_seria"`                       //
	Serial             sql.NullInt64  `bun:"column:serial" json:"serial"`                           //
	Block              int            `bun:"column:block" json:"block"`                             //
	LockPin            sql.NullString `bun:"column:lock_pin" json:"lock_pin"`                       //
	LockDate           time.Time      `bun:"column:lock_date" json:"lock_date"`                     //
	MaxConcurrent      int            `bun:"column:max_concurrent" json:"max_concurrent"`           //
	APIKey             sql.NullString `bun:"column:api_key" json:"api_key"`                         //
}

type CardInfo struct {
	ID             int64      `bun:"column:id;primary_key" json:"id"`                      //
	Creationdate   time.Time  `bun:"column:creationdate" json:"creationdate"`              //
	Firstusedate   time.Time  `bun:"column:firstusedate" json:"firstusedate"`              //
	Expirationdate time.Time  `bun:"column:expirationdate" json:"expirationdate"`          //
	Enableexpire   NullInt64  `bun:"column:enableexpire" json:"enableexpire"`              //
	Expiredays     NullInt64  `bun:"column:expiredays" json:"expiredays"`                  //
	Username       string     `bun:"column:username" json:"username"`                      //
	Useralias      string     `bun:"column:useralias" json:"useralias"`                    //
	Credit         float64    `bun:"column:credit" json:"credit"`                          //
	Activated      string     `bun:"column:activated" json:"activated"`                    //
	Status         int        `bun:"column:status" json:"status"`                          //
	Lastuse        time.Time  `bun:"column:lastuse" json:"lastuse"`                        //
	Creditlimit    NullInt64  `bun:"column:creditlimit" json:"creditlimit"`                //
	IDGroup        int        `bun:"column:id_group" json:"id_group"`                      //
	Tariff         NullInt64  `bun:"column:tariff" json:"call_plan"`                       //
	CallerIds      []CallerId `bun:"foreignKey:IDCcCard ;references:ID" json:"caller_ids"` //
}

// TableName sets the insert table name for this struct type
func (c *Card) TableName() string {
	return "cc_card"
}

func (c *CardInfo) TableName() string {
	return "cc_card"
}
