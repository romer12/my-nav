package DTO

type Group struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Links []Link `json:"links" gorm:"foreignKey:GroupID"` // 设置外键之后，插入记录时会关联插入links表
}

type Link struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	GroupID int    `json:"group_id"`
}

type GoogleJsonData struct {
	Checksum string `json:"checksum"`
	Roots    Roots  `json:"roots"`
}

type Roots struct {
	BookmarkBar BookmarkBar `json:"bookmark_bar"`
}

type BookmarkBar struct {
	DateAdded    string  `json:"date_added"`
	DateLastUsed string  `json:"date_last_used"`
	DateModified string  `json:"date_modified"`
	Guid         string  `json:"guid"`
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Children     []Child `json:"children"`
}

type Child struct {
	DateAdded    string `json:"date_added"`
	DateLastUsed string `json:"date_last_used"`
	Guid         string `json:"guid"`
	Id           string `json:"id"`
	MetaInfo     struct {
		PowerBookmarkMeta string `json:"power_bookmark_meta"`
	} `json:"meta_info"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Url      string  `json:"url"`
	Children []Child `json:"children"`
}

type ReturnJsonStruct struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}
