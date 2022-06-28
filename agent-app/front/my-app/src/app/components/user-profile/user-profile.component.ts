import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from 'src/models/user.model';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent implements OnInit {
  user: User = new User();

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    this.getUser()

    if (this.userService.getUserRole() == "ADMIN")
      this.router.navigateByUrl('enableCompany')
  }

  getUser(): void {
    if(this.userService.getUserId() != null) {
      this.userService.getUserById(this.userService.getUserId()).subscribe(
        res => {
          this.user = res
        },
        err => {
          alert('problem with loading user')
        }
      )
    }
  }

}
