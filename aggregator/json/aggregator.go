package json

import (
    k "github.com/aws/aws-sdk-go/service/kinesis"
    "sync"
    // "log"
)

/*
This is a special use case aggregator that the Kinesis Analytics understands for JSON.
It's just the append of each json string to the previous one in a buffer
*/

type Partition struct {
    partitionKey string
    data []byte
    nRecords int
    nBytes int
}

type Aggregator struct {
    sync.Mutex
    parts map[string] *Partition

}

// Size return how many bytes stored in the aggregator.
// including partition keys.
func (a *Aggregator) Size(partitionKey string) int {
    a.Lock()
    defer a.Unlock()

    if a.parts == nil {
        return 0
    }
    if partitionKey == "" {
        cnt := 0
        for k, part := range a.parts {
            if part.nBytes > 0 {
                cnt += part.nBytes + len(k)
            }
        }
        return cnt
    }

    part, ok := a.parts[partitionKey]
    if ok {
        return part.nBytes + len(partitionKey)
    }

    return 0
}

// Count return how many records stored in the aggregator.
// A key of "" returns total count
func (a *Aggregator) Count(partitionKey string) int {
    a.Lock()
    defer a.Unlock()

    if a.parts == nil {
        return 0
    }
    if partitionKey == "" {
        cnt := 0
        for _, part := range a.parts {
            cnt += part.nRecords
        }
        return cnt
    }

    part, ok := a.parts[partitionKey]
    if ok {
        return part.nRecords
    }

    return 0
}

// Put record using `data` and `partitionKey`. This method is thread-safe.
func (a *Aggregator) Put(data []byte, partitionKey string) {
    a.Lock()
    defer a.Unlock()

    if a.parts == nil {
        a.parts = make(map[string] *Partition)
    }

    var (
        part *Partition
        ok bool
        )

    safeData := make([]byte, len(data))
    copy(safeData, data)

    part, ok = a.parts[partitionKey]
    if !ok {
        part = &Partition{
            partitionKey: partitionKey,
            data: safeData,
            nRecords: 1,
            nBytes: len(data),
        }
        a.parts[partitionKey] = part
    } else {
        if part.data == nil {
            part.data = safeData
            part.nRecords = 1
            part.nBytes = len(data)

        } else {
            part.data = append(part.data, safeData...)
            part.nRecords++
            part.nBytes += len(data)
        }
    }
}

func (a *Aggregator) Drain(partitionKey string) (records []*k.PutRecordsRequestEntry, err error) {

    a.Lock()
    defer a.Unlock()

    err = nil

    if a.parts == nil {
        return
    }

    if partitionKey == "" {
        for _, part := range a.parts {

            if part.data != nil {

                entry := &k.PutRecordsRequestEntry{
                    Data:         part.data,
                    PartitionKey: &part.partitionKey,
                }
                // log.Printf("entry data is %s", string(entry.Data[:]))
                records = append(records, entry)
                part.data = nil
                part.nRecords = 0
                part.nBytes = 0
                // log.Printf("Drain `%s` key %s data %s", partitionKey, key, entry.Data)

            }
        }

    } else {
        part, ok := a.parts[partitionKey]
        if ok && part.data != nil {
            key := part.partitionKey
            entry := &k.PutRecordsRequestEntry{
                Data:         part.data,
                PartitionKey: &key,
            }
            records = append(records, entry)
            part.data = nil
            part.nRecords = 0
            part.nBytes = 0

            // log.Printf("Drain `%s` key %s data %s", partitionKey, key, entry.Data)
        }
    }



    return
}

/*
It's costly to test isAggregated from this format, so just return false
*/
func (a *Aggregator) IsAggregated(entry *k.PutRecordsRequestEntry) bool {
    return false
}

func (a *Aggregator) ExtractRecords(entry *k.PutRecordsRequestEntry) (out []*k.PutRecordsRequestEntry) {
    out = append(out, entry)
    return out
}
