package main

import (
	"log"

	"github.com/g8os/tlog/client"
)

func main() {
	var volID uint32 = 0x1f
	num_conn := 2
	clients := make([]*client.Client, num_conn)
	clientReady := make(chan int, num_conn)
	seqChan := make(chan uint64, 8)

	for i := 0; i < num_conn; i++ {
		client, err := client.New("127.0.0.1:11211")
		if err != nil {
			log.Fatal(err)
		}
		clients[i] = client
		clientReady <- i
	}
	data := make([]byte, 4096*4)

	// produce the data
	go func() {
		for i := 0; i < 25*4000; i++ {
			seqChan <- uint64(i)
		}
	}()

	for seq := range seqChan {
		idx := <-clientReady
		go func(j uint64, idx int) {
			client := clients[int(j)%num_conn]

			log.Printf("j=%v\n", j)
			err := client.Send(volID, j, uint32(j), j, j, data)
			if err != nil {
				log.Printf("client %v died\n", idx)
				return
			}
			clientReady <- idx
		}(seq, idx)
	}
}
