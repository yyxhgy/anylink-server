package handler

import (
	"github.com/bjdgyc/anylink/admin"
	"github.com/bjdgyc/anylink/base"
	"github.com/bjdgyc/anylink/dbdata"
	"github.com/bjdgyc/anylink/sessdata"
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
