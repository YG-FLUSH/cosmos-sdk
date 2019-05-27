package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/keys"
	cryptokeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys/mintkey"
	amino "github.com/tendermint/go-amino"
)

func main() {
	name := "hot-wallet"

	kb, _ := keys.NewKeyBaseFromDir("/Users/wenyeguang/.gaiacli")
	info, _ := kb.Get(name)
	linfo := cryptokeys.ConvertLocalInfo(info)
	passphrase, _ := keys.ReadPassphraseFromStdin(name)
	priv, _ := mintkey.UnarmorDecryptPrivKey(linfo.PrivKeyArmor, passphrase)
	priString := []byte(string(strings.TrimSpace(string(priv.Bytes()[amino.PrefixBytesLen:]))))
	ko, _ := cryptokeys.Bech32KeyOutput(info)
	fmt.Println("private key:", hex.EncodeToString(priString))
	fmt.Println("address:", ko.Address)
}
