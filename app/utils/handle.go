package utils

import (
	"encoding/json"
	"github.com/Esword618/MoocDownload/app/platform/mooc/model"
)

func HandleJsonStr(jsonStr string) model.MyMocTermDto {
	var STRUT model.LastLearnedMocTermDto
	var InfoStruct model.MyMocTermDto
	var contentType int
	var err = json.Unmarshal([]byte(jsonStr), &STRUT)
	if err != nil {
		panic(err)
	}
	courseName := STRUT.Result.MocTermDto.CourseName
	courseName = RemoveInvalidChar(courseName)
	InfoStruct.CourseName = courseName

	chapters := STRUT.Result.MocTermDto.Chapters
	for _, chapter := range chapters {
		var myChapter model.MyChapter
		lessons := chapter.Lessons
		contentType = chapter.ContentType
		chapterName := chapter.Name
		chapterName = RemoveInvalidChar(chapterName)

		myChapter.ContentType = contentType
		myChapter.ChapterName = chapterName

		if contentType == 2 {
			// 有 bug
			var myLesson model.MyLesson
			var myUnit model.MyUnit
			myLesson.LessonsName = chapterName
			myUnit.ContentType = contentType
			myUnit.UnitName = chapterName
			myUnit.UnitId = chapter.ID
			myUnit.ContentId = 0
			myLesson.MyUnits = append(myLesson.MyUnits, myUnit)
			myChapter.MyLessons = append(myChapter.MyLessons, myLesson)

		} else {
			var myLesson model.MyLesson
			for _, lesson := range lessons {
				LessonName := lesson.Name
				LessonName = RemoveInvalidChar(LessonName)
				myLesson.LessonsName = LessonName
				units := lesson.Units

				for _, unit := range units {
					var myUnit model.MyUnit
					contentType = unit.ContentType
					UnitName := unit.Name
					UnitName = RemoveInvalidChar(UnitName)
					switch contentType {
					case 1:
						myUnit.ContentType = contentType
						myUnit.ContentId = 0
						myUnit.UnitName = UnitName
						myUnit.UnitId = unit.ID
					case 3:
						myUnit.ContentType = contentType
						myUnit.ContentId = unit.ContentID
						myUnit.UnitName = UnitName
						myUnit.UnitId = unit.ID
					case 5:
						myUnit.ContentType = contentType
						myUnit.ContentId = unit.ContentID
						myUnit.UnitName = UnitName
						myUnit.UnitId = 0
					case 7:
						myUnit.ContentType = contentType
						myUnit.ContentId = 0
						myUnit.UnitName = UnitName
						myUnit.UnitId = unit.ID
					}
					myLesson.MyUnits = append(myLesson.MyUnits, myUnit)
				}
				myChapter.MyLessons = append(myChapter.MyLessons, myLesson)
			}
			InfoStruct.Chapters = append(InfoStruct.Chapters, myChapter)
		}
	}
	return InfoStruct
}
