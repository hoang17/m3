// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package tracepoint

// The tracepoint package is used to store operation names for tracing throughout dbnode.
// The naming convention is as follows:

// `packageName.objectName.method`

// If there isn't an object, use `packageName.method`.

const (
	// FetchTagged is the operation name for the tchannelthrift FetchTagged path.
	FetchTagged = "tchannelthrift/node.service.FetchTagged"

	// Query is the operation name for the tchannelthrift Query path.
	Query = "tchannelthrift/node.service.Query"

	// FetchReadEncoded is the operation name for the tchannelthrift FetchReadEncoded path.
	FetchReadEncoded = "tchannelthrift/node.service.FetchReadEncoded"

	// FetchReadResults is the operation name for the tchannelthrift FetchReadResults path.
	FetchReadResults = "tchannelthrift/node.service.FetchReadResults"

	// FetchReadSingleResult is the operation name for the tchannelthrift FetchReadSingleResult path.
	FetchReadSingleResult = "tchannelthrift/node.service.FetchReadSingleResult"

	// FetchReadSegment is the operation name for the tchannelthrift FetchReadSegment path.
	FetchReadSegment = "tchannelthrift/node.service.FetchReadSegment"

	// DBQueryIDs is the operation name for the db QueryIDs path.
	DBQueryIDs = "storage.db.QueryIDs"

	// DBAggregateQuery is the operation name for the db AggregateQuery path.
	DBAggregateQuery = "storage.db.AggregateQuery"

	// DBReadEncoded is the operation name for the db ReadEncoded path.
	DBReadEncoded = "storage.db.ReadEncoded"

	// DBFetchBlocks is the operation name for the db FetchBlocks path.
	DBFetchBlocks = "storage.db.FetchBlocks"

	// DBFetchBlocksMetadataV2 is the operation name for the db FetchBlocksMetadataV2 path.
	DBFetchBlocksMetadataV2 = "storage.db.FetchBlocksMetadataV2"

	// DBWriteBatch is the operation name for the db WriteBatch path.
	DBWriteBatch = "storage.db.WriteBatch"

	// NSQueryIDs is the operation name for the dbNamespace QueryIDs path.
	NSQueryIDs = "storage.dbNamespace.QueryIDs"

	// NSIdxQuery is the operation name for the nsIndex Query path.
	NSIdxQuery = "storage.nsIndex.Query"

	// NSIdxAggregateQuery is the operation name for the nsIndex AggregateQuery path.
	NSIdxAggregateQuery = "storage.nsIndex.AggregateQuery"

	// NSIdxQueryHelper is the operation name for the nsIndex query path.
	NSIdxQueryHelper = "storage.nsIndex.query"

	// NSIdxBlockQuery is the operation name for the nsIndex block query path.
	NSIdxBlockQuery = "storage.nsIndex.blockQuery"

	// NSIdxBlockAggregateQuery is the operation name for the nsIndex block aggregate query path.
	NSIdxBlockAggregateQuery = "storage.nsIndex.blockAggregateQuery"

	// BlockQuery is the operation name for the index block query path.
	BlockQuery = "storage/index.block.Query"

	// BlockAggregate is the operation name for the index block aggregate path.
	BlockAggregate = "storage/index.block.Aggregate"
)
