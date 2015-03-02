package collaboration

import (
	"fmt"
	"koding/db/mongodb/modelhelper"
	"math/rand"
	apimodels "socialapi/models"
	"socialapi/rest"
	"socialapi/workers/collaboration/models"
	"socialapi/workers/common/runner"
	"testing"
	"time"

	"socialapi/workers/helper"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCollaborationDriveService(t *testing.T) {
	r := runner.New("collaboration-drive-tests")
	err := r.Init()
	if err != nil {
		panic(err)
	}

	defer r.Close()

	modelhelper.Initialize(r.Conf.Mongo)
	defer modelhelper.Close()

	redisConn := helper.MustInitRedisConn(r.Conf)
	defer redisConn.Close()

	handler := New(r.Log, redisConn, r.Conf, r.Kite)

	Convey("while pinging collaboration", t, func() {
		// owner
		owner := apimodels.NewAccount()
		owner.OldId = AccountOldId.Hex()
		owner, err := rest.CreateAccount(owner)
		So(err, ShouldBeNil)
		So(owner, ShouldNotBeNil)

		ownerSession, err := apimodels.FetchOrCreateSession(owner.Nick)
		So(err, ShouldBeNil)
		So(ownerSession, ShouldNotBeNil)

		rand.Seed(time.Now().UnixNano())

		req := &models.Ping{
			AccountId: 1,
			FileId:    fmt.Sprintf("%d", rand.Int63()),
		}

		Convey("while testing drive operations", func() {
			req := req
			req.CreatedAt = time.Now().UTC()
			Convey("should be able to create the file", func() {
				f, err := createTestFile(handler)
				So(err, ShouldBeNil)
				req.FileId = f.Id
				Convey("should be able to get the created file", func() {
					f2, err := handler.getFile(f.Id)
					So(err, ShouldBeNil)
					So(f2, ShouldNotBeNil)
					Convey("should be able to delete the created file", func() {
						err = handler.deleteFile(req.FileId)
						So(err, ShouldBeNil)
						Convey("should not be able to get the deleted file", func() {
							f2, err = handler.getFile(f.Id)
							So(err, ShouldNotBeNil)
							So(f2, ShouldBeNil)
						})
						Convey("deleting the deleted file should not give error", func() {
							err = handler.deleteFile(req.FileId)
							So(err, ShouldBeNil)
						})
					})
				})
			})
		})
	})
}
