package producer

import (
k "github.com/aws/aws-sdk-go/service/kinesis"
)

type Aggregator interface {
// Put record using `data` and `partitionKey`. This method is thread-safe.
	Put(data []byte, partitionKey string)
// Drain create an aggregated `kinesis.PutRecordsRequestEntry` list
	Drain(partitionKey string) (records []*k.PutRecordsRequestEntry, err error)
// Count return how many records stored in the aggregator.
	Count(partitionKey string) int
// Size return how many bytes stored in the aggregator.
// including partition keys.
	Size(partitionKey string) int
// Return extracted records
	ExtractRecords(entry *k.PutRecordsRequestEntry) (out []*k.PutRecordsRequestEntry)

	IsAggregated(entry *k.PutRecordsRequestEntry) bool
}
