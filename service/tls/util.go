package tls

import (
	"encoding/json"
	"runtime"

	"github.com/gogo/protobuf/proto"
	"github.com/pierrec/lz4"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

func GetCallerFuncName(step int) string {
	// 1 = GetCallerFuncName()
	// 2 = Caller
	// 3 = Caller's caller
	pc, _, _, _ := runtime.Caller(step)
	return runtime.FuncForPC(pc).Name()
}

func CopyIncompressible(src, dst []byte) (int, error) {
	lLen, dn := len(src), len(dst)

	di := 0
	if lLen < 0xF {
		dst[di] = byte(lLen << 4)
	} else {
		dst[di] = 0xF0
		if di++; di == dn {
			return di, nil
		}
		lLen -= 0xF
		for ; lLen >= 0xFF; lLen -= 0xFF {
			dst[di] = 0xFF
			if di++; di == dn {
				return di, nil
			}
		}
		dst[di] = byte(lLen)
	}
	if di++; di+len(src) > dn {
		return di, nil
	}
	di += copy(dst[di:], src)
	return di, nil
}

func GetPutLogsBody(compressType string, logGroupList *pb.LogGroupList) ([]byte, int, error) {
	var (
		out       []byte
		outLen    int
		rawLength int
	)

	body, err := proto.Marshal(logGroupList)
	if err != nil {
		return nil, -1, err
	}

	rawLength = len(body)

	switch compressType {
	case CompressLz4:
		out = make([]byte, lz4.CompressBlockBound(len(body)))
		var hashTable [1 << 16]int
		n, err := lz4.CompressBlock(body, out, hashTable[:])
		if err != nil {
			return nil, -1, err
		}
		if n == 0 {
			n, _ = CopyIncompressible(body, out)
		}
		outLen = n
		break
	default:
		out = body
		outLen = len(out)
	}

	return out[:outLen], rawLength, nil
}

func GetWebTracksBody(compressType string, request *WebTracksRequest) ([]byte, int, error) {
	var (
		out       []byte
		outLen    int
		rawLength int
	)

	body, err := json.Marshal(request)
	if err != nil {
		return nil, -1, err
	}
	rawLength = len(body)

	switch compressType {
	case CompressLz4:
		out = make([]byte, lz4.CompressBlockBound(len(body)))
		var hashTable [1 << 16]int
		n, err := lz4.CompressBlock(body, out, hashTable[:])
		if err != nil {
			return nil, -1, err
		}
		if n == 0 {
			n, _ = CopyIncompressible(body, out)
		}
		outLen = n
	default:
		out = body
		outLen = len(out)
	}

	return out[:outLen], rawLength, nil
}

func GetLogsFromMap(logTime int64, logMap map[string]string) *pb.Log {
	log := &pb.Log{
		Time: logTime,
	}

	for k, v := range logMap {
		log.Contents = append(log.Contents, &pb.LogContent{Key: k, Value: v})
	}

	return log
}
