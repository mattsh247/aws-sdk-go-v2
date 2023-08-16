// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Provides access information used by the authorityInfoAccess and
// subjectInfoAccess extensions described in RFC 5280 (https://datatracker.ietf.org/doc/html/rfc5280)
// .
type AccessDescription struct {
	
	// The location of AccessDescription information.
	//
	// This member is required.
	AccessLocation *GeneralName
	
	// The type and format of AccessDescription information.
	//
	// This member is required.
	AccessMethod *AccessMethod
	
	noSmithyDocumentSerde
}

// Describes the type and format of extension access. Only one of
// CustomObjectIdentifier or AccessMethodType may be provided. Providing both
// results in InvalidArgsException .
type AccessMethod struct {
	
	// Specifies the AccessMethod .
	AccessMethodType AccessMethodType
	
	// An object identifier (OID) specifying the AccessMethod . The OID must satisfy
	// the regular expression shown below. For more information, see NIST's definition
	// of Object Identifier (OID) (https://csrc.nist.gov/glossary/term/Object_Identifier)
	// .
	CustomObjectIdentifier *string
	
	noSmithyDocumentSerde
}

// Contains X.509 certificate information to be placed in an issued certificate.
// An APIPassthrough or APICSRPassthrough template variant must be selected, or
// else this parameter is ignored. If conflicting or duplicate certificate
// information is supplied from other sources, Amazon Web Services Private CA
// applies order of operation rules (https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html#template-order-of-operations)
// to determine what information is used.
type ApiPassthrough struct {
	
	// Specifies X.509 extension information for a certificate.
	Extensions *Extensions
	
	// Contains information about the certificate subject. The Subject field in the
	// certificate identifies the entity that owns or controls the public key in the
	// certificate. The entity can be a user, computer, device, or service. The
	// Subject must contain an X.500 distinguished name (DN). A DN is a sequence of
	// relative distinguished names (RDNs). The RDNs are separated by commas in the
	// certificate.
	Subject *ASN1Subject
	
	noSmithyDocumentSerde
}

// Contains information about the certificate subject. The Subject field in the
// certificate identifies the entity that owns or controls the public key in the
// certificate. The entity can be a user, computer, device, or service. The
// Subject must contain an X.500 distinguished name (DN). A DN is a sequence of
// relative distinguished names (RDNs). The RDNs are separated by commas in the
// certificate.
type ASN1Subject struct {
	
	// For CA and end-entity certificates in a private PKI, the common name (CN) can
	// be any string within the length limit. Note: In publicly trusted certificates,
	// the common name must be a fully qualified domain name (FQDN) associated with the
	// certificate subject.
	CommonName *string
	
	// Two-digit code that specifies the country in which the certificate subject
	// located.
	Country *string
	
	// Contains a sequence of one or more X.500 relative distinguished names (RDNs),
	// each of which consists of an object identifier (OID) and a value. For more
	// information, see NIST’s definition of Object Identifier (OID) (https://csrc.nist.gov/glossary/term/Object_Identifier)
	// . Custom attributes cannot be used in combination with standard attributes.
	CustomAttributes []CustomAttribute
	
	// Disambiguating information for the certificate subject.
	DistinguishedNameQualifier *string
	
	// Typically a qualifier appended to the name of an individual. Examples include
	// Jr. for junior, Sr. for senior, and III for third.
	GenerationQualifier *string
	
	// First name.
	GivenName *string
	
	// Concatenation that typically contains the first letter of the GivenName, the
	// first letter of the middle name if one exists, and the first letter of the
	// Surname.
	Initials *string
	
	// The locality (such as a city or town) in which the certificate subject is
	// located.
	Locality *string
	
	// Legal name of the organization with which the certificate subject is affiliated.
	Organization *string
	
	// A subdivision or unit of the organization (such as sales or finance) with which
	// the certificate subject is affiliated.
	OrganizationalUnit *string
	
	// Typically a shortened version of a longer GivenName. For example, Jonathan is
	// often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
	Pseudonym *string
	
	// The certificate serial number.
	SerialNumber *string
	
	// State in which the subject of the certificate is located.
	State *string
	
	// Family name. In the US and the UK, for example, the surname of an individual is
	// ordered last. In Asian cultures the surname is typically ordered first.
	Surname *string
	
	// A title such as Mr. or Ms., which is pre-pended to the name to refer formally
	// to the certificate subject.
	Title *string
	
	noSmithyDocumentSerde
}

