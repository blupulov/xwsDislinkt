import { PostComment } from "./postComment.model";

export class Post {
  id: String = '';
  postComment: String = '';
  postImage: String = '';
  postOwnerId: String = '';
  postDate: Date = new Date();
  postLikeNumber: number = 0;
  postDislikeNumber: number = 0;
  fansIds: Array<String> = [];
  hatersIds: Array<String> = [];
  comments: Array<PostComment> = [];
}

//TODO: dodati reposnove za ospale metode ako treba
export class AllUserPostsResponse {
  posts: Array<Post> = [];
}