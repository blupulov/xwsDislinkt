import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NewPost } from 'src/models/newPost.model';
import { PostService } from 'src/services/post.service';

@Component({
  selector: 'app-create-post',
  templateUrl: './create-post.component.html',
  styleUrls: ['./create-post.component.css']
})
export class CreatePostComponent implements OnInit {
  image: String = '';
  comment: String = '';

  public newPost: NewPost = new NewPost();

  constructor(private postService: PostService, private router: Router) {}

  ngOnInit(): void {
  }

  onSubmit(form: any) {
    
  }

}
