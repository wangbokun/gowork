package img

// 待验证，修改
import (
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const (
	DataRoot     = "./tmp/" // 存放封面图的根目录
	TimeoutLimit = 5        // 设置超时时间
)

type VolumeCover struct {
	VolumeID int
	Url      string
	Lock     sync.Mutex
	Msg      chan string
}

func SaveImage(vc *VolumeCover) {
	res, err := http.Get(vc.Url)
	defer res.Body.Close()
	if err != nil {
		vc.Msg <- (strconv.Itoa(vc.VolumeID) + " HTTP_ERROR")
	}
	// 创建文件
	dst, err := os.Create(DataRoot + strconv.Itoa(vc.VolumeID) + ".jpg")
	if err != nil {
		vc.Msg <- (strconv.Itoa(vc.VolumeID) + " OS_ERROR")
	}
	// 生成文件
	_, err := io.Copy(dst, res.Body)
	if err != nil {
		vc.Msg <- (strconv.Itoa(vc.VolumeID) + " COPY_ERROR")
	}
	// goroutine通信
	vc.Lock.Lock()
	vc.Msg <- "in"
	vc.Lock.Unlock()
}
