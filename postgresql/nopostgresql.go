//go:build nopgx
// +build nopgx

package postgresql

import "github.com/damhau/nosql/database"

type DB = database.NotSupportedDB
