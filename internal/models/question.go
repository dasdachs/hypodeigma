package models

type QuestionType int

const (
	MultipleChoice QuestionType = iota
	TrueOrFalse
	ShortAnswer
)

type AnswerType int

const (
	String AnswerType = iota
	Int
	Boolean
)

type Question struct {
	Question     string
	QuestionType QuestionType
	Choices      []string
	AnswerType   AnswerType
	Answer       string
}
