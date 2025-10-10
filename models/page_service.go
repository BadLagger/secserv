package models

type PageService struct {
	State   int
	msgList []string
}

func NewPageService(msgList []string) *PageService {

	if len(msgList) == 0 {
		return nil
	}

	return &PageService{
		State:   0,
		msgList: msgList,
	}
}

func (s *PageService) GetCurrentPage() string {
	return s.msgList[s.State]
}

func (s *PageService) GetNextPage() string {
	if (s.State + 1) < len(s.msgList) {
		return s.msgList[s.State+1]
	}
	return ""
}

func (s *PageService) GetPreviousPage() string {
	if (s.State - 1) > 0 {
		return s.msgList[s.State-1]
	}
	return ""
}
