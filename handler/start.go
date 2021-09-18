package handler

import (
	"github.com/yyxhgy/anylink-server/admin"
	"github.com/yyxhgy/anylink-server/base"
	"github.com/yyxhgy/anylink-server/dbdata"
	"github.com/yyxhgy/anylink-server/sessdata"
)

func Start() {
	dbdata.Start()
	sessdata.Start()

	switch base.Cfg.LinkMode {
	case base.LinkModeTUN:
		checkTun()
	case base.LinkModeTAP:
		checkTap()
	case base.LinkModeMacvtap:
		checkMacvtap()
	default:
		base.Fatal("LinkMode is err")
	}

	go admin.StartAdmin()
	go startTls()
	go startDtls()
}

func Stop() {
	_ = dbdata.Stop()
	destroyVtap()
}
