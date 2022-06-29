import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CompanyService } from 'src/services/company.service';
import { UserService } from 'src/services/user.service';
import { Comment } from 'src/models/company.model';

@Component({
  selector: 'app-create-comment',
  templateUrl: './create-comment.component.html',
  styleUrls: ['./create-comment.component.css']
})
export class CreateCommentComponent implements OnInit {
  newComment:Comment = new Comment();

  constructor(private companyService: CompanyService, private router: Router, private userService: UserService) { }

  ngOnInit(): void {}

  onSubmit() {
    if (parseInt(this.newComment.grade) < 1 || parseInt(this.newComment.grade) > 10)
      return;
    
      this.newComment.commentOwnerId = this.userService.getUserId()?.toString() || ""  
    this.companyService.createComment(this.newComment).subscribe(
      res => {
        this.router.navigateByUrl('companyProfile')
      }, err => {
        alert('problem with creating comment')
      }
    )
  }

}
