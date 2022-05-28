import { Component, OnInit } from '@angular/core';
import { User } from 'src/models/user.model';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-selected-user-profile',
  templateUrl: './selected-user-profile.component.html',
  styleUrls: ['./selected-user-profile.component.css']
})
export class SelectedUserProfileComponent implements OnInit {

  user: User = new User();
 // user =   this.userService.getUserById('6270fcdf067582baca1a9a46')

  constructor(private userService: UserService) { }

  selectedUserId: String = '6270fcdf067582baca1a9a46';

  ngOnInit(): void {
    this.getUser()
  }

  getUser(): void {
    if(this.userService.getUserId() != null) {
      this.userService.getUserById(this.selectedUserId).subscribe(
        res => {
         // alert('user:' +  this.user.firstName)
          this.user = res.user
          console.log("ime: " + this.user.firstName)

        },
        err => {
          alert('problem with loading user')
        }
      )
    }
  }

}
