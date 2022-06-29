import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from 'src/models/user.model';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-selected-user-profile',
  templateUrl: './selected-user-profile.component.html',
  styleUrls: ['./selected-user-profile.component.css']
})
export class SelectedUserProfileComponent implements OnInit {
  user: User = new User();
  
  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    if(this.userService.getSelectedUserId() == this.userService.getUserId()){
      this.router.navigateByUrl('userProfile')
    }else {
      this.getUser()
    }
  }

  getUser(): void {
    this.userService.getUserById(this.userService.getSelectedUserId()).subscribe(
      res => {
        if(res.role == '1'){
          this.router.navigateByUrl('userProfile')
        }
        this.user = res
      },
      err => {
        alert('problem with loading user')
      }
    )
  }


}
