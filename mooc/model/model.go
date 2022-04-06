package model

type Status struct {
	Code   int `json:"code"`
	Result struct {
		MemberID            int         `json:"memberId"`
		SchoolName          interface{} `json:"schoolName"`
		SchoolID            interface{} `json:"schoolId"`
		NickName            string      `json:"nickName"`
		LargeFaceURL        string      `json:"largeFaceUrl"`
		Department          interface{} `json:"department"`
		MemberType          int         `json:"memberType"`
		HighestDegree       int         `json:"highestDegree"`
		JobName             interface{} `json:"jobName"`
		Description         string      `json:"description"`
		RichDescription     interface{} `json:"richDescription"`
		LectorTitle         interface{} `json:"lectorTitle"`
		RealName            interface{} `json:"realName"`
		IsTeacher           bool        `json:"isTeacher"`
		FollowCount         int         `json:"followCount"`
		FollowedCount       int         `json:"followedCount"`
		SchoolShortName     interface{} `json:"schoolShortName"`
		LogoForCertURL      interface{} `json:"logoForCertUrl"`
		SupportMooc         interface{} `json:"supportMooc"`
		SupportSpoc         interface{} `json:"supportSpoc"`
		FollowStatus        bool        `json:"followStatus"`
		SupportCommonMooc   interface{} `json:"supportCommonMooc"`
		SupportPostgradexam interface{} `json:"supportPostgradexam"`
		LectorTag           interface{} `json:"lectorTag"`
		RelType             interface{} `json:"relType"`
		IsOfflineTeacher    interface{} `json:"isOfflineTeacher"`
	} `json:"result"`
	Message string `json:"message"`
	TraceID string `json:"traceId"`
	Sampled bool   `json:"sampled"`
}

// ---------------------
type VodVideo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		VideoID        int           `json:"videoId"`
		Duration       int           `json:"duration"`
		Name           string        `json:"name"`
		DefaultQuality int           `json:"defaultQuality"`
		VideoImgURL    interface{}   `json:"videoImgUrl"`
		SrtCaptions    []interface{} `json:"srtCaptions"`
		CdnPoints      []struct {
			IP      string `json:"ip"`
			Isp     string `json:"isp"`
			IspName string `json:"ispName"`
		} `json:"cdnPoints"`
		VideoThumbnail interface{} `json:"videoThumbnail"`
		Videos         []struct {
			Quality          int    `json:"quality"`
			Size             int    `json:"size"`
			VideoURL         string `json:"videoUrl"`
			Format           string `json:"format"`
			SecondaryEncrypt bool   `json:"secondaryEncrypt"`
			E                bool   `json:"e"`
			V                string `json:"v"`
			K                string `json:"k"`
		} `json:"videos"`
	} `json:"result"`
	UUID interface{} `json:"uuid"`
}

// -------------------------------------
type Video struct {
	Code    int         `json:"code"`
	Result  VideoResult `json:"result"`
	Message string      `json:"message"`
	TraceID string      `json:"traceId"`
	Sampled bool        `json:"sampled"`
}

type VideoSignDto struct {
	Status      int         `json:"status"`
	VideoID     int         `json:"videoId"`
	Duration    int         `json:"duration"`
	VideoImgURL interface{} `json:"videoImgUrl"`
	Signature   string      `json:"signature"`
	Name        string      `json:"name"`
}

type VideoResult struct {
	VideoSignDto   VideoSignDto `json:"videoSignDto"`
	LearnVideoTime int          `json:"learnVideoTime"`
	LessonUnitID   int          `json:"lessonUnitId"`
	LessonUnitName string       `json:"lessonUnitName"`
	Code           int          `json:"code"`
	ContentType    int          `json:"contentType"`
}

// -------------------------------------
type MyMocTermDto struct {
	CourseName string
	Chapters   []MyChapter
}

