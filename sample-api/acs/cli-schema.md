# App Store Connect API

**CLI name:** `app-store-connect-api` · **Version:** 4.3
**Spec format:** OpenAPI 3.x / Swagger · 923 paths · 1208 operations
**Base URL:** https://api.appstoreconnect.apple.com

## Authentication

**Type:** `none`

## Generate

```bash
go run ./cmd/builder generate --spec sample-api/acs/openapi.oas.json --out sample-api/acs/cli-schema.yaml
```

## Usage

```bash
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml --help
```

## Example commands

List /v1/actors

```bash
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml actors list --limit 5
```

List /v1/alternativeDistributionDomains

```bash
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml alternativeDistributionDomains list --limit 5
```

List /v1/alternativeDistributionKeys

```bash
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml alternativeDistributionKeys list --limit 5
```

List /v1/alternativeDistributionPackageVersions/{id}/deltas

```bash
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml alternativeDistributionPackageVersions deltas list <id> --limit 5
```

List /v1/alternativeDistributionPackageVersions/{id}/variants

```bash
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml alternativeDistributionPackageVersions variants list <id> --limit 5
```

List /v1/alternativeDistributionPackages/{id}/versions

```bash
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml alternativeDistributionPackages versions list <id> --limit 5
```

List /v1/analyticsReportInstances/{id}/segments

```bash
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml analyticsReportInstances segments list <id> --limit 5
```

List /v1/analyticsReportRequests/{id}/reports

```bash
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml analyticsReportRequests reports list <id> --limit 5
```

List /v1/analyticsReports/{id}/instances

```bash
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml analyticsReports instances list <id> --limit 5
```

List /v2/appAvailabilities/{id}/territoryAvailabilities

```bash
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml appAvailabilitiesV2 territoryAvailabilities list <id> --limit 5
```

## Commands

This CLI defines 865 commands. Use `--help` to explore the full tree.

### accessibilityDeclarations

Manage accessibilityDeclarations

4 subcommand(s).

#### create

Create /v1/accessibilityDeclarations

#### delete

Delete /v1/accessibilityDeclarations/{id}

#### get

Get /v1/accessibilityDeclarations/{id}

#### update

Update /v1/accessibilityDeclarations/{id}

### actors

Manage actors

2 subcommand(s).

#### get

Get /v1/actors/{id}

#### list

List /v1/actors

### ageRatingDeclarations

Manage ageRatingDeclarations

1 subcommand(s).

#### update

Update /v1/ageRatingDeclarations/{id}

### alternativeDistributionDomains

Manage alternativeDistributionDomains

4 subcommand(s).

#### create

Create /v1/alternativeDistributionDomains

#### delete

Delete /v1/alternativeDistributionDomains/{id}

#### get

Get /v1/alternativeDistributionDomains/{id}

#### list

List /v1/alternativeDistributionDomains

### alternativeDistributionKeys

Manage alternativeDistributionKeys

4 subcommand(s).

#### create

Create /v1/alternativeDistributionKeys

#### delete

Delete /v1/alternativeDistributionKeys/{id}

#### get

Get /v1/alternativeDistributionKeys/{id}

#### list

List /v1/alternativeDistributionKeys

### alternativeDistributionPackageDeltas

Manage alternativeDistributionPackageDeltas

1 subcommand(s).

#### get

Get /v1/alternativeDistributionPackageDeltas/{id}

### alternativeDistributionPackageVariants

Manage alternativeDistributionPackageVariants

1 subcommand(s).

#### get

Get /v1/alternativeDistributionPackageVariants/{id}

### alternativeDistributionPackageVersions

Manage alternativeDistributionPackageVersions

3 subcommand(s).

#### deltas

Manage alternativeDistributionPackageVersions deltas

1 subcommand(s).

##### list

List /v1/alternativeDistributionPackageVersions/{id}/deltas

#### get

Get /v1/alternativeDistributionPackageVersions/{id}

#### variants

Manage alternativeDistributionPackageVersions variants

1 subcommand(s).

##### list

List /v1/alternativeDistributionPackageVersions/{id}/variants

### alternativeDistributionPackages

Manage alternativeDistributionPackages

3 subcommand(s).

#### create

Create /v1/alternativeDistributionPackages

#### get

Get /v1/alternativeDistributionPackages/{id}

#### versions

Manage alternativeDistributionPackages versions

1 subcommand(s).

##### list

List /v1/alternativeDistributionPackages/{id}/versions

### analyticsReportInstances

Manage analyticsReportInstances

2 subcommand(s).

#### get

Get /v1/analyticsReportInstances/{id}

#### segments

Manage analyticsReportInstances segments

1 subcommand(s).

##### list

List /v1/analyticsReportInstances/{id}/segments

### analyticsReportRequests

Manage analyticsReportRequests

4 subcommand(s).

#### create

Create /v1/analyticsReportRequests

#### delete

Delete /v1/analyticsReportRequests/{id}

#### get

Get /v1/analyticsReportRequests/{id}

#### reports

Manage analyticsReportRequests reports

1 subcommand(s).

##### list

List /v1/analyticsReportRequests/{id}/reports

### analyticsReportSegments

Manage analyticsReportSegments

1 subcommand(s).

#### get

Get /v1/analyticsReportSegments/{id}

### analyticsReports

Manage analyticsReports

2 subcommand(s).

#### get

Get /v1/analyticsReports/{id}

#### instances

Manage analyticsReports instances

1 subcommand(s).

##### list

List /v1/analyticsReports/{id}/instances

### androidToIosAppMappingDetails

Manage androidToIosAppMappingDetails

4 subcommand(s).

#### create

Create /v1/androidToIosAppMappingDetails

#### delete

Delete /v1/androidToIosAppMappingDetails/{id}

#### get

Get /v1/androidToIosAppMappingDetails/{id}

#### update

Update /v1/androidToIosAppMappingDetails/{id}

### appAvailabilitiesV2

Manage appAvailabilitiesV2

3 subcommand(s).

#### create

Create /v2/appAvailabilities

#### get

Get /v2/appAvailabilities/{id}

#### territoryAvailabilities

Manage appAvailabilitiesV2 territoryAvailabilities

1 subcommand(s).

##### list

List /v2/appAvailabilities/{id}/territoryAvailabilities

### appCategories

Manage appCategories

4 subcommand(s).

#### get

Get /v1/appCategories/{id}

#### list

List /v1/appCategories

#### parent

Manage appCategories parent

1 subcommand(s).

##### get

Get /v1/appCategories/{id}/parent

#### subcategories

Manage appCategories subcategories

1 subcommand(s).

##### list

List /v1/appCategories/{id}/subcategories

### appClipAdvancedExperienceImages

Manage appClipAdvancedExperienceImages

3 subcommand(s).

#### create

Create /v1/appClipAdvancedExperienceImages

#### get

Get /v1/appClipAdvancedExperienceImages/{id}

#### update

Update /v1/appClipAdvancedExperienceImages/{id}

### appClipAdvancedExperiences

Manage appClipAdvancedExperiences

3 subcommand(s).

#### create

Create /v1/appClipAdvancedExperiences

#### get

Get /v1/appClipAdvancedExperiences/{id}

#### update

Update /v1/appClipAdvancedExperiences/{id}

### appClipAppStoreReviewDetails

Manage appClipAppStoreReviewDetails

3 subcommand(s).

#### create

Create /v1/appClipAppStoreReviewDetails

#### get

Get /v1/appClipAppStoreReviewDetails/{id}

#### update

Update /v1/appClipAppStoreReviewDetails/{id}

### appClipDefaultExperienceLocalizations

Manage appClipDefaultExperienceLocalizations

5 subcommand(s).

#### appClipHeaderImage

Manage appClipDefaultExperienceLocalizations appClipHeaderImage

1 subcommand(s).

##### get

Get /v1/appClipDefaultExperienceLocalizations/{id}/appClipHeaderImage

#### create

Create /v1/appClipDefaultExperienceLocalizations

#### delete

Delete /v1/appClipDefaultExperienceLocalizations/{id}

#### get

Get /v1/appClipDefaultExperienceLocalizations/{id}

#### update

Update /v1/appClipDefaultExperienceLocalizations/{id}

### appClipDefaultExperiences

Manage appClipDefaultExperiences

7 subcommand(s).

#### appClipAppStoreReviewDetail

Manage appClipDefaultExperiences appClipAppStoreReviewDetail

1 subcommand(s).

##### get

Get /v1/appClipDefaultExperiences/{id}/appClipAppStoreReviewDetail

#### appClipDefaultExperienceLocalizations

Manage appClipDefaultExperiences appClipDefaultExperienceLocalizations

1 subcommand(s).

##### list

List /v1/appClipDefaultExperiences/{id}/appClipDefaultExperienceLocalizations

#### create

Create /v1/appClipDefaultExperiences

#### delete

Delete /v1/appClipDefaultExperiences/{id}

#### get

Get /v1/appClipDefaultExperiences/{id}

#### releaseWithAppStoreVersion

Manage appClipDefaultExperiences releaseWithAppStoreVersion

1 subcommand(s).

##### get

Get /v1/appClipDefaultExperiences/{id}/releaseWithAppStoreVersion

#### update

Update /v1/appClipDefaultExperiences/{id}

### appClipHeaderImages

Manage appClipHeaderImages

4 subcommand(s).

#### create

Create /v1/appClipHeaderImages

#### delete

Delete /v1/appClipHeaderImages/{id}

#### get

Get /v1/appClipHeaderImages/{id}

#### update

Update /v1/appClipHeaderImages/{id}

### appClips

Manage appClips

3 subcommand(s).

#### appClipAdvancedExperiences

Manage appClips appClipAdvancedExperiences

1 subcommand(s).

##### list

List /v1/appClips/{id}/appClipAdvancedExperiences

#### appClipDefaultExperiences

Manage appClips appClipDefaultExperiences

1 subcommand(s).

##### list

List /v1/appClips/{id}/appClipDefaultExperiences

#### get

Get /v1/appClips/{id}

### appCustomProductPageLocalizations

Manage appCustomProductPageLocalizations

7 subcommand(s).

#### appPreviewSets

Manage appCustomProductPageLocalizations appPreviewSets

1 subcommand(s).

##### list

List /v1/appCustomProductPageLocalizations/{id}/appPreviewSets

#### appScreenshotSets

Manage appCustomProductPageLocalizations appScreenshotSets

1 subcommand(s).

##### list

List /v1/appCustomProductPageLocalizations/{id}/appScreenshotSets

#### create

Create /v1/appCustomProductPageLocalizations

#### delete

Delete /v1/appCustomProductPageLocalizations/{id}

#### get

Get /v1/appCustomProductPageLocalizations/{id}

#### searchKeywords

Manage appCustomProductPageLocalizations searchKeywords

1 subcommand(s).

##### list

List /v1/appCustomProductPageLocalizations/{id}/searchKeywords

#### update

Update /v1/appCustomProductPageLocalizations/{id}

### appCustomProductPageVersions

Manage appCustomProductPageVersions

4 subcommand(s).

#### appCustomProductPageLocalizations

Manage appCustomProductPageVersions appCustomProductPageLocalizations

1 subcommand(s).

##### list

List /v1/appCustomProductPageVersions/{id}/appCustomProductPageLocalizations

#### create

Create /v1/appCustomProductPageVersions

#### get

Get /v1/appCustomProductPageVersions/{id}

#### update

Update /v1/appCustomProductPageVersions/{id}

### appCustomProductPages

Manage appCustomProductPages

5 subcommand(s).

#### appCustomProductPageVersions

Manage appCustomProductPages appCustomProductPageVersions

1 subcommand(s).

##### list

List /v1/appCustomProductPages/{id}/appCustomProductPageVersions

#### create

Create /v1/appCustomProductPages

#### delete

Delete /v1/appCustomProductPages/{id}

#### get

Get /v1/appCustomProductPages/{id}

#### update

Update /v1/appCustomProductPages/{id}

### appEncryptionDeclarationDocuments

Manage appEncryptionDeclarationDocuments

3 subcommand(s).

#### create

Create /v1/appEncryptionDeclarationDocuments

#### get

Get /v1/appEncryptionDeclarationDocuments/{id}

#### update

Update /v1/appEncryptionDeclarationDocuments/{id}

### appEncryptionDeclarations

Manage appEncryptionDeclarations

5 subcommand(s).

#### app

Manage appEncryptionDeclarations app

1 subcommand(s).

##### get

Get /v1/appEncryptionDeclarations/{id}/app

#### appEncryptionDeclarationDocument

Manage appEncryptionDeclarations appEncryptionDeclarationDocument

1 subcommand(s).

##### get

Get /v1/appEncryptionDeclarations/{id}/appEncryptionDeclarationDocument

#### create

Create /v1/appEncryptionDeclarations

#### get

Get /v1/appEncryptionDeclarations/{id}

#### list

List /v1/appEncryptionDeclarations

### appEventLocalizations

Manage appEventLocalizations

