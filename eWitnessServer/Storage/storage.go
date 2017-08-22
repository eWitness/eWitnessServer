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
)

import _ "github.com/go-sql-drvier/mysql"

/* 
   Here we define the constants used in this file. They are:
     1) DB - The database connection we use for I/O.
     2) InRecordStatement - The prepared statement we use to store
        records.
     3) OutRecordStatement - The prepared statement we use to 
        retrieve records.
*/
const DB, err := GetDatabase()
const InRecordStatement, err := DB.Prepare("INSERT INTO hash_records"
                                            + "VALUES (?, ?, ?, ?, ?, ?)")
const OutRecordStatement, err := DB.Prepare("SELECT hash_id, hash_data,"
                                 + "location_commitment, timestamp, "
                                 + "verified, user_id FROM hash_records "
                                 + "WHERE hash_id = ?")
const InRegisterStatement, err := DB.Prepare("INSERT INTO users VALUES"
                                  + "(?, ?, ?, ?, ?, ?)")

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


