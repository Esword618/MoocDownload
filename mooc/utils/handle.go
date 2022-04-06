package utils

import (
	"encoding/json"

	"MoocDownload/mooc/model"
)

func HandleJsonStr(jsonStr string) model.MyMocTermDto {
	var STRUT model.LastLearnedMocTermDto
	var InfoStruct model.MyMocTermDto
	var contentType int
	var err = json.Unmarshal([]byte(jsonStr), &STRUT)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v", STRUT)
	courseName := STRUT.Result.MocTermDto.CourseName
	InfoStruct.CourseName = courseName
	chapters := STRUT.Result.MocTermDto.Chapters
	for _, chapter := range chapters {
		// fmt.Println(i, "------", chapter)
		var myChapter model.MyChapter
		lessons := chapter.Lessons
		contentType = chapter.ContentType
		chapterName := chapter.Name
		myChapter.ContentType = contentType
		myChapter.ChapterName = chapterName
		if contentType == 2 {
			var myUnit model.MyUnit
			myUnit.ContentType = contentType
			myUnit.UnitName = chapterName
			myUnit.UnitId = chapter.ID
			myUnit.ContentId = 0

			myChapter.MyUnits = append(myChapter.MyUnits, myUnit)

		} else {
			for _, lesson := range lessons {
				units := lesson.Units
				for _, unit := range units {
					var myUnit model.MyUnit
					contentType = unit.ContentType
					switch contentType {
					case 1:
						myUnit.ContentType = contentType
						myUnit.ContentId = 0
						myUnit.UnitName = unit.Name
						myUnit.UnitId = unit.ID
					case 3:
						myUnit.ContentType = contentType
						myUnit.ContentId = unit.ContentID
						myUnit.UnitName = unit.Name
						myUnit.UnitId = unit.ID
					case 5:
						myUnit.ContentType = contentType
						myUnit.ContentId = unit.ContentID
						myUnit.UnitName = unit.Name
						myUnit.UnitId = 0
					case 7:
						myUnit.ContentType = contentType
						myUnit.ContentId = 0
						myUnit.UnitName = unit.Name
						myUnit.UnitId = unit.ID
					}

					myChapter.MyUnits = append(myChapter.MyUnits, myUnit)

				}
			}
			InfoStruct.Chapters = append(InfoStruct.Chapters, myChapter)
		}

	}
	// InfoStruct.Chapters = MyChaptersList
	// fmt.Println(InfoStruct)
	// for _, chapter := range InfoStruct.Chapters {
	//	fmt.Println(chapter.ChapterName)
	//	fmt.Println(chapter.ContentType)
	//	for _, i := range chapter.MyUnits {
	//		fmt.Println(i)
	//	}
	//	fmt.Println("\n\r---------\n\r")
	// }
	return InfoStruct
}
