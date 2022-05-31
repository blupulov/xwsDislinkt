import { Injectable } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { UserService } from "./user.service";
import { AllUserPostsResponse } from "src/models/post.model";
import { PostComment } from "src/models/postComment.model";
import { NewPost } from "src/models/newPost.model";

@Injectable({
  providedIn: 'root'
})

export class PostService {

  constructor (private http: HttpClient,
    private userService: UserService) { }

  private apiUrl = 'http://localhost:8001/post';

  getAllUserPosts(userId: string | null | undefined) {
    return this.http.get<AllUserPostsResponse>(this.apiUrl + '/user/' + userId);
  }

  commentPost(postId: String, comment: PostComment) {
    return this.http.put(this.apiUrl + '/' + postId + '/comment', JSON.stringify(comment));
  }

  posting(newPost: NewPost) {
    return this.http.post<any>(this.apiUrl, JSON.stringify(newPost));
  }

  like(postId: String) {
    return this.http.put(this.apiUrl + '/' + postId + '/like/' + this.userService.getUserId(), {})
  }
}