// Contains information about your private certificate authority (CA). Your
// private CA can issue and revoke X.509 digital certificates. Digital certificates
// verify that the entity named in the certificate Subject field owns or controls
// the public key contained in the Subject Public Key Info field. Call the
// CreateCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html)
// action to create your private CA. You must then call the
// GetCertificateAuthorityCertificate (https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetCertificateAuthorityCertificate.html)
// action to retrieve a private CA certificate signing request (CSR). Sign the CSR
// with your Amazon Web Services Private CA-hosted or on-premises root or
// subordinate CA certificate. Call the ImportCertificateAuthorityCertificate (https://docs.aws.amazon.com/privateca/latest/APIReference/API_ImportCertificateAuthorityCertificate.html)
// action to import the signed certificate into Certificate Manager (ACM).
type CertificateAuthority struct {
	
	// Amazon Resource Name (ARN) for your private certificate authority (CA). The
	// format is 12345678-1234-1234-1234-123456789012 .
	Arn *string
	
	// Your private CA configuration.
	CertificateAuthorityConfiguration *CertificateAuthorityConfiguration
	
	// Date and time at which your private CA was created.
	CreatedAt *time.Time
	
	// Reason the request to create your private CA failed.
	FailureReason FailureReason
	
	// Defines a cryptographic key management compliance standard used for handling CA
	// keys. Default: FIPS_140_2_LEVEL_3_OR_HIGHER Note: Amazon Web Services Region
	// ap-northeast-3 supports only FIPS_140_2_LEVEL_2_OR_HIGHER. You must explicitly
	// specify this parameter and value when creating a CA in that Region. Specifying a
	// different value (or no value) results in an InvalidArgsException with the
	// message "A certificate authority cannot be created in this region with the
	// specified security standard."
	KeyStorageSecurityStandard KeyStorageSecurityStandard
	
	// Date and time at which your private CA was last updated.
	LastStateChangeAt *time.Time
	
	// Date and time after which your private CA certificate is not valid.
	NotAfter *time.Time
	
	// Date and time before which your private CA certificate is not valid.
	NotBefore *time.Time
	
	// The Amazon Web Services account ID that owns the certificate authority.
	OwnerAccount *string
	
	// The period during which a deleted CA can be restored. For more information, see
	// the PermanentDeletionTimeInDays parameter of the
	// DeleteCertificateAuthorityRequest (https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeleteCertificateAuthorityRequest.html)
	// action.
	RestorableUntil *time.Time
	
	// Information about the Online Certificate Status Protocol (OCSP) configuration
	// or certificate revocation list (CRL) created and maintained by your private CA.
	RevocationConfiguration *RevocationConfiguration
	
	// Serial number of your private CA.
	Serial *string
	
	// Status of your private CA.
	Status CertificateAuthorityStatus
	
	// Type of your private CA.
	Type CertificateAuthorityType
	
	// Specifies whether the CA issues general-purpose certificates that typically
	// require a revocation mechanism, or short-lived certificates that may optionally
	// omit revocation because they expire quickly. Short-lived certificate validity is
	// limited to seven days. The default value is GENERAL_PURPOSE.
	UsageMode CertificateAuthorityUsageMode
	
	noSmithyDocumentSerde
}

// Contains configuration information for your private certificate authority (CA).
// This includes information about the class of public key algorithm and the key
// pair that your private CA creates when it issues a certificate. It also includes
// the signature algorithm that it uses when issuing certificates, and its X.500
// distinguished name. You must specify this information when you call the
// CreateCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html)
// action.
type CertificateAuthorityConfiguration struct {
	
	// Type of the public key algorithm and size, in bits, of the key pair that your
	// CA creates when it issues a certificate. When you create a subordinate CA, you
	// must use a key algorithm supported by the parent CA.
	//
	// This member is required.
	KeyAlgorithm KeyAlgorithm
	
	// Name of the algorithm your private CA uses to sign certificate requests. This
	// parameter should not be confused with the SigningAlgorithm parameter used to
	// sign certificates when they are issued.
	//
	// This member is required.
	SigningAlgorithm SigningAlgorithm
	
	// Structure that contains X.500 distinguished name information for your private
	// CA.
	//
	// This member is required.
	Subject *ASN1Subject
	
	// Specifies information to be added to the extension section of the certificate
	// signing request (CSR).
	CsrExtensions *CsrExtensions
	
	noSmithyDocumentSerde
}