6 subcommand(s).

#### appEventScreenshots

Manage appEventLocalizations appEventScreenshots

1 subcommand(s).

##### list

List /v1/appEventLocalizations/{id}/appEventScreenshots

#### appEventVideoClips

Manage appEventLocalizations appEventVideoClips

1 subcommand(s).

##### list

List /v1/appEventLocalizations/{id}/appEventVideoClips

#### create

Create /v1/appEventLocalizations

#### delete

Delete /v1/appEventLocalizations/{id}

#### get

Get /v1/appEventLocalizations/{id}

#### update

Update /v1/appEventLocalizations/{id}

### appEventScreenshots

Manage appEventScreenshots

4 subcommand(s).

#### create

Create /v1/appEventScreenshots

#### delete

Delete /v1/appEventScreenshots/{id}

#### get

Get /v1/appEventScreenshots/{id}

#### update

Update /v1/appEventScreenshots/{id}

### appEventVideoClips

Manage appEventVideoClips

4 subcommand(s).

#### create

Create /v1/appEventVideoClips

#### delete

Delete /v1/appEventVideoClips/{id}

#### get

Get /v1/appEventVideoClips/{id}

#### update

Update /v1/appEventVideoClips/{id}

### appEvents

Manage appEvents

5 subcommand(s).

#### create

Create /v1/appEvents

#### delete

Delete /v1/appEvents/{id}

#### get

Get /v1/appEvents/{id}

#### localizations

Manage appEvents localizations

1 subcommand(s).

##### list

List /v1/appEvents/{id}/localizations

#### update

Update /v1/appEvents/{id}

### appInfoLocalizations

Manage appInfoLocalizations

4 subcommand(s).

#### create

Create /v1/appInfoLocalizations

#### delete

Delete /v1/appInfoLocalizations/{id}

#### get

Get /v1/appInfoLocalizations/{id}

#### update

Update /v1/appInfoLocalizations/{id}

### appInfos

Manage appInfos

11 subcommand(s).

#### ageRatingDeclaration

Manage appInfos ageRatingDeclaration

1 subcommand(s).

##### get

Get /v1/appInfos/{id}/ageRatingDeclaration

#### appInfoLocalizations

Manage appInfos appInfoLocalizations

1 subcommand(s).

##### list

List /v1/appInfos/{id}/appInfoLocalizations

#### get

Get /v1/appInfos/{id}

#### primaryCategory

Manage appInfos primaryCategory

1 subcommand(s).

##### get

Get /v1/appInfos/{id}/primaryCategory

#### primarySubcategoryOne

Manage appInfos primarySubcategoryOne

1 subcommand(s).

##### get

Get /v1/appInfos/{id}/primarySubcategoryOne

#### primarySubcategoryTwo

Manage appInfos primarySubcategoryTwo

1 subcommand(s).

##### get

Get /v1/appInfos/{id}/primarySubcategoryTwo

#### secondaryCategory

Manage appInfos secondaryCategory

1 subcommand(s).

##### get

Get /v1/appInfos/{id}/secondaryCategory

#### secondarySubcategoryOne

Manage appInfos secondarySubcategoryOne

1 subcommand(s).

##### get

Get /v1/appInfos/{id}/secondarySubcategoryOne

#### secondarySubcategoryTwo

Manage appInfos secondarySubcategoryTwo

1 subcommand(s).

##### get

Get /v1/appInfos/{id}/secondarySubcategoryTwo

#### territoryAgeRatings

Manage appInfos territoryAgeRatings

1 subcommand(s).

##### list

List /v1/appInfos/{id}/territoryAgeRatings

#### update

Update /v1/appInfos/{id}

### appPreviewSets

Manage appPreviewSets

4 subcommand(s).

#### appPreviews

Manage appPreviewSets appPreviews

1 subcommand(s).

##### list

List /v1/appPreviewSets/{id}/appPreviews

#### create

Create /v1/appPreviewSets

#### delete

Delete /v1/appPreviewSets/{id}

#### get

Get /v1/appPreviewSets/{id}

### appPreviews

Manage appPreviews

4 subcommand(s).

#### create

Create /v1/appPreviews

#### delete

Delete /v1/appPreviews/{id}

#### get

Get /v1/appPreviews/{id}

#### update

Update /v1/appPreviews/{id}

### appPricePointsV3

Manage appPricePointsV3

2 subcommand(s).

#### equalizations

Manage appPricePointsV3 equalizations

1 subcommand(s).

##### list

List /v3/appPricePoints/{id}/equalizations

#### get

Get /v3/appPricePoints/{id}

### appPriceSchedules

Manage appPriceSchedules

5 subcommand(s).

#### automaticPrices

Manage appPriceSchedules automaticPrices

1 subcommand(s).

##### list

List /v1/appPriceSchedules/{id}/automaticPrices

#### baseTerritory

Manage appPriceSchedules baseTerritory

1 subcommand(s).

##### get

Get /v1/appPriceSchedules/{id}/baseTerritory

#### create

Create /v1/appPriceSchedules

#### get

Get /v1/appPriceSchedules/{id}

#### manualPrices

Manage appPriceSchedules manualPrices

1 subcommand(s).

##### list

List /v1/appPriceSchedules/{id}/manualPrices

### appScreenshotSets

Manage appScreenshotSets

4 subcommand(s).

#### appScreenshots

Manage appScreenshotSets appScreenshots

1 subcommand(s).

##### list

List /v1/appScreenshotSets/{id}/appScreenshots

#### create

Create /v1/appScreenshotSets

#### delete

Delete /v1/appScreenshotSets/{id}

#### get

Get /v1/appScreenshotSets/{id}

### appScreenshots

Manage appScreenshots

4 subcommand(s).

#### create

Create /v1/appScreenshots

#### delete

Delete /v1/appScreenshots/{id}

#### get

Get /v1/appScreenshots/{id}

#### update

Update /v1/appScreenshots/{id}

### appStoreReviewAttachments

Manage appStoreReviewAttachments

4 subcommand(s).

#### create

Create /v1/appStoreReviewAttachments

#### delete

Delete /v1/appStoreReviewAttachments/{id}

#### get

Get /v1/appStoreReviewAttachments/{id}

#### update

Update /v1/appStoreReviewAttachments/{id}

### appStoreReviewDetails

Manage appStoreReviewDetails

4 subcommand(s).

#### appStoreReviewAttachments

Manage appStoreReviewDetails appStoreReviewAttachments

1 subcommand(s).

##### list

List /v1/appStoreReviewDetails/{id}/appStoreReviewAttachments

#### create

Create /v1/appStoreReviewDetails

#### get

Get /v1/appStoreReviewDetails/{id}

#### update

Update /v1/appStoreReviewDetails/{id}

### appStoreVersionExperimentTreatmentLocalizations

Manage appStoreVersionExperimentTreatmentLocalizations

5 subcommand(s).

#### appPreviewSets

Manage appStoreVersionExperimentTreatmentLocalizations appPreviewSets

1 subcommand(s).

##### list

List /v1/appStoreVersionExperimentTreatmentLocalizations/{id}/appPreviewSets

#### appScreenshotSets

Manage appStoreVersionExperimentTreatmentLocalizations appScreenshotSets

1 subcommand(s).

##### list

List /v1/appStoreVersionExperimentTreatmentLocalizations/{id}/appScreenshotSets

#### create

Create /v1/appStoreVersionExperimentTreatmentLocalizations

#### delete

Delete /v1/appStoreVersionExperimentTreatmentLocalizations/{id}

#### get

Get /v1/appStoreVersionExperimentTreatmentLocalizations/{id}

### appStoreVersionExperimentTreatments

Manage appStoreVersionExperimentTreatments

5 subcommand(s).

#### appStoreVersionExperimentTreatmentLocalizations

Manage appStoreVersionExperimentTreatments appStoreVersionExperimentTreatmentLocalizations

1 subcommand(s).

##### list

List /v1/appStoreVersionExperimentTreatments/{id}/appStoreVersionExperimentTreatmentLocalizations

#### create

Create /v1/appStoreVersionExperimentTreatments

#### delete

Delete /v1/appStoreVersionExperimentTreatments/{id}

#### get

Get /v1/appStoreVersionExperimentTreatments/{id}

#### update

Update /v1/appStoreVersionExperimentTreatments/{id}

### appStoreVersionExperiments

Manage appStoreVersionExperiments

5 subcommand(s).

#### appStoreVersionExperimentTreatments

Manage appStoreVersionExperiments appStoreVersionExperimentTreatments

1 subcommand(s).

##### list

List /v1/appStoreVersionExperiments/{id}/appStoreVersionExperimentTreatments

#### create

Create /v1/appStoreVersionExperiments

#### delete

Delete /v1/appStoreVersionExperiments/{id}

#### get

Get /v1/appStoreVersionExperiments/{id}

#### update

Update /v1/appStoreVersionExperiments/{id}

### appStoreVersionExperimentsV2

Manage appStoreVersionExperimentsV2

5 subcommand(s).

#### appStoreVersionExperimentTreatments

Manage appStoreVersionExperimentsV2 appStoreVersionExperimentTreatments

1 subcommand(s).

##### list

List /v2/appStoreVersionExperiments/{id}/appStoreVersionExperimentTreatments

#### create

Create /v2/appStoreVersionExperiments

#### delete

Delete /v2/appStoreVersionExperiments/{id}

#### get

Get /v2/appStoreVersionExperiments/{id}

#### update

Update /v2/appStoreVersionExperiments/{id}

### appStoreVersionLocalizations

Manage appStoreVersionLocalizations

7 subcommand(s).

#### appPreviewSets

Manage appStoreVersionLocalizations appPreviewSets

1 subcommand(s).

##### list

List /v1/appStoreVersionLocalizations/{id}/appPreviewSets

#### appScreenshotSets

Manage appStoreVersionLocalizations appScreenshotSets

1 subcommand(s).

##### list

List /v1/appStoreVersionLocalizations/{id}/appScreenshotSets

#### create

Create /v1/appStoreVersionLocalizations

#### delete

Delete /v1/appStoreVersionLocalizations/{id}

#### get

Get /v1/appStoreVersionLocalizations/{id}

#### searchKeywords

Manage appStoreVersionLocalizations searchKeywords

1 subcommand(s).

##### list

List /v1/appStoreVersionLocalizations/{id}/searchKeywords

#### update

Update /v1/appStoreVersionLocalizations/{id}

### appStoreVersionPhasedReleases

Manage appStoreVersionPhasedReleases

3 subcommand(s).

#### create

Create /v1/appStoreVersionPhasedReleases

#### delete

Delete /v1/appStoreVersionPhasedReleases/{id}

#### update

Update /v1/appStoreVersionPhasedReleases/{id}

### appStoreVersionPromotions

Manage appStoreVersionPromotions

1 subcommand(s).

#### create

Create /v1/appStoreVersionPromotions

### appStoreVersionReleaseRequests

Manage appStoreVersionReleaseRequests

1 subcommand(s).

#### create

Create /v1/appStoreVersionReleaseRequests

### appStoreVersionSubmissions

Manage appStoreVersionSubmissions

1 subcommand(s).

#### delete

Delete /v1/appStoreVersionSubmissions/{id}

### appStoreVersions

Manage appStoreVersions

16 subcommand(s).

#### alternativeDistributionPackage

Manage appStoreVersions alternativeDistributionPackage

1 subcommand(s).

##### get

Get /v1/appStoreVersions/{id}/alternativeDistributionPackage

#### appClipDefaultExperience

Manage appStoreVersions appClipDefaultExperience

1 subcommand(s).

##### get

Get /v1/appStoreVersions/{id}/appClipDefaultExperience

#### appStoreReviewDetail

Manage appStoreVersions appStoreReviewDetail

1 subcommand(s).

##### get

Get /v1/appStoreVersions/{id}/appStoreReviewDetail

#### appStoreVersionExperiments

Manage appStoreVersions appStoreVersionExperiments

1 subcommand(s).

##### list

List /v1/appStoreVersions/{id}/appStoreVersionExperiments

#### appStoreVersionExperimentsV2

Manage appStoreVersions appStoreVersionExperimentsV2

1 subcommand(s).

##### list

List /v1/appStoreVersions/{id}/appStoreVersionExperimentsV2

#### appStoreVersionLocalizations

Manage appStoreVersions appStoreVersionLocalizations

1 subcommand(s).

##### list

List /v1/appStoreVersions/{id}/appStoreVersionLocalizations

#### appStoreVersionPhasedRelease

Manage appStoreVersions appStoreVersionPhasedRelease

1 subcommand(s).

##### get

Get /v1/appStoreVersions/{id}/appStoreVersionPhasedRelease

#### appStoreVersionSubmission

Manage appStoreVersions appStoreVersionSubmission

1 subcommand(s).

##### get

Get /v1/appStoreVersions/{id}/appStoreVersionSubmission

#### build

Manage appStoreVersions build

1 subcommand(s).

##### get

Get /v1/appStoreVersions/{id}/build

