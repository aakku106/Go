package try5

import "testing"

/*
This entire testing file is useless and un-profectional,
But than Why did i made one and using it
Because Go test automatically shutdown/fail in around ~600 sec.
I use TMUX and MAC so i never close my mac and tmux never kills my servers so i
intensionally made this test fiel which autokills itself in ~600 sec its saves my CPU,Memory and most imp battery
*/
func TestDummy(t *testing.T) {
	InitilizeServer4()
}
