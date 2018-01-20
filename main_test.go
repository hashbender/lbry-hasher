package main

import (
	"encoding/hex"
	"testing"
)

func TestHasher(t *testing.T) {
	data := "0100000001077cb45fc6f552c703f12717c12963c4b0e11f02818205c779feff476b9b658b010000006b483045022100c7beac4a68b9cbfb79536372492a939f9530d3ef77da97400efe954ae67e30240220292f4fe6dee7b1f3d3d6949dda32c412fbd1939a48cfcd5b23b979cec1ae49ae01210275c88cc7c78cf33899e2eb9b7f3b722d4a6fc00edfc1bbaa2b16c4c54f8f11d8feffffff0200c2eb0b000000001976a9144436223de135540f62af4464002f2cffd8db61d288ac1c6d614e030000001976a914e21762df2a9d9f9f04e3efb463186a41d925723a88aceeb10400"
	expectedHash := "a3b7cbdc533de7348c8b2ef1502b3c95b50b8eab2c20d17fe66673ee7dc84b77"

	dataBytes, _ := hex.DecodeString(data)
	hash := LbryHash(dataBytes)

	if expectedHash != hash {
		t.Errorf("Did not calculate correct hash. Got: %s, Expected: %s.", hash, expectedHash)
	}
}
