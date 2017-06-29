/* 
   File: FetchRequest/fetchRequest.go
   Authors: Kevin Gallagher

   The purpose of this library is to implement a FetchRequest, or a
   request by an eWitness investigator client to retrieve content from an
   eWitness witness client through an eWitness server. This structure is
   going to be used both to keep records of fetch requests in a database
   and to write fetch requests to the network. As such, it needs functions
   that will write it to the database and that will convert it to a JSON
   string. This library should contain all code necessary for creating and
   fulfilling fetch requests, minus networking and storage code, which
   should be placed into the appropriate files.
*/

package fetchRequest

import (
    "fmt"
    "time"
)

type FetchRequest struct {
    TransactionID   int
    RequestingUser  int
    URL             string
    RequestedDate   Time
    FulfilledDate   Time
    HashID          int
    HashRecord      Record
    RequestingUser  User
}
