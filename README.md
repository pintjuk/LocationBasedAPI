# Locations API
## Design 
* Endpoints contains http specific code
* Geo just contains coordinate code for now

* Then there are 3 bounded contexts, Location, Player and IdP in there sespective directoryes
each one consists of 3 layers with their own responsibilities
    * **Domain:**  Contacts domain objects and business logic, can't import anything
    * **Application:**  uses the domain model to create an application, data validation, data mapping and login can be done here
    * **Data:** Contains implementations of storage solutions, currently only InMemory storage is implemented
    

* Endpoints contains http setup

##Security Issues
Our home built IdP solution might not be th e most secure in the world, so we decided to separate it out into a bounded context from player to avoid mixing security and game play concerns,
and to make future migration to a commercial IdP product simpler

**Note:** that right now Admin Endpoints are not protected by authentication they are all under the path `/Admin`, 
make sure to exclude them in the API-Gateway
 
 
## How To use
### Start
```go run main.go ```
### As Client
#### Register
```bash
> curl  \
 http://localhost:1234/Client/Register \
 -H 'Content-Type: application/json' \
 -d '{"name":"Kalle","password":"1234"}'

```
response:
```Created```
#### Login

```bash
 curl -v \
 http://localhost:1234/Client/Register \
 -H 'Content-Type: application/json' \
 -d '{"name":"Kalle","password":"1234"}'
```

response:
```
{"token":"{JWT}"}
```
#### Get Player
```bash
curl localhost:1234/Client/Player \
-H "Authorization: Bearer {JWT}"
```
response:
```json
{"id":"5b5a8464-2808-4d49-8bdb-b083ae71a4a1","Name":"Kalle","score":0,"lastLong":0,"lastLat":0}
```
#### Change Name
```bash
curl -v -X PUT \
http://localhost:1234/Client/Player/Name \
-H "Authorization: Bearer {JWT}" \
-H 'Content-Type: application/json'  \
-d '{"name":"Kalle"}'

```
#### Read Metadata of a Location from the Location service

here %2810%2C10%29 is uri encoded cordinate (10, 10)

```bash
curl -v -X GET \
         http://localhost:1234/Client/Location/%2810%2C10%29 \
         -H "Authorization: Bearer {JWT}"
```
### As Admin

#### Create Location
```bash
curl -v -X POST \
http://localhost:1234/Admin/Location/%28123%2C123%29 \
-H 'Content-Type: application/json' \
-d '{"name":"Kalle", "type":"aoseuth"}'

```
#### Reahd Location
```bash
curl -v -X GET http://localhost:1234/Admin/Location/%2810%2C10%29
```
#### Updat Location
```bash
curl -v -X PUT \
http://localhost:1234/Admin/Location/%28123%2C123%29 \
-H 'Content-Type: application/json' \
-d '{"name":"New Name","updateName":true,"updateType":true}'


```
#### Delete Location
```bash
curl -v -X DELETE \
http://localhost:1234/Admin/Location/%28123%2C123%29
```
