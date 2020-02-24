package elgamal

import (
	"encoding/json"
	"math/big"

	"github.com/zerjioang/ssscomp/lib/common"
)

// ElGamal cryptosystem
// ELGamal Ciphertext
type Cypher struct {
	C1, C2, P *big.Int
}

func (cs *Cypher) Mul(cypher1, cypher2 *Cypher) *Cypher {
	cs.C1 = new(big.Int).Mod(new(big.Int).Mul(cypher1.C1, cypher2.C1), cypher1.P)
	cs.C2 = new(big.Int).Mod(new(big.Int).Mul(cypher1.C2, cypher2.C2), cypher1.P)
	cs.P = cypher1.P
	return cs
}

// Encodes crypto message as JSON
func (cs *Cypher) GetJson() ([]byte, error) {
	return json.Marshal(map[string]string{
		"c1": common.BigIntAsHex(cs.C1),
		"c2": common.BigIntAsHex(cs.C2),
		"p":  common.BigIntAsHex(cs.P),
	})
}

// Encodes public key as hexadecimal JSON map
func (cs *Cypher) ToJSON() map[string]string {
	return map[string]string{
		"c1": common.BigIntAsHex(cs.C1),
		"c2": common.BigIntAsHex(cs.C2),
		"p":  common.BigIntAsHex(cs.P),
	}
}

func (cs *Cypher) UnmarshalJSON(bytes []byte) error {
	m := make(map[string]string)
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return err
	}
	cs.FromJSON(m)
	return nil
}

func (cs *Cypher) FromJSON(json map[string]string) (*Cypher, error) {
	var err error
	cs.C1, err = common.BigIntFromHex(json["C1"])
	if err != nil {
		return nil, err
	}
	cs.C2, err = common.BigIntFromHex(json["C2"])
	if err != nil {
		return nil, err
	}
	cs.P, err = common.BigIntFromHex(json["P"])
	if err != nil {
		return nil, err
	}
	return cs, err
}
