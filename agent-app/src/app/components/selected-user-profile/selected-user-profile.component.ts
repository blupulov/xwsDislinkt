import { Component, OnInit } from '@angular/core';
import { User } from 'src/models/user.model';
import { CustomService } from 'src/services/custom.service';
import { FollowingService } from 'src/services/following.service';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-selected-user-profile',
  templateUrl: './selected-user-profile.component.html',
  styleUrls: ['./selected-user-profile.component.css']
})
export class SelectedUserProfileComponent implements OnInit {

  user: User = new User();

  constructor(private userService: UserService, private followingService: FollowingService,
    public customService: CustomService) { }

  ngOnInit(): void {
    this.getUser()
  }

  getUser(): void {
    if(this.userService.getUserId() != null) {
      this.userService.getUserById(this.userService.getSelectedUserId()).subscribe(
        res => {
          this.user = res.user
        },
        err => {
          alert('problem with loading user')
        }
      )
    }
  }

  follow(){
    this.followingService.follow().subscribe(
      res => {
        this.customService.pushFollowing()
      },
      err => {
        alert('problem with following')
      }
    )
  }

  unFollow() {
    this.followingService.unFollow().subscribe(
      res => {
        this.customService.popFollowing()
      },
      err => {
        alert('problem with following')
      }
    )
  }

}
