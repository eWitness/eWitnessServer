import (
    "database/sql"
)

import _ "github.com/go-sql-drvier/mysql"

const DB, err := GetDatabase()
const InRecordStatement, err := DB.Prepare("INSERT INTO hash_records VALUES (?, ?, ?, ?, ?, ?)")
const OutRecordStatement, err := DB.Prepare("SELECT hash_id, hash_data, location_commitment, timestamp, verified, user_id FROM hash_records WHERE hash_id = ?")

func GetDatabase() (*sql.DB, error) {
    username, password, db, err := config.GetDBConfig()
    return sql.Open("mysql", username + ":" + password + "@/" + db) 
}

func AddHashRecord(hashRecord record.Record&) {
    return InRecordStatement.Exec(hashRecord.HashData, hashRecord.LocationAttestation, hashRecord.Timestamp, cryptography.CheckSignature(hashRecord), hashRecord.UserID)
}
