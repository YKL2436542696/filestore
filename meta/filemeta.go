package meta

import (
	mydb "filestore/db"
)

// FileMeta 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta :新增/更新文件元信息
func UpdateFileMeta(fm FileMeta) {
	fileMetas[fm.FileSha1] = fm
}

// UpdateFileMetaDB 新增/更新文件元信息到mysql
func UpdateFileMetaDB(fm FileMeta) bool {
	return mydb.OnFileUploadFinished(fm.FileSha1, fm.FileName, fm.FileSize, fm.Location)
}

// GetFileMeta : 通过sha1的值获取文件元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// GetFileMetaDB  从数据库获取元信息
func GetFileMetaDB(fileSha1 string) (FileMeta, error) {
	tFile, err := mydb.GetFileMeta(fileSha1)
	if err != nil {
		return FileMeta{}, err
	}
	fMeta := FileMeta{
		FileSha1: tFile.FileHash,
		FileName: tFile.FileName.String,
		FileSize: tFile.FileSize.Int64,
		Location: tFile.FileAddr.String,
	}
	return fMeta, nil
}

// RemoveFileMeta 删除
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
