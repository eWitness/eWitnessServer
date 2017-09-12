/*
   File: Cryptography/cryptography.go
   Authors: Kevin Gallagher

   The cryptography file handles all cryptographic work done by eWitness,
   including signature checking, hashing, and key generation. In addition, 
   this file handles all generation of random values.
*/
package cryptography

import (
    "crypto/rand"
    "time"
    "encoding/base64"
)

/*
   RamdomString generates a random string for use elsewhere in the
   program. It uses the random library with the seed of the current Unix
   time.

   Args: None

   Input: None

   Output: None

   Return: A base64 encoded string of the randomly generated data.
*/

func RandomString() string {
    returnString := make([]byte, 32, 32)
    rand.Read(returnString)
    return base64.StdEncoding.EncodeToString(returnString)
}
/*
   CheckSignature takes as input a public key, a signature, and arbitrary
   data, and checks the signature on that data. It returns a boolean to 
   signify if the signature verification passed.

   Args: PublicKey
         Signature
         Data

   Input: None

   Output: None

   Return: Success or failure (type bool)
*/
