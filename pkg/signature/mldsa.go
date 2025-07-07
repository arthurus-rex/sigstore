package signature

	

import (
        "fmt"
        "crypto"
       mldsa "github.com/arthurus-rex/tob-mldsa"
)
	

// Only for testing, delete later
func print_me() {
    fmt.Println("hello world")
}

var mldsaSupportedHashFuncs = []crypto.Hash{
	crypto.SHA512,
}

// MLDSASigner is a signature.Signer that uses ML-DSA pure variant (FIPS-204) with a pre-hashing
// step similar to ed25519ph.  This is distinct from HashML-DSA and done for simplicity. 
type MLDSASigner struct {
	priv mldsa.PrivateKey
}
