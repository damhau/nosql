//go:build nobadger || nobadgerv2
// +build nobadger nobadgerv2

package badger

import "github.com/damhau/nosql/database"

type DB = database.NotSupportedDB
