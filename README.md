# mta_hosting_optimiser

Objective - The project is used to identify inefficient hosts

tech stack used - Golang, Buffalo, mysql

API - A HTTP/REST endpoint to retrieve hostnames having less or equals X active IP addresses. X(threshold is set from variable IP_CONFIG_THRESHOLD in .env )
 
GET - http://127.0.0.1:3000/hosts/inefficient
Return Json arrays of host names

output - 

[
 {
  "hostfqdn": "mta-prod-3"
 },
 {
  "hostfqdn": "mta-prod-1"
 }
]

test cases are integrated 

