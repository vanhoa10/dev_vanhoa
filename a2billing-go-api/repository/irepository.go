/*package repository

import (
	sqlclient "a2billing-go-api/internal/sql-client"
)

var SqlClient sqlclient.ISqlClientConn
*/
package repository

import (
	sqlclient "a2billing-go-api/internal/sql-client"
	"github.com/uptrace/bun"
)

var Db *bun.DB

var SqlClient sqlclient.ISqlClientConn