// Contains configuration information for a certificate revocation list (CRL).
// Your private certificate authority (CA) creates base CRLs. Delta CRLs are not
// supported. You can enable CRLs for your new or an existing private CA by setting
// the Enabled parameter to true . Your private CA writes CRLs to an S3 bucket that
// you specify in the S3BucketName parameter. You can hide the name of your bucket
// by specifying a value for the CustomCname parameter. Your private CA copies the
// CNAME or the S3 bucket name to the CRL Distribution Points extension of each
// certificate it issues. Your S3 bucket policy must give write permission to
// Amazon Web Services Private CA. Amazon Web Services Private CA assets that are
// stored in Amazon S3 can be protected with encryption. For more information, see
// Encrypting Your CRLs (https://docs.aws.amazon.com/privateca/latest/userguide/PcaCreateCa.html#crl-encryption)
// . Your private CA uses the value in the ExpirationInDays parameter to calculate
// the nextUpdate field in the CRL. The CRL is refreshed prior to a certificate's
// expiration date or when a certificate is revoked. When a certificate is revoked,
// it appears in the CRL until the certificate expires, and then in one additional
// CRL after expiration, and it always appears in the audit report. A CRL is
// typically updated approximately 30 minutes after a certificate is revoked. If
// for any reason a CRL update fails, Amazon Web Services Private CA makes further
// attempts every 15 minutes. CRLs contain the following fields:
//   - Version: The current version number defined in RFC 5280 is V2. The integer
//   value is 0x1.
//   - Signature Algorithm: The name of the algorithm used to sign the CRL.
//   - Issuer: The X.500 distinguished name of your private CA that issued the
//   CRL.
//   - Last Update: The issue date and time of this CRL.
//   - Next Update: The day and time by which the next CRL will be issued.
//   - Revoked Certificates: List of revoked certificates. Each list item contains
//   the following information.
//   - Serial Number: The serial number, in hexadecimal format, of the revoked
//   certificate.
//   - Revocation Date: Date and time the certificate was revoked.
//   - CRL Entry Extensions: Optional extensions for the CRL entry.
//   - X509v3 CRL Reason Code: Reason the certificate was revoked.
//   - CRL Extensions: Optional extensions for the CRL.
//   - X509v3 Authority Key Identifier: Identifies the public key associated with
//   the private key used to sign the certificate.
//   - X509v3 CRL Number:: Decimal sequence number for the CRL.
//   - Signature Algorithm: Algorithm used by your private CA to sign the CRL.
//   - Signature Value: Signature computed over the CRL.
// Certificate revocation lists created by Amazon Web Services Private CA are
// DER-encoded. You can use the following OpenSSL command to list a CRL. openssl
// crl -inform DER -text -in crl_path -noout For more information, see Planning a
// certificate revocation list (CRL) (https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html)
// in the Amazon Web Services Private Certificate Authority User Guide
type CrlConfiguration struct {
	
	// Boolean value that specifies whether certificate revocation lists (CRLs) are
	// enabled. You can use this value to enable certificate revocation for a new CA
	// when you call the CreateCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html)
	// action or for an existing CA when you call the UpdateCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_UpdateCertificateAuthority.html)
	// action.
	//
	// This member is required.
	Enabled *bool
	
	// Name inserted into the certificate CRL Distribution Points extension that
	// enables the use of an alias for the CRL distribution point. Use this value if
	// you don't want the name of your S3 bucket to be public. The content of a
	// Canonical Name (CNAME) record must conform to RFC2396 (https://www.ietf.org/rfc/rfc2396.txt)
	// restrictions on the use of special characters in URIs. Additionally, the value
	// of the CNAME must not include a protocol prefix such as "http://" or "https://".
	CustomCname *string
	
	// Validity period of the CRL in days.
	ExpirationInDays *int32
	
	// Name of the S3 bucket that contains the CRL. If you do not provide a value for
	// the CustomCname argument, the name of your S3 bucket is placed into the CRL
	// Distribution Points extension of the issued certificate. You can change the name
	// of your bucket by calling the UpdateCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_UpdateCertificateAuthority.html)
	// operation. You must specify a bucket policy (https://docs.aws.amazon.com/privateca/latest/userguide/PcaCreateCa.html#s3-policies)
	// that allows Amazon Web Services Private CA to write the CRL to your bucket. The
	// S3BucketName parameter must conform to the S3 bucket naming rules (https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html)
	// .
	S3BucketName *string
	
	// Determines whether the CRL will be publicly readable or privately held in the
	// CRL Amazon S3 bucket. If you choose PUBLIC_READ, the CRL will be accessible over
	// the public internet. If you choose BUCKET_OWNER_FULL_CONTROL, only the owner of
	// the CRL S3 bucket can access the CRL, and your PKI clients may need an
	// alternative method of access. If no value is specified, the default is
	// PUBLIC_READ . Note: This default can cause CA creation to fail in some
	// circumstances. If you have have enabled the Block Public Access (BPA) feature in
	// your S3 account, then you must specify the value of this parameter as
	// BUCKET_OWNER_FULL_CONTROL , and not doing so results in an error. If you have
	// disabled BPA in S3, then you can specify either BUCKET_OWNER_FULL_CONTROL or
	// PUBLIC_READ as the value. For more information, see Blocking public access to
	// the S3 bucket (https://docs.aws.amazon.com/privateca/latest/userguide/PcaCreateCa.html#s3-bpa)
	// .
	S3ObjectAcl S3ObjectAcl
	
	noSmithyDocumentSerde
}