type MyChapter struct {
	ChapterName string
	MyUnits     []MyUnit
	ContentType int
}

type MyUnit struct {
	UnitName    string
	ContentType int
	ContentId   int
	UnitId      int
}

// ---------------------------------------
type LastLearnedMocTermDto struct {
	Code    int    `json:"code"`
	Result  Result `json:"result"`
	Message string `json:"message"`
	TraceID string `json:"traceId"`
	Sampled bool   `json:"sampled"`
}

type Units struct {
	ID                 int         `json:"id"`
	GmtCreate          int64       `json:"gmtCreate"`
	GmtModified        int64       `json:"gmtModified"`
	Name               string      `json:"name"`
	Position           int         `json:"position"`
	LessonID           int         `json:"lessonId"`
	ChapterID          int         `json:"chapterId"`
	TermID             int         `json:"termId"`
	ContentType        int         `json:"contentType"`
	ContentID          int         `json:"contentId"`
	UnitID             interface{} `json:"unitId"`
	Live               interface{} `json:"live"`
	FreePreview        interface{} `json:"freePreview"`
	DurationInSeconds  interface{} `json:"durationInSeconds"`
	LearnCount         interface{} `json:"learnCount"`
	ViewStatus         int         `json:"viewStatus"`
	NameAlias          interface{} `json:"nameAlias"`
	PhotoURL           interface{} `json:"photoUrl"`
	ResourceInfo       interface{} `json:"resourceInfo"`
	Attachments        interface{} `json:"attachments"`
	AnchorQuestions    interface{} `json:"anchorQuestions"`
	JSONContent        interface{} `json:"jsonContent"`
	LiveInfoDto        interface{} `json:"liveInfoDto"`
	YktRelatedLiveInfo interface{} `json:"yktRelatedLiveInfo"`
}

type Lessons struct {
	ID              int         `json:"id"`
	GmtCreate       int64       `json:"gmtCreate"`
	GmtModified     int64       `json:"gmtModified"`
	Name            string      `json:"name"`
	Position        int         `json:"position"`
	TermID          int         `json:"termId"`
	ChapterID       int         `json:"chapterId"`
	ContentType     int         `json:"contentType"`
	ContentID       interface{} `json:"contentId"`
	IsTestChecked   bool        `json:"isTestChecked"`
	Units           []Units     `json:"units"`
	ReleaseTime     int64       `json:"releaseTime"`
	ViewStatus      int         `json:"viewStatus"`
	TestDraftStatus int         `json:"testDraftStatus"`
	Test            interface{} `json:"test"`
}

type Chapters struct {
	ID                  int           `json:"id"`
	GmtCreate           int64         `json:"gmtCreate"`
	GmtModified         int64         `json:"gmtModified"`
	Name                string        `json:"name"`
	Position            int           `json:"position"`
	TermID              int           `json:"termId"`
	ContentType         int           `json:"contentType"`
	ContentID           interface{}   `json:"contentId"`
	ReleaseTime         int64         `json:"releaseTime"`
	Published           bool          `json:"published"`
	Lessons             []Lessons     `json:"lessons"`
	Homeworks           []interface{} `json:"homeworks"`
	Quizs               []interface{} `json:"quizs"`
	HasFreePreviewVideo bool          `json:"hasFreePreviewVideo"`
	Exam                interface{}   `json:"exam"`
	DraftStatus         int           `json:"draftStatus"`
}

