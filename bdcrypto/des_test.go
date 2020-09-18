package bdcrypto_test

import (
	"fmt"
	"github.com/Erope/Baidu-Login/bdcrypto"
	"testing"
)

func TestDESCBCDecrypt3(t *testing.T) {
	data := bdcrypto.Base64Decode([]byte(
		"4D2a4IsjSyD1GHh+w2IfMCTdO0uxHBFbzKVP4l81z+UZY2LzKqhovEcymeYeFW2lDWZSa56cS16lHjbN7woyyFNu+DMAeUaFA+4JWI/7zNXgqicIYpAcPtPZgdiLRpZzKB55q4+df9BSKC+dHge7dIK+Jv2fmV8vIWwvDL/1JYLB/LSB55IZ+BQQS9kPSCLju/snryQFPzHMlgHDNsTjqZWGDmOfMfXR/WriX+9Qkh1fu1zld28T7H+GgVSkyD1ldzDsTZU6vwxJk1Fyqzl2BtoB7OAFxbPWAGipKtTycN+ss/4sQpDLKPrnUJWcWfKb635v3wK+t1bfS7U+shWPKU5DoPLfXch/xEakF2Q7mXSGFxSs//ajvpS5YnovGCTd4iX2FsMOqB9sdCAx8l5CTGYWXuyDsrwxaJKeENav0lBLLqtd5X83VIv2YAs86pJUQEUPBI0wNlPNkEVFH+XFku/xIMavUvXC1hhwEwQy2apntyd+jiq0rwgDd+tJpcDTAzUa6SvgmMT97DN0FogPuqQ3Dar9HsvjNkv87RI+TCJSABWUVJvw+mzjfgaJL5cF",
	))
	ciphertext, _ := bdcrypto.DESCBCDecrypt3(data, []byte("fHSgdfH326XJ@328%521DSAg"), []byte("01234567"))
	fmt.Printf("%s\n", ciphertext)

	data = bdcrypto.Base64Decode([]byte(
		"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",
	))

	ciphertext, _ = bdcrypto.DESCBCDecrypt3(data, []byte("common3DESCode$%^_wangsu"), []byte("01234567"))
	fmt.Printf("%s\n", ciphertext)
}