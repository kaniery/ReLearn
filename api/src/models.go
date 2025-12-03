package main

import "time"

// ----------------------------------------------------
// DBテーブルに対応する構造体
// ----------------------------------------------------

// User: ユーザー情報 (usersテーブル)
type User struct {
	UserID        int       `json:"user_id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	PasswordHash  string    `json:"password_hash"`
	CreatedAt     time.Time `json:"created_at"`
	LastLogin     time.Time `json:"last_login"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Qualification: 資格・試験情報 (qualificationsテーブル)
type Qualification struct {
	QualificationID int       `json:"qualification_id"`
	Name            string    `json:"name"`
	Provider        string    `json:"provider"`
	ExamDate        time.Time `json:"exam_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Topic: トピック・ジャンル (topicsテーブル)
type Topic struct {
	TopicID         int    `json:"topic_id"`
	QualificationID int    `json:"qualification_id"`
	Name            string `json:"name"`
	ParentTopicID   int    `json:"parent_topic_id"` // 外部キーとして扱うが、JSONではint
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Tag: タグ (tagsテーブル)
type Tag struct {
	TagID int    `json:"tag_id"`
	Name  string `json:"name"`
}

// Question: 質問データ (questionsテーブル)
// QuestionDataフィールドは、実際の質問内容（マークダウン、HTML、JSONなど）を含むためTEXT
type Question struct {
	QuestionID      int       `json:"question_id"`
	QualificationID int       `json:"qualification_id"`
	TopicID         int       `json:"topic_id"`
	AuthorUserID    int       `json:"author_user_id"`
	QuestionData    string    `json:"question_data"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// StudySet: 学習セット (study_setsテーブル)
type StudySet struct {
	SetID         int       `json:"set_id"`
	OwnerUserID   int       `json:"owner_user_id"`
	Name          string    `json:"name"`
	IsPublic      bool      `json:"is_public"` // TINYINT(0/1)はGoではboolで扱うのが一般的
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Response: ユーザーの回答履歴 (responsesテーブル)
type Response struct {
	ResponseID  int       `json:"response_id"`
	UserID      int       `json:"user_id"`
	QuestionID  int       `json:"question_id"`
	IsCorrect   bool      `json:"is_correct"` // TINYINT(0/1)をboolで
	ElapsedMs   int       `json:"elapsed_ms"`
	AnsweredAt  time.Time `json:"answered_at"`
}

// Review: 復習間隔システム (reviewsテーブル)
type Review struct {
	UserID         int       `json:"user_id"`
	QuestionID     int       `json:"question_id"`
	LastReviewAt   time.Time `json:"last_review_at"`
	NextReviewAt   time.Time `json:"next_review_at"`
}

// Attachment: 添付ファイル (attachmentsテーブル)
type Attachment struct {
	AttachmentID int       `json:"attachment_id"`
	QuestionID   int       `json:"question_id"`
	Kind         string    `json:"kind"` // image|audio|video|pdf|link
	URL          string    `json:"url"`
	MetaJSON     string    `json:"meta_json"` // JSON文字列として保存
	CreatedAt    time.Time `json:"created_at"`
}

// ----------------------------------------------------
// 中間テーブル/特殊なテーブル (JSON APIでは直接使われないことが多いが定義)
// ----------------------------------------------------

// QuestionTag: 質問とタグの中間テーブル (question_tagsテーブル)
type QuestionTag struct {
	QuestionID int `json:"question_id"`
	TagID      int `json:"tag_id"`
}

// StudySetQuestion: 学習セット内の質問 (study_set_questionsテーブル)
type StudySetQuestion struct {
	SetID      int `json:"set_id"`
	QuestionID int `json:"question_id"`
	SortOrder  int `json:"sort_order"`
}

// QualificationEnrollment: ユーザーの資格登録状況 (qualification_enrollmentsテーブル)
type QualificationEnrollment struct {
	UserID          int       `json:"user_id"`
	QualificationID int       `json:"qualification_id"`
	TargetDate      time.Time `json:"target_date"`
	Priority        int       `json:"priority"`
	CreatedAt       time.Time `json:"created_at"`
}

// UserFavorite: ユーザーフォロー (user_favoritesテーブル)
type UserFavorite struct {
	FollowerUserID int       `json:"follower_user_id"`
	FavoriteUserID int       `json:"favorite_user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

// UserDataHash: データ変更確認用ハッシュ (user_data_hashesテーブル)
type UserDataHash struct {
	UserID    int       `json:"user_id"`
	HashValue string    `json:"hash_value"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ----------------------------------------------------
// APIリクエスト/レスポンス用のシンプルな構造体（main.goからの移行）
// ----------------------------------------------------

// UserPostRequest: ユーザー登録用 (main.goのUserDataをリネーム)
type UserPostRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

// QuestionPostRequest: 質問登録用 (main.goのQuestionDataをリネーム)
type QuestionPostRequest struct {
	QualificationID int    `json:"qualification_id"`
	TopicID         int    `json:"topic_id"`
	AuthorUserID    int    `json:"author_user_id"`
	QuestionData    string `json:"question_data"`
}
