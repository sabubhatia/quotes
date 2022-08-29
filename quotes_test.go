package quotes

import "testing"

func TestFav(t *testing.T) {

	wl := 3
	got := Favs()
	if len(got) != wl {
		t.Errorf("Got %d. Want %d\n", len(got), wl)
	}
}