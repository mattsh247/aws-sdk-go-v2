// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Describes the timestamp range and timestamp origin of a range of fragments.
// Fragments that have duplicate producer timestamps are deduplicated. This means
// that if producers are producing a stream of fragments with producer timestamps
// that are approximately equal to the true clock time, the clip will contain all
// of the fragments within the requested timestamp range. If some fragments are
// ingested within the same time range and very different points in time, only the
// oldest ingested collection of fragments are returned.
type ClipFragmentSelector struct {
	
	// The origin of the timestamps to use (Server or Producer).
	//
	// This member is required.
	FragmentSelectorType ClipFragmentSelectorType
	
	// The range of timestamps to return.
	//
	// This member is required.
	TimestampRange *ClipTimestampRange
	
	noSmithyDocumentSerde
}

// The range of timestamps for which to return fragments.
type ClipTimestampRange struct {
	
	// The end of the timestamp range for the requested media. This value must be
	// within 24 hours of the specified StartTimestamp , and it must be later than the
	// StartTimestamp value. If FragmentSelectorType for the request is
	// SERVER_TIMESTAMP , this value must be in the past. This value is inclusive. The
	// EndTimestamp is compared to the (starting) timestamp of the fragment. Fragments
	// that start before the EndTimestamp value and continue past it are included in
	// the session.
	//
	// This member is required.
	EndTimestamp *time.Time
	
	// The starting timestamp in the range of timestamps for which to return
	// fragments. Only fragments that start exactly at or after StartTimestamp are
	// included in the session. Fragments that start before StartTimestamp and
	// continue past it aren't included in the session. If FragmentSelectorType is
	// SERVER_TIMESTAMP , the StartTimestamp must be later than the stream head.
	//
	// This member is required.
	StartTimestamp *time.Time
	
	noSmithyDocumentSerde
}

// Contains the range of timestamps for the requested media, and the source of the
// timestamps.
type DASHFragmentSelector struct {
	
	// The source of the timestamps for the requested media. When FragmentSelectorType
	// is set to PRODUCER_TIMESTAMP and GetDASHStreamingSessionURLInput$PlaybackMode
	// is ON_DEMAND or LIVE_REPLAY , the first fragment ingested with a producer
	// timestamp within the specified FragmentSelector$TimestampRange is included in
	// the media playlist. In addition, the fragments with producer timestamps within
	// the TimestampRange ingested immediately following the first fragment (up to the
	// GetDASHStreamingSessionURLInput$MaxManifestFragmentResults value) are included.
	// Fragments that have duplicate producer timestamps are deduplicated. This means
	// that if producers are producing a stream of fragments with producer timestamps
	// that are approximately equal to the true clock time, the MPEG-DASH manifest will
	// contain all of the fragments within the requested timestamp range. If some
	// fragments are ingested within the same time range and very different points in
	// time, only the oldest ingested collection of fragments are returned. When
	// FragmentSelectorType is set to PRODUCER_TIMESTAMP and
	// GetDASHStreamingSessionURLInput$PlaybackMode is LIVE , the producer timestamps
	// are used in the MP4 fragments and for deduplication. But the most recently
	// ingested fragments based on server timestamps are included in the MPEG-DASH
	// manifest. This means that even if fragments ingested in the past have producer
	// timestamps with values now, they are not included in the HLS media playlist. The
	// default is SERVER_TIMESTAMP .
	FragmentSelectorType DASHFragmentSelectorType
	
	// The start and end of the timestamp range for the requested media. This value
	// should not be present if PlaybackType is LIVE .
	TimestampRange *DASHTimestampRange
	
	noSmithyDocumentSerde
}

// The start and end of the timestamp range for the requested media. This value
// should not be present if PlaybackType is LIVE . The values in DASHimestampRange
// are inclusive. Fragments that start exactly at or after the start time are
// included in the session. Fragments that start before the start time and continue
// past it are not included in the session.
type DASHTimestampRange struct {
	
	// The end of the timestamp range for the requested media. This value must be
	// within 24 hours of the specified StartTimestamp , and it must be later than the
	// StartTimestamp value. If FragmentSelectorType for the request is
	// SERVER_TIMESTAMP , this value must be in the past. The EndTimestamp value is
	// required for ON_DEMAND mode, but optional for LIVE_REPLAY mode. If the
	// EndTimestamp is not set for LIVE_REPLAY mode then the session will continue to
	// include newly ingested fragments until the session expires. This value is
	// inclusive. The EndTimestamp is compared to the (starting) timestamp of the
	// fragment. Fragments that start before the EndTimestamp value and continue past
	// it are included in the session.
	EndTimestamp *time.Time
	
	// The start of the timestamp range for the requested media. If the
	// DASHTimestampRange value is specified, the StartTimestamp value is required.
	// Only fragments that start exactly at or after StartTimestamp are included in
	// the session. Fragments that start before StartTimestamp and continue past it
	// aren't included in the session. If FragmentSelectorType is SERVER_TIMESTAMP ,
	// the StartTimestamp must be later than the stream head.
	StartTimestamp *time.Time
	
	noSmithyDocumentSerde
}

