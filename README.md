# common-library
General-purpose common libraries

## MongoDB
mongosh -u "root" -p "root" --authenticationDatabase "admin"


DB and user creation flow
-------------------------
use myDatabase

db.createUser({
  user: "dbUser",                  // Replace with the username
  pwd: "dbPassword123",            // Replace with a secure password
  roles: [{ role: "readWrite", db: "myDatabase" }]
})

mongosh -u "dbUser" -p "dbPassword123" --authenticationDatabase "myDatabase"
