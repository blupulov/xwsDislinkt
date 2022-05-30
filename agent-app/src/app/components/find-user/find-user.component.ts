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

  ngOnInit(): void {

  }

  onSubmit() {
    this.userService.GetUserByUsername(this.username).subscribe(
      res => {
          this.username = '';
          if(res.user.id == this.userService.getUserId())
            this.router.navigateByUrl('profile')
          else {
            this.userService.setSelectedUserId(res.user.id)
            this.router.navigateByUrl('selectedUserProfile')
            window.location.reload()
          }
      }, err => {
        alert('user not found')
      }
    )
  }
}
