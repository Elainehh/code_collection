package Structural

import "testing"

func TestAdapter(t *testing.T) {
	MAdapter := NewMediaAdapter()
	res := MAdapter.Play("mp4", "mp4 file")
	if res != "Play MP4 mp4 file" {
		t.Errorf("Wrong MP4 play: %s", res)
	}

	res = MAdapter.Play("vlc", "vlc file")
	if res != "Play VLC vlc file" {
		t.Errorf("Wrong VLC play: %s", res)
	}
}
