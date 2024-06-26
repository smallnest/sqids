package sqids

import (
	"reflect"
	"testing"
)

func TestMultibyteAlphabet(t *testing.T) {
	_, err := New(WithAlphabet("ë1092"))

	if err != errAlphabetMultibyte {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAlphabetSimple(t *testing.T) {
	numbers := []uint64{1, 2, 3}
	id := "489158"

	alphabet := "0123456789abcdef"
	s, err := New(WithAlphabet(alphabet))
	if err != nil {
		t.Fatal(err)
	}

	generatedID, err := s.Encode(numbers...)
	if err != nil {
		t.Fatal(err)
	}

	if id != generatedID {
		t.Errorf("Encoding `%v` should produce `%v`, but instead produced `%v`", numbers, id, generatedID)
	}

	decodedNumbers := s.Decode(generatedID)
	if !reflect.DeepEqual(numbers, decodedNumbers) {
		t.Errorf("Decoding `%v` should produce `%v`, but instead produced `%v`", id, numbers, decodedNumbers)
	}
}

func TestAlphabetShort(t *testing.T) {
	s, err := New(WithAlphabet("abc"))
	if err != nil {
		t.Fatal(err)
	}

	numbers := []uint64{1, 2, 3}

	generatedID, err := s.Encode(numbers...)
	if err != nil {
		t.Fatal(err)
	}

	decodedNumbers := s.Decode(generatedID)
	if !reflect.DeepEqual(numbers, decodedNumbers) {
		t.Errorf("Decoding `%v` should produce `%v`, but instead produced `%v`", generatedID, numbers, decodedNumbers)
	}
}

func TestAlphabetLong(t *testing.T) {
	s, err := New(WithAlphabet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+|{}[];:'\"/?.>,<`~"))

	if err != nil {
		t.Fatal(err)
	}

	numbers := []uint64{1, 2, 3}

	generatedID, err := s.Encode(numbers...)
	if err != nil {
		t.Fatal(err)
	}

	decodedNumbers := s.Decode(generatedID)
	if !reflect.DeepEqual(numbers, decodedNumbers) {
		t.Errorf("Decoding `%v` should produce `%v`, but instead produced `%v`", generatedID, numbers, decodedNumbers)
	}
}

func TestRepeatingAlphabetCharacters(t *testing.T) {
	if _, err := New(WithAlphabet("aabcdefg")); err == nil {
		t.Errorf("Should not accept alphabet with repeating characters")
	}
}

func TestTooShortOfAnAlphabet(t *testing.T) {
	if _, err := New(WithAlphabet("ab")); err == nil {
		t.Errorf("Should not accept too short of an alphabet")
	}
}
