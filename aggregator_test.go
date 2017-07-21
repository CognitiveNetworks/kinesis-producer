package producer

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"github.com/CognitiveNetworks/kinesis-producer/aggregator/kpl"
	"github.com/CognitiveNetworks/kinesis-producer/aggregator/json"
	"log"
)

func assert(t *testing.T, val bool, msg string) {
	if !val {
		t.Error(msg)
	}
}

func TestKPLSizeAndCount(t *testing.T) {
	a := new(kpl.Aggregator)
	assert(t, a.Size("")+a.Count("") == 0, "size and count should equal to 0 at the beginning")
	data := []byte("hello")
	pkey := "world"
	n := rand.Intn(100)
	for i := 0; i < n; i++ {
		a.Put(data, pkey)
	}
	assert(t, a.Size(pkey) == 5*n+5, "size should equal to the data and the partition-key")
	assert(t, a.Count(pkey) == n, "count should be equal to the number of Put calls")
}

func TestKPLAggregation(t *testing.T) {
	var wg sync.WaitGroup
	a := new(kpl.Aggregator)
	n := 50
	wg.Add(n)
	for i := 0; i < n; i++ {
		c := strconv.Itoa(i)
		data := []byte("hello-" + c)
		a.Put(data, c)
		wg.Done()
	}
	wg.Wait()
	records, err := a.Drain("")
	if err != nil {
		t.Error(err)
	}
	for _, record := range records {
		assert(t, a.IsAggregated(record), "should return an agregated record")
		records := a.ExtractRecords(record)
		for i := 0; i < n; i++ {
			c := strconv.Itoa(i)
			found := false
			for _, record := range records {
				if string(record.Data) == "hello-"+c {
					assert(t, string(record.Data) == "hello-"+c, "`Data` field contains invalid value")
					found = true
				}
			}
			assert(t, found, "record not found after extracting: "+c)
		}
	}
}

func TestJSONSizeAndCount(t *testing.T) {
	a := new(json.Aggregator)
	assert(t, a.Size("")+a.Count("") == 0, "size and count should equal to 0 at the beginning")
	data := []byte("hello")
	pkey := "world"
	n := rand.Intn(100)
	for i := 0; i < n; i++ {
		a.Put(data, pkey)
	}
	assert(t, a.Size(pkey) == 5*n+5, "size should equal to the data and the partition-key")
	assert(t, a.Count(pkey) == n, "count should be equal to the number of Put calls")

	assert(t, a.Size("") == 5*n+5, "size should equal to the data and the partition-key")
	assert(t, a.Count("") == n, "count should be equal to the number of Put calls")
}

func TestJSONAggregation(t *testing.T) {
	var wg sync.WaitGroup
	a := new(json.Aggregator)
	n := 50
	wg.Add(n)
	for i := 0; i < n; i++ {
		c := strconv.Itoa(i)
		data := []byte("hello-" + c)
		a.Put(data, c)
		wg.Done()
	}
	wg.Wait()
	record, err := a.Drain("")
	if err != nil {
		t.Error(err)
	}

	log.Printf("records is %v", record)
}
