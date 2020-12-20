package Structural

import "fmt"

type AdvancedPlayer interface {
	PlayVLC(string) string
	PlayMP4(string) string
}

type MediaPlayer interface {
	Play(mediaType string, mediaName string) string
}

type VLCPlayer struct {
}

type MP4Player struct {
}

func (vlc *VLCPlayer) PlayVLC(filename string) string {
	return fmt.Sprintf("Play VLC %s", filename)
}

func (vlc *VLCPlayer) PlayMP4(filename string) string {
	return ""
}

func (mp4 *MP4Player) PlayMP4(filename string) string {
	return fmt.Sprintf("Play MP4 %s", filename)
}

func (mp4 *MP4Player) PlayVLC(filename string) string {
	return fmt.Sprintf("Play MP4 %s", filename)
}

type MediaAdapter struct {
	mp4 *MP4Player
	vlc *VLCPlayer
}

func NewMediaAdapter() *MediaAdapter {
	return &MediaAdapter{
		mp4: &MP4Player{},
		vlc: &VLCPlayer{},
	}
}

func (adapter *MediaAdapter) Play(MediaType string, MediaName string) string {
	if MediaType == "mp4" {
		return adapter.mp4.PlayMP4(MediaName)
	}

	if MediaType == "vlc" {
		return adapter.vlc.PlayVLC(MediaName)
	}

	return ""
}
