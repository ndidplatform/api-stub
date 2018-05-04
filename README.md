# API-STUB

## Docker build

`git clone https://github.com/ndidplatform/api-stub`

`cd api-stub`

`docker build -t "api-stub:latest" .`

## Docker Run 
`docker run -it -p 8000:8000 api-stub:latest`

## Example 
### Identity API
#### POST /identity
`curl -X POST  http://localhost:8000/identity  -H 'Content-Type: application/json'  -d '{"namespace": "cid", "identifier": "1234567890130", "secret": "This is my secret", "accessor_type": "ed25519", "accessor_key": "pub_key", "accessor_id": "aid" }"'`
#### GET /identity/{namespace}/{identifier}
`curl -X GET \
  http://localhost:8000/identity/cid/1234567890123 \
  -H 'Cache-Control: no-cache' \
  -H 'Postman-Token: d7e0430b-de8c-ff09-4ff4-7f9da6807a04'`
