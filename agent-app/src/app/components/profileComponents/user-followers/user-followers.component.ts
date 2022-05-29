import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from 'src/models/user.model';
import { CustomService } from 'src/services/custom.service';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-user-followers',
  templateUrl: './user-followers.component.html',
  styleUrls: ['./user-followers.component.css']
})
export class UserFollowersComponent implements OnInit {
  users: User[] = [];
  
  constructor(private customService: CustomService, private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    this.loadUsers()
  }

  loadUsers() {
    this.customService.getUserFollowers(this.userService.getUserId()?.toString()).subscribe(
      res => {
        this.users = res
        console.log(this.users)
      }, err => {
        alert('problem with loading users')
      }
    )
  }

  showProfile(userId: String) {
    this.userService.setSelectedUserId(userId)
    this.router.navigateByUrl('selectedUserProfile')
  }

}
