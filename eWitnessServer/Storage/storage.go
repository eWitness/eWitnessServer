/*
    File: Storage/storage.go
    Authors: Kevin Gallagher

    Storage is a package that contains all funcitons necessary for
    database storage. This includes statements for interaction with the
    SQL database, including fetching information and inserting
    information.

    NOTE: I expect this package to change significantly once consensus
    is implimented.
*/

import (
    "database/sql"
    "time"
)

import _ "github.com/go-sql-drvier/mysql"

/* 
   Here we define the constants used in this file.
*/
const DB, err := GetDatabase()
const InRecordStatement, err := DB.Prepare("INSERT INTO hash_records"
                                 + "VALUES (?, ?, ?, ?, ?, ?)")
const OutRecordStatement, err := DB.Prepare("SELECT hash_data,"
                                 + " location_commitment, timestamp, "
                                 + " verified, user_id FROM hash_records "
                                 + " WHERE hash_id = ?")
const InRegisterStatement, err := DB.Prepare("INSERT INTO users VALUES"
                                 + " (?, ?, ?, ?, ?, ?)")
const OutKeyStatement, err := DB.Prepare("SELECT public_key FROM users"
                                 + " WHERE user_id = ?")
const OutSessionStatement, err := DB.Prepare("SELECT session_id,"
                                 + " session_exp FROM users WHERE"
                                 + " user_id = ?")
const UpdateSessionStatement, err := DB.Prepare("UPDATE users SET"
                                 + " session_id = ?, session_exp = ? WHERE"
                                 + " user_id = ?")
const UpdateKeyStatement, err := DB.Prepare("UPDATE users SET "
                                 + " public_key=?, key_exp = ? WHERE"
                                 + " user_id = ?")
const GetMaxUserID, err := DB.Prepare("SELECT MAX(user_id) FROM users")
const RegisterFetchRequest, err := DB.Prepare("INSERT INTO fetch_requests"
                                 + " VALUES (?, ?, ?)")
const GetFetchRequests, err := DB.Prepare("SELECT hash_id, requesting_user"
                                 + " FROM fetch_requests WHERE"
                                 + " client_user=? and fulfilled=false")
/*
   GetDatabase will fetch the database instance stored in the config so
   the functions in this file can interact with the database.

   Args: None
   
   Input: 1) Username, password, and pointer to the db file are gotten
             from a call to the config.GetDBConfig() function.
   
   Output: None

   Return: A DB instance pointer usable by other functions for insertion
           fetching DB data. (type *sql.DB)

*/
func GetDatabase() (*sql.DB, error) {
    username, password, db, err := config.GetDBConfig()
    return sql.Open("mysql", username + ":" + password + "@/" + db)
}

/*
   AddHashRecord takes a record received from the network and adds it
   to the database using the prepared statements from the declared
   constants.

   Args: A hash record (type record.Record)

   Input: Prepared statement (declared constant type *Stmt)

   Output: None

   Returns: An SQL result instance (type sql.Result)
*/
func AddHashRecord(hashRecord record.Record&) {
    return InRecordStatement.Exec(hashRecord.HashData,
                                  hashRecord.LocationAttestation,
                                  hashRecord.Timestamp,
                                  cryptography.CheckSignature(hashRecord),
                                  hashRecord.UserID)
}

/*
   GetHashRecord takes a hash_id and responds with the record data stored
   in the database.

   Args: A hashId (type string)

   Input: Prepared Statement (declared constant type *Stmt)

   Output: None

   Returns: A record containing the data (type record.Record)
*/
func GetHashRecord(hashID string) (record.Record, error){
    var hashData [crypto.SHA3_512]byte
    var locationAttestation [LOCATION_SIZE]byte
    var verified bool
    var timestamp time.Time
    var userID int

    err := OutRecordStatement.QueryRow(hashID).Scan(hashData,
                                              locationAttestation,
                                              timestamp,
                                              verified,
                                              userID)
    if err != nil {
        return (nil, error.New("Database Error: No such hash exists"))

    returnRecord := record.NewHashRecord(hashID, hashData, userID,
                                     locationAttestation, verified,
                                     timestamp)

    return returnRecord, nil
}

/*
   RegisterUser takes a record received from the network and uses the
   information contained to register a user to the eWitness network.

   Args: A hash record (type record.Record)
 
   Input: Prepared statement (declared constant type *Stmt)
          Max current user ID from the database.

   Output: None

   Returns: A session ID (type string)
*/
func RegisterUser(registrationRecord record.Record&) {
    var lastUserID int
    now := time.Now()
    err:= GetMaxUserID.QueryRow().Scan(lastUserID)
    if err != nil {
       return (nil, err)
    }
    sessionString := cryptography.RandomString()
    result, err:= InRegisterRecord.Exec(lastUserID + 1,
                                    registrationRecord.PublicKey,
                                    now,
                                    now.AddDate(1,0,0),
                                    sessionString),
                                    now.AddDate(0,0,3))
    if err != nil {
        return (nil, err)
    }
    return (sessionString, nil)
}

/*
   GetKey returns the public key of a user who has registered with the
   eWitness system. If the desired user doesn't exist, an error is
   returned.

   Args: A userID (type int)

   Input: Prepared Statement (declared constant type *Stmt)

   Output: None

   Returns: A public key for the desired user.
*/
/* TODO: Change the return type to fit a real key. Worry about
   conversion code. */
func GetKey(userID int) (PublicKey, error) {
    var returnKey PublicKey
    err := OutKeyStatement.QueryRow(userID).Scan(returnKey)
    if err != nil {
        return (nil, error.New("Database Error: User does not exist!"))
    }
    return (returnKey, nil)
}
/*
   CheckSession retrieves the session_id of a user and checks if it is
   still valid. If so, it compares it to the passed session key. If they
   are equal, it returns true.

   Args: A passed session_id (type string)
         A passed user_id (type int)

   Input: A session_id taken from the database.

   Output: None

   Returns: A boolean that signifies if sesion_id is valid. (type bool)
*/
func CheckSession(session_id string, user_id int) (bool, error) {
    var currSessionID string
    var expTime time.Time
    now := time.Now()
    err := SessionOutStatement.QueryRow(user_id).Scan(currSessionID, expTime)
    if err != nil {
        return false, error.new("User Error: No such user exists.")
    }
    if session_id == currSessionID && expTime <= now {
        return true, nil
    }
    return false, nil
}

/*
   UpdateSession replaces one session key with a new one. This function
   should only be called after the user is authenticated.

   Args: A userID (type int)

   Input: A prepared Statement

   Output: A database row is updated with a new sessionID.

   Returns: A session identifier (type string)
*/
func UpdateSession(userID int) (string, error) {
    sessionString := cryptography.RandomString()
    now = time.Now()
    result, err := UpdateSessionStatement.Exec(sessionString,
                                              now.AddDate(0,0,3),
                                              userID)
    if err != nil{
        return (nil, error.New("Database error: user does not exist!"))
   return sessionString
}
