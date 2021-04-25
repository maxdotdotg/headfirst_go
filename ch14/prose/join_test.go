// part of the prose package
package prose

// must import testing
import "testing"

// test functions must start with "Test"
// and take a pointer to testing.T
func TestTwoElements(t *testing.T) {
	elements := []string{"apple", "orange"}
	if JoinWithCommas(elements) != "apple and orange" {
		t.Error("didn't match expected value")
	}
}
func TestThreeElements(t *testing.T) {
	elements := []string{"apple", "orange", "pear"}
	if JoinWithCommas(elements) != "apple, orange, and pear" {
		t.Error("didn't match expected value")
	}
}
