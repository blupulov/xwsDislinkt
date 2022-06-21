import { Component, OnInit } from '@angular/core';
import { User } from 'src/models/user.model';
import { UserService } from 'src/services/user.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-post-haters',
  templateUrl: './post-haters.component.html',
  styleUrls: ['./post-haters.component.css']
})
export class PostHatersComponent implements OnInit {

  haters: User[] = [];

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    this.loadHaters()
  }

  loadHaters() {
    var hatersIds = this.userService.getHatersIds()

    this.userService.getManyUsersById(hatersIds).subscribe(
      res => {
        this.haters = res.users
      }, err => {
        alert('problem with loading haters')
      }
    )
  }

  showProfile(userId: String) {
    if(userId == this.userService.getUserId()){
      this.router.navigateByUrl('profile')
    } else {
      this.userService.setSelectedUserId(userId)
      this.router.navigateByUrl('selectedUserProfile')
    }
  }

}
