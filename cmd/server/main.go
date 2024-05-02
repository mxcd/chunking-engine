package main

import (
	"context"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/mxcd/chunking-engine/ent"
	"github.com/mxcd/chunking-engine/ent/chunk"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vmihailenco/msgpack"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:sqlite.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	http.HandleFunc("/get", getChunks(client))
	go http.ListenAndServe(":8080", nil)

	for {
		startTime := time.Now()
		iData, dData := generateData()

		iDataChunk, err := msgpack.Marshal(iData)
		if err != nil {
			log.Fatalf("failed to marshal iData: %v", err)
		}
		log.Default().Printf("iDataChunk size: %d", len(iDataChunk))
		_, err = client.Chunk.Create().SetT(time.Now()).SetName("i").SetData(iDataChunk).Save(ctx)
		if err != nil {
			log.Fatalf("failed to save iData: %v", err)
		}

		dDataChunk, err := msgpack.Marshal(dData)
		if err != nil {
			log.Fatalf("failed to marshal dData: %v", err)
		}
		client.Chunk.Create().SetT(time.Now()).SetName("d").SetData(dDataChunk).Save(ctx)

		log.Printf("Time taken: %v", time.Since(startTime))

		time.Sleep(time.Second * 1)
	}
}

func generateData() ([]uint8, []uint8) {
	var iData = make([]uint8, 2000000)
	var dData = make([]uint8, 2000000)

	for i := 0; i < 2000000; i++ {
		iData[i] = uint8(math.Sin(float64(i)) * 100)
		dData[i] = uint8(math.Cos(float64(i)) * 100)
	}

	return iData, dData
}

func getChunks(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx := r.Context()

		startTimestamp := time.Now().Add(-time.Minute)

		iDataChunk, err := client.Chunk.Query().Where(chunk.And(chunk.TGT(startTimestamp), chunk.Name("i"))).Order(ent.Asc(chunk.FieldT)).All(ctx)
		if err != nil {
			log.Fatalf("failed to get iData: %v", err)
		}

		// content disposition
		w.Header().Set("Content-Disposition", "attachment; filename=iData.msgpack")
		w.Header().Set("Content-Type", "octet/stream")
		for _, chunk := range iDataChunk {
			w.Write(chunk.Data)
		}

		log.Default().Printf("Served request in: %v", time.Since(startTime))
	}
}
