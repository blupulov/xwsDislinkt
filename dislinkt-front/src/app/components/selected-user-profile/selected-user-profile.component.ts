import { TOUCH_BUFFER_MS } from '@angular/cdk/a11y/input-modality/input-modality-detector';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
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
    public customService: CustomService, private router: Router) { }

  ngOnInit(): void {
    this.getUser()
  }

  getUser(): void {
    this.userService.getUserById(this.userService.getSelectedUserId()).subscribe(
      res => {
        if(res.user.role == "ADMIN"){
          this.router.navigateByUrl('profile')
        }
        this.user = res.user
      },
      err => {
        alert('problem with loading user')
      }
    )
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
