/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package xaddr

import (
	"testing"
)

func TestNewFromB5_AllZeroes_ReturnsAllZeroes(t *testing.T) {
	expected := [5]byte{0x00, 0x00, 0x00, 0x00, 0x00}

	actual := NewFromB5([5]byte{0x00, 0x00, 0x00, 0x00, 0x00})

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewFromB5_AllOnes_ReturnsAllOnes(t *testing.T) {
	expected := [5]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	actual := NewFromB5([5]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF})

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewFromB5_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := [5]byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}

	actual := NewFromB5([5]byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA})

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewFromB_AllZeroes_ReturnsAllZeroes(t *testing.T) {
	expected := [5]byte{0x00, 0x00, 0x00, 0x00, 0x00}

	actual := NewFromB(expected[:])

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewFromB_AllOnes_ReturnsAllOnes(t *testing.T) {
	expected := [5]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	actual := NewFromB(expected[:])

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewFromB_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := [5]byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}

	actual := NewFromB(expected[:])

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewFromB_LessThanFiveBytes_SetLSBytes(t *testing.T) {
	expected := [5]byte{0x00, 0x00, 0xAA, 0xBB, 0xCC}
	
	actual := NewFromB([]byte{0xAA, 0xBB, 0xCC})

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	} 
}

func TestNewFromB_MoreThanFiveBytes_UsesOnlyLastFiveBytes(t *testing.T) {
	expected := [5]byte{0xBB, 0xCC, 0xDD, 0xEE, 0xFF} // 5 bytes

	actual := NewFromB([]byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}) // 6 bytes

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewFromI_AllZeroes_ReturnsAllZeroes(t *testing.T) {
	expected := [5]byte{0x00, 0x00, 0x00, 0x00, 0x00}

	actual := NewFromI(0x0000000000)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewFromI_AllOnes_ReturnsAllOnes(t *testing.T) {
	expected := [5]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	actual := NewFromI(0xFFFFFFFFFF)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewFromI_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := [5]byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}

	actual := NewFromI(0xAAAAAAAAAA)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestNewFromI_LessThanFiveBytes_SetLSBytes(t *testing.T) {
	expected := [5]byte{0x00, 0x00, 0xAA, 0xBB, 0xCC}
	
	actual := NewFromI(0xAABBCC)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	} 
}

func TestNewFromI_MoreThanFiveBytes_UsesFiveLSBytes(t *testing.T) {
	expected := [5]byte{0xBB, 0xCC, 0xDD, 0xEE, 0xFF} // 5 bytes

	actual := NewFromI(0xAABBCCDDEEFF) // 6 bytes

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestToB5_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := [5]byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}
	a := NewFromB5(expected)

	actual := a.ToB5()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestToI_MixedOnesAndZeroes_ReturnsRightBits(t *testing.T) {
	expected := uint64(0xAAAAAAAAAA)
	a := NewFromI(expected)
	
	actual := a.ToI()

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestIToB5_Zeroes_ReturnsZeroes(t *testing.T) {
	expected := [5]byte{0, 0, 0, 0, 0}

	actual := iToB5(0x0000000000)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestIToB5_Ones_ReturnsMaxBytes(t *testing.T) {
	expected := [5]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	actual := iToB5(0xFFFFFFFFFF)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestIToB5_MixedZeroesAndOnes_ReturnsRightBytes(t *testing.T) {
	expected := [5]byte{0xBB, 0xCC, 0xDD, 0xEE, 0xFF}

	actual := iToB5(0xBBCCDDEEFF)

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestIToB5_Overflow5Bytes_ReturnsTruncated(t *testing.T) {
	expected := [5]byte{0xBB, 0xCC, 0xDD, 0xEE, 0xFF} // 5 bytes

	actual := iToB5(0xAABBCCDDEEFF) // 6 bytes

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestB5ToI_Zeroes_ReturnsZeroes(t *testing.T) {
	expected := uint64(0x0000000000)

	actual := b5ToI([5]byte{0, 0, 0, 0, 0})

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestB5ToI_Ones_ReturnsMaxBytes(t *testing.T) {
	expected := uint64(0xFFFFFFFFFF)

	actual := b5ToI([5]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF})

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}

func TestB5ToI_MixedZeroesAndOnes_ReturnsRightBytes(t *testing.T) {
	expected := uint64(0xAABBCCDDEE)

	actual := b5ToI([5]byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE})

	if actual != expected {
		t.Errorf("expected '%b' but found '%b'", expected, actual)
	}
}
