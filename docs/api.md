


# Resume
A Test API for testing lambda/go-swagger integration
  

## Informations

### Version

1.0.0

### License

[Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### Contact

 spam@ChrisScotMartin.com 

## Tags

  ### <span id="tag-open"></span>open

Calls which are unrestricted

  ### <span id="tag-user"></span>user

Calls which require authentication

  ### <span id="tag-admin"></span>admin

Calls which require admin rights

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  open

  

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /resume | [create resume](#create-resume) | Create a new resume |
| DELETE | /resume/{id} | [delete resume](#delete-resume) | Delete a resume |
| GET | / | [get Api identifier](#get-api-identifier) | API Identifier endpoint |
| GET | /resume/{id} | [get resume](#get-resume) | Get a resume by ID |
| PUT | /resume/{id} | [update resume](#update-resume) | Update an existing resume |
  


## Paths

### <span id="create-resume"></span> Create a new resume (*createResume*)

```
POST /resume
```

Create a new resume
---


#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Resume | `body` | [Resume](#resume) | `models.Resume` | |  | | Resume to create |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-resume-201) | Created | Resume created |  | [schema](#create-resume-201-schema) |
| [400](#create-resume-400) | Bad Request | Invalid input |  | [schema](#create-resume-400-schema) |
| [500](#create-resume-500) | Internal Server Error | General Failure |  | [schema](#create-resume-500-schema) |

#### Responses


##### <span id="create-resume-201"></span> 201 - Resume created
Status: Created

###### <span id="create-resume-201-schema"></span> Schema
   
  

[UUIDResponse](#uuid-response)

##### <span id="create-resume-400"></span> 400 - Invalid input
Status: Bad Request

###### <span id="create-resume-400-schema"></span> Schema
   
  

[SimpleErrorResponse](#simple-error-response)

##### <span id="create-resume-500"></span> 500 - General Failure
Status: Internal Server Error

###### <span id="create-resume-500-schema"></span> Schema
   
  

[SimpleErrorResponse](#simple-error-response)

### <span id="delete-resume"></span> Delete a resume (*deleteResume*)

```
DELETE /resume/{id}
```

Delete a resume
---


#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | ID of resume to delete |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-resume-200) | OK | Resume deleted |  | [schema](#delete-resume-200-schema) |
| [404](#delete-resume-404) | Not Found | Resume not found |  | [schema](#delete-resume-404-schema) |
| [500](#delete-resume-500) | Internal Server Error | General Failure |  | [schema](#delete-resume-500-schema) |

#### Responses


##### <span id="delete-resume-200"></span> 200 - Resume deleted
Status: OK

###### <span id="delete-resume-200-schema"></span> Schema
   
  

[UUIDResponse](#uuid-response)

##### <span id="delete-resume-404"></span> 404 - Resume not found
Status: Not Found

###### <span id="delete-resume-404-schema"></span> Schema
   
  

[SimpleErrorResponse](#simple-error-response)

##### <span id="delete-resume-500"></span> 500 - General Failure
Status: Internal Server Error

###### <span id="delete-resume-500-schema"></span> Schema
   
  

[SimpleErrorResponse](#simple-error-response)

### <span id="get-api-identifier"></span> API Identifier endpoint (*getApiIdentifier*)

```
GET /
```

Endpoint which returns the API version and the running backend version

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-api-identifier-200) | OK | Returns API version and running backend version |  | [schema](#get-api-identifier-200-schema) |
| [500](#get-api-identifier-500) | Internal Server Error | General Failure |  | [schema](#get-api-identifier-500-schema) |

#### Responses


##### <span id="get-api-identifier-200"></span> 200 - Returns API version and running backend version
Status: OK

###### <span id="get-api-identifier-200-schema"></span> Schema
   
  

[SimpleMessageResponse](#simple-message-response)

##### <span id="get-api-identifier-500"></span> 500 - General Failure
Status: Internal Server Error

###### <span id="get-api-identifier-500-schema"></span> Schema

### <span id="get-resume"></span> Get a resume by ID (*getResume*)

```
GET /resume/{id}
```

Get a resume by ID
---


#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | ID of resume to return |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-resume-200) | OK | successful operation |  | [schema](#get-resume-200-schema) |
| [404](#get-resume-404) | Not Found | Resume not found |  | [schema](#get-resume-404-schema) |
| [500](#get-resume-500) | Internal Server Error | General Failure |  | [schema](#get-resume-500-schema) |

#### Responses


##### <span id="get-resume-200"></span> 200 - successful operation
Status: OK

###### <span id="get-resume-200-schema"></span> Schema
   
  

[Resume](#resume)

##### <span id="get-resume-404"></span> 404 - Resume not found
Status: Not Found

###### <span id="get-resume-404-schema"></span> Schema
   
  

[SimpleErrorResponse](#simple-error-response)

##### <span id="get-resume-500"></span> 500 - General Failure
Status: Internal Server Error

###### <span id="get-resume-500-schema"></span> Schema

### <span id="update-resume"></span> Update an existing resume (*updateResume*)

```
PUT /resume/{id}
```

Update an existing resume
---


#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | ID of resume to update |
| resume | `body` | [Resume](#resume) | `models.Resume` | | ✓ | | Resume object |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-resume-200) | OK | Resume updated |  | [schema](#update-resume-200-schema) |
| [400](#update-resume-400) | Bad Request | Invalid input |  | [schema](#update-resume-400-schema) |
| [404](#update-resume-404) | Not Found | Resume not found |  | [schema](#update-resume-404-schema) |
| [500](#update-resume-500) | Internal Server Error | General Failure |  | [schema](#update-resume-500-schema) |

#### Responses


##### <span id="update-resume-200"></span> 200 - Resume updated
Status: OK

###### <span id="update-resume-200-schema"></span> Schema
   
  

[UUIDResponse](#uuid-response)

##### <span id="update-resume-400"></span> 400 - Invalid input
Status: Bad Request

###### <span id="update-resume-400-schema"></span> Schema
   
  

[SimpleErrorResponse](#simple-error-response)

##### <span id="update-resume-404"></span> 404 - Resume not found
Status: Not Found

###### <span id="update-resume-404-schema"></span> Schema
   
  

[SimpleErrorResponse](#simple-error-response)

##### <span id="update-resume-500"></span> 500 - General Failure
Status: Internal Server Error

###### <span id="update-resume-500-schema"></span> Schema
   
  

[SimpleErrorResponse](#simple-error-response)

## Models

### <span id="resume"></span> Resume


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| awards | [][ResumeAwardsItems0](#resume-awards-items0)| `[]*ResumeAwardsItems0` |  | |  |  |
| basics | [ResumeBasics](#resume-basics)| `ResumeBasics` |  | |  |  |
| certificates | [][ResumeCertificatesItems0](#resume-certificates-items0)| `[]*ResumeCertificatesItems0` |  | |  |  |
| education | [][ResumeEducationItems0](#resume-education-items0)| `[]*ResumeEducationItems0` |  | |  |  |
| interests | [][ResumeInterestsItems0](#resume-interests-items0)| `[]*ResumeInterestsItems0` |  | |  |  |
| languages | [][ResumeLanguagesItems0](#resume-languages-items0)| `[]*ResumeLanguagesItems0` |  | |  |  |
| projects | [][ResumeProjectsItems0](#resume-projects-items0)| `[]*ResumeProjectsItems0` |  | |  |  |
| publications | [][ResumePublicationsItems0](#resume-publications-items0)| `[]*ResumePublicationsItems0` |  | |  |  |
| references | [][ResumeReferencesItems0](#resume-references-items0)| `[]*ResumeReferencesItems0` |  | |  |  |
| skills | [][ResumeSkillsItems0](#resume-skills-items0)| `[]*ResumeSkillsItems0` |  | |  |  |
| volunteer | [][ResumeVolunteerItems0](#resume-volunteer-items0)| `[]*ResumeVolunteerItems0` |  | |  |  |
| work | [][ResumeWorkItems0](#resume-work-items0)| `[]*ResumeWorkItems0` |  | |  |  |



#### Inlined models

**<span id="resume-awards-items0"></span> ResumeAwardsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| awarder | string| `string` |  | |  | `Company` |
| date | date (formatted string)| `strfmt.Date` |  | |  | `2014-11-01` |
| summary | string| `string` |  | |  | `There is no spoon.` |
| title | string| `string` |  | |  | `Award` |



**<span id="resume-basics"></span> ResumeBasics**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  | `john@gmail.com` |
| image | string| `string` |  | |  | `https://johndoe.com/me.jpg` |
| label | string| `string` |  | |  | `Programmer` |
| location | [ResumeBasicsLocation](#resume-basics-location)| `ResumeBasicsLocation` |  | |  |  |
| name | string| `string` | ✓ | |  | `John Doe` |
| phone | string| `string` | ✓ | |  | `(912) 555-4321` |
| profiles | [][ResumeBasicsProfilesItems0](#resume-basics-profiles-items0)| `[]*ResumeBasicsProfilesItems0` |  | |  |  |
| summary | string| `string` | ✓ | |  | `A summary of John Doe…` |
| url | string| `string` |  | |  | `https://johndoe.com` |



**<span id="resume-basics-location"></span> ResumeBasicsLocation**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| address | string| `string` |  | |  | `2712 Broadway St` |
| city | string| `string` |  | |  | `San Francisco` |
| countryCode | string| `string` |  | |  | `US` |
| postalCode | string| `string` |  | |  | `CA 94115` |
| region | string| `string` |  | |  | `California` |



**<span id="resume-basics-profiles-items0"></span> ResumeBasicsProfilesItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| network | string| `string` |  | |  | `Twitter` |
| url | string| `string` |  | |  | `https://twitter.com/john` |
| username | string| `string` |  | |  | `john` |



**<span id="resume-certificates-items0"></span> ResumeCertificatesItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| date | date (formatted string)| `strfmt.Date` |  | |  | `2021-11-07` |
| issuer | string| `string` |  | |  | `Company` |
| name | string| `string` |  | |  | `Certificate` |
| url | string| `string` |  | |  | `https://certificate.com` |



**<span id="resume-education-items0"></span> ResumeEducationItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| area | string| `string` |  | |  | `Software Development` |
| courses | []string| `[]string` |  | |  |  |
| endDate | date (formatted string)| `strfmt.Date` |  | |  | `2013-01-01` |
| institution | string| `string` |  | |  | `University` |
| score | string| `string` |  | |  | `4.0` |
| startDate | date (formatted string)| `strfmt.Date` |  | |  | `2011-01-01` |
| studyType | string| `string` |  | |  | `Bachelor` |
| url | string| `string` |  | |  | `https://institution.com/` |



**<span id="resume-interests-items0"></span> ResumeInterestsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| keywords | []string| `[]string` |  | |  |  |
| name | string| `string` |  | |  | `Wildlife` |



**<span id="resume-languages-items0"></span> ResumeLanguagesItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| fluency | string| `string` |  | |  | `Native speaker` |
| language | string| `string` |  | |  | `English` |



**<span id="resume-projects-items0"></span> ResumeProjectsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| description | string| `string` |  | |  | `Description...` |
| endDate | date (formatted string)| `strfmt.Date` |  | |  | `2021-01-01` |
| highlights | []string| `[]string` |  | |  |  |
| name | string| `string` |  | |  | `Project` |
| startDate | date (formatted string)| `strfmt.Date` |  | |  | `2019-01-01` |
| url | string| `string` |  | |  | `https://project.com/` |



**<span id="resume-publications-items0"></span> ResumePublicationsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | |  | `Publication` |
| publisher | string| `string` |  | |  | `Company` |
| releaseDate | date (formatted string)| `strfmt.Date` |  | |  | `2014-10-01` |
| summary | string| `string` |  | |  | `Description…` |
| url | string| `string` |  | |  | `https://publication.com` |



**<span id="resume-references-items0"></span> ResumeReferencesItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | |  | `Jane Doe` |
| reference | string| `string` |  | |  | `Reference…` |



**<span id="resume-skills-items0"></span> ResumeSkillsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| keywords | []string| `[]string` |  | |  |  |
| level | string| `string` |  | |  | `Master` |
| name | string| `string` |  | |  | `Web Development` |



**<span id="resume-volunteer-items0"></span> ResumeVolunteerItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| endDate | date (formatted string)| `strfmt.Date` |  | |  | `2013-01-01` |
| highlights | []string| `[]string` |  | |  |  |
| organization | string| `string` |  | |  | `Organization` |
| position | string| `string` |  | |  | `Volunteer` |
| startDate | date (formatted string)| `strfmt.Date` |  | |  | `2012-01-01` |
| summary | string| `string` |  | |  | `Description…` |
| url | string| `string` |  | |  | `https://organization.com/` |



**<span id="resume-work-items0"></span> ResumeWorkItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| endDate | date (formatted string)| `strfmt.Date` |  | |  | `2014-01-01` |
| highlights | []string| `[]string` |  | |  |  |
| name | string| `string` |  | |  | `Company` |
| position | string| `string` |  | |  | `President` |
| startDate | date (formatted string)| `strfmt.Date` |  | |  | `2013-01-01` |
| summary | string| `string` |  | |  | `Description…` |
| url | string| `string` |  | |  | `https://company.com` |



### <span id="resume-id"></span> ResumeID


> ID of resume to return
  



| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| ResumeID | uuid (formatted string)| strfmt.UUID | | ID of resume to return |  |



### <span id="resume-metadata"></span> ResumeMetadata


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| created | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| tags | map of string| `map[string]string` |  | |  | `{"key1":"value1","key2":"value2"}` |
| updated | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| user | string| `string` |  | |  | `John Doe` |



### <span id="simple-error-response"></span> SimpleErrorResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` | ✓ | |  |  |



### <span id="simple-message-response"></span> SimpleMessageResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| message | string| `string` | ✓ | |  |  |



### <span id="uuid-response"></span> UUIDResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | uuid (formatted string)| `strfmt.UUID` | ✓ | |  |  |