#### create

Create /v1/appStoreVersions

#### customerReviews

Manage appStoreVersions customerReviews

1 subcommand(s).

##### list

List /v1/appStoreVersions/{id}/customerReviews

#### delete

Delete /v1/appStoreVersions/{id}

#### gameCenterAppVersion

Manage appStoreVersions gameCenterAppVersion

1 subcommand(s).

##### get

Get /v1/appStoreVersions/{id}/gameCenterAppVersion

#### get

Get /v1/appStoreVersions/{id}

#### routingAppCoverage

Manage appStoreVersions routingAppCoverage

1 subcommand(s).

##### get

Get /v1/appStoreVersions/{id}/routingAppCoverage

#### update

Update /v1/appStoreVersions/{id}

### appTags

Manage appTags

2 subcommand(s).

#### territories

Manage appTags territories

1 subcommand(s).

##### list

List /v1/appTags/{id}/territories

#### update

Update /v1/appTags/{id}

### apps

Manage apps

45 subcommand(s).

#### accessibilityDeclarations

Manage apps accessibilityDeclarations

1 subcommand(s).

##### list

List /v1/apps/{id}/accessibilityDeclarations

#### alternativeDistributionKey

Manage apps alternativeDistributionKey

1 subcommand(s).

##### get

Get /v1/apps/{id}/alternativeDistributionKey

#### analyticsReportRequests

Manage apps analyticsReportRequests

1 subcommand(s).

##### list

List /v1/apps/{id}/analyticsReportRequests

#### androidToIosAppMappingDetails

Manage apps androidToIosAppMappingDetails

1 subcommand(s).

##### list

List /v1/apps/{id}/androidToIosAppMappingDetails

#### appAvailabilityV2

Manage apps appAvailabilityV2

1 subcommand(s).

##### get

Get /v1/apps/{id}/appAvailabilityV2

#### appClips

Manage apps appClips

1 subcommand(s).

##### list

List /v1/apps/{id}/appClips

#### appCustomProductPages

Manage apps appCustomProductPages

1 subcommand(s).

##### list

List /v1/apps/{id}/appCustomProductPages

#### appEncryptionDeclarations

Manage apps appEncryptionDeclarations

1 subcommand(s).

##### list

List /v1/apps/{id}/appEncryptionDeclarations

#### appEvents

Manage apps appEvents

1 subcommand(s).

##### list

List /v1/apps/{id}/appEvents

#### appInfos

Manage apps appInfos

1 subcommand(s).

##### list

List /v1/apps/{id}/appInfos

#### appPricePoints

Manage apps appPricePoints

1 subcommand(s).

##### list

List /v1/apps/{id}/appPricePoints

#### appPriceSchedule

Manage apps appPriceSchedule

1 subcommand(s).

##### get

Get /v1/apps/{id}/appPriceSchedule

#### appStoreVersionExperimentsV2

Manage apps appStoreVersionExperimentsV2

1 subcommand(s).

##### list

List /v1/apps/{id}/appStoreVersionExperimentsV2

#### appStoreVersions

Manage apps appStoreVersions

1 subcommand(s).

##### list

List /v1/apps/{id}/appStoreVersions

#### appTags

Manage apps appTags

1 subcommand(s).

##### list

List /v1/apps/{id}/appTags

#### backgroundAssets

Manage apps backgroundAssets

1 subcommand(s).

##### list

List /v1/apps/{id}/backgroundAssets

#### betaAppLocalizations

Manage apps betaAppLocalizations

1 subcommand(s).

##### list

List /v1/apps/{id}/betaAppLocalizations

#### betaAppReviewDetail

Manage apps betaAppReviewDetail

1 subcommand(s).

##### get

Get /v1/apps/{id}/betaAppReviewDetail

#### betaFeedbackCrashSubmissions

Manage apps betaFeedbackCrashSubmissions

1 subcommand(s).

##### list

List /v1/apps/{id}/betaFeedbackCrashSubmissions

#### betaFeedbackScreenshotSubmissions

Manage apps betaFeedbackScreenshotSubmissions

1 subcommand(s).

##### list

List /v1/apps/{id}/betaFeedbackScreenshotSubmissions

#### betaGroups

Manage apps betaGroups

1 subcommand(s).

##### list

List /v1/apps/{id}/betaGroups

#### betaLicenseAgreement

Manage apps betaLicenseAgreement

1 subcommand(s).

##### get

Get /v1/apps/{id}/betaLicenseAgreement

#### betaTesterUsages

Manage apps betaTesterUsages

1 subcommand(s).

##### metrics

Metrics /v1/apps/{id}/metrics/betaTesterUsages

#### buildUploads

Manage apps buildUploads

1 subcommand(s).

##### list

List /v1/apps/{id}/buildUploads

#### builds

Manage apps builds

1 subcommand(s).

##### list

List /v1/apps/{id}/builds

#### ciProduct

Manage apps ciProduct

1 subcommand(s).

##### get

Get /v1/apps/{id}/ciProduct

#### customerReviewSummarizations

Manage apps customerReviewSummarizations

1 subcommand(s).

##### list

List /v1/apps/{id}/customerReviewSummarizations

#### customerReviews

Manage apps customerReviews

1 subcommand(s).

##### list

List /v1/apps/{id}/customerReviews

#### endUserLicenseAgreement

Manage apps endUserLicenseAgreement

1 subcommand(s).

##### get

Get /v1/apps/{id}/endUserLicenseAgreement

#### gameCenterDetail

Manage apps gameCenterDetail

1 subcommand(s).

##### get

Get /v1/apps/{id}/gameCenterDetail

#### gameCenterEnabledVersions

Manage apps gameCenterEnabledVersions

1 subcommand(s).

##### list

List /v1/apps/{id}/gameCenterEnabledVersions

#### get

Get /v1/apps/{id}

#### inAppPurchases

Manage apps inAppPurchases

1 subcommand(s).

##### list

List /v1/apps/{id}/inAppPurchases

#### inAppPurchasesV2

Manage apps inAppPurchasesV2

1 subcommand(s).

##### list

List /v1/apps/{id}/inAppPurchasesV2

#### list

List /v1/apps

#### marketplaceSearchDetail

Manage apps marketplaceSearchDetail

1 subcommand(s).

##### get

Get /v1/apps/{id}/marketplaceSearchDetail

#### perfPowerMetrics

Manage apps perfPowerMetrics

1 subcommand(s).

##### list

List /v1/apps/{id}/perfPowerMetrics

#### preReleaseVersions

Manage apps preReleaseVersions

1 subcommand(s).

##### list

List /v1/apps/{id}/preReleaseVersions

#### promotedPurchases

Manage apps promotedPurchases

1 subcommand(s).

##### list

List /v1/apps/{id}/promotedPurchases

#### reviewSubmissions

Manage apps reviewSubmissions

1 subcommand(s).

##### list

List /v1/apps/{id}/reviewSubmissions

#### searchKeywords

Manage apps searchKeywords

1 subcommand(s).

##### list

List /v1/apps/{id}/searchKeywords

#### subscriptionGracePeriod

Manage apps subscriptionGracePeriod

1 subcommand(s).

##### get

Get /v1/apps/{id}/subscriptionGracePeriod

#### subscriptionGroups

Manage apps subscriptionGroups

1 subcommand(s).

##### list

List /v1/apps/{id}/subscriptionGroups

#### update

Update /v1/apps/{id}

#### webhooks

Manage apps webhooks

1 subcommand(s).

##### list

List /v1/apps/{id}/webhooks

### backgroundAssetUploadFiles

Manage backgroundAssetUploadFiles

3 subcommand(s).

#### create

Create /v1/backgroundAssetUploadFiles

#### get

Get /v1/backgroundAssetUploadFiles/{id}

#### update

Update /v1/backgroundAssetUploadFiles/{id}

### backgroundAssetVersionAppStoreReleases

Manage backgroundAssetVersionAppStoreReleases

1 subcommand(s).

#### get

Get /v1/backgroundAssetVersionAppStoreReleases/{id}

### backgroundAssetVersionExternalBetaReleases

Manage backgroundAssetVersionExternalBetaReleases

1 subcommand(s).

#### get

Get /v1/backgroundAssetVersionExternalBetaReleases/{id}

### backgroundAssetVersionInternalBetaReleases

Manage backgroundAssetVersionInternalBetaReleases

1 subcommand(s).

#### get

Get /v1/backgroundAssetVersionInternalBetaReleases/{id}

### backgroundAssetVersions

Manage backgroundAssetVersions

3 subcommand(s).

#### backgroundAssetUploadFiles

Manage backgroundAssetVersions backgroundAssetUploadFiles

1 subcommand(s).

##### list

List /v1/backgroundAssetVersions/{id}/backgroundAssetUploadFiles

#### create

Create /v1/backgroundAssetVersions

#### get

Get /v1/backgroundAssetVersions/{id}

### backgroundAssets

Manage backgroundAssets

4 subcommand(s).

#### create

Create /v1/backgroundAssets

#### get

Get /v1/backgroundAssets/{id}

#### update

Update /v1/backgroundAssets/{id}

#### versions

Manage backgroundAssets versions

1 subcommand(s).

##### list

List /v1/backgroundAssets/{id}/versions

### betaAppClipInvocationLocalizations

Manage betaAppClipInvocationLocalizations

3 subcommand(s).

#### create

Create /v1/betaAppClipInvocationLocalizations

#### delete

Delete /v1/betaAppClipInvocationLocalizations/{id}

#### update

Update /v1/betaAppClipInvocationLocalizations/{id}

### betaAppClipInvocations

Manage betaAppClipInvocations

4 subcommand(s).

#### create

Create /v1/betaAppClipInvocations

#### delete

Delete /v1/betaAppClipInvocations/{id}

#### get

Get /v1/betaAppClipInvocations/{id}

#### update

Update /v1/betaAppClipInvocations/{id}

### betaAppLocalizations

Manage betaAppLocalizations

6 subcommand(s).

#### app

Manage betaAppLocalizations app

1 subcommand(s).

##### get

Get /v1/betaAppLocalizations/{id}/app

#### create

Create /v1/betaAppLocalizations

#### delete

Delete /v1/betaAppLocalizations/{id}

#### get

Get /v1/betaAppLocalizations/{id}

#### list

List /v1/betaAppLocalizations

#### update

Update /v1/betaAppLocalizations/{id}

### betaAppReviewDetails

Manage betaAppReviewDetails

4 subcommand(s).

#### app

Manage betaAppReviewDetails app

1 subcommand(s).

##### get

Get /v1/betaAppReviewDetails/{id}/app

#### get

Get /v1/betaAppReviewDetails/{id}

#### list

List /v1/betaAppReviewDetails

#### update

Update /v1/betaAppReviewDetails/{id}

### betaAppReviewSubmissions

Manage betaAppReviewSubmissions

4 subcommand(s).

#### build

Manage betaAppReviewSubmissions build

1 subcommand(s).

##### get

Get /v1/betaAppReviewSubmissions/{id}/build

#### create

Create /v1/betaAppReviewSubmissions

#### get

Get /v1/betaAppReviewSubmissions/{id}

#### list

List /v1/betaAppReviewSubmissions

### betaBuildLocalizations

Manage betaBuildLocalizations

6 subcommand(s).

#### build

Manage betaBuildLocalizations build

1 subcommand(s).

##### get

Get /v1/betaBuildLocalizations/{id}/build

#### create

Create /v1/betaBuildLocalizations

#### delete

Delete /v1/betaBuildLocalizations/{id}

#### get

Get /v1/betaBuildLocalizations/{id}

#### list

List /v1/betaBuildLocalizations

#### update

Update /v1/betaBuildLocalizations/{id}

### betaCrashLogs

Manage betaCrashLogs

1 subcommand(s).

#### get

Get /v1/betaCrashLogs/{id}

### betaFeedbackCrashSubmissions

Manage betaFeedbackCrashSubmissions

3 subcommand(s).

#### crashLog

Manage betaFeedbackCrashSubmissions crashLog

1 subcommand(s).

##### get

Get /v1/betaFeedbackCrashSubmissions/{id}/crashLog

#### delete

Delete /v1/betaFeedbackCrashSubmissions/{id}

#### get

Get /v1/betaFeedbackCrashSubmissions/{id}

### betaFeedbackScreenshotSubmissions

Manage betaFeedbackScreenshotSubmissions

2 subcommand(s).

#### delete

Delete /v1/betaFeedbackScreenshotSubmissions/{id}

#### get

Get /v1/betaFeedbackScreenshotSubmissions/{id}

### betaGroups

Manage betaGroups

12 subcommand(s).

#### app

Manage betaGroups app

1 subcommand(s).

##### get

Get /v1/betaGroups/{id}/app

#### betaRecruitmentCriteria

Manage betaGroups betaRecruitmentCriteria

1 subcommand(s).

##### get

Get /v1/betaGroups/{id}/betaRecruitmentCriteria

#### betaRecruitmentCriterionCompatibleBuildCheck

