package json

import (
    k "github.com/aws/aws-sdk-go/service/kinesis"
)

type Partition struct {
    partitionKey string
    data []byte
    nRecords int
    nBytes int
}

type Aggregator struct {
    parts map[string] *Partition
}

// Size return how many bytes stored in the aggregator.
// including partition keys.
func (a *Aggregator) Size(partitionKey string) int {
    return 0
}

// Count return how many records stored in the aggregator.
func (a *Aggregator) Count(partitionKey string) int {
    return 0
}

// Put record using `data` and `partitionKey`. This method is thread-safe.
func (a *Aggregator) Put(data []byte, partitionKey string) {
}

func (a *Aggregator) Drain(partitionKey string) (records []*k.PutRecordsRequestEntry, err error) {
    key := "test"
    data := []byte("{alpha}")
    entry := &k.PutRecordsRequestEntry{
        Data:         data,
        PartitionKey: &key,
    }
    records = append(records, entry)
    err = nil
    return
}

func (a *Aggregator) IsAggregated(entry *k.PutRecordsRequestEntry) bool {
    return true
}

func (a *Aggregator) ExtractRecords(entry *k.PutRecordsRequestEntry) (out []*k.PutRecordsRequestEntry) {
    out = append(out, entry)
    return out
}
