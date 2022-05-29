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

  openFans() {
    alert('adksahdkjash')
  }

  onSubmit() {
    //datum treba prilagotiti
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