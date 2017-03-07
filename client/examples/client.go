package main

import (
	"log"
	"sync"

	"github.com/g8os/tlog/client"
)

func main() {
	var volID uint32 = 0x1f
	num_conn := 2
	clients := make([]*client.Client, num_conn)

	for i := 0; i < num_conn; i++ {
		client, err := client.New("127.0.0.1:11211")
		if err != nil {
			log.Fatal(err)
		}
		clients[i] = client
	}
	data := make([]byte, 4096*4)

	var wg sync.WaitGroup

	for i := 0; i < 30; i++ {
		wg.Add(1)
		act := func(j int) {
			defer wg.Done()
			client := clients[j%num_conn]

			err := client.Send(volID, uint64(j), uint32(j), uint64(j), uint64(j), data)
			if err != nil {
				log.Fatalf("failed to send :%v", err)
			}

		}
		//go act(i)
		act(i)
	}
	wg.Wait()
}
