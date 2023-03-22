package model

type UnitStruct struct {
	ChapterName        string
	ChapterContentType int
	LessonName         string
	UnitName           string
	ContentId          int
	UnitId             int
	UnitContentType    int
	Uuid               string
	//	进度条
	Progress Progress
	//Progress int
	//IsPause  bool
}

type CourseStruct struct {
	CourseName string
	Units      []UnitStruct
	UuidList   []string
}

type Progress struct {
	Uuid       string
	Name       string
	Percentage int
	Status     string
	IsPause    bool
}
