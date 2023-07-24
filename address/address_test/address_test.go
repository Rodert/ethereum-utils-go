package addresstest_test

import (
	"ethereum-utils-go/address"
	"testing"
)

func TestAddressTransform(t *testing.T) {
	address.AddressTransform("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	address.LongAddressTransform("0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f")
}
