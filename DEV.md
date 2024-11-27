# DEV

curl -X POST http://localhost:8000/login -d '{"username":"guionardo","password":"1234"}' -w "%{http_code}" -H "Content-type:application/json"