package conversion

import "testing"

var decimalTestCases = map[int]string{
	0: "0", 1: "1", 2: "10", 3: "11", 4: "100",
	5: "101", 6: "110", 7: "111", 8: "1000", 9: "1001",
	10: "1010", 11: "1011", 12: "1100", 13: "1101", 14: "1110",
	15: "1111", 16: "10000", 17: "10001", 18: "10010", 19: "10011",
	20: "10100", 21: "10101", 22: "10110", 23: "10111", 24: "11000",
	25: "11001", 26: "11010", 27: "11011", 28: "11100", 29: "11101",
	30: "11110", 31: "11111", 32: "100000", 33: "100001", 34: "100010",
	35: "100011", 36: "100100", 37: "100101", 38: "100110", 39: "100111",
	40: "101000", 41: "101001", 42: "101010", 43: "101011", 44: "101100",
	45: "101101", 46: "101110", 47: "101111", 48: "110000", 49: "110001",
	50: "110010", 51: "110011", 52: "110100", 53: "110101", 54: "110110",
	55: "110111", 56: "111000", 57: "111001", 58: "111010", 59: "111011",
	60: "111100", 61: "111101", 62: "111110", 63: "111111", 64: "1000000",
	65: "1000001", 66: "1000010", 67: "1000011", 68: "1000100", 69: "1000101",
	70: "1000110", 71: "1000111", 72: "1001000", 73: "1001001", 74: "1001010",
	75: "1001011", 76: "1001100", 77: "1001101", 78: "1001110", 79: "1001111",
	80: "1010000", 81: "1010001", 82: "1010010", 83: "1010011", 84: "1010100",
	85: "1010101", 86: "1010110", 87: "1010111", 88: "1011000", 89: "1011001",
	90: "1011010", 91: "1011011", 92: "1011100", 93: "1011101", 94: "1011110",
	95: "1011111", 96: "1100000", 97: "1100001", 98: "1100010", 99: "1100011",
	100: "1100100",
}

func TestConvert(t *testing.T) {
	for input, expected := range decimalTestCases {
		out, err := Convert(input)
		if err != nil {
			t.Errorf("Convert(%d) returned an error %s", input, err.Error())
		}
		if out != expected {
			t.Errorf("Convert(%d) = %s; want %s", input, out, expected)
		}
	}
}

func BenchmarkConvert(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = Convert(100)
	}
}
