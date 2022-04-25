package handler

import (
	pb "github.com/blupulov/xwsDislinkt/common/proto/services/post-service"
	"github.com/blupulov/xwsDislinkt/post-service/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapPbFromPost(post *model.Post) *pb.Post {
	pbPost := &pb.Post{
		Id:                post.Id.Hex(),
		PostComment:       post.PostComment,
		PostImage:         post.PostImage,
		PostOwnerId:       post.PostOwnerId,
		PostDate:          timestamppb.New(post.PostDate),
		PostLikeNumber:    int32(post.LikeNumber),
		PostDislikeNumber: int32(post.DislikeNumber),
		FansIds:           post.FansIds,
		HatersIds:         post.HatersIds,
	}

	for _, c := range post.Comments {
		pbComment := &pb.Comment{
			CommentOwnerId: c.CommentOwnerId,
			CommentContent: c.CommentContent,
			CreationDate:   timestamppb.New(c.CreationDate).String(),
		}
		pbPost.Comments = append(pbPost.Comments, pbComment)
	}

	return pbPost
}

func mapPostFromPb(pbPost *pb.Post) *model.Post {
	return &model.Post{
		PostComment:   pbPost.PostComment,
		PostImage:     pbPost.PostImage,
		PostOwnerId:   pbPost.PostOwnerId,
		Comments:      make([]model.Comment, 0),
		FansIds:       make([]string, 0),
		HatersIds:     make([]string, 0),
		LikeNumber:    0,
		DislikeNumber: 0,
	}
}
