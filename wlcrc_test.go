package wlcrc

import "testing"


type tester struct {
  value byte
  want uint16
}

func TestCrc(t *testing.T) {
	crc := New()

	if crc.Sum16() != 0 {
		t.Error("Acc Start not 0")
	}

  testers := []tester{
    tester{0xc6, 0xb98a},
    tester{0xce, 0x8470},
    tester{0xa2, 0x34a4},
    tester{0x03, 0xe2b4},
    tester{0xe2, 0xb400},
    tester{0xb4, 0x0000},
  }

  for _, test := range testers {
    crc.Write([]byte{test.value})
    if want, got := uint16(test.want), crc.Sum16(); want != got {
      t.Errorf("Not Valid checksum, got %v, want %v", got, want)
    }
  }
}
