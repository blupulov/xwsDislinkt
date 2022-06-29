import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-find-user',
  templateUrl: './find-user.component.html',
  styleUrls: ['./find-user.component.css']
})
export class FindUserComponent implements OnInit {
  username : String = '';

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void { }

  onSubmit() {
    this.userService.getUserByUsername(this.username).subscribe(
      res => {
          this.username = '';
          if(res.id == this.userService.getUserId())
            this.router.navigateByUrl('userProfile')
          else {
            this.userService.setSelectedUserId(res.id)
            this.router.navigateByUrl('mid')
          }
      }, err => {
        alert('user not found')
      }
    )
  }

}
