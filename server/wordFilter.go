package server

import (
	"sensitive_word/protobuf/grpc/wordFilter"

	"github.com/importcjj/sensitive"

	"golang.org/x/net/context"
)

type WordFilterGRPCStruct struct {
	S *sensitive.Filter
}

// 敏感词 过滤
func (wf *WordFilterGRPCStruct) Filter(ctx context.Context, Content *wordFilter.ContentReq) (*wordFilter.ContentRes, error) {
	rc := wf.S.Filter(Content.Content)

	return &wordFilter.ContentRes{Content: rc}, nil
}

// 敏感词验证 返回结果与第一个敏感词
func (wf *WordFilterGRPCStruct) Validate(ctx context.Context, Content *wordFilter.ContentReq) (*wordFilter.ValidateRes, error) {
	res, word := wf.S.Validate(Content.Content)

	return &wordFilter.ValidateRes{Res: res, Word: word}, nil
}

// 敏感词 匹配返回敏感词组
func (wf *WordFilterGRPCStruct) FindAll(ctx context.Context, Content *wordFilter.ContentReq) (*wordFilter.FindAllRes, error) {
	words := wf.S.FindAll(Content.Content)
	if words == nil {
		words = make([]string, 0)
	}

	return &wordFilter.FindAllRes{Word: words}, nil
}

// 敏感词替换字符
func (wf *WordFilterGRPCStruct) Replace(ctx context.Context, Content *wordFilter.ContentReq) (*wordFilter.ContentRes, error) {
	rc := wf.S.Replace(Content.Content, '*')

	return &wordFilter.ContentRes{Content: rc}, nil
}
