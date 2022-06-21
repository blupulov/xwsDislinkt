import { DatePipe } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ChangeUserInfo } from 'src/models/changeUserInfo.model';
import { User } from 'src/models/user.model';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-change-info',
  templateUrl: './change-info.component.html',
  styleUrls: ['./change-info.component.css']
})
export class ChangeInfoComponent implements OnInit {

  newUser: ChangeUserInfo = new ChangeUserInfo();
  

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    this.getUser()
  }

  getUser(): void {
    if(this.userService.getUserId() != null) {
      this.userService.getUserById(this.userService.getUserId()).subscribe(
        res => {
          this.mapUserToChangeUserInfo(res.user)
        },
        err => {
          alert('problem with loading user')
        }
      )
    }
  }

  onSubmit() {
    this.userService.changeUser(this.newUser).subscribe(
      res => {
        this.router.navigateByUrl('profile')
      }, err => {
        alert('problem with changing user')
      }
    )
  }

  mapUserToChangeUserInfo(user: User) {
    let pipe = new DatePipe('en-US');
    this.newUser.birthDate = pipe.transform(user.birthDate, 'yyyy-MM-ddThh:mm:ss')
    this.newUser.username = user.username
    this.newUser.firstName = user.firstName
    this.newUser.lastName = user.lastName
    this.newUser.email = user.email
    this.newUser.phoneNumber = user.phoneNumber
    this.newUser.biography = user.biography
  }

}
