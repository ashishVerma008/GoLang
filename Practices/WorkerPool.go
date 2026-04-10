package main

import (
	"fmt"
	"sync"
	"time"
)


var items = []string{
	"item1", "item2", "item3", "item4", "item5",
	"item6", "item7", "item8", "item9", "item10",
	"item11", "item12", "item13", "item14", "item15",
	"item16", "item17", "item18", "item19", "item20",
}

func worker(jobsChannel chan string, wg *sync.WaitGroup,resultChan chan string) {
	
	defer wg.Done()
	for job:= range jobsChannel{                            // As long as the channel is open the worker will listen the job 
	    time.Sleep(time.Millisecond*50)
		fmt.Printf("Job done %s\n",job)
		resultChan<- job
	}
	fmt.Printf("Worker Shut down.\n")
	
}

func main() {
var wg sync.WaitGroup
totalWorkers:=5
resultChannel:=make(chan string,50) 
jobsChannel:=make(chan string,len(items))
startTime:=time.Now()

// 1) Create the workers
for i:=1;i<=totalWorkers;i++ {
	wg.Add(1)
	go worker(jobsChannel , &wg, resultChannel)
}

// ask the wait logic to wait on some other goRoutine
go func(){
	wg.Wait()
	close(resultChannel)
}()

// Send the jobs to the jobChannel
for i:=0;i<len(items);i++ {
	jobsChannel<- items[i]
}
close(jobsChannel)

for result := range resultChannel {	
		_ = result 
}
	
fmt.Printf("Time taken: %s  ",time.Since(startTime))
}


/*


*/