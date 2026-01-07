package http_proxy_test

import (
	"server/boot"
	"testing"
)

var serverName = boot.RootPath() + "/PMsgLogicServer"

func TestMain(m *testing.M) {
	_ = boot.Boot([]string{serverName}, serverName)
	m.Run()
}
