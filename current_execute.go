package gojenkins

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type CurrentExecute struct {
	Raw     *string
	Resp    *CurrentExecuteResponse
	Jenkins *Jenkins
	Base    string
}

type CurrentExecuteResponse struct {
	TotalTaskNum int                   `json:"totalTaskNum"`
	Tasks        []*CurrentExecuteTask `json:"tasks"`
}

type CurrentExecuteTask struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	IsIdle  bool   `json:"isIdle"`
}

func (ce *CurrentExecute) Get() (int, error) {
	//ce.Resp = new(CurrentExecuteResponse)
	//ce.Resp.Tasks = make([]*CurrentExecuteTask, 0)
	//if code, err := ce.Poll(); err != nil {
	//	return code, err
	//}
	//doc, err := goquery.NewDocumentFromReader(strings.NewReader(*ce.Raw))
	//if err != nil {
	//	return 500, err
	//}
	//doc.Find("tr").Each(func(i int, selection *goquery.Selection) {
	//	if selection.Text() != "" {
	//		task := new(CurrentExecuteTask)
	//		selection.Find(".pane").Each(func(i int, selection *goquery.Selection) {
	//			switch i {
	//			case 0:
	//				task.ID = selection.Text()
	//				break
	//			case 1:
	//				task.Name = selection.Text()
	//				if task.Name == "Idle" {
	//					task.IsIdle = true
	//				} else {
	//					task.IsIdle = false
	//				}
	//				break
	//			case 2:
	//				task.Content = selection.Text()
	//				break
	//			}
	//		})
	//		ce.Resp.Tasks = append(ce.Resp.Tasks, task)
	//	}
	//	ce.Resp.TotalTaskNum = len(ce.Resp.Tasks)
	//})
	return 200, nil
}

func (n *CurrentExecute) Poll() (int, error) {
	response, err := n.Jenkins.Requester.Get(n.Base, n.Raw, nil)
	if err != nil {
		return 0, err
	}
	return response.StatusCode, nil
}
