package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
	oneLayer := map[string]interface{}{
		"flatField": true,
	}
	twoLayer := Clone(oneLayer)
	twoLayer["oneLayer"] = oneLayer

	threeLayer := Clone(twoLayer)
	threeLayer["twoLayer"] = twoLayer

	fourLayer := Clone(threeLayer)
	fourLayer["threeLayer"] = threeLayer

	flatLayer := map[string]interface{}{
		"flatField": true,

		"oneLayer:flatField": true,

		"twoLayer:flatField":          true,
		"twoLayer:oneLayer:flatField": true,

		"threeLayer:flatField":                   true,
		"threeLayer:oneLayer:flatField":          true,
		"threeLayer:twoLayer:flatField":          true,
		"threeLayer:twoLayer:oneLayer:flatField": true,
	}

	t.Run("Base example of deeply nested", func(t *testing.T) {
		assert.EqualValues(t, flatLayer, Flatten(fourLayer))
	})
}
