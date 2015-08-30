package archive

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

/******************************************************************************
 @brief
 	文件包遍函文件接口的回调函数,当遍历到一个文件时，会触发回调此函数
 @author
 	chenzhiguo
 @param
	a					文件包
	f					当前遍历到的文件
	curren				当前遍历的文件位置
	total				总文件数
 @return
 	bool				如果返回false表示终断上层遍历操作，否则继续
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
type FileForeachHandler func(a *Archive, f *FileInfo, current, total uint64) bool

/******************************************************************************
 @brief
 	文件包结构，支持多协程访问
 @author
 	chenzhiguo
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
type Archive struct {
	mu                sync.RWMutex         //线程锁
	filepath          string               //路径
	offset            int64                //当前偏移位置
	files             map[string]*FileInfo //文件列表
	local_count       int                  //本地文件数量
	replication_count int                  //复制的文件数量
}

/******************************************************************************
 @brief
 	创建一个文件包
 @author
 	chenzhiguo
 @param
	filepath  			文件包路径
 @return
 	*Archive			返回创建的文件包
 	error				如果返回nil，表示创建成功，否则，表示创建失败
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func Create(filepath string) (*Archive, error) {

	//创建文件
	fp, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_EXCL, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	//生成实例
	var archive Archive
	archive.filepath = filepath
	archive.offset = 0
	archive.files = make(map[string]*FileInfo)

	return &archive, nil
}

/******************************************************************************
 @brief
 	打开一个文件包
 @author
 	chenzhiguo
 @param
	filepath  			文件包路径
 @return
 	*Archive			返回打开的文件包
 	error				如果返回nil，表示创建成功，否则，表示创建失败
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func Open(filepath string) (*Archive, error) {

	//打开文件
	fp, err := os.OpenFile(filepath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	files := make(map[string]*FileInfo)
	offset := int64(0)
	local_count := 0
	replication_count := 0

	//加载文件
	for {
		var f FileInfo
		if err := f.Read(fp, offset); err != nil {
			if err == io.EOF {
				break
			}
			files = nil
			return nil, err
		}

		//更新列表和位移
		files[f.FID] = &f
		fp.Seek(f.Size, os.SEEK_CUR)
		offset = f.offset + f.Size

		if f.IsReplication == 1 {
			replication_count += 1
		} else {
			local_count += 1
		}
	}

	//生成实例
	var archive Archive
	archive.filepath = filepath
	archive.offset = offset
	archive.files = files
	archive.replication_count = replication_count
	archive.local_count = local_count

	//返回结果
	return &archive, nil
}

/******************************************************************************
 @brief
 	返回文件包的路径
 @author
 	chenzhiguo
 @param
	-
 @return
 	string				返回文件包的路径
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) GetPath() string {
	this.mu.RLock()
	defer this.mu.RUnlock()

	return this.filepath
}

/******************************************************************************
 @brief
 	返回文件包的大小
 @author
 	chenzhiguo
 @param
	-
 @return
 	string				返回文件包的大小
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) GetSize() int64 {
	this.mu.RLock()
	defer this.mu.RUnlock()

	return this.offset
}

/******************************************************************************
 @brief
 	判断文件包中是否存在文件
 @author
 	chenzhiguo
 @param
	fid					文件id,即在文件包中的文件名
 @return
 	bool				返回true表示文件存在，否则不存在。
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) IsFileExists(fid string) bool {

	this.mu.RLock()
	defer this.mu.RUnlock()

	_, ok := this.files[fid]
	if ok {
		return true
	}

	return false
}

/******************************************************************************
 @brief
 	判断文件包中的文件是否已经删除
 @author
 	chenzhiguo
 @param
	fid					文件id,即在文件包中的文件名
 @return
 	bool				返回true表示文件已经删除，否则没有。
 	error				返回nil表示成功，否则表示错误原因
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) IsFileDelete(fid string) (bool, error) {

	this.mu.RLock()
	defer this.mu.RUnlock()

	file, ok := this.files[fid]
	if !ok {
		return false, fmt.Errorf("Archive.IsDelete failed! err: fid=%s not exists!", fid)
	}

	if file.IsDelete == 1 {
		return true, nil
	}

	return false, nil
}

/******************************************************************************
 @brief
 	判断文件包中的文件是否已经被复制
 @author
 	chenzhiguo
 @param
	fid					文件id,即在文件包中的文件名
 @return
 	bool				返回true表示文件已经复制，否则没有。
 	error				返回nil表示成功，否则表示错误原因
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) IsFileReplication(fid string) (bool, error) {

	this.mu.RLock()
	defer this.mu.RUnlock()

	file, ok := this.files[fid]
	if !ok {
		return false, fmt.Errorf("Archive.IsFileReplication failed! err: fid=%s not exists!", fid)
	}

	if file.IsReplication == 1 {
		return true, nil
	}

	return false, nil
}

/******************************************************************************
 @brief
 	取得文件包中复制的文件数量
 @author
 	chenzhiguo
 @param
	-
 @return
 	int					返回数量
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) GetFileReplicationCount() int {
	this.mu.RLock()
	defer this.mu.RUnlock()

	return this.replication_count
}

/******************************************************************************
 @brief
 	取得文件包中本地的文件数量，就是不包含复制的数量
 @author
 	chenzhiguo
 @param
	-
 @return
 	int					返回数量
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) GetFileLocalCount() int {
	this.mu.RLock()
	defer this.mu.RUnlock()

	return this.local_count
}

/******************************************************************************
 @brief
 	取得文件包中总的文件数量，总数量＝本地的数量＋复制的数量
 @author
 	chenzhiguo
 @param
	-
 @return
 	int					返回数量
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) GetFileTotalCount() int {
	this.mu.RLock()
	defer this.mu.RUnlock()

	return len(this.files)
}

/******************************************************************************
 @brief
 	遍历文件包中的文件
 @author
 	chenzhiguo
 @param
	f					遍历处理函数
 @return
 	-
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) ForeachFile(f FileForeachHandler) {

	this.mu.RLock()
	defer this.mu.RUnlock()

	current := uint64(0)
	total := uint64(len(this.files))

	for _, file := range this.files {
		if !f(this, file, current, total) {
			break
		}

		current++
	}
}

/******************************************************************************
 @brief
 	取得文件信息
 @author
 	chenzhiguo
 @param
	fid					文件id,即在文件包中的文件名
 @return
 	*FileInfo			返回文件信息
 @history
 	2015-05-28_19:19 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) GetFileInfo(fid string) *FileInfo {
	this.mu.RLock()
	defer this.mu.RUnlock()

	file, ok := this.files[fid]
	if !ok {
		return nil
	}

	return file
}

/******************************************************************************
 @brief
 	从文件包中提取文件
 @author
 	chenzhiguo
 @param
	fid					文件id,即在文件包中的文件名
	force				强制开关，如果force=true，那么即使文件删除了，也要读出来数据，如果force=false表示删除了或其他原因就不会实际读取文件数据
 @return
 	[]byte				文件数据
 	error				如果为nil表示提取成功，否则表示提取失败
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) ExtractFile(fid string, force bool) ([]byte, error) {

	this.mu.RLock()
	defer this.mu.RUnlock()

	file, ok := this.files[fid]
	if !ok {
		return nil, fmt.Errorf("Archive.Extract failed! err : fid=%s not exists!", fid)
	}

	//检查是否是强制解压
	if force != false {
		if file.IsDelete == 1 {
			return nil, fmt.Errorf("Archive.Extract failed! err : fid=%s be delete!", fid)
		}
	}

	if file.Size == 0 {
		return []byte{}, nil
	}

	fp, err := os.OpenFile(this.filepath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	buf := make([]byte, file.Size)
	n, err := fp.ReadAt(buf, file.offset)
	if err != nil {
		return nil, err
	}

	if n != int(file.Size) || n > int(file.Size) {
		return nil, fmt.Errorf("Archive.Extract failed! err : fid=%s n=%d file.Size=%d", fid, n, file.Size)
	}

	return buf[:n], nil
}

/******************************************************************************
 @brief
 	向文件包中以追加的方式追加文件
 @author
 	chenzhiguo
 @param
	fid					文件id,即在文件包中的文件名
	filename			本地路径
	isReplication		可选项，标记此文件是复制的，还是本地的，默认为本地的
 @return
 	error				如果为nil表示成功，否则表示失败
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) AppendFile(fid, filename string, isReplication ...bool) error {

	data, err := ReadFile(filename)
	if err != nil {
		return err
	}

	return this.AppendData(fid, data, isReplication...)
}

/******************************************************************************
 @brief
 	向文件包中以追加的方式追加文件
 @author
 	chenzhiguo
 @param
	fid					文件id,即在文件包中的文件名
	data				要追加的文件内容
	isReplication		标记此文件是复制的，还是本地的
 @return
 	error				如果为nil表示成功，否则表示失败
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) AppendData(fid string, data []byte, isReplication ...bool) error {

	this.mu.Lock()
	defer this.mu.Unlock()

	_, ok := this.files[fid]
	if ok {
		return fmt.Errorf("Archive.Append failed! err: fid=%s exists!", fid)
	}

	var f FileInfo
	f.FID = fid                     //2
	f.MD5 = MakeByteMd5(data)       //2
	f.Size = int64(len(data))       //8
	f.IsDelete = 0                  //1
	f.IsReplication = 0             //1
	f.TimeStamp = time.Now().Unix() //8

	if len(isReplication) > 0 && isReplication[0] {
		f.IsReplication = 1
	}

	return this.writeFileData(&f, data)
}

/******************************************************************************
 @brief
 	将文件信息序列化到文件包文件中
 @author
 	chenzhiguo
 @param
	f					文件信息头（fid,md5,size,isdelete,isreplication,timestamp)
	data				文件内容
 @return
 	error				如果为nil表示成功，否则表示失败原因
 @history
 	2015-05-16_08:57 	chenzhiguo		创建
*******************************************************************************/
func (this *Archive) writeFileData(f *FileInfo, data []byte) error {

	fp, err := os.OpenFile(this.filepath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer fp.Close()

	pos := this.offset
	n, err := f.Write(fp, pos)
	if err != nil {
		return err
	}
	pos += int64(n)
	f.offset = pos //记录位置

	//注意如果数据为0的时候的情况
	n = 0
	if data != nil && len(data) > 0 {
		n, err = fp.WriteAt(data, pos)
		if err != nil {
			return err
		}
	}

	if n != len(data) {
		return fmt.Errorf("Archive.writeFileData data filed! n=%d size=%d f=%#v", n, len(data), f)
	}
	pos += int64(n)

	this.files[f.FID] = f
	this.offset = pos

	if f.IsReplication == 1 {
		this.replication_count += 1
	} else {
		this.local_count += 1
	}

	return nil
}