Manage betaGroups betaRecruitmentCriterionCompatibleBuildCheck

1 subcommand(s).

##### get

Get /v1/betaGroups/{id}/betaRecruitmentCriterionCompatibleBuildCheck

#### betaTesterUsages

Manage betaGroups betaTesterUsages

1 subcommand(s).

##### metrics

Metrics /v1/betaGroups/{id}/metrics/betaTesterUsages

#### betaTesters

Manage betaGroups betaTesters

1 subcommand(s).

##### list

List /v1/betaGroups/{id}/betaTesters

#### builds

Manage betaGroups builds

1 subcommand(s).

##### list

List /v1/betaGroups/{id}/builds

#### create

Create /v1/betaGroups

#### delete

Delete /v1/betaGroups/{id}

#### get

Get /v1/betaGroups/{id}

#### list

List /v1/betaGroups

#### publicLinkUsages

Manage betaGroups publicLinkUsages

1 subcommand(s).

##### metrics

Metrics /v1/betaGroups/{id}/metrics/publicLinkUsages

#### update

Update /v1/betaGroups/{id}

### betaLicenseAgreements

Manage betaLicenseAgreements

4 subcommand(s).

#### app

Manage betaLicenseAgreements app

1 subcommand(s).

##### get

Get /v1/betaLicenseAgreements/{id}/app

#### get

Get /v1/betaLicenseAgreements/{id}

#### list

List /v1/betaLicenseAgreements

#### update

Update /v1/betaLicenseAgreements/{id}

### betaRecruitmentCriteria

Manage betaRecruitmentCriteria

3 subcommand(s).

#### create

Create /v1/betaRecruitmentCriteria

#### delete

Delete /v1/betaRecruitmentCriteria/{id}

#### update

Update /v1/betaRecruitmentCriteria/{id}

### betaRecruitmentCriterionOptions

Manage betaRecruitmentCriterionOptions

1 subcommand(s).

#### list

List /v1/betaRecruitmentCriterionOptions

### betaTesterInvitations

Manage betaTesterInvitations

1 subcommand(s).

#### create

Create /v1/betaTesterInvitations

### betaTesters

Manage betaTesters

8 subcommand(s).

#### apps

Manage betaTesters apps

1 subcommand(s).

##### list

List /v1/betaTesters/{id}/apps

#### betaGroups

Manage betaTesters betaGroups

1 subcommand(s).

##### list

List /v1/betaTesters/{id}/betaGroups

#### betaTesterUsages

Manage betaTesters betaTesterUsages

1 subcommand(s).

##### metrics

Metrics /v1/betaTesters/{id}/metrics/betaTesterUsages

#### builds

Manage betaTesters builds

1 subcommand(s).

##### list

List /v1/betaTesters/{id}/builds

#### create

Create /v1/betaTesters

#### delete

Delete /v1/betaTesters/{id}

#### get

Get /v1/betaTesters/{id}

#### list

List /v1/betaTesters

### buildBetaDetails

Manage buildBetaDetails

4 subcommand(s).

#### build

Manage buildBetaDetails build

1 subcommand(s).

##### get

Get /v1/buildBetaDetails/{id}/build

#### get

Get /v1/buildBetaDetails/{id}

#### list

List /v1/buildBetaDetails

#### update

Update /v1/buildBetaDetails/{id}

### buildBetaNotifications

Manage buildBetaNotifications

1 subcommand(s).

#### create

Create /v1/buildBetaNotifications

### buildBundles

Manage buildBundles

4 subcommand(s).

#### appClipDomainCacheStatus

Manage buildBundles appClipDomainCacheStatus

1 subcommand(s).

##### get

Get /v1/buildBundles/{id}/appClipDomainCacheStatus

#### appClipDomainDebugStatus

Manage buildBundles appClipDomainDebugStatus

1 subcommand(s).

##### get

Get /v1/buildBundles/{id}/appClipDomainDebugStatus

#### betaAppClipInvocations

Manage buildBundles betaAppClipInvocations

1 subcommand(s).

##### list

List /v1/buildBundles/{id}/betaAppClipInvocations

#### buildBundleFileSizes

Manage buildBundles buildBundleFileSizes

1 subcommand(s).

##### list

List /v1/buildBundles/{id}/buildBundleFileSizes

### buildUploadFiles

Manage buildUploadFiles

3 subcommand(s).

#### create

Create /v1/buildUploadFiles

#### get

Get /v1/buildUploadFiles/{id}

#### update

Update /v1/buildUploadFiles/{id}

### buildUploads

Manage buildUploads

4 subcommand(s).

#### buildUploadFiles

Manage buildUploads buildUploadFiles

1 subcommand(s).

##### list

List /v1/buildUploads/{id}/buildUploadFiles

#### create

Create /v1/buildUploads

#### delete

Delete /v1/buildUploads/{id}

#### get

Get /v1/buildUploads/{id}

### builds

Manage builds

15 subcommand(s).

#### app

Manage builds app

1 subcommand(s).

##### get

Get /v1/builds/{id}/app

#### appEncryptionDeclaration

Manage builds appEncryptionDeclaration

1 subcommand(s).

##### get

Get /v1/builds/{id}/appEncryptionDeclaration

#### appStoreVersion

Manage builds appStoreVersion

1 subcommand(s).

##### get

Get /v1/builds/{id}/appStoreVersion

#### betaAppReviewSubmission

Manage builds betaAppReviewSubmission

1 subcommand(s).

##### get

Get /v1/builds/{id}/betaAppReviewSubmission

#### betaBuildLocalizations

Manage builds betaBuildLocalizations

1 subcommand(s).

##### list

List /v1/builds/{id}/betaBuildLocalizations

#### betaBuildUsages

Manage builds betaBuildUsages

1 subcommand(s).

##### metrics

Metrics /v1/builds/{id}/metrics/betaBuildUsages

#### buildBetaDetail

Manage builds buildBetaDetail

1 subcommand(s).

##### get

Get /v1/builds/{id}/buildBetaDetail

#### diagnosticSignatures

Manage builds diagnosticSignatures

1 subcommand(s).

##### list

List /v1/builds/{id}/diagnosticSignatures

#### get

Get /v1/builds/{id}

#### icons

Manage builds icons

1 subcommand(s).

##### list

List /v1/builds/{id}/icons

#### individualTesters

Manage builds individualTesters

1 subcommand(s).

##### list

List /v1/builds/{id}/individualTesters

#### list

List /v1/builds

#### perfPowerMetrics

Manage builds perfPowerMetrics

1 subcommand(s).

##### list

List /v1/builds/{id}/perfPowerMetrics

#### preReleaseVersion

Manage builds preReleaseVersion

1 subcommand(s).

##### get

Get /v1/builds/{id}/preReleaseVersion

#### update

Update /v1/builds/{id}

### bundleIdCapabilities

Manage bundleIdCapabilities

3 subcommand(s).

#### create

Create /v1/bundleIdCapabilities

#### delete

Delete /v1/bundleIdCapabilities/{id}

#### update

Update /v1/bundleIdCapabilities/{id}

### bundleIds

Manage bundleIds

8 subcommand(s).

#### app

Manage bundleIds app

1 subcommand(s).

##### get

Get /v1/bundleIds/{id}/app

#### bundleIdCapabilities

Manage bundleIds bundleIdCapabilities

1 subcommand(s).

##### list

List /v1/bundleIds/{id}/bundleIdCapabilities

#### create

Create /v1/bundleIds

#### delete

Delete /v1/bundleIds/{id}

#### get

Get /v1/bundleIds/{id}

#### list

List /v1/bundleIds

#### profiles

Manage bundleIds profiles

1 subcommand(s).

##### list

List /v1/bundleIds/{id}/profiles

#### update

Update /v1/bundleIds/{id}

### certificates

Manage certificates

6 subcommand(s).

#### create

Create /v1/certificates

#### delete

Delete /v1/certificates/{id}

#### get

Get /v1/certificates/{id}

#### list

List /v1/certificates

#### passTypeId

Manage certificates passTypeId

1 subcommand(s).

##### get

Get /v1/certificates/{id}/passTypeId

#### update

Update /v1/certificates/{id}

### ciArtifacts

Manage ciArtifacts

1 subcommand(s).

#### get

Get /v1/ciArtifacts/{id}

### ciBuildActions

Manage ciBuildActions

5 subcommand(s).

#### artifacts

Manage ciBuildActions artifacts

1 subcommand(s).

##### list

List /v1/ciBuildActions/{id}/artifacts

#### buildRun

Manage ciBuildActions buildRun

1 subcommand(s).

##### get

Get /v1/ciBuildActions/{id}/buildRun

#### get

Get /v1/ciBuildActions/{id}

#### issues

Manage ciBuildActions issues

1 subcommand(s).

##### list

List /v1/ciBuildActions/{id}/issues

#### testResults

Manage ciBuildActions testResults

1 subcommand(s).

##### list

List /v1/ciBuildActions/{id}/testResults

### ciBuildRuns

Manage ciBuildRuns

4 subcommand(s).

#### actions

Manage ciBuildRuns actions

1 subcommand(s).

##### list

List /v1/ciBuildRuns/{id}/actions

#### builds

Manage ciBuildRuns builds

1 subcommand(s).

##### list

List /v1/ciBuildRuns/{id}/builds

#### create

Create /v1/ciBuildRuns

#### get

Get /v1/ciBuildRuns/{id}

### ciIssues

Manage ciIssues

1 subcommand(s).

#### get

Get /v1/ciIssues/{id}

### ciMacOsVersions

Manage ciMacOsVersions

3 subcommand(s).

#### get

Get /v1/ciMacOsVersions/{id}

#### list

List /v1/ciMacOsVersions

#### xcodeVersions

Manage ciMacOsVersions xcodeVersions

1 subcommand(s).

##### list

List /v1/ciMacOsVersions/{id}/xcodeVersions

### ciProducts

Manage ciProducts

8 subcommand(s).

#### additionalRepositories

Manage ciProducts additionalRepositories

1 subcommand(s).

##### list

List /v1/ciProducts/{id}/additionalRepositories

#### app

Manage ciProducts app

1 subcommand(s).

##### get

Get /v1/ciProducts/{id}/app

#### buildRuns

Manage ciProducts buildRuns

1 subcommand(s).

##### list

List /v1/ciProducts/{id}/buildRuns

#### delete

Delete /v1/ciProducts/{id}

#### get

Get /v1/ciProducts/{id}

#### list

List /v1/ciProducts

#### primaryRepositories

Manage ciProducts primaryRepositories

1 subcommand(s).

##### list

List /v1/ciProducts/{id}/primaryRepositories

#### workflows

Manage ciProducts workflows

1 subcommand(s).

##### list

List /v1/ciProducts/{id}/workflows

### ciTestResults

Manage ciTestResults

1 subcommand(s).

#### get

Get /v1/ciTestResults/{id}

### ciWorkflows

Manage ciWorkflows

6 subcommand(s).

#### buildRuns

Manage ciWorkflows buildRuns

1 subcommand(s).

##### list

List /v1/ciWorkflows/{id}/buildRuns

#### create

Create /v1/ciWorkflows

#### delete

Delete /v1/ciWorkflows/{id}

#### get

Get /v1/ciWorkflows/{id}

#### repository

Manage ciWorkflows repository

1 subcommand(s).

##### get

Get /v1/ciWorkflows/{id}/repository

#### update

Update /v1/ciWorkflows/{id}

### ciXcodeVersions

Manage ciXcodeVersions

3 subcommand(s).

#### get

Get /v1/ciXcodeVersions/{id}

#### list

List /v1/ciXcodeVersions

#### macOsVersions

Manage ciXcodeVersions macOsVersions

1 subcommand(s).

##### list

List /v1/ciXcodeVersions/{id}/macOsVersions

### customerReviewResponses

Manage customerReviewResponses

3 subcommand(s).

#### create

Create /v1/customerReviewResponses

#### delete

Delete /v1/customerReviewResponses/{id}

#### get

Get /v1/customerReviewResponses/{id}

### customerReviews

Manage customerReviews

2 subcommand(s).

#### get

Get /v1/customerReviews/{id}

#### response

Manage customerReviews response

1 subcommand(s).

##### get

Get /v1/customerReviews/{id}/response

### devices

Manage devices

4 subcommand(s).

#### create

Create /v1/devices

#### get

Get /v1/devices/{id}

#### list

List /v1/devices

#### update

Update /v1/devices/{id}

### diagnosticSignatures

Manage diagnosticSignatures

1 subcommand(s).

#### logs

Manage diagnosticSignatures logs

1 subcommand(s).

##### list

List /v1/diagnosticSignatures/{id}/logs

### endAppAvailabilityPreOrders

Manage endAppAvailabilityPreOrders

1 subcommand(s).

#### create

Create /v1/endAppAvailabilityPreOrders

### endUserLicenseAgreements

Manage endUserLicenseAgreements

5 subcommand(s).

#### create

Create /v1/endUserLicenseAgreements

