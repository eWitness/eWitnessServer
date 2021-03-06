package user

import (
    "fmt"
    "Time"
)

/* 
   Here we create the user struct. It  contains:
        1) User ID: a unique integer for each user
        2) PublicKey: a cryptographic public key used to verify signatures
        3) RegistrationDate: the date the user registered
        4) ExpirationDate: the date the user's key expires.
        5) SessionID: used to keep track of logins.
        6) SessionExpiration: used to determine when we should make the 
           client log in again
   This struct will mainly interact with the database and network. 
*/

type User struct {
   UserID            int
   PublicKey         []byte
   RegistrationDate  Time
   ExpirationDate    Time
   SessionID         string
   SessionExpiration Time
}