// Describes the certificate extensions to be added to the certificate signing
// request (CSR).
type CsrExtensions struct {
	
	// Indicates the purpose of the certificate and of the key contained in the
	// certificate.
	KeyUsage *KeyUsage
	
	// For CA certificates, provides a path to additional information pertaining to
	// the CA, such as revocation and policy. For more information, see Subject
	// Information Access (https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.2.2)
	// in RFC 5280.
	SubjectInformationAccess []AccessDescription
	
	noSmithyDocumentSerde
}

// Defines the X.500 relative distinguished name (RDN).
type CustomAttribute struct {
	
	// Specifies the object identifier (OID) of the attribute type of the relative
	// distinguished name (RDN).
	//
	// This member is required.
	ObjectIdentifier *string
	
	// Specifies the attribute value of relative distinguished name (RDN).
	//
	// This member is required.
	Value *string
	
	noSmithyDocumentSerde
}

// Specifies the X.509 extension information for a certificate. Extensions present
// in CustomExtensions follow the ApiPassthrough template rules (https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html#template-order-of-operations)
// .
type CustomExtension struct {
	
	// Specifies the object identifier (OID) of the X.509 extension. For more
	// information, see the Global OID reference database. (https://oidref.com/2.5.29)
	//
	// This member is required.
	ObjectIdentifier *string
	
	// Specifies the base64-encoded value of the X.509 extension.
	//
	// This member is required.
	Value *string
	
	// Specifies the critical flag of the X.509 extension.
	Critical *bool
	
	noSmithyDocumentSerde
}

// Describes an Electronic Data Interchange (EDI) entity as described in as
// defined in Subject Alternative Name (https://datatracker.ietf.org/doc/html/rfc5280)
// in RFC 5280.
type EdiPartyName struct {
	
	// Specifies the party name.
	//
	// This member is required.
	PartyName *string
	
	// Specifies the name assigner.
	NameAssigner *string
	
	noSmithyDocumentSerde
}

// Specifies additional purposes for which the certified public key may be used
// other than basic purposes indicated in the KeyUsage extension.
type ExtendedKeyUsage struct {
	
	// Specifies a custom ExtendedKeyUsage with an object identifier (OID).
	ExtendedKeyUsageObjectIdentifier *string
	
	// Specifies a standard ExtendedKeyUsage as defined as in RFC 5280 (https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12)
	// .
	ExtendedKeyUsageType ExtendedKeyUsageType
	
	noSmithyDocumentSerde
}

