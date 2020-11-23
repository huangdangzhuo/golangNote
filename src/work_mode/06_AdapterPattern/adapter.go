package adapter

// MediaPlayer 多媒体播放器
type MediaPlayer interface {
	Play(audioType,fileName string)
}

// AdvanceMediaPlayer 新多媒体播放器接口
type AdvanceMediaPlayer interface {
	PlayRmvb()
}

