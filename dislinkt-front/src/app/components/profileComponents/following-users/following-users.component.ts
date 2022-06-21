import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from 'src/models/user.model';
import { CustomService } from 'src/services/custom.service';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-following-users',
  templateUrl: './following-users.component.html',
  styleUrls: ['./following-users.component.css']
})

export class FollowingUsersComponent implements OnInit {
  users: User[] = [];
  
  constructor(private customService: CustomService, private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    this.loadUsers()
  }

  loadUsers() {
    //this.userService.getSelectedUserId()
    this.customService.getFollowingUsers(this.userService.getUserId()?.toString()).subscribe(
      res => {
        this.users = res
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
