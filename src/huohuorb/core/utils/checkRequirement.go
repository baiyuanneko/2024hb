package utils

import (
	"chihuo2104.dev/huohuorb/config"
	"chihuo2104.dev/huohuorb/core/db"
)

func CheckRequirements(cid string, reqid string) bool {
	conn := db.New(config.DB)
	// requires challenge0-in otherwise it is not authorized.
	resp := conn.Query("SELECT * FROM `clientids` WHERE `cid`=\"" + cid + "\" AND `action`=\"" + reqid + "\"")
	if resp.Next() {
		return true
	} else {
		return false
	}
}
