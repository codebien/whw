package magic

import "testing"

func TestAbracadabra(t *testing.T) {
	sut := Abracadabra()
	expected := "Wonderful Hello World!"
	if sut != expected {
		t.Errorf("expected %s instead got: %s", expected, sut)
	}
}
