package bucket

import "strings"

type HashFactor struct {
	TransId string

	OrderId string

	ModelId string

	Version string
}

func (factor HashFactor) Hash() uint32 {

	keyArray := []string{factor.TransId, factor.OrderId, factor.ModelId, factor.Version}

	key := strings.Join(keyArray, "")

	result := Murmur32([]byte(key), 0)

	residue := result%100 + 1

	if residue > 100 {

		residue = 100

	}

	return residue

}
