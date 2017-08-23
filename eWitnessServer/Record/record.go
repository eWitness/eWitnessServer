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
    "encoding/json"
    "crypto/sha3"
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
    Expiration          time.Time //May not compile depending on type
    PublicKey           string //May change to byte array
    ClientIP            string
    ClientPort          int
    Nonce               [NONCE_SIZE]byte
    Timestamp           time.Time
    SessionKey          string
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
