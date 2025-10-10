package models

type ButtonUrl struct {
	text string
	url  string
}

func NewButtonUrl(text string, url string) *ButtonUrl {
	return &ButtonUrl{
		text: text,
		url:  url,
	}
}

func (b *ButtonUrl) GetText() string {
	return b.text
}

func (b *ButtonUrl) GetURL() string {
	return b.url
}

type PageMain struct {
	title   string
	btnList []*ButtonUrl
}

func NewPageMain(title string, btnList []*ButtonUrl) *PageMain {
	return &PageMain{
		title:   title,
		btnList: btnList,
	}
}

func (p *PageMain) GetTitle() string {
	return p.title
}

func (p *PageMain) GetButton() []*ButtonUrl {
	return p.btnList
}

type PageService struct {
	State int
	pages []*PageMain
}

func NewPageService(pageList []*PageMain) *PageService {

	if len(pageList) == 0 {
		return nil
	}

	return &PageService{
		State: 0,
		pages: pageList,
	}
}

func (s *PageService) GetPageById(id int) *PageMain {
	if id < len(s.pages) {
		return s.pages[id]
	}
	return nil
}

/*func (s *PageService) GetCurrentPage() *ButtonUrl {
	return s.msgList[s.State]
}

func (s *PageService) GetNextPage() *ButtonUrl {
	if (s.State + 1) < len(s.msgList) {
		return s.msgList[s.State+1]
	}
	return nil
}

func (s *PageService) GetPreviousPage() *ButtonUrl {
	if (s.State - 1) > 0 {
		return s.msgList[s.State-1]
	}
	return nil
}*/