// Contains X.509 extension information for a certificate.
type Extensions struct {
	
	// Contains a sequence of one or more policy information terms, each of which
	// consists of an object identifier (OID) and optional qualifiers. For more
	// information, see NIST's definition of Object Identifier (OID) (https://csrc.nist.gov/glossary/term/Object_Identifier)
	// . In an end-entity certificate, these terms indicate the policy under which the
	// certificate was issued and the purposes for which it may be used. In a CA
	// certificate, these terms limit the set of policies for certification paths that
	// include this certificate.
	CertificatePolicies []PolicyInformation
	
	// Contains a sequence of one or more X.509 extensions, each of which consists of
	// an object identifier (OID), a base64-encoded value, and the critical flag. For
	// more information, see the Global OID reference database. (https://oidref.com/2.5.29)
	CustomExtensions []CustomExtension
	
	// Specifies additional purposes for which the certified public key may be used
	// other than basic purposes indicated in the KeyUsage extension.
	ExtendedKeyUsage []ExtendedKeyUsage
	
	// Defines one or more purposes for which the key contained in the certificate can
	// be used. Default value for each option is false.
	KeyUsage *KeyUsage
	
	// The subject alternative name extension allows identities to be bound to the
	// subject of the certificate. These identities may be included in addition to or
	// in place of the identity in the subject field of the certificate.
	SubjectAlternativeNames []GeneralName
	
	noSmithyDocumentSerde
}

// Describes an ASN.1 X.400 GeneralName as defined in RFC 5280 (https://datatracker.ietf.org/doc/html/rfc5280)
// . Only one of the following naming options should be provided. Providing more
// than one option results in an InvalidArgsException error.
type GeneralName struct {
	
	// Contains information about the certificate subject. The Subject field in the
	// certificate identifies the entity that owns or controls the public key in the
	// certificate. The entity can be a user, computer, device, or service. The
	// Subject must contain an X.500 distinguished name (DN). A DN is a sequence of
	// relative distinguished names (RDNs). The RDNs are separated by commas in the
	// certificate.
	DirectoryName *ASN1Subject
	
	// Represents GeneralName as a DNS name.
	DnsName *string
	
	// Represents GeneralName as an EdiPartyName object.
	EdiPartyName *EdiPartyName
	
	// Represents GeneralName as an IPv4 or IPv6 address.
	IpAddress *string
	
	// Represents GeneralName using an OtherName object.
	OtherName *OtherName
	
	// Represents GeneralName as an object identifier (OID).
	RegisteredId *string
	
	// Represents GeneralName as an RFC 822 (https://datatracker.ietf.org/doc/html/rfc822)
	// email address.
	Rfc822Name *string
	
	// Represents GeneralName as a URI.
	UniformResourceIdentifier *string
	
	noSmithyDocumentSerde
}

// Defines one or more purposes for which the key contained in the certificate can
// be used. Default value for each option is false.
type KeyUsage struct {
	
	// Key can be used to sign CRLs.
	CRLSign bool
	
	// Key can be used to decipher data.
	DataEncipherment bool
	
	// Key can be used only to decipher data.
	DecipherOnly bool
	
	// Key can be used for digital signing.
	DigitalSignature bool
	
	// Key can be used only to encipher data.
	EncipherOnly bool
	
	// Key can be used in a key-agreement protocol.
	KeyAgreement bool
	
	// Key can be used to sign certificates.
	KeyCertSign bool
	
	// Key can be used to encipher data.
	KeyEncipherment bool
	
	// Key can be used for non-repudiation.
	NonRepudiation bool
	
	noSmithyDocumentSerde
}

