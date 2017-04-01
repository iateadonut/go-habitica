package habitica_test

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/wfernandes/go-habitica"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	mux    *http.ServeMux
	ts     *httptest.Server
	client *habitica.HabiticaClient
	ctx    context.Context
)

var _ = Describe("Task Service", func() {
	Context("endpoints", func() {
		BeforeEach(func() {
			var err error
			mux = http.NewServeMux()
			ts = httptest.NewServer(mux)
			client, err = habitica.New(
				"b0413351-405f-416f-8787-947ec1c85199",
				"api",
				habitica.WithBaseURL(ts.URL),
			)
			Expect(err).ToNot(HaveOccurred())
			ctx = context.Background()
		})

		AfterEach(func() {
			ts.Close()
		})

		Context("Get", func() {
			It("creates request with correct headers", func() {
				request := &http.Request{}
				mux.HandleFunc("/tasks/some-task-id", func(w http.ResponseWriter, r *http.Request) {
					request = r
					w.WriteHeader(http.StatusOK)
					w.Write(taskResponse)
				})
				client.Tasks.Get(ctx, "some-task-id")
				Expect(request.Method).To(Equal(http.MethodGet))
				Expect(request.UserAgent()).To(Equal(habitica.UserAgent))
				Expect(request.Header.Get("x-api-user")).To(Equal("b0413351-405f-416f-8787-947ec1c85199"))
				Expect(request.Header.Get("x-api-key")).To(Equal("api"))
				Expect(request.Header.Get("Content-Type")).To(Equal("application/json"))
			})

			It("returns a task for get", func() {
				mux.HandleFunc("/tasks/some-task-id", func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write(taskResponse)
				})

				task, err := client.Tasks.Get(ctx, "some-task-id")
				Expect(err).ToNot(HaveOccurred())
				Expect(task).ToNot(BeNil())
				Expect(task.Success).To(BeTrue())
				Expect(task.Data.Text).To(Equal("API Trial"))
				Expect(task.Data.ID).To(Equal("2b774d70-ec8b-41c1-8967-eb6b13d962ba"))
			})

			It("returns an error if it cannot decode a respose", func() {
				mux.HandleFunc("/tasks/some-task-id", func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{some bad response}`))
				})
<<<<<<< HEAD

				task, err := client.Tasks.Get(ctx, "some-task-id")
				Expect(err).To(HaveOccurred())
				Expect(task).To(BeNil())
			})
		})

		Context("User Tasks", func() {
			It("returns the tasks from the user in the header", func() {
				mux.HandleFunc("/tasks/user", func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write(userTasksResponse)
				})

				resp, err := client.Tasks.List(ctx)
				Expect(err).ToNot(HaveOccurred())
				Expect(resp).ToNot(BeNil())
				Expect(resp.Data).To(HaveLen(1))
				Expect(resp.Data[0].Text).To(Equal("Practice Task 31"))
				Expect(resp.Data[0].UserID).To(Equal("b0413351-405f-416f-8787-947ec1c85199"))
			})
		})
=======
>>>>>>> d66df3d... Adds endpoint for getting the user's tasks

				task, err := client.Tasks.Get(ctx, "some-task-id")
				Expect(err).To(HaveOccurred())
				Expect(task).To(BeNil())
			})
		})

		Context("User Tasks", func() {
			It("returns the tasks from the user in the header", func() {
				mux.HandleFunc("/tasks/user", func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write(userTasksResponse)
				})

				resp, err := client.Tasks.List(ctx)
				Expect(err).ToNot(HaveOccurred())
				Expect(resp).ToNot(BeNil())
				Expect(resp.Data).To(HaveLen(1))
				Expect(resp.Data[0].Text).To(Equal("Practice Task 31"))
				Expect(resp.Data[0].UserID).To(Equal("b0413351-405f-416f-8787-947ec1c85199"))
			})
		})
	})

})

var userTasksResponse = []byte(`
{
	"success": true,
	"data": [{
		"_id": "84c2e874-a8c9-4673-bd31-d97a1a42e9a3",
		"userId": "b0413351-405f-416f-8787-947ec1c85199",
		"alias": "prac31",
		"text": "Practice Task 31",
		"type": "daily",
		"notes": "",
		"tags": [],
		"value": 1,
		"priority": 1,
		"attribute": "str",
		"challenge": {},
		"group": {
			"assignedUsers": [],
			"approval": {
				"required": false,
				"approved": false,
				"requested": false
			}
		},
		"reminders": [{
			"time": "2017-01-13T16:21:00.074Z",
			"startDate": "2017-01-13T16:20:00.074Z",
			"id": "b8b549c4-8d56-4e49-9b38-b4dcde9763b9"
		}],
		"createdAt": "2017-01-13T16:34:06.632Z",
		"updatedAt": "2017-01-13T16:49:35.762Z",
		"checklist": [],
		"collapseChecklist": false,
		"completed": true,
		"history": [],
		"streak": 1,
		"repeat": {
			"su": false,
			"s": false,
			"f": true,
			"th": true,
			"w": true,
			"t": true,
			"m": true
		},
		"startDate": "2017-01-13T00:00:00.000Z",
		"everyX": 1,
		"frequency": "weekly",
		"id": "84c2e874-a8c9-4673-bd31-d97a1a42e9a3"
	}],
	"notifications": []
}`)

var taskResponse = []byte(`
{
    "success": true,
    "data": {
        "_id": "2b774d70-ec8b-41c1-8967-eb6b13d962ba",
        "userId": "b0413351-405f-416f-8787-947ec1c85199",
        "text": "API Trial",
        "alias": "apiTrial",
        "type": "habit",
        "notes": "",
        "tags": [],
        "value": 11.996661122825959,
        "priority": 1.5,
        "attribute": "str",
        "challenge": {
            "taskId": "5f12bfba-da30-4733-ad01-9c42f9817975",
            "id": "f23c12f2-5830-4f15-9c36-e17fd729a812"
        },
        "group": {
            "assignedUsers": [],
            "approval": {
                "required": false,
                "approved": false,
                "requested": false
            }
        },
        "reminders": [],
        "createdAt": "2017-01-12T19:03:33.495Z",
        "updatedAt": "2017-01-13T20:52:02.927Z",
        "history": [
            {
                "value": 1,
                "date": 1484248053486
            }
        ],
        "down": false,
        "up": true,
        "id": "2b774d70-ec8b-41c1-8967-eb6b13d962ba"
    },
    "notifications": []
}`)
