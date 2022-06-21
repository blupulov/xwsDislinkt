import { Component, OnInit } from '@angular/core';
import { User } from 'src/models/user.model';
import { FollowingService } from 'src/services/following.service';
import { UserService } from 'src/services/user.service';
import { CustomService } from'src/services/custom.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  user: User = new User();

  constructor(private userService: UserService, private followingService: FollowingService,
    private customService: CustomService, private router: Router) { }

  ngOnInit(): void {
    this.getUser()

    if (this.userService.getUserRole() != "ADMIN"){
      this.loadFollowingUsersIds()
    } else {
      this.router.navigateByUrl('enableCompany')
    }
    
  }

  getUser(): void {
    if(this.userService.getUserId() != null) {
      this.userService.getUserById(this.userService.getUserId()).subscribe(
        res => {
          this.user = res.user
        },
        err => {
          alert('problem with loading user')
        }
      )
    }
  }

  loadFollowingUsersIds() {
    this.followingService.getAllFollowingUsersIds(this.userService.getUserId()).subscribe(
      res => {
        this.customService.setFollowing(res.usersIds)
      }, err => {
        alert('problem with loading following users ids')
      }
    )
  }
}
