// Code generated by smithy-go-codegen DO NOT EDIT.

// Package medicalimaging provides the API client, operations, and parameter types
// for AWS Health Imaging.
//
// This is the AWS HealthImaging API Reference. AWS HealthImaging is an AWS
// service for storing, accessing, and analyzing medical images. For an
// introduction to the service, see the AWS HealthImaging Developer Guide  (https://docs.aws.amazon.com/medical-imaging/latest/devguide)
// . We recommend using one of the AWS Software Development Kits (SDKs) for your
// programming language, as they take care of request authentication,
// serialization, and connection management. For more information, see Tools to
// build on AWS (http://aws.amazon.com/developer/tools) . For information about
// using AWS HealthImaging API actions in one of the language-specific AWS SDKs,
// refer to the See Also link at the end of each section that describes an API
// action or data type. The following sections list AWS HealthImaging API actions
// categorized according to functionality. Links are provided to actions within
// this Reference, along with links back to corresponding sections in the AWS
// HealthImaging Developer Guide so you can view console procedures and CLI/SDK
// code examples. Data store actions
//   - CreateDatastore (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_CreateDatastore.html)
//   – See Creating a data store (https://docs.aws.amazon.com/medical-imaging/latest/devguide/create-data-store.html)
//   .
//   - GetDatastore (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_GetDatastore.html)
//   – See Getting data store properties (https://docs.aws.amazon.com/medical-imaging/latest/devguide/get-data-store.html)
//   .
//   - ListDatastores (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_ListDatastores.html)
//   – See Listing data stores (https://docs.aws.amazon.com/medical-imaging/latest/devguide/list-data-stores.html)
//   .
//   - DeleteDatastore (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_DeleteDatastore.html)
//   – See Deleting a data store (https://docs.aws.amazon.com/medical-imaging/latest/devguide/delete-data-store.html)
//   .
// Import job actions
//   - StartDICOMImportJob (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_StartDICOMImportJob.html)
//   – See Starting an import job (https://docs.aws.amazon.com/medical-imaging/latest/devguide/start-dicom-import-job.html)
//   .
//   - GetDICOMImportJob (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_GetDICOMImportJob.html)
//   – See Getting import job properties (https://docs.aws.amazon.com/medical-imaging/latest/devguide/get-dicom-import-job.html)
//   .
//   - ListDICOMImportJobs (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_ListDICOMImportJobs.html)
//   – See Listing import jobs (https://docs.aws.amazon.com/medical-imaging/latest/devguide/list-dicom-import-jobs.html)
//   .
// Image set access actions
//   - SearchImageSets (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_SearchImageSets.html)
//   – See Searching image sets (https://docs.aws.amazon.com/medical-imaging/latest/devguide/search-image-sets.html)
//   .
//   - GetImageSet (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_GetImageSet.html)
//   – See Getting image set properties (https://docs.aws.amazon.com/medical-imaging/latest/devguide/get-image-set-properties.html)
//   .
//   - GetImageSetMetadata (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_GetImageSetMetadata.html)
//   – See Getting image set metadata (https://docs.aws.amazon.com/medical-imaging/latest/devguide/get-image-set-metadata.html)
//   .
//   - GetImageFrame (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_GetImageFrame.html)
//   – See Getting image set pixel data (https://docs.aws.amazon.com/medical-imaging/latest/devguide/get-image-frame.html)
//   .
// Image set modification actions
//   - ListImageSetVersions (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_ListImageSetVersions.html)
//   – See Listing image set versions (https://docs.aws.amazon.com/medical-imaging/latest/devguide/list-image-set-versions.html)
//   .
//   - UpdateImageSetMetadata (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_UpdateImageSetMetadata.html)
//   – See Updating image set metadata (https://docs.aws.amazon.com/medical-imaging/latest/devguide/update-image-set-metadata.html)
//   .
//   - CopyImageSet (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_CopyImageSet.html)
//   – See Copying an image set (https://docs.aws.amazon.com/medical-imaging/latest/devguide/copy-image-set.html)
//   .
//   - DeleteImageSet (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_DeleteImageSet.html)
//   – See Deleting an image set (https://docs.aws.amazon.com/medical-imaging/latest/devguide/delete-image-set.html)
//   .
// Tagging actions
//   - TagResource (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_TagResource.html)
//   – See Tagging a data store (https://docs.aws.amazon.com/medical-imaging/latest/devguide/tag-list-untag-data-store.html)
//   and Tagging an image set (https://docs.aws.amazon.com/medical-imaging/latest/devguide/tag-list-untag-image-set.html)
//   .
//   - ListTagsForResource (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_ListTagsForResource.html)
//   – See Tagging a data store (https://docs.aws.amazon.com/medical-imaging/latest/devguide/tag-list-untag-data-store.html)
//   and Tagging an image set (https://docs.aws.amazon.com/medical-imaging/latest/devguide/tag-list-untag-image-set.html)
//   .
//   - UntagResource (https://docs.aws.amazon.com/medical-imaging/latest/APIReference/API_UntagResource.html)
//   – See Tagging a data store (https://docs.aws.amazon.com/medical-imaging/latest/devguide/tag-list-untag-data-store.html)
//   and Tagging an image set (https://docs.aws.amazon.com/medical-imaging/latest/devguide/tag-list-untag-image-set.html)
//   .
package medicalimaging


