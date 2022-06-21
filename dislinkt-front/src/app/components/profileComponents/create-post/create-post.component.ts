import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NewPost } from 'src/models/newPost.model';
import { PostService } from 'src/services/post.service';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-create-post',
  templateUrl: './create-post.component.html',
  styleUrls: ['./create-post.component.css']
})
export class CreatePostComponent implements OnInit {
  public newPost: NewPost = new NewPost();
  url = "./assets/welcome.jpg"

  constructor(private postService: PostService, private userService:UserService, private router: Router) {}

  ngOnInit(): void {
    this.newPost.postOwnerId = this.userService.getUserId() || "";
  }

  onSelectFile(image: any) {
    if(image.target.files) {
      var reader = new FileReader();
      reader.readAsDataURL(image.target.files[0]);
      reader.onload = (event: any) => {
        this.url = event.target.result;
        this.newPost.postImage = event.target.result;
      }
    }
  }

  onSubmit(form: any) {
    this.postService.posting(this.newPost).subscribe(
      res => {
        alert('post created')
        this.router.navigateByUrl('profile')
      },
      err => {
        alert('problem with creating post')
      }
    )
  }

}
