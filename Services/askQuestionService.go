package Services

import (
	"time"

	"github.com/bigby/project/Models"
	"github.com/bigby/project/Utils"
)

type AskQuestion struct {
	AQid      int64     `gorm:"primary_key;column:AQid;AUTO_INCREMENT"`
	Sid       int64     `gorm:"column:Sid"`
	Eid       int64     `gorm:"column:Eid"`
	AQtime    time.Time `gorm:"column:AQtime"`
	AQremark  string    `gorm:"column:AQremark"`
	Snickname string
	Etitle    string
}

func (aq *AskQuestion) Insert() (AQid int64, err error) {
	var aqModel Models.AskQuestion

	aqModel.Sid = aq.Sid
	aqModel.Eid = aq.Eid
	aqModel.AQtime = aq.AQtime
	aqModel.AQremark = aq.AQremark

	AQid, err = aqModel.Insert()

	return

}

func (aq *AskQuestion) QueryByAQid(AQid int64) (result AskQuestion, err error) {
	var aqModel Models.AskQuestion

	tmpAQ, err := aqModel.QueryByAQid(AQid)

	if err != nil {
		return
	}

	err = Utils.CopyFields(&result, tmpAQ)
	if err != nil {
		return
	}

	return
}

func (aq *AskQuestion) QueryFromTeacher(Tid int64) (result []AskQuestion, err error) {
	var aqModel Models.AskQuestion

	tmpAQ, err := aqModel.QueryFromTeacher(Tid)

	result = make([]AskQuestion, len(tmpAQ))
	// 把 askquestion model 转为 askquestion service
	for i := 0; i < len(tmpAQ); i++ {
		err = Utils.CopyFields(&result[i], tmpAQ[i])
		if err != nil {
			return
		}
	}

	var studentModel Models.Student
	var exerciseModel Models.Exercise
	for i := 0; i < len(result); i++ {
		student, err := studentModel.QueryBySid(result[i].Sid)
		if err != nil {
			return nil, err
		}

		result[i].Snickname = student.Snickname

		exercise, err := exerciseModel.QueryByEid(result[i].Eid)
		if err != nil {
			return nil, err
		}

		result[i].Etitle = exercise.Etitle
	}

	if err != nil {
		return nil, err
	}

	return
}