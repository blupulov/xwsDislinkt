import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Post } from 'src/models/post.model';
import { PostComment } from 'src/models/postComment.model';
import { PostService } from 'src/services/post.service';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-user-posts',
  templateUrl: './user-posts.component.html',
  styleUrls: ['./user-posts.component.css']
})
export class UserPostsComponent implements OnInit {
  posts: Array<Post> = [];
  selectedPostId: String = '';
  comment: String = '';

  constructor(private postService: PostService, private userService: UserService, private router:Router) { }

  ngOnInit(): void {
    this.getAllUserPosts()
  }

  alertPost(post: Post) {
    alert(JSON.stringify(post.postImage))
  }

  getAllUserPosts() {
    this.postService.getAllUserPosts(this.userService.getUserId()?.toString()).subscribe(
      res => {
        this.posts = res.posts
      },
      err => {
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

  showFans(fansIds: String[]) {
    if(fansIds.length > 0) {
      this.userService.setFansIds(fansIds)
      this.router.navigateByUrl('postFans')
    } else {
      alert('there is no fans')
    }
  }

  showHaters(hatersIds: String[]) {
    if(hatersIds.length > 0) {
      this.userService.setHatersIds(hatersIds)
      this.router.navigateByUrl('postHaters')
    } else {
      alert('there is no haters')
    }

  }

  deletePost(postId: String) {
    this.postService.deleteUserPost(postId).subscribe(
      res => {
        alert('post deleted')
        this.getAllUserPosts()
      },err =>{
        alert('problem with deleting')
      }
    )
  }

  onSubmit() {
    let comment = new PostComment();
    comment.commentOwnerId = new String(this.userService.getUserId());
    comment.commentContent = this.comment;

    this.postService.commentPost(this.selectedPostId, comment).subscribe(
      res => {
        this.posts.find(post => post.id === this.selectedPostId)?.comments.push(comment);
        this.comment = '';
      }, err => {
        alert('problem with sending comment')
      }
    )
  }
}
