
package main

import (
	"fmt"
	"time"
	"math/rand"
)

type BirdLocation struct {
	latitude     float64
	longitude    float64
}

type BirdTrack struct {
	id 				int
	speed 			int // mph
	location 		BirdLocation
}


func main() {
	logChannel := make(chan BirdTrack, 10)

	go birdTracksGenerator(logChannel)
	go birdTracksSpeedController(logChannel)
	time.Sleep(time.Second * 20)
	close(logChannel)
}

func birdTracksSpeedController(logChannel <-chan BirdTrack) {
	for bird := range logChannel {
		if bird.speed > 10 {
			bird.PrintOverSpeedingAlert()
		} 
	}
}

func birdTracksGenerator (logChannel chan<- BirdTrack) {
	for i:=1 ; true ; i++ {
		newBird := BirdTrack {
			id: i,
			speed: genRandomSpeed(),
			location: BirdLocation{
				latitude: generateLatitude(),
				longitude: generateLongitude(),
			},
		}
		time.Sleep(time.Millisecond * 1000)
		logChannel <- newBird
	}
}

func genRandomSpeed() int{
	rand.Seed(time.Now().UnixNano())
    min := 2
    max := 18
	speed := rand.Intn(max - min + 1) + min
	return speed
}

func generateLatitude() float64 {
	rnd := rand.Float64()
    return rnd
}
var generateLongitude = generateLatitude

func (bird *BirdTrack) PrintOverSpeedingAlert () {
	fmt.Printf("[ALERT] -- The BirdTrack with ID: %v is going at %vmph\n", bird.id, bird.speed)
}