// Contains information to enable and configure Online Certificate Status Protocol
// (OCSP) for validating certificate revocation status. When you revoke a
// certificate, OCSP responses may take up to 60 minutes to reflect the new status.
type OcspConfiguration struct {
	
	// Flag enabling use of the Online Certificate Status Protocol (OCSP) for
	// validating certificate revocation status.
	//
	// This member is required.
	Enabled *bool
	
	// By default, Amazon Web Services Private CA injects an Amazon Web Services
	// domain into certificates being validated by the Online Certificate Status
	// Protocol (OCSP). A customer can alternatively use this object to define a CNAME
	// specifying a customized OCSP domain. The content of a Canonical Name (CNAME)
	// record must conform to RFC2396 (https://www.ietf.org/rfc/rfc2396.txt)
	// restrictions on the use of special characters in URIs. Additionally, the value
	// of the CNAME must not include a protocol prefix such as "http://" or "https://".
	// For more information, see Customizing Online Certificate Status Protocol (OCSP)  (https://docs.aws.amazon.com/privateca/latest/userguide/ocsp-customize.html)
	// in the Amazon Web Services Private Certificate Authority User Guide.
	OcspCustomCname *string
	
	noSmithyDocumentSerde
}

// Defines a custom ASN.1 X.400 GeneralName using an object identifier (OID) and
// value. The OID must satisfy the regular expression shown below. For more
// information, see NIST's definition of Object Identifier (OID) (https://csrc.nist.gov/glossary/term/Object_Identifier)
// .
type OtherName struct {
	
	// Specifies an OID.
	//
	// This member is required.
	TypeId *string
	
	// Specifies an OID value.
	//
	// This member is required.
	Value *string
	
	noSmithyDocumentSerde
}

// Permissions designate which private CA actions can be performed by an Amazon
// Web Services service or entity. In order for ACM to automatically renew private
// certificates, you must give the ACM service principal all available permissions
// ( IssueCertificate , GetCertificate , and ListPermissions ). Permissions can be
// assigned with the CreatePermission (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreatePermission.html)
// action, removed with the DeletePermission (https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeletePermission.html)
// action, and listed with the ListPermissions (https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListPermissions.html)
// action.
type Permission struct {
	
	// The private CA actions that can be performed by the designated Amazon Web
	// Services service.
	Actions []ActionType
	
	// The Amazon Resource Number (ARN) of the private CA from which the permission
	// was issued.
	CertificateAuthorityArn *string
	
	// The time at which the permission was created.
	CreatedAt *time.Time
	
	// The name of the policy that is associated with the permission.
	Policy *string
	
	// The Amazon Web Services service or entity that holds the permission. At this
	// time, the only valid principal is acm.amazonaws.com .
	Principal *string
	
	// The ID of the account that assigned the permission.
	SourceAccount *string
	
	noSmithyDocumentSerde
}

// Defines the X.509 CertificatePolicies extension.
type PolicyInformation struct {
	
	// Specifies the object identifier (OID) of the certificate policy under which the
	// certificate was issued. For more information, see NIST's definition of Object
	// Identifier (OID) (https://csrc.nist.gov/glossary/term/Object_Identifier) .
	//
	// This member is required.
	CertPolicyId *string
	
	// Modifies the given CertPolicyId with a qualifier. Amazon Web Services Private
	// CA supports the certification practice statement (CPS) qualifier.
	PolicyQualifiers []PolicyQualifierInfo
	
	noSmithyDocumentSerde
}

// Modifies the CertPolicyId of a PolicyInformation object with a qualifier.
// Amazon Web Services Private CA supports the certification practice statement
// (CPS) qualifier.
type PolicyQualifierInfo struct {
	
	// Identifies the qualifier modifying a CertPolicyId .
	//
	// This member is required.
	PolicyQualifierId PolicyQualifierId
	
	// Defines the qualifier type. Amazon Web Services Private CA supports the use of
	// a URI for a CPS qualifier in this field.
	//
	// This member is required.
	Qualifier *Qualifier
	
	noSmithyDocumentSerde
}

// Defines a PolicyInformation qualifier. Amazon Web Services Private CA supports
// the certification practice statement (CPS) qualifier (https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.4)
// defined in RFC 5280.
type Qualifier struct {
	
	// Contains a pointer to a certification practice statement (CPS) published by the
	// CA.
	//
	// This member is required.
	CpsUri *string
	
	noSmithyDocumentSerde
}

