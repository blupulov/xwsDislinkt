import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from 'src/models/user.model';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-post-fans',
  templateUrl: './post-fans.component.html',
  styleUrls: ['./post-fans.component.css']
})
export class PostFansComponent implements OnInit {

  fans: User[] = [];

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    this.loadFans()
  }

  loadFans() {
    var fansIds = this.userService.getFansIds()

    this.userService.getManyUsersById(fansIds).subscribe(
      res => {
        this.fans = res.users
      }, err => {
        alert('problem with loading fans')
      }
    )
  }

  showProfile(userId: String) {
    this.userService.setSelectedUserId(userId)
    this.router.navigateByUrl('selectedUserProfile')
  }
  
}
