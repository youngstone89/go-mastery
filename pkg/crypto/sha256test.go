package crypto

import (
	"crypto/sha256"
	"fmt"
)

func DoCryptoTest()  {

	s := "NMDPTRIAL_kim_seok_nuance_com_20210427T044351526748"
	hsha2 := sha256.Sum256([]byte(s))
	fmt.Printf("SHA256: %x\n", hsha2)

}
