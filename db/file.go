package db

import (
	"database/sql"
	mydb "filestore/db/mysql"
	"fmt"
)

// OnFileUploadFinished 文件上传完成，保存meta
func OnFileUploadFinished(fileHash string, fileName string, fileSize int64, fileAddr string) bool {
	stmt, err := mydb.DbConn().Prepare("INSERT IGNORE INTO tbl_file(`file_sha1`,`file_name`,`file_size`,`file_addr`,`status`) values(?,?,?,?,1)")
	if err != nil {
		fmt.Println("Failed to prepare statement,err:" + err.Error())
		return false
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("stmt 对象关闭失败")
		}
	}(stmt)

	ret, err := stmt.Exec(fileHash, fileName, fileSize, fileAddr)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("File with hash:%s has been uploaded before", fileHash)
		}
		return true
	}
	return false
}

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

// GetFileMeta 从数据库获取元信息
func GetFileMeta(fileHash string) (*TableFile, error) {
	stmt, err := mydb.DbConn().Prepare("SELECT file_sha1,file_addr,file_name,file_size FROM`tbl_file` WHERE file_sha1 = ? AND STATUS = 1 LIMIT 1")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("stmt 对象关闭失败")
		}
	}(stmt)

	tFile := TableFile{}

	err = stmt.QueryRow(fileHash).Scan(&tFile.FileHash, &tFile.FileAddr, &tFile.FileName, &tFile.FileSize)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &tFile, nil
}
