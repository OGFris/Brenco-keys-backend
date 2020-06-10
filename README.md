# Brenco-keys-backend
A key generation backend for the brenco app written in Go.

## API Documentation

    GET:/api/keys (takes no parameters) to retrieve all the keys in the database.
    
    POST:/api/keys/create (takes one parameter "name" a string) to create a new key in the database.
    
    POST:/api/keys/remove/:id (takes the "id" an int from the route) to remove the key from the database.
