// Code generated by hertz generator. DO NOT EDIT.

package Video

import (
	video "gitee.com/derrickball/douyin/app/api/biz/handler/video"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		_douyin.GET("/feed", append(_feedMw(), video.Feed)...)
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			{
				_action := _publish.Group("/action", _actionMw()...)
				_action.POST("/", append(_publish_ctionMw(), video.PublishAction)...)
			}
			{
				_list := _publish.Group("/list", _listMw()...)
				_list.GET("/", append(_publishlistMw(), video.PublishList)...)
			}
		}
	}
}
