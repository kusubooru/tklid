// +build stats

package tklid

import (
	"fmt"
	"testing"
)

func TestStats(t *testing.T) {
	fmt.Println("Possible combos:", combos())
	fmt.Println("Longest ID length:", len(findLongestID()))
}
