/*
    Messagining funcitons to handle the network communication
*/

package Messaging

import (
    "fmt"
    "net"
    "io/ioutil"
)

/*
    Below are the INT assignments to given functions that will appear
    in the messaging struct.
*/

const ADD_HASH              := 1
const FETCH_DATA            := 2
const REGISTRATION          := 3
const DISCONNECT            := 4
const ACK                   := 5
const REQUEST_BLOCK_NUMBER  := 6
const AUTH_REQ              := 7
const AUTH_REPLY            := 8
const FAIL                  := 9

// TODO: Need to change the databaseConnection type.

/*
  This function will handle a connection when it comes in.

  Input:  1) A connection to the remote client. (type net.Conn)
          2) A database connection to run queries on. (type string)
  
  Output: 1) Network output to the client. This can be an acknowledge
             message, a failure message, or an authentication challenge.
          2) Storage output to the database referred to by input 2. This
             can be a hash record, a user record, a fetch request update,
             or more.
          3) An integer to indicate success or failure of the function's
             execution.

  TODO: Change the reading mechanism to ioutil readall
*/
func HandleSocket(connection net.Conn, databaseConnection string) {
   defer connection.Close()

   var buffer [512]byte
   var jsonString := ""
   for {
       n, err = conn.Read(buff[0:])
       if err != nil {
           break
       }
       jsonString += string(buff[0:])
   }
   /*
      Here we convert the jsonString to a record and then switch on the
      record's RequestType. Depending on the result we call the 
      appropriate function.
   */
    receivedData = record.DecodeRecord([]byte(jsonString))
    switch receivedData.RequestType {
    case ADD_HASH:
        status, err :=  storage.AddHash(receivedData, databaseConnection)
    case REGISTRATION:
        status, err := storage.RegisterUser(receivedData,
                                            databaseConnection)
    case AUTH_REQ:
        status, err := storage.AuthenticateUser(receivedData,
                                                databaseConnection,
                                                conn)
    case FAIL:
        status, err := alertFailure(receivedData, conn)
    }
    if status > 0 {
        alertAckowledge(receivedData, conn)
    } else {
        alerFailure(receivedData, conn)
    }
}


/*
  This function will acknowledge a successful communication event.

    Input:  1) A record struct containing the client's data.
            2) A connection to the remote client. (type net.Conn)

    Output: 1) Network output to the client. This can be an acknowledge
               message, a failure message, or an authentication challenge.

    Return: 1) An integer, with -1 meaning failure and 0 meaning success.
*/

func alertAkcnowledge(receievedData record.Record, conn net.Conn) int {
    receivedData.RequestType = ACK
    _, err := conn.Write(record.EncodeRecord(receivedData))
    if err == nil {
        return 0
    }
    return -1
}


/*

*/
func alertFailure(receivedData record.Record, conn net.Conn) int {
    receivedData.RequestType = FAIL
    _, err := conn.Write(record.EncodedRecord(receivedData))
    if err == nil {
        return 0
    }
    return -1
}
