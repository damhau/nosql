//go:build nobadger || nobadgerv1
// +build nobadger nobadgerv1

package badger

import "github.com/damhau/nosql/database"

type DB = database.NotSupportedDB