type MocTermDto struct {
	Times                 int           `json:"times"`
	ID                    int           `json:"id"`
	GmtCreate             interface{}   `json:"gmtCreate"`
	GmtModified           interface{}   `json:"gmtModified"`
	CourseID              int           `json:"courseId"`
	CloseVisableStatus    int           `json:"closeVisableStatus"`
	StartTime             int64         `json:"startTime"`
	Duration              interface{}   `json:"duration"`
	EndTime               int64         `json:"endTime"`
	PublishStatus         interface{}   `json:"publishStatus"`
	CourseLoad            interface{}   `json:"courseLoad"`
	SmallPhoto            interface{}   `json:"smallPhoto"`
	BigPhoto              interface{}   `json:"bigPhoto"`
	FirstPublishTime      interface{}   `json:"firstPublishTime"`
	EnrollCount           interface{}   `json:"enrollCount"`
	LessonsCount          interface{}   `json:"lessonsCount"`
	CourseName            string        `json:"courseName"`
	CoverPhoto            string        `json:"coverPhoto"`
	Chapters              []Chapters    `json:"chapters"`
	Exams                 []interface{} `json:"exams"`
	Mode                  int           `json:"mode"`
	FromTermID            interface{}   `json:"fromTermId"`
	SchoolID              int           `json:"schoolId"`
	HasFreePreviewVideo   bool          `json:"hasFreePreviewVideo"`
	VideoID               interface{}   `json:"videoId"`
	Description           interface{}   `json:"description"`
	BgKnowledge           interface{}   `json:"bgKnowledge"`
	Outline               interface{}   `json:"outline"`
	OutlineStructure      interface{}   `json:"outlineStructure"`
	ReommendRead          interface{}   `json:"reommendRead"`
	CourseStyle           interface{}   `json:"courseStyle"`
	Faq                   interface{}   `json:"faq"`
	JSONContent           interface{}   `json:"jsonContent"`
	Requirements          interface{}   `json:"requirements"`
	RequirementsForCert   interface{}   `json:"requirementsForCert"`
	DescriptionForCert    interface{}   `json:"descriptionForCert"`
	Target                interface{}   `json:"target"`
	MobDescription        interface{}   `json:"mobDescription"`
	ChiefLectorDto        interface{}   `json:"chiefLectorDto"`
	StaffLectorDtos       interface{}   `json:"staffLectorDtos"`
	StaffAssistDtos       interface{}   `json:"staffAssistDtos"`
	ChiefLector           interface{}   `json:"chiefLector"`
	StaffLectors          interface{}   `json:"staffLectors"`
	StaffAssists          interface{}   `json:"staffAssists"`
	ChargeableCert        interface{}   `json:"chargeableCert"`
	SpecialChargeableTerm bool          `json:"specialChargeableTerm"`
	TimeToFreeze          interface{}   `json:"timeToFreeze"`
	PreviousCourseDtos    interface{}   `json:"previousCourseDtos"`
	AnnouncementDtos      interface{}   `json:"announcementDtos"`
	Enrolled              interface{}   `json:"enrolled"`
	HasPaid               interface{}   `json:"hasPaid"`
	AchievementStatus     interface{}   `json:"achievementStatus"`
	FromTermMode          interface{}   `json:"fromTermMode"`
	ApplyMoocStatus       interface{}   `json:"applyMoocStatus"`
	OriginCopyRightTermID interface{}   `json:"originCopyRightTermId"`
	ExtraInfo             interface{}   `json:"extraInfo"`
	NeedPassword          bool          `json:"needPassword"`
	Copied                interface{}   `json:"copied"`
	CopyTime              interface{}   `json:"copyTime"`
	Price                 interface{}   `json:"price"`
	IsStart               bool          `json:"isStart"`
	IsEnd                 bool          `json:"isEnd"`
	OriginalPrice         interface{}   `json:"originalPrice"`
	OutLineStructureDtos  interface{}   `json:"outLineStructureDtos"`
	ProductType           interface{}   `json:"productType"`
	Channel               int           `json:"channel"`
	PositionStatus        interface{}   `json:"positionStatus"`
	DetailDraftStatus     int           `json:"detailDraftStatus"`
	WebVisible            interface{}   `json:"webVisible"`
}

type Result struct {
	LastLearnUnitID int        `json:"lastLearnUnitId"`
	MocTermDto      MocTermDto `json:"mocTermDto"`
}
