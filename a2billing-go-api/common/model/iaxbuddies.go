package model

import "database/sql"

type IaxBuddies struct {
	ID                         int            `bun:"column:id;primary_key" json:"id"`                                       //
	IDCcCard                   int            `bun:"column:id_cc_card" json:"id_cc_card"`                                   //
	Name                       string         `bun:"column:name" json:"name"`                                               //
	Accountcode                string         `bun:"column:accountcode" json:"accountcode"`                                 //
	Regexten                   string         `bun:"column:regexten" json:"regexten"`                                       //
	Amaflags                   sql.NullString `bun:"column:amaflags" json:"amaflags"`                                       //
	Callerid                   string         `bun:"column:callerid" json:"callerid"`                                       //
	Context                    string         `bun:"column:context" json:"context"`                                         //
	DEFAULTip                  sql.NullString `bun:"column:DEFAULTip" json:"DEFAULTip"`                                     //
	Host                       string         `bun:"column:host" json:"host"`                                               //
	Language                   sql.NullString `bun:"column:language" json:"language"`                                       //
	Deny                       string         `bun:"column:deny" json:"deny"`                                               //
	Permit                     sql.NullString `bun:"column:permit" json:"permit"`                                           //
	Mask                       string         `bun:"column:mask" json:"mask"`                                               //
	Port                       string         `bun:"column:port" json:"port"`                                               //
	Qualify                    sql.NullString `bun:"column:qualify" json:"qualify"`                                         //
	Secret                     string         `bun:"column:secret" json:"secret"`                                           //
	Type                       string         `bun:"column:type" json:"type"`                                               //
	Username                   string         `bun:"column:username" json:"username"`                                       //
	Disallow                   string         `bun:"column:disallow" json:"disallow"`                                       //
	Allow                      string         `bun:"column:allow" json:"allow"`                                             //
	Regseconds                 int            `bun:"column:regseconds" json:"regseconds"`                                   //
	Ipaddr                     string         `bun:"column:ipaddr" json:"ipaddr"`                                           //
	Trunk                      sql.NullString `bun:"column:trunk" json:"trunk"`                                             //
	Dbsecret                   string         `bun:"column:dbsecret" json:"dbsecret"`                                       //
	Regcontext                 string         `bun:"column:regcontext" json:"regcontext"`                                   //
	Sourceaddress              string         `bun:"column:sourceaddress" json:"sourceaddress"`                             //
	Mohinterpret               string         `bun:"column:mohinterpret" json:"mohinterpret"`                               //
	Mohsuggest                 string         `bun:"column:mohsuggest" json:"mohsuggest"`                                   //
	Inkeys                     string         `bun:"column:inkeys" json:"inkeys"`                                           //
	Outkey                     string         `bun:"column:outkey" json:"outkey"`                                           //
	CidNumber                  string         `bun:"column:cid_number" json:"cid_number"`                                   //
	Sendani                    string         `bun:"column:sendani" json:"sendani"`                                         //
	Fullname                   string         `bun:"column:fullname" json:"fullname"`                                       //
	Auth                       string         `bun:"column:auth" json:"auth"`                                               //
	Maxauthreq                 string         `bun:"column:maxauthreq" json:"maxauthreq"`                                   //
	Encryption                 string         `bun:"column:encryption" json:"encryption"`                                   //
	Transfer                   string         `bun:"column:transfer" json:"transfer"`                                       //
	Jitterbuffer               string         `bun:"column:jitterbuffer" json:"jitterbuffer"`                               //
	Forcejitterbuffer          string         `bun:"column:forcejitterbuffer" json:"forcejitterbuffer"`                     //
	Codecpriority              string         `bun:"column:codecpriority" json:"codecpriority"`                             //
	Qualifysmoothing           string         `bun:"column:qualifysmoothing" json:"qualifysmoothing"`                       //
	Qualifyfreqok              string         `bun:"column:qualifyfreqok" json:"qualifyfreqok"`                             //
	Qualifyfreqnotok           string         `bun:"column:qualifyfreqnotok" json:"qualifyfreqnotok"`                       //
	Timezone                   string         `bun:"column:timezone" json:"timezone"`                                       //
	Adsi                       string         `bun:"column:adsi" json:"adsi"`                                               //
	Setvar                     string         `bun:"column:setvar" json:"setvar"`                                           //
	Requirecalltoken           string         `bun:"column:requirecalltoken" json:"requirecalltoken"`                       //
	Maxcallnumbers             string         `bun:"column:maxcallnumbers" json:"maxcallnumbers"`                           //
	MaxcallnumbersNonvalidated string         `bun:"column:maxcallnumbers_nonvalidated" json:"maxcallnumbers_nonvalidated"` //
}

// TableName sets the insert table name for this struct type
func (i *IaxBuddies) TableName() string {
	return "cc_iax_buddies"
}
