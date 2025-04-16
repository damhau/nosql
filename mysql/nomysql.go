//go:build nomysql
// +build nomysql

package mysql

import "github.com/damhau/nosql/database"

type DB = database.NotSupportedDB
