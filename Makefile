up:
	echo "TODO"

request:
	curl -v -X POST http://localhost:8080/users
	curl -v --header "Content-Type: application/json" --request POST --data '{"id":1,"first_name":"alp","password":"xyz"}' http://localhost:8080/users