// Certificate revocation information used by the CreateCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html)
// and UpdateCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_UpdateCertificateAuthority.html)
// actions. Your private certificate authority (CA) can configure Online
// Certificate Status Protocol (OCSP) support and/or maintain a certificate
// revocation list (CRL). OCSP returns validation information about certificates as
// requested by clients, and a CRL contains an updated list of certificates revoked
// by your CA. For more information, see RevokeCertificate (https://docs.aws.amazon.com/privateca/latest/APIReference/API_RevokeCertificate.html)
// and Setting up a certificate revocation method (https://docs.aws.amazon.com/privateca/latest/userguide/revocation-setup.html)
// in the Amazon Web Services Private Certificate Authority User Guide.
type RevocationConfiguration struct {
	
	// Configuration of the certificate revocation list (CRL), if any, maintained by
	// your private CA. A CRL is typically updated approximately 30 minutes after a
	// certificate is revoked. If for any reason a CRL update fails, Amazon Web
	// Services Private CA makes further attempts every 15 minutes.
	CrlConfiguration *CrlConfiguration
	
	// Configuration of Online Certificate Status Protocol (OCSP) support, if any,
	// maintained by your private CA. When you revoke a certificate, OCSP responses may
	// take up to 60 minutes to reflect the new status.
	OcspConfiguration *OcspConfiguration
	
	noSmithyDocumentSerde
}

// Tags are labels that you can use to identify and organize your private CAs.
// Each tag consists of a key and an optional value. You can associate up to 50
// tags with a private CA. To add one or more tags to a private CA, call the
// TagCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_TagCertificateAuthority.html)
// action. To remove a tag, call the UntagCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_UntagCertificateAuthority.html)
// action.
type Tag struct {
	
	// Key (name) of the tag.
	//
	// This member is required.
	Key *string
	
	// Value of the tag.
	Value *string
	
	noSmithyDocumentSerde
}

// Validity specifies the period of time during which a certificate is valid.
// Validity can be expressed as an explicit date and time when the validity of a
// certificate starts or expires, or as a span of time after issuance, stated in
// days, months, or years. For more information, see Validity (https://tools.ietf.org/html/rfc5280#section-4.1.2.5)
// in RFC 5280. Amazon Web Services Private CA API consumes the Validity data type
// differently in two distinct parameters of the IssueCertificate action. The
// required parameter IssueCertificate : Validity specifies the end of a
// certificate's validity period. The optional parameter IssueCertificate :
// ValidityNotBefore specifies a customized starting time for the validity period.
type Validity struct {
	
	// Determines how Amazon Web Services Private CA interprets the Value parameter,
	// an integer. Supported validity types include those listed below. Type
	// definitions with values include a sample input value and the resulting output.
	// END_DATE : The specific date and time when the certificate will expire,
	// expressed using UTCTime (YYMMDDHHMMSS) or GeneralizedTime (YYYYMMDDHHMMSS)
	// format. When UTCTime is used, if the year field (YY) is greater than or equal to
	// 50, the year is interpreted as 19YY. If the year field is less than 50, the year
	// is interpreted as 20YY.
	//   - Sample input value: 491231235959 (UTCTime format)
	//   - Output expiration date/time: 12/31/2049 23:59:59
	// ABSOLUTE : The specific date and time when the validity of a certificate will
	// start or expire, expressed in seconds since the Unix Epoch.
	//   - Sample input value: 2524608000
	//   - Output expiration date/time: 01/01/2050 00:00:00
	// DAYS , MONTHS , YEARS : The relative time from the moment of issuance until the
	// certificate will expire, expressed in days, months, or years. Example if DAYS ,
	// issued on 10/12/2020 at 12:34:54 UTC:
	//   - Sample input value: 90
	//   - Output expiration date: 01/10/2020 12:34:54 UTC
	// The minimum validity duration for a certificate using relative time ( DAYS ) is
	// one day. The minimum validity for a certificate using absolute time ( ABSOLUTE
	// or END_DATE ) is one second.
	//
	// This member is required.
	Type ValidityPeriodType
	
	// A long integer interpreted according to the value of Type , below.
	//
	// This member is required.
	Value *int64
	
	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
