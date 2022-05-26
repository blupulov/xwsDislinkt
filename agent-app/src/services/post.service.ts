import { Injectable } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { UserService } from "./user.service";
import { AllUserPostsResponse } from "src/models/post.model";
import { PostComment } from "src/models/postComment.model";

@Injectable({
  providedIn: 'root'
})

export class PostService {

  constructor (private http: HttpClient,
    private userService: UserService) { }

  private apiUrl = 'http://localhost:8001/post';

  getAllUserPosts() {
    return this.http.get<AllUserPostsResponse>(this.apiUrl + '/user/' + this.userService.getUserId());
  }

  commentPost(postId: String, comment: PostComment) {
    return this.http.put(this.apiUrl + '/' + postId + '/comment', JSON.stringify(comment));
  }
}