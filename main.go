package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/ripemd160"
)

func main() {

}

func LbryHash(data []byte) string {
	sha1 := sha256.Sum256(data)
	sha2 := sha256.Sum256(sha1[:])
	sha3 := sha512.Sum512(sha2[:])
	mdHasher := ripemd160.New()
	mdHasher.Write(sha3[0:32])
	hashB := mdHasher.Sum(nil)
	mdHasher = ripemd160.New()
	mdHasher.Write(sha3[32:64])
	hashC := mdHasher.Sum(nil)
	ctx := sha256.New()
	ctx.Write(hashB[0:20])
	ctx.Write(hashC[0:20])
	sha4 := ctx.Sum(nil)
	sha5 := sha256.Sum256(sha4)
	return hex.EncodeToString(sha5[:])
}
