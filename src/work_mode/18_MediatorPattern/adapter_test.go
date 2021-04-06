package adapter

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("audio test:", AudioPlayerTest)
}

func AudioPlayerTest(t *testing.T){
	audio := NewAudioPlayer()
	audio.Play("mp4","yellow.mp4")
	audio.Play("rmvb","yellow.rmvb")
	audio.Play("mp3","yellow.mp3")
	audio.Play("avi","yellow.avi")
}
