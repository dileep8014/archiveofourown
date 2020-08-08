package model

import (
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"github.com/shyptr/sqlex"
	"time"
)

const (
	CN = iota + 1
	EN
)

// Article example
type Article struct {
	ID              int64     `json:"id" example:"1" format:"int64"`
	Title           string    `json:"title" example:"article title"`
	SubTitle        string    `json:"subTitle" example:"desc about article"`
	UserId          int64     `json:"userId" example:"1" format:"int64"`
	MarkId          int64     `json:"markId" example:"1" format:"int64"`
	Language        int8      `json:"language" example:"1"`
	Words           int64     `json:"words" example:"1000" format:"int64"`
	ViewNums        int64     `json:"viewNums" example:"100" format:"int64"`
	TalkNums        int64     `json:"talkNums" example:"100" format:"int64"`
	ShareNums       int64     `json:"shareNums" example:"100" format:"int64"`
	DownloadNums    int64     `json:"downloadNums" example:"100" format:"int64"`
	ChapterNums     int64     `json:"chapterNums" example:"100" format:"int64"`
	ChapterRealNums int64     `json:"chapterRealNums" example:"100" format:"int64"`
	CreatedAt       time.Time `json:"createdAt" example:"2000-01-01 00:00:00" format:"datetime"`
	UpdatedAt       time.Time `json:"updatedAt" example:"2000-01-01 00:00:00" format:"datetime"`
}

// Article list example
type ArticleSwagger struct {
	List  []Article `json:"list"`
	Total int       `json:"total"`
}

func (a Article) Create(tx sqlex.BaseRunner) (article Article, err error) {
	defer errwrap.Wrap(&err, "model.article.Create")

	result, err := sqlex.Insert("article").Columns("title,sub_title,user_id,mark_id,language,chapter_nums").
		Values(a.Title, a.SubTitle, a.UserId, a.MarkId, a.Language, a.ChapterNums).
		RunWith(tx).Exec()
	if err != nil {
		return
	}
	a.ID, _ = result.LastInsertId()
	article, err = a.Get(tx)
	return
}

func (a Article) Get(tx sqlex.BaseRunner) (article Article, err error) {
	defer errwrap.Wrap(&err, "model.article.Get")

	queryRow := sqlex.Select("title, sub_title, user_id, mark_id, language, words, view_nums, talk_nums, share_nums",
		"download_nums, chapter_nums, chapter_real_nums, created_at, updated_at").
		From("article").Where("id=?", a.ID).
		RunWith(tx).QueryRow()
	err = queryRow.Scan(&article.Title, &article.SubTitle, &article.UserId, &article.MarkId, &article.Language,
		&article.Words, &article.ViewNums, &article.TalkNums, &article.ShareNums, &article.DownloadNums, &article.ChapterNums,
		&article.ChapterRealNums, &article.CreatedAt, &article.UpdatedAt)
	return
}

func (a Article) Update(tx sqlex.BaseRunner) (err error) {
	defer errwrap.Wrap(&err, "model.article.Update")

	setMap := make(map[string]interface{})
	if a.Title != "" {
		setMap["title"] = a.Title
	}
	if a.SubTitle != "" {
		setMap["sub_title"] = a.SubTitle
	}
	if a.MarkId != 0 {
		setMap["mark_id"] = a.MarkId
	}
	if a.ChapterNums != 0 {
		setMap["chapter_nums"] = a.ChapterNums
	}
	_, err = sqlex.Update("article").SetMap(setMap).Where("id=?", a.ID).RunWith(tx).Exec()
	return
}

func (a Article) Delete(tx sqlex.BaseRunner) (err error) {
	defer errwrap.Wrap(&err, "model.article.Delete")

	_, err = sqlex.Delete("article").Where("id=?", a.ID).RunWith(tx).Exec()
	return
}

func (a Article) List(tx sqlex.BaseRunner, pageSize, pageOffset uint64) (articles []Article, err error) {
	defer errwrap.Wrap(&err, "model.article.List")

	rows, err := sqlex.Select("id,title, sub_title, user_id, mark_id, language, words, view_nums, talk_nums, share_nums",
		"download_nums, chapter_nums, chapter_real_nums, created_at, updated_at").
		From("article").Offset(pageOffset).Limit(pageSize).RunWith(tx).Query()
	if err != nil {
		return
	}
	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID, &article.Title, &article.SubTitle, &article.UserId, &article.MarkId, &article.Language,
			&article.Words, &article.ViewNums, &article.TalkNums, &article.ShareNums, &article.DownloadNums, &article.ChapterNums,
			&article.ChapterRealNums, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			return
		}
	}
	return
}
