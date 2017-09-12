/*
    File: Record/record.go
    Authors: Kevin Gallagher

    Record defines a struct for holding data that arrives from the
    network. It also contains functions for operating on this data.
    Last, this file contains constants needed to determine what operations
    to perform on the record structure.

*/

package record

import (
    "fmt"
    "time"
    "encoding/pem"
    "encoding/json"
    "crypto/sha3"
    "crypto"
)

/*
   The following are constants for parameters that will be used frequently
   like the size of a nonce, or a location attestation.
*/

const NONCE_SIZE int = 32
const LOCATION_SIZE int = 32

/* 
   The following are shortcuts for ints that are expected in the 
   RequestType field.
*/

const ADD_HASH int = 1
const FETCH_DATA int = 2
const REGISTRATION int = 3
const DISCONNECT int = 4
const ACK int = 5

/* 
   This struct stores the content of a record. An eWitness record must 
   contain five things: 
        1) a hash of evidence content
        2) a location attestation committment
        3) a digital signature
        4) the timestamp of when the record was received
        5) a proxy-URL for evidence retreival
    IDs and other fields may be used for administration, but are not 
    strictly necessary. Other fiends such as nonces, etc, may be necessary
    for other applications. 
*/

type Record struct {
    RequestType         int
    HashData            [crypto.SHA3_512]byte // for use with sha3
    LocationAttestation [LOCATION_SIZE]byte //could change
    HashID              int
    UserID              int
    Signature           string //could change to byte array
    Expiration          time.Time
    PublicKey           crypto.PublicKey
    KeyType             string
    ClientIP            string
    ClientPort          int
    Nonce               [NONCE_SIZE]byte
    Timestamp           time.Time
    SessionKey          string
    Verified            bool
}

/* 
   The struct must support conversion into a valid JSON format.
   The functions applied here will be used to:
        1) Convert a record to a valid JSON format
        2) Convert a valid JSON string to a record
*/

/*
  EncodeRecord will convert a Record to a JSON String.

    Args:   1) A record struct to convert. (type Record)

    Input:  None
                  
    Output: None

    Return: A byte array that represents the json String. 
*/

func EncodeRecord(convertMe Record) []byte {
    returnArray, err := json.Marshal(convertMe)
    if err != nil {
        //Implement error handling here.
        //for now print a line and panic
        fmt.Println("ERROR:", err)
        panic(err)
    }
    return returnArray
}


/*
  DecodeRecord will convert a JSON String into a Record.

    Args:   1) A jsonString to convert. (type byte[])

    Input:  None
                            
    Output: None

    Return: A record containing the information in the JSON string.
*/

func DecodeRecord(jsonString byte[]) Record {
    returnRecord := Record{}
    json.Unmarshal([]byte(jsonString), &returnRecord)
    return returnRecord
}


/*
   The functions below will help in constructing different types of
   records for different uses. 
*/
/*
   NewHashRecord will create a new record that contains the data of a hash
   submission. The present variables represent the columns of the
   hash_record table.

   Args: hashID, the identifier of the hash (type int)
         hashData, the data of the hash (type []byte)
         userID, the identifier of the user who submitted the hash 
             (type int)
         locationAttestation, the location commitment of the photo 
             (type []byte)
         verified, which signifies if the ownership has been verified by 
             checking the cryptographic signature (type bool)
         timestamp, which signifies what time the hash was submitted
             (type time.Time)

   Input: None

   Output: None

   Returns: A new record struct containing this data (type Record)
*/
func NewHashRecord(hashID int, hashData [crypto.SHA3_512]byte,
                   userID, int, locationAttestation [LOCATION_SIZE]byte,
                   verified bool, timestamp time.Time) (Record, error){

    var returnRecord record.Record
    returnRecord.UserID = userID
    returnRecord.HashID = hashID
    returnRecord.HashData = hashData
    returnRecord.LocationAttestation = locationAttestation
    returnRecord.Verified = verified
    returnRecord.Timestamp = timestamp

    return returnRecord
}
