# ATTOM API Reference


## Property Expanded Profile

A detailed profile of a property.  Requires a address from the PropertyDetails

Address format:
 - ADDRESS1: 4529 Winona Court
 - ADDRES2: Denver, CO

```bash
curl -X GET --header 'Accept: application/json' --header 'apikey: [REDACTED]' --header 'accept: application/json' 'https://api.gateway.attomdata.com/propertyapi/v1.0.0/property/expandedprofile?address1=[ADDRESS1]&address2=[ADDRESS2]'
```

Payload: <a href="property/expandedprofile.json">property/expandedprofile.json</a>


## Neighborhood

Requires a GEOIDv4 from the PropertyDetails

```bash
curl -X GET --header 'Accept: application/json' --header 'ApiKey: [REDACTED]' --header 'accept: application/json' 'https://api.gateway.attomdata.com/v4/neighborhood/community?geoIdv4=[GEOIDv4]'
```

Payload: <a href="neighborhood/community.json">neighborhood/community.json</a>

## Events Detail

A historical record of events for a property.  Requires a ATTOM_ID from the PropertyDetails

```bash
curl -X GET --header 'Accept: application/json' --header 'apikey: [REDACTED]' --header 'accept: application/json' 'https://api.gateway.attomdata.com/propertyapi/v1.0.0/allevents/detail?id=[ATTOM_ID]'
```

Payload: <a href="allevents/detail.json">allevents/detail.json</a>

## Assessment History Detail

A list of historical assessments from the county.  Requires a ATTOM_ID from the PropertyDetails

```bash
curl -X GET --header 'Accept: application/json' --header 'apikey: [REDACTED]' --header 'accept: application/json' 'https://api.gateway.attomdata.com/propertyapi/v1.0.0/assessmenthistory/detail?attomid=[ATTOM_ID]'
```

Payload: <a href="assessmenthistory/detail.json">assessmenthistory/detail.json</a>

## Building Permits

A list of building permits for a property. Requires an ATTOM_ID from the PropertyDetails

```bash
curl -X GET --header 'Accept: application/json' --header 'apikey: [REDACTED]' --header 'accept: application/json' 'https://api.gateway.attomdata.com/propertyapi/v1.0.0/property/buildingpermits?attomid=[ATTOM_ID]'
```

Payload: <a href="property/buildingpermits.json">property/buildingpermits.json</a>

## Transportation Noise

A list of transportation noise sources for a property.  Requires a address from the PropertyDetails 

Address format: `4301 MURRAY ST, FLUSHING, NY 11355`

```bash
curl -X GET --header 'Accept: application/json' --header 'apikey: [REDACTED]' --header 'accept: application/json' 'https://api.gateway.attomdata.com/transportationnoise?address=[ADDRESS]'
```

Payload: <a href="transportationnoise.json">transportationnoise.json</a>

## Sales Trend

A list of sales trends for a property.  Requires a ATTOM_ID from the PropertyDetails

```bash
curl -X GET --header 'Accept: application/json' --header 'apikey: [REDACTED]' --header 'accept: application/json' 'https://api.gateway.attomdata.com/v4/transaction/salestrend?geoIdV4=[GEOIDV4]&interval=yearly&startyear=[STARTYEAR]&endyear=[ENDYEAR]&propertytype=[PROPERTYTYPE]&transfertype=ALL'
```

Property Types
```
ALL
CONDO/TOWNHOME
COOPERATIVE
FARM/RANCH
INDUSTRIAL
MANUFACTURED
MISCELLANEOUS
MULTI%20(2-4%20Units)
MUTLTI%205+
OFFICE
OTHER
RETAIL
SIGLE%20FAMILY%20RESIDENCE
VACANT%20LAND
```