package ip_helper

import "testing"

func TestIpConvert(t *testing.T) {
	for _, unit := range []struct{
		input string
		expected int
	} {
		{
			"172.168.5.1",
			2896692481,
		},
		{
			"172 .168.5.1",
			2896692481,
		},
		{
			"172. 168.5.1",
			2896692481,
		},
		{
			" 172.168.5.1",
			2896692481,
		},
	} {
		res, _ := IpConvert(unit.input)
		if res != unit.expected {
			t.Error("input = ", unit.input, ", expected = ", unit.expected, ", result = ", res)
		}
	}
}