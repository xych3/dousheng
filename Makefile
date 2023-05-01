rpc-user-gen:
	kitex -I idl/rpc -module gitee.com/derrickball/douyin idl/rpc/userrpc.proto
rpc-video-gen:
	kitex -I idl/rpc -module gitee.com/derrickball/douyin idl/rpc/videorpc.proto

rpc-user-server:
	cd app/user && kitex -I ../../idl/rpc -use gitee.com/derrickball/douyin/kitex_gen -service video ../../idl/rpc/userrpc.proto
rpc-video-server:
	cd app/video && kitex -I ../../idl/rpc -use gitee.com/derrickball/douyin/kitex_gen -service video ../../idl/rpc/videorpc.proto

rpc-user-server:
	kitex -I ../../idl/rpc/ -service user_rpc -module gitee.com/derrickball/douyin user_rpc.proto

rpc-comment-server:
	kitex -I ../../idl/rpc/ -service comment_rpc -module gitee.com/derrickball/douyin comment_rpc.proto

rpc-relation-server:
	kitex -I ../../idl/rpc/ -service relationt_rpc -module gitee.com/derrickball/douyin relation_rpc.proto

rpc-favorite-server:
	cd app/favorite && kitex -I ../../idl/rpc/ -service favoritet_rpc -module gitee.com/derrickball/douyin ../../idl/rpc/favorite_rpc.proto

rpc-video-server:
	kitex -I ../../idl/rpc/ -service video_rpc -module gitee.com/derrickball/douyin video_rpc.proto

hz-video-model:
	cd app/api && hz update --no_recurse -I ../../idl --model_dir gitee.com/derrickball/douyin/app/api/biz/model -idl ../../idl/api/video.proto
	kitex -I ../../idl/rpc/ -service favoritet_rpc -module gitee.com/derrickball/douyin favorite_rpc.proto

hz-video-model:
	cd app/api && hz update --no_recurse -I ../../idl --model_dir gitee.com/derrickball/douyin/app/api/biz/model -idl ../../idl/api/user.proto
	kitex -I ../../idl/rpc/ -service favoritet_rpc -module gitee.com/derrickball/douyin favorite_rpc.proto

rpc-chat-server:
	kitex -I ../../idl/rpc/ -service chat_rpc -module gitee.com/derrickball/douyin chat_rpc.proto
