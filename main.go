
package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
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

const BIRDS = 100
var wg = sync.WaitGroup{}

func main() {
	birdSpeedControllerChannel := make(chan BirdTrack, 2)

	wg.Add(2)
	go birdTracksGenerator(birdSpeedControllerChannel)
	go birdTracksSpeedController(birdSpeedControllerChannel)
	wg.Wait()
}

func birdTracksSpeedController(birdSpeedControllerChannel <-chan BirdTrack) {
	for bird := range birdSpeedControllerChannel {
		if bird.speed > 10 {
			bird.PrintOverSpeedingAlert()
		} 
	}
	wg.Done()
}

func birdTracksGenerator (birdSpeedControllerChannel chan<- BirdTrack) {
	for i:=1 ; i < BIRDS ; i++ {
		newBird := BirdTrack {
			id: i,
			speed: genRandomSpeed(),
			location: BirdLocation{
				latitude: generateLatitude(),
				longitude: generateLongitude(),
			},
		}
		time.Sleep(time.Millisecond * 1000)
		birdSpeedControllerChannel <- newBird
	}
	close(birdSpeedControllerChannel)
	wg.Done()
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