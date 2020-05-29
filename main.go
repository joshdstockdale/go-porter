package main

import (
	"go-porter/porter"
)

func main() {
	tks := []porter.Task{porter.Task{
		Id: 1, Action: "PUT", Recipient: "post",
		Package: []byte(`{"title": "New Title 1"}`),
	}, porter.Task{
		Id: 2, Action: "PUT", Recipient: "post",
		Package: []byte(`{"title": "New Title 2"}`),
	}, porter.Task{
		Id: 3, Action: "PUT", Recipient: "post",
		Package: []byte(`{"title": "New Title 3"}`),
	}, porter.Task{
		Id: 4, Action: "PUT", Recipient: "post",
		Package: []byte(`{"title": "New Title 4"}`),
	}, porter.Task{
		Id: 5, Action: "PUT", Recipient: "post",
		Package: []byte(`{"title": "New Title 5"}`),
	}, porter.Task{
		Id: 6, Action: "PUT", Recipient: "post",
		Package: []byte(`{"title": "New Title 6"}`),
	}, porter.Task{
		Id: 7, Action: "PUT", Recipient: "post",
		Package: []byte(`{"title": "New Title 7"}`),
	}, porter.Task{
		Id: 8, Action: "PUT", Recipient: "post",
		Package: []byte(`{"title": "New Title 8"}`),
	}, porter.Task{
		Id: 9, Action: "PUT", Recipient: "post",
		Package: []byte(`{"title": "New Title 9"}`),
	}, porter.Task{
		Id: 10, Action: "PUT", Recipient: "post",
		Package: []byte(`{"title": "New Title 10"}`),
	}}

	porter.WakePorters(3, tks)
}
