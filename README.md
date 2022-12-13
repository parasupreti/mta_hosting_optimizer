# mta_hosting_optimiser

Description

Currently, about 35 physical servers host 482 mail transfer agents (MTAs) each having a dedicated public IP address. To be economical while hosting MTAs as a software engineer, we want to create a service that uncovers the inefficient servers hosting only few active MTAs.

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

Sample Data

IP Hostname Active
127.0.0.1 mta-prod-1 true
127.0.0.2 mta-prod-1 false
127.0.0.3 mta-prod-2 true
127.0.0.4 mta-prod-2 true
127.0.0.5 mta-prod-2 false
127.0.0.6 mta-prod-3 false

Expected Resultset
(Using default value for threshold/X of 1).

mta-prod-1
mta-prod-3