#### delete

Delete /v1/endUserLicenseAgreements/{id}

#### get

Get /v1/endUserLicenseAgreements/{id}

#### territories

Manage endUserLicenseAgreements territories

1 subcommand(s).

##### list

List /v1/endUserLicenseAgreements/{id}/territories

#### update

Update /v1/endUserLicenseAgreements/{id}

### financeReports

Manage financeReports

1 subcommand(s).

#### list

List /v1/financeReports

### gameCenterAchievementImages

Manage gameCenterAchievementImages

4 subcommand(s).

#### create

Create /v1/gameCenterAchievementImages

#### delete

Delete /v1/gameCenterAchievementImages/{id}

#### get

Get /v1/gameCenterAchievementImages/{id}

#### update

Update /v1/gameCenterAchievementImages/{id}

### gameCenterAchievementImagesV2

Manage gameCenterAchievementImagesV2

4 subcommand(s).

#### create

Create /v2/gameCenterAchievementImages

#### delete

Delete /v2/gameCenterAchievementImages/{id}

#### get

Get /v2/gameCenterAchievementImages/{id}

#### update

Update /v2/gameCenterAchievementImages/{id}

### gameCenterAchievementLocalizations

Manage gameCenterAchievementLocalizations

6 subcommand(s).

#### create

Create /v1/gameCenterAchievementLocalizations

#### delete

Delete /v1/gameCenterAchievementLocalizations/{id}

#### gameCenterAchievement

Manage gameCenterAchievementLocalizations gameCenterAchievement

1 subcommand(s).

##### get

Get /v1/gameCenterAchievementLocalizations/{id}/gameCenterAchievement

#### gameCenterAchievementImage

Manage gameCenterAchievementLocalizations gameCenterAchievementImage

1 subcommand(s).

##### get

Get /v1/gameCenterAchievementLocalizations/{id}/gameCenterAchievementImage

#### get

Get /v1/gameCenterAchievementLocalizations/{id}

#### update

Update /v1/gameCenterAchievementLocalizations/{id}

### gameCenterAchievementLocalizationsV2

Manage gameCenterAchievementLocalizationsV2

5 subcommand(s).

#### create

Create /v2/gameCenterAchievementLocalizations

#### delete

Delete /v2/gameCenterAchievementLocalizations/{id}

#### get

Get /v2/gameCenterAchievementLocalizations/{id}

#### image

Manage gameCenterAchievementLocalizationsV2 image

1 subcommand(s).

##### get

Get /v2/gameCenterAchievementLocalizations/{id}/image

#### update

Update /v2/gameCenterAchievementLocalizations/{id}

### gameCenterAchievementReleases

Manage gameCenterAchievementReleases

3 subcommand(s).

#### create

Create /v1/gameCenterAchievementReleases

#### delete

Delete /v1/gameCenterAchievementReleases/{id}

#### get

Get /v1/gameCenterAchievementReleases/{id}

### gameCenterAchievementVersionsV2

Manage gameCenterAchievementVersionsV2

3 subcommand(s).

#### create

Create /v2/gameCenterAchievementVersions

#### get

Get /v2/gameCenterAchievementVersions/{id}

#### localizations

Manage gameCenterAchievementVersionsV2 localizations

1 subcommand(s).

##### list

List /v2/gameCenterAchievementVersions/{id}/localizations

### gameCenterAchievements

Manage gameCenterAchievements

7 subcommand(s).

#### create

Create /v1/gameCenterAchievements

#### delete

Delete /v1/gameCenterAchievements/{id}

#### get

Get /v1/gameCenterAchievements/{id}

#### groupAchievement

Manage gameCenterAchievements groupAchievement

1 subcommand(s).

##### get

Get /v1/gameCenterAchievements/{id}/groupAchievement

#### localizations

Manage gameCenterAchievements localizations

1 subcommand(s).

##### list

List /v1/gameCenterAchievements/{id}/localizations

#### releases

Manage gameCenterAchievements releases

1 subcommand(s).

##### list

List /v1/gameCenterAchievements/{id}/releases

#### update

Update /v1/gameCenterAchievements/{id}

### gameCenterAchievementsV2

Manage gameCenterAchievementsV2

5 subcommand(s).

#### create

Create /v2/gameCenterAchievements

#### delete

Delete /v2/gameCenterAchievements/{id}

#### get

Get /v2/gameCenterAchievements/{id}

#### update

Update /v2/gameCenterAchievements/{id}

#### versions

Manage gameCenterAchievementsV2 versions

1 subcommand(s).

##### list

List /v2/gameCenterAchievements/{id}/versions

### gameCenterActivities

Manage gameCenterActivities

5 subcommand(s).

#### create

Create /v1/gameCenterActivities

#### delete

Delete /v1/gameCenterActivities/{id}

#### get

Get /v1/gameCenterActivities/{id}

#### update

Update /v1/gameCenterActivities/{id}

#### versions

Manage gameCenterActivities versions

1 subcommand(s).

##### list

List /v1/gameCenterActivities/{id}/versions

### gameCenterActivityImages

Manage gameCenterActivityImages

4 subcommand(s).

#### create

Create /v1/gameCenterActivityImages

#### delete

Delete /v1/gameCenterActivityImages/{id}

#### get

Get /v1/gameCenterActivityImages/{id}

#### update

Update /v1/gameCenterActivityImages/{id}

### gameCenterActivityLocalizations

Manage gameCenterActivityLocalizations

5 subcommand(s).

#### create

Create /v1/gameCenterActivityLocalizations

#### delete

Delete /v1/gameCenterActivityLocalizations/{id}

#### get

Get /v1/gameCenterActivityLocalizations/{id}

#### image

Manage gameCenterActivityLocalizations image

1 subcommand(s).

##### get

Get /v1/gameCenterActivityLocalizations/{id}/image

#### update

Update /v1/gameCenterActivityLocalizations/{id}

### gameCenterActivityVersionReleases

Manage gameCenterActivityVersionReleases

3 subcommand(s).

#### create

Create /v1/gameCenterActivityVersionReleases

#### delete

Delete /v1/gameCenterActivityVersionReleases/{id}

#### get

Get /v1/gameCenterActivityVersionReleases/{id}

### gameCenterActivityVersions

Manage gameCenterActivityVersions

5 subcommand(s).

#### create

Create /v1/gameCenterActivityVersions

#### defaultImage

Manage gameCenterActivityVersions defaultImage

1 subcommand(s).

##### get

Get /v1/gameCenterActivityVersions/{id}/defaultImage

#### get

Get /v1/gameCenterActivityVersions/{id}

#### localizations

Manage gameCenterActivityVersions localizations

1 subcommand(s).

##### list

List /v1/gameCenterActivityVersions/{id}/localizations

#### update

Update /v1/gameCenterActivityVersions/{id}

### gameCenterAppVersions

Manage gameCenterAppVersions

5 subcommand(s).

#### appStoreVersion

Manage gameCenterAppVersions appStoreVersion

1 subcommand(s).

##### get

Get /v1/gameCenterAppVersions/{id}/appStoreVersion

#### compatibilityVersions

Manage gameCenterAppVersions compatibilityVersions

1 subcommand(s).

##### list

List /v1/gameCenterAppVersions/{id}/compatibilityVersions

#### create

Create /v1/gameCenterAppVersions

#### get

Get /v1/gameCenterAppVersions/{id}

#### update

Update /v1/gameCenterAppVersions/{id}

### gameCenterChallengeImages

Manage gameCenterChallengeImages

4 subcommand(s).

#### create

Create /v1/gameCenterChallengeImages

#### delete

Delete /v1/gameCenterChallengeImages/{id}

#### get

Get /v1/gameCenterChallengeImages/{id}

#### update

Update /v1/gameCenterChallengeImages/{id}

### gameCenterChallengeLocalizations

Manage gameCenterChallengeLocalizations

5 subcommand(s).

#### create

Create /v1/gameCenterChallengeLocalizations

#### delete

Delete /v1/gameCenterChallengeLocalizations/{id}

#### get

Get /v1/gameCenterChallengeLocalizations/{id}

#### image

Manage gameCenterChallengeLocalizations image

1 subcommand(s).

##### get

Get /v1/gameCenterChallengeLocalizations/{id}/image

#### update

Update /v1/gameCenterChallengeLocalizations/{id}

### gameCenterChallengeVersionReleases

Manage gameCenterChallengeVersionReleases

3 subcommand(s).

#### create

Create /v1/gameCenterChallengeVersionReleases

#### delete

Delete /v1/gameCenterChallengeVersionReleases/{id}

#### get

Get /v1/gameCenterChallengeVersionReleases/{id}

### gameCenterChallengeVersions

Manage gameCenterChallengeVersions

4 subcommand(s).

#### create

Create /v1/gameCenterChallengeVersions

#### defaultImage

Manage gameCenterChallengeVersions defaultImage

1 subcommand(s).

##### get

Get /v1/gameCenterChallengeVersions/{id}/defaultImage

#### get

Get /v1/gameCenterChallengeVersions/{id}

#### localizations

Manage gameCenterChallengeVersions localizations

1 subcommand(s).

##### list

List /v1/gameCenterChallengeVersions/{id}/localizations

### gameCenterChallenges

Manage gameCenterChallenges

5 subcommand(s).

#### create

Create /v1/gameCenterChallenges

#### delete

Delete /v1/gameCenterChallenges/{id}

#### get

Get /v1/gameCenterChallenges/{id}

#### update

Update /v1/gameCenterChallenges/{id}

#### versions

Manage gameCenterChallenges versions

1 subcommand(s).

##### list

List /v1/gameCenterChallenges/{id}/versions

### gameCenterDetails

Manage gameCenterDetails

20 subcommand(s).

#### achievementReleases

Manage gameCenterDetails achievementReleases

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/achievementReleases

#### activityReleases

Manage gameCenterDetails activityReleases

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/activityReleases

#### challengeReleases

Manage gameCenterDetails challengeReleases

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/challengeReleases

#### classicMatchmakingRequests

Manage gameCenterDetails classicMatchmakingRequests

1 subcommand(s).

##### metrics

Metrics /v1/gameCenterDetails/{id}/metrics/classicMatchmakingRequests

#### create

Create /v1/gameCenterDetails

#### gameCenterAchievements

Manage gameCenterDetails gameCenterAchievements

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/gameCenterAchievements

#### gameCenterAchievementsV2

Manage gameCenterDetails gameCenterAchievementsV2

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/gameCenterAchievementsV2

#### gameCenterActivities

Manage gameCenterDetails gameCenterActivities

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/gameCenterActivities

#### gameCenterAppVersions

Manage gameCenterDetails gameCenterAppVersions

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/gameCenterAppVersions

#### gameCenterChallenges

Manage gameCenterDetails gameCenterChallenges

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/gameCenterChallenges

#### gameCenterGroup

Manage gameCenterDetails gameCenterGroup

1 subcommand(s).

##### get

Get /v1/gameCenterDetails/{id}/gameCenterGroup

#### gameCenterLeaderboardSets

Manage gameCenterDetails gameCenterLeaderboardSets

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/gameCenterLeaderboardSets

#### gameCenterLeaderboardSetsV2

Manage gameCenterDetails gameCenterLeaderboardSetsV2

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/gameCenterLeaderboardSetsV2

#### gameCenterLeaderboards

Manage gameCenterDetails gameCenterLeaderboards

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/gameCenterLeaderboards

#### gameCenterLeaderboardsV2

Manage gameCenterDetails gameCenterLeaderboardsV2

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/gameCenterLeaderboardsV2

#### get

Get /v1/gameCenterDetails/{id}

#### leaderboardReleases

Manage gameCenterDetails leaderboardReleases

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/leaderboardReleases

#### leaderboardSetReleases

Manage gameCenterDetails leaderboardSetReleases

1 subcommand(s).

##### list

List /v1/gameCenterDetails/{id}/leaderboardSetReleases

#### ruleBasedMatchmakingRequests

Manage gameCenterDetails ruleBasedMatchmakingRequests

1 subcommand(s).

##### metrics

Metrics /v1/gameCenterDetails/{id}/metrics/ruleBasedMatchmakingRequests

#### update

Update /v1/gameCenterDetails/{id}

### gameCenterEnabledVersions

Manage gameCenterEnabledVersions

1 subcommand(s).

#### compatibleVersions

Manage gameCenterEnabledVersions compatibleVersions

1 subcommand(s).

##### list

List /v1/gameCenterEnabledVersions/{id}/compatibleVersions

### gameCenterGroups

Manage gameCenterGroups

14 subcommand(s).

#### create

Create /v1/gameCenterGroups

#### delete

Delete /v1/gameCenterGroups/{id}

#### gameCenterAchievements

Manage gameCenterGroups gameCenterAchievements

1 subcommand(s).

##### list

List /v1/gameCenterGroups/{id}/gameCenterAchievements

#### gameCenterAchievementsV2

Manage gameCenterGroups gameCenterAchievementsV2

1 subcommand(s).

##### list

