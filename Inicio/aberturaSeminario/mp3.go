package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	// Abra o arquivo MP3
	f, err := os.Open("/caminho/para/audio/GoPowerRangers.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Decodifique o arquivo MP3
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Inicialize o alto-falante
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Reproduza o áudio
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	// Aguarde até que o áudio termine de tocar
	<-done
	fmt.Println("Áudio reproduzido com sucesso!")
}

