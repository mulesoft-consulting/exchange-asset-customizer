# Exchange Asset Customizer

This is a simple CLI to automate actions on exchange assets, specifically: 
  * updatable standard attributes
  * custom attributes
  * custom categories


## How to use ? 

### Use pre-built binaries 
Navigate to the version of your choice and download the binary depending on your OS.

### Compile code
Use the following to compile the source code: 
```
$ make build
```
### Run

The project is a CLI, and as such, you can always use `--help` option to view usage information. 

```
Usage of ./exchange-asset-customizer:
  -attrfile string
        path to attributes csv file
  -catfile string
        path to categories csv file
  -o string
        anypoint organization id
  -p string
        anypoint password
  -r string
        anypoint region, by default eu (eu/us) (default "eu")
  -tagfile string
        path to tags csv file
  -u string
        anypoint username
```

## Input files
You will find in the `examples/csv` folder examples of csv files format you need to provide for each of the supported types. 

### Attributes

| Name | description | example |
|------|-------------|---------|
|orgId | the business group id where the asset is actually hosted | a83c3c15-a780-499e-8765-36a586ecf211 |
|assetName | the asset name in exchange | system-api-xx |
|fieldKey | the attribute key (or name) from its standard updatable attributes | contactName |
|fieldValue | the value you wish to put in the attribute | user1 |

If you seek to list exchange attributes, checkout [this](https://anypoint.mulesoft.com/exchange/portals/anypoint-platform/f1e97bc6-315a-4490-82a7-23abe036327a.anypoint-platform/exchange-experience-api/minor/2.1/console/method/%231900/) api.

### Custom Attributes

| Name | description | example |
|------|-------------|---------|
|orgId | the business group id where the asset is actually hosted | a83c3c15-a780-499e-8765-36a586ecf211 |
|assetName | the asset name in exchange | system-api-xx |
|version   | the asset version you wish to update | 1.0.0 |
|fieldKey   | the category's key | API Type |
|fieldValue | the value you wish to put in the catagory | val1 |
|fieldType  | the type of the value (date, text) | text |

### Categories

| Name | description | example |
|------|-------------|---------|
|orgId | the business group id where the asset is actually hosted | a83c3c15-a780-499e-8765-36a586ecf211 |
|assetName | the asset name in exchange | system-api-xx |
|version   | the asset version you wish to update | 1.0.0 |
|fieldKey   | the category's key | API Type |
|fieldValue | the values (separated by `;`) you wish to put in the catagory | val1;val2;val3 |