List /v1/gameCenterGroups/{id}/gameCenterAchievementsV2

#### gameCenterActivities

Manage gameCenterGroups gameCenterActivities

1 subcommand(s).

##### list

List /v1/gameCenterGroups/{id}/gameCenterActivities

#### gameCenterChallenges

Manage gameCenterGroups gameCenterChallenges

1 subcommand(s).

##### list

List /v1/gameCenterGroups/{id}/gameCenterChallenges

#### gameCenterDetails

Manage gameCenterGroups gameCenterDetails

1 subcommand(s).

##### list

List /v1/gameCenterGroups/{id}/gameCenterDetails

#### gameCenterLeaderboardSets

Manage gameCenterGroups gameCenterLeaderboardSets

1 subcommand(s).

##### list

List /v1/gameCenterGroups/{id}/gameCenterLeaderboardSets

#### gameCenterLeaderboardSetsV2

Manage gameCenterGroups gameCenterLeaderboardSetsV2

1 subcommand(s).

##### list

List /v1/gameCenterGroups/{id}/gameCenterLeaderboardSetsV2

#### gameCenterLeaderboards

Manage gameCenterGroups gameCenterLeaderboards

1 subcommand(s).

##### list

List /v1/gameCenterGroups/{id}/gameCenterLeaderboards

#### gameCenterLeaderboardsV2

Manage gameCenterGroups gameCenterLeaderboardsV2

1 subcommand(s).

##### list

List /v1/gameCenterGroups/{id}/gameCenterLeaderboardsV2

#### get

Get /v1/gameCenterGroups/{id}

#### list

List /v1/gameCenterGroups

#### update

Update /v1/gameCenterGroups/{id}

### gameCenterLeaderboardEntrySubmissions

Manage gameCenterLeaderboardEntrySubmissions

1 subcommand(s).

#### create

Create /v1/gameCenterLeaderboardEntrySubmissions

### gameCenterLeaderboardImages

Manage gameCenterLeaderboardImages

4 subcommand(s).

#### create

Create /v1/gameCenterLeaderboardImages

#### delete

Delete /v1/gameCenterLeaderboardImages/{id}

#### get

Get /v1/gameCenterLeaderboardImages/{id}

#### update

Update /v1/gameCenterLeaderboardImages/{id}

### gameCenterLeaderboardImagesV2

Manage gameCenterLeaderboardImagesV2

4 subcommand(s).

#### create

Create /v2/gameCenterLeaderboardImages

#### delete

Delete /v2/gameCenterLeaderboardImages/{id}

#### get

Get /v2/gameCenterLeaderboardImages/{id}

#### update

Update /v2/gameCenterLeaderboardImages/{id}

### gameCenterLeaderboardLocalizations

Manage gameCenterLeaderboardLocalizations

5 subcommand(s).

#### create

Create /v1/gameCenterLeaderboardLocalizations

#### delete

Delete /v1/gameCenterLeaderboardLocalizations/{id}

#### gameCenterLeaderboardImage

Manage gameCenterLeaderboardLocalizations gameCenterLeaderboardImage

1 subcommand(s).

##### get

Get /v1/gameCenterLeaderboardLocalizations/{id}/gameCenterLeaderboardImage

#### get

Get /v1/gameCenterLeaderboardLocalizations/{id}

#### update

Update /v1/gameCenterLeaderboardLocalizations/{id}

### gameCenterLeaderboardLocalizationsV2

Manage gameCenterLeaderboardLocalizationsV2

5 subcommand(s).

#### create

Create /v2/gameCenterLeaderboardLocalizations

#### delete

Delete /v2/gameCenterLeaderboardLocalizations/{id}

#### get

Get /v2/gameCenterLeaderboardLocalizations/{id}

#### image

Manage gameCenterLeaderboardLocalizationsV2 image

1 subcommand(s).

##### get

Get /v2/gameCenterLeaderboardLocalizations/{id}/image

#### update

Update /v2/gameCenterLeaderboardLocalizations/{id}

### gameCenterLeaderboardReleases

Manage gameCenterLeaderboardReleases

3 subcommand(s).

#### create

Create /v1/gameCenterLeaderboardReleases

#### delete

Delete /v1/gameCenterLeaderboardReleases/{id}

#### get

Get /v1/gameCenterLeaderboardReleases/{id}

### gameCenterLeaderboardSetImages

Manage gameCenterLeaderboardSetImages

4 subcommand(s).

#### create

Create /v1/gameCenterLeaderboardSetImages

#### delete

Delete /v1/gameCenterLeaderboardSetImages/{id}

#### get

Get /v1/gameCenterLeaderboardSetImages/{id}

#### update

Update /v1/gameCenterLeaderboardSetImages/{id}

### gameCenterLeaderboardSetImagesV2

Manage gameCenterLeaderboardSetImagesV2

4 subcommand(s).

#### create

Create /v2/gameCenterLeaderboardSetImages

#### delete

Delete /v2/gameCenterLeaderboardSetImages/{id}

#### get

Get /v2/gameCenterLeaderboardSetImages/{id}

#### update

Update /v2/gameCenterLeaderboardSetImages/{id}

### gameCenterLeaderboardSetLocalizations

Manage gameCenterLeaderboardSetLocalizations

5 subcommand(s).

#### create

Create /v1/gameCenterLeaderboardSetLocalizations

#### delete

Delete /v1/gameCenterLeaderboardSetLocalizations/{id}

#### gameCenterLeaderboardSetImage

Manage gameCenterLeaderboardSetLocalizations gameCenterLeaderboardSetImage

1 subcommand(s).

##### get

Get /v1/gameCenterLeaderboardSetLocalizations/{id}/gameCenterLeaderboardSetImage

#### get

Get /v1/gameCenterLeaderboardSetLocalizations/{id}

#### update

Update /v1/gameCenterLeaderboardSetLocalizations/{id}

### gameCenterLeaderboardSetLocalizationsV2

Manage gameCenterLeaderboardSetLocalizationsV2

5 subcommand(s).

#### create

Create /v2/gameCenterLeaderboardSetLocalizations

#### delete

Delete /v2/gameCenterLeaderboardSetLocalizations/{id}

#### get

Get /v2/gameCenterLeaderboardSetLocalizations/{id}

#### image

Manage gameCenterLeaderboardSetLocalizationsV2 image

1 subcommand(s).

##### get

Get /v2/gameCenterLeaderboardSetLocalizations/{id}/image

#### update

Update /v2/gameCenterLeaderboardSetLocalizations/{id}

### gameCenterLeaderboardSetMemberLocalizations

Manage gameCenterLeaderboardSetMemberLocalizations

6 subcommand(s).

#### create

Create /v1/gameCenterLeaderboardSetMemberLocalizations

#### delete

Delete /v1/gameCenterLeaderboardSetMemberLocalizations/{id}

#### gameCenterLeaderboard

Manage gameCenterLeaderboardSetMemberLocalizations gameCenterLeaderboard

1 subcommand(s).

##### get

Get /v1/gameCenterLeaderboardSetMemberLocalizations/{id}/gameCenterLeaderboard

#### gameCenterLeaderboardSet

Manage gameCenterLeaderboardSetMemberLocalizations gameCenterLeaderboardSet

1 subcommand(s).

##### get

Get /v1/gameCenterLeaderboardSetMemberLocalizations/{id}/gameCenterLeaderboardSet

#### list

List /v1/gameCenterLeaderboardSetMemberLocalizations

#### update

Update /v1/gameCenterLeaderboardSetMemberLocalizations/{id}

### gameCenterLeaderboardSetReleases

Manage gameCenterLeaderboardSetReleases

3 subcommand(s).

#### create

Create /v1/gameCenterLeaderboardSetReleases

#### delete

Delete /v1/gameCenterLeaderboardSetReleases/{id}

#### get

Get /v1/gameCenterLeaderboardSetReleases/{id}

### gameCenterLeaderboardSetVersionsV2

Manage gameCenterLeaderboardSetVersionsV2

3 subcommand(s).

#### create

Create /v2/gameCenterLeaderboardSetVersions

#### get

Get /v2/gameCenterLeaderboardSetVersions/{id}

#### localizations

Manage gameCenterLeaderboardSetVersionsV2 localizations

1 subcommand(s).

##### list

List /v2/gameCenterLeaderboardSetVersions/{id}/localizations

### gameCenterLeaderboardSets

Manage gameCenterLeaderboardSets

8 subcommand(s).

#### create

Create /v1/gameCenterLeaderboardSets

#### delete

Delete /v1/gameCenterLeaderboardSets/{id}

#### gameCenterLeaderboards

Manage gameCenterLeaderboardSets gameCenterLeaderboards

1 subcommand(s).

##### list

List /v1/gameCenterLeaderboardSets/{id}/gameCenterLeaderboards

#### get

Get /v1/gameCenterLeaderboardSets/{id}

#### groupLeaderboardSet

Manage gameCenterLeaderboardSets groupLeaderboardSet

1 subcommand(s).

##### get

Get /v1/gameCenterLeaderboardSets/{id}/groupLeaderboardSet

#### localizations

Manage gameCenterLeaderboardSets localizations

1 subcommand(s).

##### list

List /v1/gameCenterLeaderboardSets/{id}/localizations

#### releases

Manage gameCenterLeaderboardSets releases

1 subcommand(s).

##### list

List /v1/gameCenterLeaderboardSets/{id}/releases

#### update

Update /v1/gameCenterLeaderboardSets/{id}

### gameCenterLeaderboardSetsV2

Manage gameCenterLeaderboardSetsV2

6 subcommand(s).

#### create

Create /v2/gameCenterLeaderboardSets

#### delete

Delete /v2/gameCenterLeaderboardSets/{id}

#### gameCenterLeaderboards

Manage gameCenterLeaderboardSetsV2 gameCenterLeaderboards

1 subcommand(s).

##### list

List /v2/gameCenterLeaderboardSets/{id}/gameCenterLeaderboards

#### get

Get /v2/gameCenterLeaderboardSets/{id}

#### update

Update /v2/gameCenterLeaderboardSets/{id}

#### versions

Manage gameCenterLeaderboardSetsV2 versions

1 subcommand(s).

##### list

List /v2/gameCenterLeaderboardSets/{id}/versions

### gameCenterLeaderboardVersionsV2

Manage gameCenterLeaderboardVersionsV2

3 subcommand(s).

#### create

Create /v2/gameCenterLeaderboardVersions

#### get

Get /v2/gameCenterLeaderboardVersions/{id}

#### localizations

Manage gameCenterLeaderboardVersionsV2 localizations

1 subcommand(s).

##### list

List /v2/gameCenterLeaderboardVersions/{id}/localizations

### gameCenterLeaderboards

Manage gameCenterLeaderboards

7 subcommand(s).

#### create

Create /v1/gameCenterLeaderboards

#### delete

Delete /v1/gameCenterLeaderboards/{id}

#### get

Get /v1/gameCenterLeaderboards/{id}

#### groupLeaderboard

Manage gameCenterLeaderboards groupLeaderboard

1 subcommand(s).

##### get

Get /v1/gameCenterLeaderboards/{id}/groupLeaderboard

#### localizations

Manage gameCenterLeaderboards localizations

1 subcommand(s).

##### list

List /v1/gameCenterLeaderboards/{id}/localizations

#### releases

Manage gameCenterLeaderboards releases

1 subcommand(s).

##### list

List /v1/gameCenterLeaderboards/{id}/releases

#### update

Update /v1/gameCenterLeaderboards/{id}

### gameCenterLeaderboardsV2

Manage gameCenterLeaderboardsV2

5 subcommand(s).

#### create

Create /v2/gameCenterLeaderboards

#### delete

Delete /v2/gameCenterLeaderboards/{id}

#### get

Get /v2/gameCenterLeaderboards/{id}

#### update

Update /v2/gameCenterLeaderboards/{id}

#### versions

Manage gameCenterLeaderboardsV2 versions

1 subcommand(s).

##### list

List /v2/gameCenterLeaderboards/{id}/versions

### gameCenterMatchmakingQueues

Manage gameCenterMatchmakingQueues

10 subcommand(s).

#### create

Create /v1/gameCenterMatchmakingQueues

#### delete

Delete /v1/gameCenterMatchmakingQueues/{id}

#### experimentMatchmakingQueueSizes

Manage gameCenterMatchmakingQueues experimentMatchmakingQueueSizes

1 subcommand(s).

##### metrics

Metrics /v1/gameCenterMatchmakingQueues/{id}/metrics/experimentMatchmakingQueueSizes

#### experimentMatchmakingRequests

Manage gameCenterMatchmakingQueues experimentMatchmakingRequests

1 subcommand(s).

##### metrics

Metrics /v1/gameCenterMatchmakingQueues/{id}/metrics/experimentMatchmakingRequests

#### get

Get /v1/gameCenterMatchmakingQueues/{id}

#### list

List /v1/gameCenterMatchmakingQueues

#### matchmakingQueueSizes

