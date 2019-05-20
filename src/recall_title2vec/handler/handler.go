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

	"recall_title2vec/proto/recall_title2vec"
)

func HandleRequest(ctx context.Context, in *recall_title2vec.RecallTitle2VecRequest) (*recall_title2vec.RecallTitle2VecResponse, error) {
    //fmt.Println("HandleRequest")

	rsp := &recall_title2vec.RecallTitle2VecResponse{}

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
