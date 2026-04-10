package main

import (
	"fmt"
	"sync"
	"time"
)

type taskConfig struct{
	Name string
	Rating int
	Err error
}

func worker(url string, wg *sync.WaitGroup,resultChan chan taskConfig) {
	defer wg.Done()
	time.Sleep(time.Millisecond*50)
	// fmt.Printf("image processed: %s\n",url)
	result:=taskConfig {
		Name: url,
		Rating: 2,
		Err : nil,
	}
	resultChan <- result
}

func main() {
var wg sync.WaitGroup
resultChannel:=make(chan taskConfig,5) 
startTime:=time.Now()

wg.Add(5)
go worker("task 1",&wg,resultChannel)
go worker("task 2",&wg,resultChannel)
go worker("task 3",&wg,resultChannel)
go worker("task 4",&wg,resultChannel)
go worker("task 5",&wg,resultChannel)
wg.Wait()
close(resultChannel)
for res:=range resultChannel{
	fmt.Printf("Data received %v\n ",res)
	if res.Err!=nil{
		// You can setup here a dead letter queue or any other logic
	}
}
fmt.Printf("Time taken: %s  ",time.Since(startTime))

}





/*
FAN OUT MODEL --> When multiple workers are created or multiple goRoutine created and each goRoutine has assignes some work , a kind of consumer.
FAN IN MODEL  --> When multiple workers gives some data ,and we need to aggregate these data
                 ex: suppose there are 3 goRoutine all are fecting some info form the DB , and all will return their data , and we need to aggregate these data in our main fuction
                 For this one we need a channel 
				CHANNEL (act as a queue)
				If we want to send some data from a goRoutine to another goRoutine or main routine then , we need to use the CHANNEL
				NOTE: Whenever you careate any channel then simply set a close of the channel

				You can create multiple channels for the different goRoutines workers. 
				Channel:
				A) Write --> Inside the worker you will write in the channel , like  channel_name <- {data you want to send}
				B) Read --> Read by any one by the channel name , simple used by the consumer function (consumer might be run concurrently )


				Full Notes:
				A) Consumer --> a function which consumes some data form the channel or any queue
				B) Producer --> a function which produces or reads some data like from the database or the any file and do some process and then send/write to the channel 
				
				Consumer and Producers might run on the multiple goRoutines. This is based on the requirement.

*/