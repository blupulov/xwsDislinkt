import { VariableAst } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import { Post } from 'src/models/post.model';
import { PostComment } from 'src/models/postComment.model';
import { PostService } from 'src/services/post.service';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-selected-user-posts',
  templateUrl: './selected-user-posts.component.html',
  styleUrls: ['./selected-user-posts.component.css']
})
export class SelectedUserPostsComponent implements OnInit {
  posts: Array<Post> = [];
  selectedPostId: String = '';
  comment: String = '';

  constructor(private postService: PostService, private userService: UserService) { }

  ngOnInit(): void {
    this.getAllUserPosts()
  }

  getAllUserPosts() {
    this.postService.getAllUserPosts(this.userService.getSelectedUserId()).subscribe(
      res => {
        this.posts = res.posts
      },
      () => {
        alert('loading posts problem')
      }
    )
  }

  showPostComments(postId: String) {
    if(this.selectedPostId === postId) {
      this.selectedPostId = '';
    } else {
      this.selectedPostId = postId;
    }
  }

  openFans() {
    alert('adksahdkjash')
  }

  openHaters() {
    alert('proba')
  }

  onSubmit() {
    let comment = new PostComment();
    comment.commentOwnerId = new String(this.userService.getUserId());
    comment.commentContent = this.comment;

    this.postService.commentPost(this.selectedPostId, comment).subscribe(
      () => {
        this.posts.find(post => post.id === this.selectedPostId)?.comments.push(comment);
        this.comment = '';
      }, () => {
        alert('problem with sending comment')
      }
    )
  }

  findPostById(postId: String): Post {
    for(let post of this.posts)
      if (post.id == postId)
        return post
    return new Post()
  }

  isUserLikePost(postId: String): boolean {
    let post = this.findPostById(postId)

    for (let fanId of post.fansIds)
      if(fanId == this.userService.getUserId())
        return true;

    return false;
  }

  isUserDislikePost(postId: String): boolean {
    let post = this.findPostById(postId)
    
    for (let haterId of post.hatersIds)
      if(haterId == this.userService.getUserId())
        return true;

    return false;
  }

  like(postId: String){
    this.postService.like(postId).subscribe(
      res => {
        this.addUserIdToPostFans(postId)
      }, err => {
        alert('problem with liking post')
      }
    )
  }

  addUserIdToPostFans(postId: String) {
    for(let post in this.posts){
      if(this.posts[post].id == postId) {
        this.posts[post].fansIds.push(this.userService.getUserId() || "")
        this.posts[post].postLikeNumber++
        for(let haterId in this.posts[post].hatersIds){
          if(this.posts[post].hatersIds[haterId] == this.userService.getUserId()){
            this.posts[post].hatersIds.splice(parseInt(haterId), 1)
            this.posts[post].postDislikeNumber--
          }
        }
      }
    }
  }

  dislike(postId: String) {
    this.postService.dislike(postId).subscribe(
      res => {
        this.addUserIdToPostHaters(postId)
      }, err => {
        alert('problem with disliking post')
      }
    )
  }

  addUserIdToPostHaters(postId: String) {
    for(let post in this.posts) {
      if(this.posts[post].id == postId) {
        this.posts[post].hatersIds.push(this.userService.getUserId() || "")
        this.posts[post].postDislikeNumber++
        for(let fanId in this.posts[post].fansIds) {
          if(this.posts[post].fansIds[fanId] == this.userService.getUserId()) {
            this.posts[post].fansIds.splice(parseInt(fanId), 1)
            this.posts[post].postLikeNumber--
          }
        }
      }
    }
  }
  
}
