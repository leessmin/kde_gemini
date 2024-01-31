package notice

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func PlayError() {
	f, err := os.Open("./resource/voice/violin-lose-4-185125.mp3")
	if err != nil {
		log.Print("播放声音失败:", err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Print("播放声音失败:", err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
