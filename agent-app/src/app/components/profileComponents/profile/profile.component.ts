import { Component, OnInit } from '@angular/core';
import { User } from 'src/models/user.model';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  user: User = new User();

  constructor(private userService: UserService) { }

  ngOnInit(): void {
    this.getUser()
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
}
