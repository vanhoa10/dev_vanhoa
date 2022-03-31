package model

import "database/sql"

type SipBuddies struct {
	ID                int            `bun:"column:id;primary_key" json:"id"`                   //
	IDCcCard          int            `bun:"column:id_cc_card" json:"id_cc_card"`               //
	Name              string         `bun:"column:name" json:"name"`                           //
	Accountcode       string         `bun:"column:accountcode" json:"accountcode"`             //
	Regexten          string         `bun:"column:regexten" json:"regexten"`                   //
	Amaflags          sql.NullString `bun:"column:amaflags" json:"amaflags"`                   //
	Callgroup         sql.NullString `bun:"column:callgroup" json:"callgroup"`                 //
	Callerid          string         `bun:"column:callerid" json:"callerid"`                   //
	Canreinvite       string         `bun:"column:canreinvite" json:"canreinvite"`             //
	Context           string         `bun:"column:context" json:"context"`                     //
	DEFAULTip         sql.NullString `bun:"column:DEFAULTip" json:"DEFAULTip"`                 //
	Dtmfmode          string         `bun:"column:dtmfmode" json:"dtmfmode"`                   //
	Fromuser          string         `bun:"column:fromuser" json:"fromuser"`                   //
	Fromdomain        string         `bun:"column:fromdomain" json:"fromdomain"`               //
	Host              string         `bun:"column:host" json:"host"`                           //
	Insecure          string         `bun:"column:insecure" json:"insecure"`                   //
	Language          sql.NullString `bun:"column:language" json:"language"`                   //
	Mailbox           string         `bun:"column:mailbox" json:"mailbox"`                     //
	Md5secret         string         `bun:"column:md5secret" json:"md5secret"`                 //
	Nat               sql.NullString `bun:"column:nat" json:"nat"`                             //
	Deny              string         `bun:"column:deny" json:"deny"`                           //
	Permit            sql.NullString `bun:"column:permit" json:"permit"`                       //
	Mask              string         `bun:"column:mask" json:"mask"`                           //
	Pickupgroup       sql.NullString `bun:"column:pickupgroup" json:"pickupgroup"`             //
	Port              string         `bun:"column:port" json:"port"`                           //
	Qualify           sql.NullString `bun:"column:qualify" json:"qualify"`                     //
	Restrictcid       sql.NullString `bun:"column:restrictcid" json:"restrictcid"`             //
	Rtptimeout        sql.NullString `bun:"column:rtptimeout" json:"rtptimeout"`               //
	Rtpholdtimeout    sql.NullString `bun:"column:rtpholdtimeout" json:"rtpholdtimeout"`       //
	Secret            string         `bun:"column:secret" json:"secret"`                       //
	Type              string         `bun:"column:type" json:"type"`                           //
	Username          string         `bun:"column:username" json:"username"`                   //
	Disallow          string         `bun:"column:disallow" json:"disallow"`                   //
	Allow             string         `bun:"column:allow" json:"allow"`                         //
	Musiconhold       string         `bun:"column:musiconhold" json:"musiconhold"`             //
	Regseconds        int            `bun:"column:regseconds" json:"regseconds"`               //
	Ipaddr            string         `bun:"column:ipaddr" json:"ipaddr"`                       //
	Cancallforward    sql.NullString `bun:"column:cancallforward" json:"cancallforward"`       //
	Fullcontact       string         `bun:"column:fullcontact" json:"fullcontact"`             //
	Setvar            string         `bun:"column:setvar" json:"setvar"`                       //
	Regserver         sql.NullString `bun:"column:regserver" json:"regserver"`                 //
	Lastms            sql.NullString `bun:"column:lastms" json:"lastms"`                       //
	Defaultuser       string         `bun:"column:defaultuser" json:"defaultuser"`             //
	Auth              string         `bun:"column:auth" json:"auth"`                           //
	Subscribemwi      string         `bun:"column:subscribemwi" json:"subscribemwi"`           //
	Vmexten           string         `bun:"column:vmexten" json:"vmexten"`                     //
	CidNumber         string         `bun:"column:cid_number" json:"cid_number"`               //
	Callingpres       string         `bun:"column:callingpres" json:"callingpres"`             //
	Usereqphone       string         `bun:"column:usereqphone" json:"usereqphone"`             //
	Incominglimit     string         `bun:"column:incominglimit" json:"incominglimit"`         //
	Subscribecontext  string         `bun:"column:subscribecontext" json:"subscribecontext"`   //
	Musicclass        string         `bun:"column:musicclass" json:"musicclass"`               //
	Mohsuggest        string         `bun:"column:mohsuggest" json:"mohsuggest"`               //
	Allowtransfer     string         `bun:"column:allowtransfer" json:"allowtransfer"`         //
	Autoframing       string         `bun:"column:autoframing" json:"autoframing"`             //
	Maxcallbitrate    string         `bun:"column:maxcallbitrate" json:"maxcallbitrate"`       //
	Outboundproxy     string         `bun:"column:outboundproxy" json:"outboundproxy"`         //
	Rtpkeepalive      string         `bun:"column:rtpkeepalive" json:"rtpkeepalive"`           //
	Useragent         sql.NullString `bun:"column:useragent" json:"useragent"`                 //
	Callbackextension sql.NullString `bun:"column:callbackextension" json:"callbackextension"` //
}

// TableName sets the insert table name for this struct type
func (s *SipBuddies) TableName() string {
	return "cc_sip_buddies"
}
