package problem16

import (
	"math/big"
	"strconv"
)

func SumPowerOfTwo(exp int64) (int, error) {
	gros := new(big.Int)
	gros = gros.Exp(big.NewInt(2), big.NewInt(exp), nil)
	s := gros.String()
	sum := 0
	for _, x := range s {
		c, err := strconv.Unquote(strconv.QuoteRune(x))
		if err != nil {
			return 0, err
		}
		i, err := strconv.Atoi(c)
		if err != nil {
			return 0, err
		}
		sum = sum + i
	}
	return sum, nil
}
