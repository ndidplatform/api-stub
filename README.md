# api-stub

## Docker build

`git clone https://github.com/ndidplatform/api-stub`

`cd api-stub`

`docker build -t "api-stub:latest" .`

## Docker Run 
`docker run -it -p 8000:8000 api-stub:latest`

## Example 
### Identity API
`curl -X POST \
  http://localhost:8000/identity \
  -H 'Content-Type: application/json' \
  -d '{
"namespace": "cid",
"identifier": "1234567890129",
"secret": "This is my secret",
"accessor_type": "ed25519",
"accessor_key": "pub_key",
"accessor_id": "aid"
}"'`