Manage gameCenterMatchmakingQueues matchmakingQueueSizes

1 subcommand(s).

##### metrics

Metrics /v1/gameCenterMatchmakingQueues/{id}/metrics/matchmakingQueueSizes

#### matchmakingRequests

Manage gameCenterMatchmakingQueues matchmakingRequests

1 subcommand(s).

##### metrics

Metrics /v1/gameCenterMatchmakingQueues/{id}/metrics/matchmakingRequests

#### matchmakingSessions

Manage gameCenterMatchmakingQueues matchmakingSessions

1 subcommand(s).

##### metrics

Metrics /v1/gameCenterMatchmakingQueues/{id}/metrics/matchmakingSessions

#### update

Update /v1/gameCenterMatchmakingQueues/{id}

### gameCenterMatchmakingRuleSetTests

Manage gameCenterMatchmakingRuleSetTests

1 subcommand(s).

#### create

Create /v1/gameCenterMatchmakingRuleSetTests

### gameCenterMatchmakingRuleSets

Manage gameCenterMatchmakingRuleSets

8 subcommand(s).

#### create

Create /v1/gameCenterMatchmakingRuleSets

#### delete

Delete /v1/gameCenterMatchmakingRuleSets/{id}

#### get

Get /v1/gameCenterMatchmakingRuleSets/{id}

#### list

List /v1/gameCenterMatchmakingRuleSets

#### matchmakingQueues

Manage gameCenterMatchmakingRuleSets matchmakingQueues

1 subcommand(s).

##### list

List /v1/gameCenterMatchmakingRuleSets/{id}/matchmakingQueues

#### rules

Manage gameCenterMatchmakingRuleSets rules

1 subcommand(s).

##### list

List /v1/gameCenterMatchmakingRuleSets/{id}/rules

#### teams

Manage gameCenterMatchmakingRuleSets teams

1 subcommand(s).

##### list

List /v1/gameCenterMatchmakingRuleSets/{id}/teams

#### update

Update /v1/gameCenterMatchmakingRuleSets/{id}

### gameCenterMatchmakingRules

Manage gameCenterMatchmakingRules

6 subcommand(s).

#### create

Create /v1/gameCenterMatchmakingRules

#### delete

Delete /v1/gameCenterMatchmakingRules/{id}

#### matchmakingBooleanRuleResults

Manage gameCenterMatchmakingRules matchmakingBooleanRuleResults

1 subcommand(s).

##### metrics

Metrics /v1/gameCenterMatchmakingRules/{id}/metrics/matchmakingBooleanRuleResults

#### matchmakingNumberRuleResults

Manage gameCenterMatchmakingRules matchmakingNumberRuleResults

1 subcommand(s).

##### metrics

Metrics /v1/gameCenterMatchmakingRules/{id}/metrics/matchmakingNumberRuleResults

#### matchmakingRuleErrors

Manage gameCenterMatchmakingRules matchmakingRuleErrors

1 subcommand(s).

##### metrics

Metrics /v1/gameCenterMatchmakingRules/{id}/metrics/matchmakingRuleErrors

#### update

Update /v1/gameCenterMatchmakingRules/{id}

### gameCenterMatchmakingTeams

Manage gameCenterMatchmakingTeams

3 subcommand(s).

#### create

Create /v1/gameCenterMatchmakingTeams

#### delete

Delete /v1/gameCenterMatchmakingTeams/{id}

#### update

Update /v1/gameCenterMatchmakingTeams/{id}

### gameCenterPlayerAchievementSubmissions

Manage gameCenterPlayerAchievementSubmissions

1 subcommand(s).

#### create

Create /v1/gameCenterPlayerAchievementSubmissions

### inAppPurchaseAppStoreReviewScreenshots

Manage inAppPurchaseAppStoreReviewScreenshots

4 subcommand(s).

#### create

Create /v1/inAppPurchaseAppStoreReviewScreenshots

#### delete

Delete /v1/inAppPurchaseAppStoreReviewScreenshots/{id}

#### get

Get /v1/inAppPurchaseAppStoreReviewScreenshots/{id}

#### update

Update /v1/inAppPurchaseAppStoreReviewScreenshots/{id}

### inAppPurchaseAvailabilities

Manage inAppPurchaseAvailabilities

3 subcommand(s).

#### availableTerritories

Manage inAppPurchaseAvailabilities availableTerritories

1 subcommand(s).

##### list

List /v1/inAppPurchaseAvailabilities/{id}/availableTerritories

#### create

Create /v1/inAppPurchaseAvailabilities

#### get

Get /v1/inAppPurchaseAvailabilities/{id}

### inAppPurchaseContents

Manage inAppPurchaseContents

1 subcommand(s).

#### get

Get /v1/inAppPurchaseContents/{id}

### inAppPurchaseImages

Manage inAppPurchaseImages

4 subcommand(s).

#### create

Create /v1/inAppPurchaseImages

#### delete

Delete /v1/inAppPurchaseImages/{id}

#### get

Get /v1/inAppPurchaseImages/{id}

#### update

Update /v1/inAppPurchaseImages/{id}

### inAppPurchaseLocalizations

Manage inAppPurchaseLocalizations

4 subcommand(s).

#### create

Create /v1/inAppPurchaseLocalizations

#### delete

Delete /v1/inAppPurchaseLocalizations/{id}

#### get

Get /v1/inAppPurchaseLocalizations/{id}

#### update

Update /v1/inAppPurchaseLocalizations/{id}

### inAppPurchaseOfferCodeCustomCodes

Manage inAppPurchaseOfferCodeCustomCodes

3 subcommand(s).

#### create

Create /v1/inAppPurchaseOfferCodeCustomCodes

#### get

Get /v1/inAppPurchaseOfferCodeCustomCodes/{id}

#### update

Update /v1/inAppPurchaseOfferCodeCustomCodes/{id}

### inAppPurchaseOfferCodeOneTimeUseCodes

Manage inAppPurchaseOfferCodeOneTimeUseCodes

4 subcommand(s).

#### create

Create /v1/inAppPurchaseOfferCodeOneTimeUseCodes

#### get

Get /v1/inAppPurchaseOfferCodeOneTimeUseCodes/{id}

#### update

Update /v1/inAppPurchaseOfferCodeOneTimeUseCodes/{id}

#### values

Manage inAppPurchaseOfferCodeOneTimeUseCodes values

1 subcommand(s).

##### get

Get /v1/inAppPurchaseOfferCodeOneTimeUseCodes/{id}/values

### inAppPurchaseOfferCodes

Manage inAppPurchaseOfferCodes

6 subcommand(s).

#### create

Create /v1/inAppPurchaseOfferCodes

#### customCodes

Manage inAppPurchaseOfferCodes customCodes

1 subcommand(s).

##### list

List /v1/inAppPurchaseOfferCodes/{id}/customCodes

#### get

Get /v1/inAppPurchaseOfferCodes/{id}

#### oneTimeUseCodes

Manage inAppPurchaseOfferCodes oneTimeUseCodes

1 subcommand(s).

##### list

List /v1/inAppPurchaseOfferCodes/{id}/oneTimeUseCodes

#### prices

Manage inAppPurchaseOfferCodes prices

1 subcommand(s).

##### list

List /v1/inAppPurchaseOfferCodes/{id}/prices

#### update

Update /v1/inAppPurchaseOfferCodes/{id}

### inAppPurchasePricePoints

Manage inAppPurchasePricePoints

1 subcommand(s).

#### equalizations

Manage inAppPurchasePricePoints equalizations

1 subcommand(s).

##### list

List /v1/inAppPurchasePricePoints/{id}/equalizations

### inAppPurchasePriceSchedules

Manage inAppPurchasePriceSchedules

5 subcommand(s).

#### automaticPrices

Manage inAppPurchasePriceSchedules automaticPrices

1 subcommand(s).

##### list

List /v1/inAppPurchasePriceSchedules/{id}/automaticPrices

#### baseTerritory

Manage inAppPurchasePriceSchedules baseTerritory

1 subcommand(s).

##### get

Get /v1/inAppPurchasePriceSchedules/{id}/baseTerritory

#### create

Create /v1/inAppPurchasePriceSchedules

#### get

Get /v1/inAppPurchasePriceSchedules/{id}

#### manualPrices

Manage inAppPurchasePriceSchedules manualPrices

1 subcommand(s).

##### list

List /v1/inAppPurchasePriceSchedules/{id}/manualPrices

### inAppPurchaseSubmissions

Manage inAppPurchaseSubmissions

1 subcommand(s).

#### create

Create /v1/inAppPurchaseSubmissions

### inAppPurchases

Manage inAppPurchases

1 subcommand(s).

#### get

Get /v1/inAppPurchases/{id}

### inAppPurchasesV2

Manage inAppPurchasesV2

13 subcommand(s).

#### appStoreReviewScreenshot

Manage inAppPurchasesV2 appStoreReviewScreenshot

1 subcommand(s).

##### get

Get /v2/inAppPurchases/{id}/appStoreReviewScreenshot

#### content

Manage inAppPurchasesV2 content

1 subcommand(s).

##### get

Get /v2/inAppPurchases/{id}/content

#### create

Create /v2/inAppPurchases

#### delete

Delete /v2/inAppPurchases/{id}

#### get

Get /v2/inAppPurchases/{id}

#### iapPriceSchedule

Manage inAppPurchasesV2 iapPriceSchedule

1 subcommand(s).

##### get

Get /v2/inAppPurchases/{id}/iapPriceSchedule

#### images

Manage inAppPurchasesV2 images

1 subcommand(s).

##### list

List /v2/inAppPurchases/{id}/images

#### inAppPurchaseAvailability

Manage inAppPurchasesV2 inAppPurchaseAvailability

1 subcommand(s).

##### get

Get /v2/inAppPurchases/{id}/inAppPurchaseAvailability

#### inAppPurchaseLocalizations

Manage inAppPurchasesV2 inAppPurchaseLocalizations

1 subcommand(s).

##### list

List /v2/inAppPurchases/{id}/inAppPurchaseLocalizations

#### offerCodes

Manage inAppPurchasesV2 offerCodes

1 subcommand(s).

##### list

List /v2/inAppPurchases/{id}/offerCodes

#### pricePoints

Manage inAppPurchasesV2 pricePoints

1 subcommand(s).

##### list

List /v2/inAppPurchases/{id}/pricePoints

#### promotedPurchase

Manage inAppPurchasesV2 promotedPurchase

1 subcommand(s).

##### get

Get /v2/inAppPurchases/{id}/promotedPurchase

#### update

Update /v2/inAppPurchases/{id}

### marketplaceSearchDetails

Manage marketplaceSearchDetails

3 subcommand(s).

#### create

Create /v1/marketplaceSearchDetails

#### delete

Delete /v1/marketplaceSearchDetails/{id}

#### update

Update /v1/marketplaceSearchDetails/{id}

### marketplaceWebhooks

Manage marketplaceWebhooks

4 subcommand(s).

#### create

Create /v1/marketplaceWebhooks

#### delete

Delete /v1/marketplaceWebhooks/{id}

#### list

List /v1/marketplaceWebhooks

#### update

Update /v1/marketplaceWebhooks/{id}

### merchantIds

Manage merchantIds

6 subcommand(s).

#### certificates

Manage merchantIds certificates

1 subcommand(s).

##### list

List /v1/merchantIds/{id}/certificates

#### create

Create /v1/merchantIds

#### delete

Delete /v1/merchantIds/{id}

#### get

Get /v1/merchantIds/{id}

#### list

List /v1/merchantIds

#### update

Update /v1/merchantIds/{id}

### nominations

Manage nominations

5 subcommand(s).

#### create

Create /v1/nominations

#### delete

Delete /v1/nominations/{id}

#### get

Get /v1/nominations/{id}

#### list

List /v1/nominations

#### update

Update /v1/nominations/{id}

### passTypeIds

Manage passTypeIds

6 subcommand(s).

#### certificates

Manage passTypeIds certificates

1 subcommand(s).

##### list

List /v1/passTypeIds/{id}/certificates

#### create

Create /v1/passTypeIds

#### delete

Delete /v1/passTypeIds/{id}

#### get

Get /v1/passTypeIds/{id}

#### list

List /v1/passTypeIds

#### update

Update /v1/passTypeIds/{id}

### preReleaseVersions

Manage preReleaseVersions

4 subcommand(s).

#### app

Manage preReleaseVersions app

1 subcommand(s).

##### get

Get /v1/preReleaseVersions/{id}/app

#### builds

Manage preReleaseVersions builds

1 subcommand(s).

##### list

List /v1/preReleaseVersions/{id}/builds

#### get

Get /v1/preReleaseVersions/{id}

#### list

List /v1/preReleaseVersions

### profiles

Manage profiles

7 subcommand(s).

#### bundleId

Manage profiles bundleId

1 subcommand(s).

##### get

Get /v1/profiles/{id}/bundleId

#### certificates

Manage profiles certificates

1 subcommand(s).

##### list

List /v1/profiles/{id}/certificates