// Represents a segment of video or other time-delimited data.
type Fragment struct {
	
	// The playback duration or other time value associated with the fragment.
	FragmentLengthInMilliseconds int64
	
	// The unique identifier of the fragment. This value monotonically increases based
	// on the ingestion order.
	FragmentNumber *string
	
	// The total fragment size, including information about the fragment and contained
	// media data.
	FragmentSizeInBytes int64
	
	// The timestamp from the producer corresponding to the fragment.
	ProducerTimestamp *time.Time
	
	// The timestamp from the Amazon Web Services server corresponding to the fragment.
	ServerTimestamp *time.Time
	
	noSmithyDocumentSerde
}

// Describes the timestamp range and timestamp origin of a range of fragments.
// Only fragments with a start timestamp greater than or equal to the given start
// time and less than or equal to the end time are returned. For example, if a
// stream contains fragments with the following start timestamps:
//   - 00:00:00
//   - 00:00:02
//   - 00:00:04
//   - 00:00:06
// A fragment selector range with a start time of 00:00:01 and end time of
// 00:00:04 would return the fragments with start times of 00:00:02 and 00:00:04.
type FragmentSelector struct {
	
	// The origin of the timestamps to use (Server or Producer).
	//
	// This member is required.
	FragmentSelectorType FragmentSelectorType
	
	// The range of timestamps to return.
	//
	// This member is required.
	TimestampRange *TimestampRange
	
	noSmithyDocumentSerde
}

// Contains the range of timestamps for the requested media, and the source of the
// timestamps.
type HLSFragmentSelector struct {
	
	// The source of the timestamps for the requested media. When FragmentSelectorType
	// is set to PRODUCER_TIMESTAMP and GetHLSStreamingSessionURLInput$PlaybackMode is
	// ON_DEMAND or LIVE_REPLAY , the first fragment ingested with a producer timestamp
	// within the specified FragmentSelector$TimestampRange is included in the media
	// playlist. In addition, the fragments with producer timestamps within the
	// TimestampRange ingested immediately following the first fragment (up to the
	// GetHLSStreamingSessionURLInput$MaxMediaPlaylistFragmentResults value) are
	// included. Fragments that have duplicate producer timestamps are deduplicated.
	// This means that if producers are producing a stream of fragments with producer
	// timestamps that are approximately equal to the true clock time, the HLS media
	// playlists will contain all of the fragments within the requested timestamp
	// range. If some fragments are ingested within the same time range and very
	// different points in time, only the oldest ingested collection of fragments are
	// returned. When FragmentSelectorType is set to PRODUCER_TIMESTAMP and
	// GetHLSStreamingSessionURLInput$PlaybackMode is LIVE , the producer timestamps
	// are used in the MP4 fragments and for deduplication. But the most recently
	// ingested fragments based on server timestamps are included in the HLS media
	// playlist. This means that even if fragments ingested in the past have producer
	// timestamps with values now, they are not included in the HLS media playlist. The
	// default is SERVER_TIMESTAMP .
	FragmentSelectorType HLSFragmentSelectorType
	
	// The start and end of the timestamp range for the requested media. This value
	// should not be present if PlaybackType is LIVE .
	TimestampRange *HLSTimestampRange
	
	noSmithyDocumentSerde
}

// The start and end of the timestamp range for the requested media. This value
// should not be present if PlaybackType is LIVE .
type HLSTimestampRange struct {
	
	// The end of the timestamp range for the requested media. This value must be
	// within 24 hours of the specified StartTimestamp , and it must be later than the
	// StartTimestamp value. If FragmentSelectorType for the request is
	// SERVER_TIMESTAMP , this value must be in the past. The EndTimestamp value is
	// required for ON_DEMAND mode, but optional for LIVE_REPLAY mode. If the
	// EndTimestamp is not set for LIVE_REPLAY mode then the session will continue to
	// include newly ingested fragments until the session expires. This value is
	// inclusive. The EndTimestamp is compared to the (starting) timestamp of the
	// fragment. Fragments that start before the EndTimestamp value and continue past
	// it are included in the session.
	EndTimestamp *time.Time
	
	// The start of the timestamp range for the requested media. If the
	// HLSTimestampRange value is specified, the StartTimestamp value is required.
	// Only fragments that start exactly at or after StartTimestamp are included in
	// the session. Fragments that start before StartTimestamp and continue past it
	// aren't included in the session. If FragmentSelectorType is SERVER_TIMESTAMP ,
	// the StartTimestamp must be later than the stream head.
	StartTimestamp *time.Time
	
	noSmithyDocumentSerde
}

// A structure that contains the Timestamp , Error , and ImageContent .
type Image struct {
	
	// The error message shown when the image for the provided timestamp was not
	// extracted due to a non-tryable error. An error will be returned if:
	//   - There is no media that exists for the specified Timestamp .
	//
	//   - The media for the specified time does not allow an image to be extracted.
	//   In this case the media is audio only, or the incorrect media has been ingested.
	Error ImageError
	
	// An attribute of the Image object that is Base64 encoded.
	ImageContent *string
	
	// An attribute of the Image object that is used to extract an image from the
	// video stream. This field is used to manage gaps on images or to better
	// understand the pagination window.
	TimeStamp *time.Time
	
	noSmithyDocumentSerde
}

// The range of timestamps for which to return fragments.
type TimestampRange struct {
	
	// The ending timestamp in the range of timestamps for which to return fragments.
	//
	// This member is required.
	EndTimestamp *time.Time
	
	// The starting timestamp in the range of timestamps for which to return fragments.
	//
	// This member is required.
	StartTimestamp *time.Time
	
	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
