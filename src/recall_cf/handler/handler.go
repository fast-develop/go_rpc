package handler

import (
    //"sync"
	"context"
	//"errors"
	//"sort"
	//"strings"

	//"strconv"

	//"encoding/json"
	//"errors"
	//"fmt"
	log "github.com/golang/glog"

	"recall_cf/proto/recall_cf"
)

func HandleRequest(ctx context.Context, in *recall_cf.RecallCFRequest) (*recall_cf.RecallCFResponse, error) {
    //fmt.Println("HandleRequest")

	rsp := &recall_cf.RecallCFResponse{}

    userId := in.GetUserId()
    docIdList := in.GetDocIdList()
    fromList := in.GetFromList()
    //var finalMap sync.Map
    docMap := make(map[int32]int32)
    for i, docid := range docIdList {
        docMap[docid] = fromList[i]
    }
    
    log.Infof("handle request userid:%s", userId)

    //fmt.Printf("%v",finalMap)

    errId := int32(0)
	rsp.ErrorId = &errId
	//fmt.Println(rsp)

	return rsp, nil
}