#### create

Create /v1/profiles

#### delete

Delete /v1/profiles/{id}

#### devices

Manage profiles devices

1 subcommand(s).

##### list

List /v1/profiles/{id}/devices

#### get

Get /v1/profiles/{id}

#### list

List /v1/profiles

### promotedPurchases

Manage promotedPurchases

4 subcommand(s).

#### create

Create /v1/promotedPurchases

#### delete

Delete /v1/promotedPurchases/{id}

#### get

Get /v1/promotedPurchases/{id}

#### update

Update /v1/promotedPurchases/{id}

### reviewSubmissionItems

Manage reviewSubmissionItems

3 subcommand(s).

#### create

Create /v1/reviewSubmissionItems

#### delete

Delete /v1/reviewSubmissionItems/{id}

#### update

Update /v1/reviewSubmissionItems/{id}

### reviewSubmissions

Manage reviewSubmissions

5 subcommand(s).

#### create

Create /v1/reviewSubmissions

#### get

Get /v1/reviewSubmissions/{id}

#### items

Manage reviewSubmissions items

1 subcommand(s).

##### list

List /v1/reviewSubmissions/{id}/items

#### list

List /v1/reviewSubmissions

#### update

Update /v1/reviewSubmissions/{id}

### routingAppCoverages

Manage routingAppCoverages

4 subcommand(s).

#### create

Create /v1/routingAppCoverages

#### delete

Delete /v1/routingAppCoverages/{id}

#### get

Get /v1/routingAppCoverages/{id}

#### update

Update /v1/routingAppCoverages/{id}

### salesReports

Manage salesReports

1 subcommand(s).

#### list

List /v1/salesReports

### sandboxTestersClearPurchaseHistoryRequestV2

Manage sandboxTestersClearPurchaseHistoryRequestV2

1 subcommand(s).

#### create

Create /v2/sandboxTestersClearPurchaseHistoryRequest

### sandboxTestersV2

Manage sandboxTestersV2

2 subcommand(s).

#### list

List /v2/sandboxTesters

#### update

Update /v2/sandboxTesters/{id}

### scmGitReferences

Manage scmGitReferences

1 subcommand(s).

#### get

Get /v1/scmGitReferences/{id}

### scmProviders

Manage scmProviders

3 subcommand(s).

#### get

Get /v1/scmProviders/{id}

#### list

List /v1/scmProviders

#### repositories

Manage scmProviders repositories

1 subcommand(s).

##### list

List /v1/scmProviders/{id}/repositories

### scmPullRequests

Manage scmPullRequests

1 subcommand(s).

#### get

Get /v1/scmPullRequests/{id}

### scmRepositories

Manage scmRepositories

4 subcommand(s).

#### get

Get /v1/scmRepositories/{id}

#### gitReferences

Manage scmRepositories gitReferences

1 subcommand(s).

##### list

List /v1/scmRepositories/{id}/gitReferences

#### list

List /v1/scmRepositories

#### pullRequests

Manage scmRepositories pullRequests

1 subcommand(s).

##### list

List /v1/scmRepositories/{id}/pullRequests

### subscriptionAppStoreReviewScreenshots

Manage subscriptionAppStoreReviewScreenshots

4 subcommand(s).

#### create

Create /v1/subscriptionAppStoreReviewScreenshots

#### delete

Delete /v1/subscriptionAppStoreReviewScreenshots/{id}

#### get

Get /v1/subscriptionAppStoreReviewScreenshots/{id}

#### update

Update /v1/subscriptionAppStoreReviewScreenshots/{id}

### subscriptionAvailabilities

Manage subscriptionAvailabilities

3 subcommand(s).

#### availableTerritories

Manage subscriptionAvailabilities availableTerritories

1 subcommand(s).

##### list

List /v1/subscriptionAvailabilities/{id}/availableTerritories

#### create

Create /v1/subscriptionAvailabilities

#### get

Get /v1/subscriptionAvailabilities/{id}

### subscriptionGracePeriods

Manage subscriptionGracePeriods

2 subcommand(s).

#### get

Get /v1/subscriptionGracePeriods/{id}

#### update

Update /v1/subscriptionGracePeriods/{id}

### subscriptionGroupLocalizations

Manage subscriptionGroupLocalizations

4 subcommand(s).

#### create

Create /v1/subscriptionGroupLocalizations

#### delete

Delete /v1/subscriptionGroupLocalizations/{id}

#### get

Get /v1/subscriptionGroupLocalizations/{id}

#### update

Update /v1/subscriptionGroupLocalizations/{id}

### subscriptionGroupSubmissions

Manage subscriptionGroupSubmissions

1 subcommand(s).

#### create

Create /v1/subscriptionGroupSubmissions

### subscriptionGroups

Manage subscriptionGroups

6 subcommand(s).

#### create

Create /v1/subscriptionGroups

#### delete

Delete /v1/subscriptionGroups/{id}

#### get

Get /v1/subscriptionGroups/{id}

#### subscriptionGroupLocalizations

Manage subscriptionGroups subscriptionGroupLocalizations

1 subcommand(s).

##### list

List /v1/subscriptionGroups/{id}/subscriptionGroupLocalizations

#### subscriptions

Manage subscriptionGroups subscriptions

1 subcommand(s).

##### list

List /v1/subscriptionGroups/{id}/subscriptions

#### update

Update /v1/subscriptionGroups/{id}

### subscriptionImages

Manage subscriptionImages

4 subcommand(s).

#### create

Create /v1/subscriptionImages

#### delete

Delete /v1/subscriptionImages/{id}

#### get

Get /v1/subscriptionImages/{id}

#### update

Update /v1/subscriptionImages/{id}

### subscriptionIntroductoryOffers

Manage subscriptionIntroductoryOffers

3 subcommand(s).

#### create

Create /v1/subscriptionIntroductoryOffers

#### delete

Delete /v1/subscriptionIntroductoryOffers/{id}

#### update

Update /v1/subscriptionIntroductoryOffers/{id}

### subscriptionLocalizations

Manage subscriptionLocalizations

4 subcommand(s).

#### create

Create /v1/subscriptionLocalizations

#### delete

Delete /v1/subscriptionLocalizations/{id}

#### get

Get /v1/subscriptionLocalizations/{id}

#### update

Update /v1/subscriptionLocalizations/{id}

### subscriptionOfferCodeCustomCodes

Manage subscriptionOfferCodeCustomCodes

3 subcommand(s).

#### create

Create /v1/subscriptionOfferCodeCustomCodes

#### get

Get /v1/subscriptionOfferCodeCustomCodes/{id}

#### update

Update /v1/subscriptionOfferCodeCustomCodes/{id}

### subscriptionOfferCodeOneTimeUseCodes

Manage subscriptionOfferCodeOneTimeUseCodes

4 subcommand(s).

#### create

Create /v1/subscriptionOfferCodeOneTimeUseCodes

#### get

Get /v1/subscriptionOfferCodeOneTimeUseCodes/{id}

#### update

Update /v1/subscriptionOfferCodeOneTimeUseCodes/{id}

#### values

Manage subscriptionOfferCodeOneTimeUseCodes values

1 subcommand(s).

##### get

Get /v1/subscriptionOfferCodeOneTimeUseCodes/{id}/values

### subscriptionOfferCodes

Manage subscriptionOfferCodes

6 subcommand(s).

#### create

Create /v1/subscriptionOfferCodes

#### customCodes

Manage subscriptionOfferCodes customCodes

1 subcommand(s).

##### list

List /v1/subscriptionOfferCodes/{id}/customCodes

#### get

Get /v1/subscriptionOfferCodes/{id}

#### oneTimeUseCodes

Manage subscriptionOfferCodes oneTimeUseCodes

1 subcommand(s).

##### list

List /v1/subscriptionOfferCodes/{id}/oneTimeUseCodes

#### prices

Manage subscriptionOfferCodes prices

1 subcommand(s).

##### list

List /v1/subscriptionOfferCodes/{id}/prices

#### update

Update /v1/subscriptionOfferCodes/{id}

### subscriptionPricePoints

Manage subscriptionPricePoints

2 subcommand(s).

#### equalizations

Manage subscriptionPricePoints equalizations

1 subcommand(s).

##### list

List /v1/subscriptionPricePoints/{id}/equalizations

#### get

Get /v1/subscriptionPricePoints/{id}

### subscriptionPrices

Manage subscriptionPrices

2 subcommand(s).

#### create

Create /v1/subscriptionPrices

#### delete

Delete /v1/subscriptionPrices/{id}

### subscriptionPromotionalOffers

Manage subscriptionPromotionalOffers

5 subcommand(s).

#### create

Create /v1/subscriptionPromotionalOffers

#### delete

Delete /v1/subscriptionPromotionalOffers/{id}

#### get

Get /v1/subscriptionPromotionalOffers/{id}

#### prices

Manage subscriptionPromotionalOffers prices

1 subcommand(s).

##### list

List /v1/subscriptionPromotionalOffers/{id}/prices

#### update

Update /v1/subscriptionPromotionalOffers/{id}

### subscriptionSubmissions

Manage subscriptionSubmissions

1 subcommand(s).

#### create

Create /v1/subscriptionSubmissions

### subscriptions

Manage subscriptions

15 subcommand(s).

#### appStoreReviewScreenshot

Manage subscriptions appStoreReviewScreenshot

1 subcommand(s).

##### get

Get /v1/subscriptions/{id}/appStoreReviewScreenshot

#### create

Create /v1/subscriptions

#### delete

Delete /v1/subscriptions/{id}

#### get

Get /v1/subscriptions/{id}

#### images

Manage subscriptions images

1 subcommand(s).

##### list

List /v1/subscriptions/{id}/images

#### introductoryOffers

Manage subscriptions introductoryOffers

1 subcommand(s).

##### list

List /v1/subscriptions/{id}/introductoryOffers

#### offerCodes

Manage subscriptions offerCodes

1 subcommand(s).

##### list

List /v1/subscriptions/{id}/offerCodes

#### pricePoints

Manage subscriptions pricePoints

1 subcommand(s).

##### list

List /v1/subscriptions/{id}/pricePoints

#### prices

Manage subscriptions prices

1 subcommand(s).

##### list

List /v1/subscriptions/{id}/prices

#### promotedPurchase

Manage subscriptions promotedPurchase

1 subcommand(s).

##### get

Get /v1/subscriptions/{id}/promotedPurchase

#### promotionalOffers

Manage subscriptions promotionalOffers

1 subcommand(s).

##### list

List /v1/subscriptions/{id}/promotionalOffers

#### subscriptionAvailability

Manage subscriptions subscriptionAvailability

1 subcommand(s).

##### get

Get /v1/subscriptions/{id}/subscriptionAvailability

#### subscriptionLocalizations

Manage subscriptions subscriptionLocalizations

1 subcommand(s).

##### list

List /v1/subscriptions/{id}/subscriptionLocalizations

#### update

Update /v1/subscriptions/{id}

#### winBackOffers

Manage subscriptions winBackOffers

1 subcommand(s).

##### list

List /v1/subscriptions/{id}/winBackOffers

### territories

Manage territories

1 subcommand(s).

#### list

List /v1/territories

### territoryAvailabilities

Manage territoryAvailabilities

1 subcommand(s).

#### update

Update /v1/territoryAvailabilities/{id}

### userInvitations

Manage userInvitations

5 subcommand(s).

#### create

Create /v1/userInvitations

#### delete

Delete /v1/userInvitations/{id}

#### get

Get /v1/userInvitations/{id}

#### list

List /v1/userInvitations

#### visibleApps

Manage userInvitations visibleApps

1 subcommand(s).

##### list

List /v1/userInvitations/{id}/visibleApps

### users

Manage users

5 subcommand(s).

#### delete

Delete /v1/users/{id}

#### get

Get /v1/users/{id}

#### list

List /v1/users

#### update

Update /v1/users/{id}

#### visibleApps

Manage users visibleApps

1 subcommand(s).

##### list

List /v1/users/{id}/visibleApps

### webhookDeliveries

Manage webhookDeliveries

1 subcommand(s).

#### create

Create /v1/webhookDeliveries

### webhookPings

Manage webhookPings

1 subcommand(s).

#### create

Create /v1/webhookPings

### webhooks

Manage webhooks

5 subcommand(s).

#### create

Create /v1/webhooks

#### delete

Delete /v1/webhooks/{id}

#### deliveries

Manage webhooks deliveries

1 subcommand(s).

##### list

List /v1/webhooks/{id}/deliveries

#### get

Get /v1/webhooks/{id}

#### update

Update /v1/webhooks/{id}

### winBackOffers

Manage winBackOffers

5 subcommand(s).

#### create

Create /v1/winBackOffers

#### delete

Delete /v1/winBackOffers/{id}

#### get

Get /v1/winBackOffers/{id}

#### prices

Manage winBackOffers prices

1 subcommand(s).

##### list

List /v1/winBackOffers/{id}/prices

#### update

Update /v1/winBackOffers/{id}

