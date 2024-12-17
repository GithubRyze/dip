package proxy

import (
	"testing"
)

func TestDipProxyManager_MatchPath(t *testing.T) {
	extractMap := make(map[string]DipProxy)
	extractMap["/dev/test/001"] = DipProxy{Path: "/dev/test/001"}
	extractMap["/uat/abc/001"] = DipProxy{Path: "/uat/abc/001"}
	extractMap["/prod/efd/001"] = DipProxy{Path: "/prod/efd/001"}
	prefixMap := make(map[string]DipProxy)
	prefixMap["/dcms/dev/"] = DipProxy{ServiceName: "dcms", Path: "/dcms/dev/"}
	prefixMap["/ehrss/uat/"] = DipProxy{ServiceName: "ehrss", Path: "/ehrss/uat/"}
	prefixMap["/ewell/"] = DipProxy{ServiceName: "ewell", Path: "/ewell/"}

	manager := DipProxyManager{
		extractProxyCache: extractMap,
		prefixProxyCache:  prefixMap,
	}
	_, err := manager.MatchPath("/uat")
	if err == nil {
		t.Logf("Expected err, but get nil")
		t.Fail()
	}
	_, err = manager.MatchPath("//uat/test_011")
	if err == nil {
		t.Logf("Expected err, but get nil")
		t.Fail()
	}
	_, err = manager.MatchPath("//dev/test/001")
	if err == nil {
		t.Logf("Expected err, but get nil")
		t.Fail()
	}
	dip, _ := manager.MatchPath("/dev/test/001")
	if dip.Path != "/dev/test/001" {
		t.Logf("Expected /dev/test/001, but get %s", dip.Path)
		t.Fail()
	}
	dip, _ = manager.MatchPath("/uat/abc/001")
	if dip.Path != "/uat/abc/001" {
		t.Logf("Expected /uat/abc/001, but get %s", dip.Path)
		t.Fail()
	}
	dip, _ = manager.MatchPath("/ewell/test")
	if dip.ServiceName != "ewell" {
		t.Logf("Expected ewell, but get %s", dip.ServiceName)
		t.Fail()
	}
	dip, _ = manager.MatchPath("/ewell/001")
	if dip.ServiceName != "ewell" {
		t.Logf("Expected ewell, but get %s", dip.ServiceName)
		t.Fail()
	}

	dip, _ = manager.MatchPath("/dcms/dev/TEST/0001")
	if dip.ServiceName != "dcms" {
		t.Logf("Expected dcms, but get %s", dip.ServiceName)
		t.Fail()
	}
	dip, _ = manager.MatchPath("/ehrss/uat/TEST/0001")
	if dip.ServiceName != "ehrss" {
		t.Logf("Expected ehrss, but get %s", dip.ServiceName)
		t.Fail()
	}

}
