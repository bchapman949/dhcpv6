package dhcpv6

import (
	"bytes"
	"net"
	"strings"
	"testing"
	"time"
)

func TestOptionTypeString(t *testing.T) {
	tests := []struct {
		in  OptionType
		out string
	}{
		{0, "Unknown (0)"},
		{OptionTypeClientID, "Client Identifier (1)"},
		{OptionTypeServerID, "Server Identifier (2)"},
		{OptionTypeIANA, "Identity Association for Non-temporary Addresses (3)"},
		{OptionTypeIATA, "Identity Association for Temporary Addresses (4)"},
		{OptionTypeIAAddress, "Identity Association Address (5)"},
		{OptionTypeOptionRequest, "Option Request (6)"},
		{OptionTypePreference, "Preference (7)"},
		{OptionTypeElapsedTime, "Elapsed Time (8)"},
		{OptionTypeRelayMessage, "Relay Message (9)"},
		{OptionTypeAuthentication, "Authentication (11)"},
		{OptionTypeServerUnicast, "Server Unicast (12)"},
		{OptionTypeStatusCode, "Status Code (13)"},
		{OptionTypeRapidCommit, "Rapid Commit (14)"},
		{OptionTypeUserClass, "User Class (15)"},
		{OptionTypeVendorClass, "Vendor Class (16)"},
		{OptionTypeVendorOption, "Vendor-specific Information (17)"},
		{OptionTypeInterfaceID, "Interface-ID (18)"},
		{OptionTypeReconfigureMessage, "Reconfigure Message (19)"},
		{OptionTypeReconfigureAccept, "Reconfigure Accept (20)"},
		{OptionTypeDNSServer, "DNS Server (23)"},
		{OptionTypeDNSSearchList, "DNS Search List (24)"},
	}

	for _, test := range tests {
		if strings.Compare(test.in.String(), test.out) != 0 {
			t.Errorf("expected %s but got %s", test.out, test.in.String())
		}
	}
}

// test OptionClientID
func TestOptionClientID(t *testing.T) {
	var opt *OptionClientID

	fixtbyte := []byte{0, 1, 0, 14, 0, 1, 0, 1, 29, 205, 101, 0, 170, 187, 204, 221, 238, 255}
	// test decoding bytes to []Option
	if list, err := ParseOptions(fixtbyte); err != nil {
		t.Errorf("could not parse fixture: %s", err)
	} else if len(list) != 1 {
		t.Errorf("expected exactly 1 option, got %d", len(list))
	} else {
		opt = list[0].(*OptionClientID)
	}

	// check contents of Option
	if opt.Type() != OptionTypeClientID {
		t.Errorf("unexpected type: %s", opt.Type())
	}
	// the DUID of this Option is tested separately
	// for now, just check the type of the DUID
	if opt.DUID.Type() != DUIDTypeLLT {
		t.Errorf("unexpected DUID type: %s", opt.DUID.Type())
	}

	// test matching output for String()
	fixtstr := "client-ID hwaddr/time type 1 time 500000000 aa:bb:cc:dd:ee:ff"
	if fixtstr != opt.String() {
		t.Errorf("unexpected String() output: %s", opt.String())
	}

	// test if marshalled bytes match fixture
	if mshByte, err := opt.Marshal(); err != nil {
		t.Errorf("error marshalling ClientID: %s", err)
	} else if bytes.Compare(fixtbyte, mshByte) != 0 {
		t.Errorf("marshalled ClientID didn't match fixture!\nfixture: %v\nmarshal: %v", fixtbyte, mshByte)
	}

	// create same OptionClientID and see if its marshal matches fixture
	opt = &OptionClientID{
		OptionBase: &OptionBase{
			OptionType: OptionTypeClientID,
		},
		DUID: &DUIDLLT{
			DUIDBase: &DUIDBase{
				DUIDType: DUIDTypeLLT,
			},
			HardwareType: 1,
			Time:         time.Unix(1446771200, 0),
		},
	}
	opt.DUID.(*DUIDLLT).LinkLayerAddress, _ = net.ParseMAC("aa:bb:cc:dd:ee:ff")
	if mshByte, err := opt.Marshal(); err != nil {
		t.Errorf("error marshalling ClientID: %s", err)
	} else if bytes.Compare(mshByte, fixtbyte) != 0 {
		t.Errorf("marshalled ClientID didn't match fixture!\nfixture: %v\nmarshal: %v", fixtbyte, mshByte)
	}
}

// test OptionServerID
func TestOptionServerID(t *testing.T) {
	var opt *OptionServerID

	fixtbyte := []byte{0, 2, 0, 14, 0, 1, 0, 1, 29, 205, 101, 0, 170, 187, 204, 221, 238, 255}
	// test decoding bytes to []Option
	if list, err := ParseOptions(fixtbyte); err != nil {
		t.Errorf("could not parse fixture: %s", err)
	} else if len(list) != 1 {
		t.Errorf("expected exactly 1 option, got %d", len(list))
	} else {
		opt = list[0].(*OptionServerID)
	}

	// check contents of Option
	if opt.Type() != OptionTypeServerID {
		t.Errorf("unexpected type: %s", opt.Type())
	}
	// the DUID of this Option is tested separately
	// for now, just check the type of the DUID
	if opt.DUID.Type() != DUIDTypeLLT {
		t.Errorf("unexpected DUID type: %s", opt.DUID.Type())
	}

	// test matching output for String()
	fixtstr := "server-ID hwaddr/time type 1 time 500000000 aa:bb:cc:dd:ee:ff"
	if fixtstr != opt.String() {
		t.Errorf("unexpected String() output: %s", opt.String())
	}

	// test if marshalled bytes match fixture
	if mshByte, err := opt.Marshal(); err != nil {
		t.Errorf("error marshalling ServerID: %s", err)
	} else if bytes.Compare(fixtbyte, mshByte) != 0 {
		t.Errorf("marshalled ServerID didn't match fixture!\nfixture: %v\nmarshal: %v", fixtbyte, mshByte)
	}

	// create same OptionClientID and see if its marshal matches fixture
	opt = &OptionServerID{
		OptionBase: &OptionBase{
			OptionType: OptionTypeServerID,
		},
		DUID: &DUIDLLT{
			DUIDBase: &DUIDBase{
				DUIDType: DUIDTypeLLT,
			},
			HardwareType: 1,
			Time:         time.Unix(1446771200, 0),
		},
	}
	opt.DUID.(*DUIDLLT).LinkLayerAddress, _ = net.ParseMAC("aa:bb:cc:dd:ee:ff")
	if mshByte, err := opt.Marshal(); err != nil {
		t.Errorf("error marshalling ServerID: %s", err)
	} else if bytes.Compare(mshByte, fixtbyte) != 0 {
		t.Errorf("marshalled ServerID didn't match fixture!\nfixture: %v\nmarshal: %v", fixtbyte, mshByte)
	}
}
