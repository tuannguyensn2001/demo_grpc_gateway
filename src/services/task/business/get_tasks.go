package taskbusiness

import "context"

type biz struct {
}

func NewBiz() *biz {
	return &biz{}
}

func (b *biz) GetTasks(ctx context.Context) {

}
